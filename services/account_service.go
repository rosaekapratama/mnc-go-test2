package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/rosaekapratama/go-starter/constant/location"
	"github.com/rosaekapratama/go-starter/database"
	"github.com/rosaekapratama/go-starter/log"
	"github.com/rosaekapratama/go-starter/response"
	"github.com/rosaekapratama/mnc-go-test2/constants"
	repoModel "github.com/rosaekapratama/mnc-go-test2/models/repo"
	restModel "github.com/rosaekapratama/mnc-go-test2/models/rest"
	"github.com/rosaekapratama/mnc-go-test2/repositories"
	"time"
)

func NewAccountService(_ context.Context, accountRepository repositories.AccountRepository, transactionRepository repositories.TransactionRepository) AccountService {
	return &accountServiceImpl{
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
	}
}

func (s *accountServiceImpl) Topup(ctx context.Context, req *restModel.TopupRequest) (res *restModel.BaseResponse[*restModel.TopUpResponse], err error) {
	res = &restModel.BaseResponse[*restModel.TopUpResponse]{}

	// Init tx
	tx, err := database.Manager.Begin(ctx, "playground")
	if err != nil {
		res.Message = response.GeneralError.Description()
		return
	}
	defer tx.Rollback()

	if req.Amount <= 0 {
		res.Message = "Amount must be greater than zero"
		return
	}

	userIdStr := ctx.Value("userId").(string)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		log.Errorf(ctx, err, "Invalid user ID, userId=%s", userIdStr)
		res.Message = response.GeneralError.Description()
		return
	}

	// Get saving account by user id from token
	account, err := s.accountRepository.FindSavingByUserId(ctx, userId)
	if err != nil {
		log.Errorf(ctx, err, "error on FindSavingByUserId(), userId=%s", userId.String())
		res.Message = response.GeneralError.Description()
		return
	}

	if account == nil {
		res.Message = "User doesn't have a SAVING account"
		return
	}

	// Topup account
	balanceBefore := account.Balance
	account, err = s.accountRepository.Topup(ctx, tx, account.ID, req.Amount)
	if err != nil {
		log.Error(ctx, err, "error on s.accountRepository.Topup, accountId=%s", account.ID.String())
		res.Message = response.GeneralError.Description()
		return
	}

	// Add transaction history
	transaction := &repoModel.Transaction{
		ID:        uuid.New(),
		AccountID: account.ID,
		Type:      "TOPUP",
		CR:        false,
		Amount:    req.Amount,
		Balance:   account.Balance,
		CreatedDt: time.Now().In(location.AsiaJakarta),
	}
	err = s.transactionRepository.Topup(ctx, tx, transaction)
	if err != nil {
		log.Error(ctx, err, "error on s.transactionRepository.Topup, accountId=%s", account.ID.String())
		res.Message = response.GeneralError.Description()
		return
	}

	res.Status = "SUCCESS"
	res.Result = &restModel.TopUpResponse{
		TopupID:       transaction.ID.String(),
		AmountTopup:   req.Amount,
		BalanceBefore: balanceBefore,
		BalanceAfter:  transaction.Balance,
		CreatedDate:   transaction.CreatedDt.Format(constants.LayoutDt),
	}

	tx.Commit()
	return
}
