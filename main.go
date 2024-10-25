package main

import (
	"context"
	"github.com/rosaekapratama/go-starter/app"
	v1 "github.com/rosaekapratama/mnc-go-test2/controllers/rest/v1"
	"github.com/rosaekapratama/mnc-go-test2/repositories"
	"github.com/rosaekapratama/mnc-go-test2/routers"
	"github.com/rosaekapratama/mnc-go-test2/services"
)

func main() {
	ctx := context.Background()

	// Init repos
	userRepository := repositories.NewUserRepository(ctx)

	// Init services
	userService := services.NewUserService(ctx, userRepository)

	// Init rest controller
	userRestController := v1.NewUserRestController(ctx, userService)

	// Init rest router
	routers.Init(userRestController)

	app.Run()
}
