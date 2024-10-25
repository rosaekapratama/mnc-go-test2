package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rosaekapratama/mnc-go-test2/services"
)

type UserRestController interface {
	Register(c *gin.Context)
}

type userRestControllerImpl struct {
	userService services.UserService
}
