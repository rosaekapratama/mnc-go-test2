package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/rosaekapratama/go-starter/constant/location"
	"github.com/rosaekapratama/go-starter/constant/str"
	"github.com/rosaekapratama/go-starter/database"
	"github.com/rosaekapratama/go-starter/log"
	"github.com/rosaekapratama/go-starter/response"
	"github.com/rosaekapratama/go-starter/utils"
	"github.com/rosaekapratama/mnc-go-test2/constants"
	"github.com/rosaekapratama/mnc-go-test2/crypto"
	repoModel "github.com/rosaekapratama/mnc-go-test2/models/repo"
	restModel "github.com/rosaekapratama/mnc-go-test2/models/rest"
	"github.com/rosaekapratama/mnc-go-test2/repositories"
	"strings"
	"time"
)

func NewUserService(_ context.Context, secret string, userRepository repositories.UserRepository, accountRepository repositories.AccountRepository) UserService {
	return &userServiceImpl{
		secret:            secret,
		userRepository:    userRepository,
		accountRepository: accountRepository,
	}
}

func (s *userServiceImpl) Register(ctx context.Context, req *restModel.RegisterRequest) (res *restModel.BaseResponse[*restModel.RegisterResponse], err error) {
	res = &restModel.BaseResponse[*restModel.RegisterResponse]{}

	// Init tx
	tx, err := database.Manager.Begin(ctx, "playground")
	if err != nil {
		res.Message = response.GeneralError.Description()
		return
	}
	defer tx.Rollback()

	if req.PhoneNumber == str.Empty {
		res.Message = "Phone number is empty"
		return
	}

	if req.Pin == str.Empty {
		res.Message = "Pin is empty"
		return
	}

	req.PhoneNumber = strings.TrimSpace(req.PhoneNumber)
	req.Pin = strings.TrimSpace(req.Pin)

	user, err := s.userRepository.FindByPhoneNo(ctx, req.PhoneNumber)
	if err != nil {
		log.Error(ctx, err, "error on s.userRepository.FindByPhoneNo, phoneNo=%s", req.PhoneNumber)
		res.Message = response.GeneralError.Description()
		return
	}

	if user != nil {
		res.Message = "Phone number already exists"
		return
	}

	var firstName *string
	if req.FirstName != str.Empty {
		firstName = utils.StringP(strings.TrimSpace(req.FirstName))
	}

	var lastName *string
	if req.LastName != str.Empty {
		lastName = utils.StringP(strings.TrimSpace(req.LastName))
	}

	var address *string
	if req.Address != str.Empty {
		address = utils.StringP(strings.TrimSpace(req.Address))
	}

	user = &repoModel.User{
		ID:          uuid.New(),
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: req.PhoneNumber,
		Address:     address,
		Pin:         crypto.Hash(req.Pin),
		CreatedDt:   time.Now().In(location.AsiaJakarta),
	}

	// Create user
	err = s.userRepository.Save(ctx, tx, user)
	if err != nil {
		log.Error(ctx, err, "error on s.userRepository.Save")
		res.Message = response.GeneralError.Description()
		return
	}

	res = &restModel.BaseResponse[*restModel.RegisterResponse]{
		Status: "SUCCESS",
		Result: &restModel.RegisterResponse{
			UserID:      user.ID,
			FirstName:   utils.PString(user.FirstName),
			LastName:    utils.PString(user.LastName),
			PhoneNumber: user.PhoneNumber,
			Address:     utils.PString(user.Address),
			CreatedDate: user.CreatedDt.Format(constants.LayoutDt),
		},
	}

	// Create saving account
	err = s.accountRepository.CreateSaving(ctx, tx, &repoModel.Account{
		ID:        uuid.New(),
		UserID:    user.ID,
		Type:      "SAVING",
		Balance:   0,
		CreatedDt: time.Now().In(location.AsiaJakarta),
	})
	if err != nil {
		log.Error(ctx, err, "error on s.accountRepository.CreateSaving")
		res.Message = response.GeneralError.Description()
		return
	}

	tx.Commit()
	return
}

func (s *userServiceImpl) Login(ctx context.Context, req *restModel.LoginRequest) (res *restModel.BaseResponse[*restModel.LoginResponse], err error) {
	res = &restModel.BaseResponse[*restModel.LoginResponse]{}

	// Init tx
	tx, err := database.Manager.Begin(ctx, "playground")
	if err != nil {
		res.Message = response.GeneralError.Description()
		return
	}
	defer tx.Rollback()

	if req.PhoneNumber == str.Empty {
		res.Message = "Phone number is empty"
		return
	}

	if req.Pin == str.Empty {
		res.Message = "Pin is empty"
		return
	}

	req.PhoneNumber = strings.TrimSpace(req.PhoneNumber)
	req.Pin = strings.TrimSpace(req.Pin)

	user, err := s.userRepository.FindByPhoneNo(ctx, req.PhoneNumber)
	if err != nil {
		log.Error(ctx, err, "error on s.userRepository.FindByPhoneNo, phoneNo=%s", req.PhoneNumber)
		res.Message = response.GeneralError.Description()
		return
	}

	if user == nil {
		res.Message = "User not found"
		return
	}

	req.Pin = crypto.Hash(req.Pin)
	if req.PhoneNumber == user.PhoneNumber && req.Pin == user.Pin {
		accessToken, refreshToken, errIf := crypto.GenerateTokens(ctx, user.ID.String(), req.PhoneNumber, s.secret)
		if errIf != nil {
			log.Error(ctx, err)
			err = errIf
			res.Message = response.GeneralError.Description()
			return
		}

		updatedDt := time.Now().In(location.AsiaJakarta)
		user.UpdatedDt = &updatedDt
		err = s.userRepository.Save(ctx, tx, user)
		if err != nil {
			log.Error(ctx, err)
			res.Message = response.GeneralError.Description()
			return
		}

		res.Status = "SUCCESS"
		res.Result = &restModel.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		tx.Commit()
		return
	}

	res.Message = "Phone Number and PIN doesn’t match."
	return
}
