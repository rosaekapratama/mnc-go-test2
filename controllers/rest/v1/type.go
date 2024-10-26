package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rosaekapratama/mnc-go-test2/services"
)

type UserRestController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type FinanceRestController interface {
	Topup(c *gin.Context)
	Payment(c *gin.Context)
	Transfer(c *gin.Context)
	Transactions(c *gin.Context)
}

type userRestControllerImpl struct {
	userService services.UserService
}

type financeRestControllerImpl struct {
	financeService services.FinanceService
}
