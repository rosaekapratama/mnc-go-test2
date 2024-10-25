package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/rosaekapratama/go-starter/constant/location"
	"github.com/rosaekapratama/go-starter/constant/str"
	"github.com/rosaekapratama/go-starter/log"
	"github.com/rosaekapratama/go-starter/response"
	"github.com/rosaekapratama/go-starter/utils"
	"github.com/rosaekapratama/mnc-go-test2/constants"
	"github.com/rosaekapratama/mnc-go-test2/crypto"
	"github.com/rosaekapratama/mnc-go-test2/models/repo"
	"github.com/rosaekapratama/mnc-go-test2/models/rest"
	"github.com/rosaekapratama/mnc-go-test2/repositories"
	"strings"
	"time"
)

func NewUserService(_ context.Context, secret string, userRepository repositories.UserRepository) UserService {
	return &userServiceImpl{
		secret:         secret,
		userRepository: userRepository,
	}
}

func (s *userServiceImpl) Register(ctx context.Context, req *rest.RegisterRequest) (res *rest.BaseResponse[*rest.RegisterResponse], err error) {
	res = &rest.BaseResponse[*rest.RegisterResponse]{}
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

	user = &repo.User{
		ID:          uuid.New(),
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: req.PhoneNumber,
		Address:     address,
		Pin:         crypto.Hash(req.Pin),
		CreatedDt:   time.Now().In(location.AsiaJakarta),
	}
	err = s.userRepository.Save(ctx, user)
	if err != nil {
		log.Error(ctx, err, "error on s.userRepository.Save")
		res.Message = response.GeneralError.Description()
		return
	}

	res = &rest.BaseResponse[*rest.RegisterResponse]{
		Status: "SUCCESS",
		Result: &rest.RegisterResponse{
			UserID:      user.ID,
			FirstName:   utils.PString(user.FirstName),
			LastName:    utils.PString(user.LastName),
			PhoneNumber: user.PhoneNumber,
			Address:     utils.PString(user.Address),
			CreatedDate: user.CreatedDt.Format(constants.LayoutDt),
		},
	}

	return
}

func (s *userServiceImpl) Login(ctx context.Context, req *rest.LoginRequest) (res *rest.BaseResponse[*rest.LoginResponse], err error) {
	res = &rest.BaseResponse[*rest.LoginResponse]{}
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
		user.AccessToken = accessToken
		user.RefreshToken = refreshToken
		user.UpdatedDt = &updatedDt
		err = s.userRepository.Save(ctx, user)
		if err != nil {
			log.Error(ctx, err)
			res.Message = response.GeneralError.Description()
			return
		}

		res.Status = "SUCCESS"
		res.Result = &rest.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		return
	}

	res.Message = "Phone Number and PIN doesn’t match."
	return
}
