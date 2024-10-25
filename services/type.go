package services

import (
	"context"
	"github.com/rosaekapratama/mnc-go-test2/models/rest"
	"github.com/rosaekapratama/mnc-go-test2/repositories"
)

type UserService interface {
	Register(ctx context.Context, req *rest.RegisterRequest) (res *rest.BaseResponse[*rest.RegisterResponse], err error)
	Login(ctx context.Context, req *rest.LoginRequest) (res *rest.BaseResponse[*rest.LoginResponse], err error)
}

type AccountService interface {
	Topup(ctx context.Context, req *rest.TopupRequest) (res *rest.BaseResponse[*rest.TopUpResponse], err error)
}

type userServiceImpl struct {
	secret            string
	userRepository    repositories.UserRepository
	accountRepository repositories.AccountRepository
}

type accountServiceImpl struct {
	accountRepository     repositories.AccountRepository
	transactionRepository repositories.TransactionRepository
}
