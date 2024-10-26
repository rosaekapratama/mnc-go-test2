package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/rosaekapratama/go-starter/constant/location"
	"github.com/rosaekapratama/go-starter/constant/str"
	"github.com/rosaekapratama/go-starter/database"
	"github.com/rosaekapratama/go-starter/log"
	"github.com/rosaekapratama/go-starter/response"
	"github.com/rosaekapratama/mnc-go-test2/constants"
	repoModel "github.com/rosaekapratama/mnc-go-test2/models/repo"
	restModel "github.com/rosaekapratama/mnc-go-test2/models/rest"
	"github.com/rosaekapratama/mnc-go-test2/repositories"
	"time"
)

func NewFinanceService(_ context.Context, accountRepository repositories.AccountRepository, transactionRepository repositories.TransactionRepository) FinanceService {
	return &financeServiceImpl{
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
	}
}

func (s *financeServiceImpl) Topup(ctx context.Context, req *restModel.TopupRequest) (res *restModel.BaseResponse[*restModel.TopUpResponse], err error) {
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

	// Get user ID from context
	userIdStr := ctx.Value("userId").(string)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		log.Errorf(ctx, err, "Invalid user ID, userId=%s", userIdStr)
		res.Message = response.GeneralError.Description()
		return
	}

	// Get saving account by user ID from token
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

	// Credit account
	balanceBefore := account.Balance
	account, err = s.accountRepository.Credit(ctx, tx, account.ID, req.Amount)
	if err != nil {
		log.Error(ctx, err, "error on s.accountRepository.Credit, accountId=%s", account.ID.String())
		res.Message = response.GeneralError.Description()
		return
	}

	// Add transaction history
	transaction := &repoModel.Transaction{
		ID:            uuid.New(),
		AccountID:     account.ID,
		Type:          "TOPUP",
		Status:        "SUCCESS",
		CR:            true,
		Amount:        req.Amount,
		BalanceBefore: balanceBefore,
		BalanceAfter:  account.Balance,
		CreatedDt:     time.Now().In(location.AsiaJakarta),
	}
	err = s.transactionRepository.Save(ctx, tx, transaction)
	if err != nil {
		log.Error(ctx, err, "error on s.transactionRepository.Credit, accountId=%s", account.ID.String())
		res.Message = response.GeneralError.Description()
		return
	}

	res.Status = "SUCCESS"
	res.Result = &restModel.TopUpResponse{
		TopupID:       transaction.ID.String(),
		AmountTopup:   req.Amount,
		BalanceBefore: transaction.BalanceBefore,
		BalanceAfter:  transaction.BalanceAfter,
		CreatedDate:   transaction.CreatedDt.Format(constants.LayoutDt),
	}

	tx.Commit()
	return
}

func (s *financeServiceImpl) Payment(ctx context.Context, req *restModel.PaymentRequest) (res *restModel.BaseResponse[*restModel.PaymentResponse], err error) {
	res = &restModel.BaseResponse[*restModel.PaymentResponse]{}

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

	// Get user ID from context
	userIdStr := ctx.Value("userId").(string)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		log.Errorf(ctx, err, "Invalid user ID, userId=%s", userIdStr)
		res.Message = response.GeneralError.Description()
		return
	}

	// Get saving account by user ID from token
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

	// Debit account
	balanceBefore := account.Balance
	account, err = s.accountRepository.Debit(ctx, tx, account.ID, req.Amount)
	if err != nil && err.Error() == "not enough balance" {
		err = nil
		res.Message = "not enough balance"
		return
	} else if err != nil {
		log.Error(ctx, err, "error on s.accountRepository.Debit, accountId=%s", account.ID.String())
		res.Message = response.GeneralError.Description()
		return
	}

	// Add transaction history
	var remark *string
	if req.Remarks != str.Empty {
		remark = &req.Remarks
	}
	transaction := &repoModel.Transaction{
		ID:            uuid.New(),
		AccountID:     account.ID,
		Type:          "PAYMENT",
		Status:        "SUCCESS",
		CR:            false,
		Amount:        req.Amount,
		BalanceBefore: balanceBefore,
		BalanceAfter:  account.Balance,
		Remark:        remark,
		CreatedDt:     time.Now().In(location.AsiaJakarta),
	}
	err = s.transactionRepository.Save(ctx, tx, transaction)
	if err != nil {
		log.Error(ctx, err, "error on s.transactionRepository.Save, accountId=%s", account.ID.String())
		res.Message = response.GeneralError.Description()
		return
	}

	res.Status = "SUCCESS"
	res.Result = &restModel.PaymentResponse{
		PaymentID:     transaction.ID.String(),
		Amount:        req.Amount,
		Remarks:       req.Remarks,
		BalanceBefore: transaction.BalanceBefore,
		BalanceAfter:  transaction.BalanceAfter,
		CreatedDate:   transaction.CreatedDt.Format(constants.LayoutDt),
	}

	tx.Commit()
	return
}

