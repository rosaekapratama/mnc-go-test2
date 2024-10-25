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
	"github.com/rosaekapratama/mnc-go-test2/models/repo"
	"github.com/rosaekapratama/mnc-go-test2/models/rest"
	"github.com/rosaekapratama/mnc-go-test2/repositories"
	utils2 "github.com/rosaekapratama/mnc-go-test2/utils"
	"time"
)

func NewUserService(_ context.Context, userRepository repositories.UserRepository) UserService {
	return &userServiceImpl{
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
		firstName = utils.StringP(req.FirstName)
	}

	var lastName *string
	if req.LastName != str.Empty {
		lastName = utils.StringP(req.LastName)
	}

	var address *string
	if req.Address != str.Empty {
		address = utils.StringP(req.Address)
	}

	user = &repo.User{
		ID:          uuid.New(),
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: req.PhoneNumber,
		Address:     address,
		Pin:         utils2.Hash(req.Pin),
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
