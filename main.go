package main

import (
	"context"
	"github.com/rosaekapratama/go-starter/app"
	"github.com/rosaekapratama/go-starter/config"
	v1 "github.com/rosaekapratama/mnc-go-test2/controllers/rest/v1"
	"github.com/rosaekapratama/mnc-go-test2/repositories"
	"github.com/rosaekapratama/mnc-go-test2/routers"
	"github.com/rosaekapratama/mnc-go-test2/services"
)

func main() {
	ctx := context.Background()

	// Init repos
	userRepository := repositories.NewUserRepository(ctx)
	accountRepository := repositories.NewAccountRepository(ctx)
	transactionRepository := repositories.NewTransactionRepository(ctx)

	// Init services
	secret := config.Instance.GetStringAndThrowFatalIfEmpty("app.secret")
	userService := services.NewUserService(ctx, secret, userRepository, accountRepository)
	financeService := services.NewFinanceService(ctx, accountRepository, transactionRepository)

	// Init rest controller
	userRestController := v1.NewUserRestController(ctx, userService)
	financeRestController := v1.NewFinanceRestController(ctx, financeService)

	// Init rest router
	routers.Init(userRestController, financeRestController)

	app.Run()
}
