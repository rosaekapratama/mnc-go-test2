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

type FinanceService interface {
	Topup(ctx context.Context, req *rest.TopupRequest) (res *rest.BaseResponse[*rest.TopUpResponse], err error)
	Payment(ctx context.Context, req *rest.PaymentRequest) (res *rest.BaseResponse[*rest.PaymentResponse], err error)
	Transfer(ctx context.Context, req *rest.TransferRequest) (res *rest.BaseResponse[*rest.TransferResponse], err error)
	FindAllTransaction(ctx context.Context) (res *rest.BaseResponse[*[]*rest.TransactionDetailResponse], err error)
}

type userServiceImpl struct {
	secret            string
	userRepository    repositories.UserRepository
	accountRepository repositories.AccountRepository
}

type financeServiceImpl struct {
	accountRepository     repositories.AccountRepository
	transactionRepository repositories.TransactionRepository
}
