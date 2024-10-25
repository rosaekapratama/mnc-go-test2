package services

import (
	"context"
	"github.com/rosaekapratama/mnc-go-test2/models/rest"
	"github.com/rosaekapratama/mnc-go-test2/repositories"
)

type UserService interface {
	Register(ctx context.Context, req *rest.RegisterRequest) (res *rest.BaseResponse[*rest.RegisterResponse], err error)
}

type userServiceImpl struct {
	userRepository repositories.UserRepository
}