func (s *financeServiceImpl) Transfer(ctx context.Context, req *restModel.TransferRequest) (res *restModel.BaseResponse[*restModel.TransferResponse], err error) {
	res = &restModel.BaseResponse[*restModel.TransferResponse]{}

	// Init tx
	tx, err := database.Manager.Begin(ctx, "playground")
	if err != nil {
		res.Message = response.GeneralError.Description()
		return
	}
	defer tx.Rollback()

	// Validate transfer amount
	if req.Amount <= 0 {
		res.Message = "Amount must be greater than zero"
		return
	}

	// Validate destination user ID
	destUserId, err := uuid.Parse(req.TargetUser)
	if err != nil {
		log.Tracef(ctx, "Invalid destination user ID, destUserId=%s", req.TargetUser)
		res.Message = response.InvalidArgument.Description()
		err = nil
		return
	}

	// Check destination user has SAVING account
	destAcc, err := s.accountRepository.FindSavingByUserId(ctx, destUserId)
	if err != nil {
		log.Errorf(ctx, err, "error on FindSavingByUserId(), destUserId=%s", destUserId.String())
		res.Message = response.GeneralError.Description()
		return
	}

	if destAcc == nil {
		res.Message = "Unknown destination user"
		return
	}

	// Get user ID from context
	srcUserIdStr := ctx.Value("userId").(string)
	srcUserId, err := uuid.Parse(srcUserIdStr)
	if err != nil {
		log.Errorf(ctx, err, "Invalid user ID, userId=%s", srcUserIdStr)
		res.Message = response.GeneralError.Description()
		return
	}

	// Cant transfer to own self
	if srcUserId == destUserId {
		res.Message = "destination user cannot be the same with source user"
		return
	}

	// Get saving source account by user ID from token
	srcAcc, err := s.accountRepository.FindSavingByUserId(ctx, srcUserId)
	if err != nil {
		log.Errorf(ctx, err, "error on FindSavingByUserId(), userId=%s", srcUserId.String())
		res.Message = response.GeneralError.Description()
		return
	}

	if srcAcc == nil {
		res.Message = "User doesn't have a SAVING srcAcc"
		return
	}

	// Debit srcAcc
	srcBalanceBefore := srcAcc.Balance
	srcAcc, err = s.accountRepository.Debit(ctx, tx, srcAcc.ID, req.Amount)
	if err != nil && err.Error() == "not enough balance" {
		err = nil
		res.Message = "not enough balance"
		return
	} else if err != nil {
		log.Error(ctx, err, "error on s.accountRepository.Debit, srcAccId=%s", srcAcc.ID.String())
		res.Message = response.GeneralError.Description()
		return
	}

	// Credit destAcc
	destBalanceBefore := destAcc.Balance
	destAcc, err = s.accountRepository.Credit(ctx, tx, destAcc.ID, req.Amount)
	if err != nil {
		log.Error(ctx, err, "error on s.accountRepository.Credit, destAccId=%s", destAcc.ID.String())
		res.Message = response.GeneralError.Description()
		return
	}

	// Add transaction history for source account
	now := time.Now().In(location.AsiaJakarta)
	var remark *string
	if req.Remarks != str.Empty {
		remark = &req.Remarks
	}
	srcTransaction := &repoModel.Transaction{
		ID:            uuid.New(),
		AccountID:     srcAcc.ID,
		Type:          "TRANSFER",
		Status:        "SUCCESS",
		CR:            false,
		Amount:        req.Amount,
		BalanceBefore: srcBalanceBefore,
		BalanceAfter:  srcAcc.Balance,
		Remark:        remark,
		CreatedDt:     now,
	}
	err = s.transactionRepository.Save(ctx, tx, srcTransaction)
	if err != nil {
		log.Error(ctx, err, "error on s.transactionRepository.Save, srcAccId=%s", srcAcc.ID.String())
		res.Message = response.GeneralError.Description()
		return
	}

	// Add transaction history for destination account
	destTransaction := &repoModel.Transaction{
		ID:            uuid.New(),
		AccountID:     destAcc.ID,
		Type:          "TRANSFER",
		Status:        "SUCCESS",
		CR:            true,
		Amount:        req.Amount,
		BalanceBefore: destBalanceBefore,
		BalanceAfter:  destAcc.Balance,
		Remark:        remark,
		CreatedDt:     now,
	}
	err = s.transactionRepository.Save(ctx, tx, destTransaction)
	if err != nil {
		log.Errorf(ctx, err, "error on s.transactionRepository.Save, destAccId=%s", destAcc.ID.String())
		res.Message = response.GeneralError.Description()
		return
	}

	res.Status = "SUCCESS"
	res.Result = &restModel.TransferResponse{
		TransferID:    srcTransaction.ID.String(),
		Amount:        req.Amount,
		Remarks:       req.Remarks,
		BalanceBefore: srcTransaction.BalanceBefore,
		BalanceAfter:  srcTransaction.BalanceAfter,
		CreatedDate:   srcTransaction.CreatedDt.Format(constants.LayoutDt),
	}

	tx.Commit()
	return
}

func (s *financeServiceImpl) FindAllTransaction(ctx context.Context) (res *restModel.BaseResponse[*[]*restModel.TransactionDetailResponse], err error) {
	res = &restModel.BaseResponse[*[]*restModel.TransactionDetailResponse]{}

	// Get user ID from context
	userIdStr := ctx.Value("userId").(string)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		log.Errorf(ctx, err, "Invalid user ID, userId=%s", userIdStr)
		res.Message = response.GeneralError.Description()
		return
	}

	queryList, err := s.transactionRepository.FindAllByUserId(ctx, userId)
	if err != nil {
		log.Error(ctx, err, "error on s.transactionRepository.FindAll(ctx)")
		res.Message = response.GeneralError.Description()
		return
	}

	jsonList := make([]*restModel.TransactionDetailResponse, len(queryList))
	for i, l := range queryList {
		jsonList[i] = &restModel.TransactionDetailResponse{
			ID:              l.ID,
			UserID:          l.UserID,
			TransactionType: l.TransactionType,
			Status:          l.Status,
			Cr:              l.Cr,
			Amount:          l.Amount,
			Remarks:         l.Remarks,
			BalanceBefore:   l.BalanceBefore,
			BalanceAfter:    l.BalanceAfter,
			CreatedDate:     l.CreatedDt.Format(constants.LayoutDt),
		}
	}

	res.Status = "SUCCESS"
	res.Result = &jsonList
	return
}
