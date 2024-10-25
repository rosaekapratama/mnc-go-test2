package routers

import (
	"github.com/rosaekapratama/go-starter/transport/restserver"
	v1 "github.com/rosaekapratama/mnc-go-test2/controllers/rest/v1"
)

func Init(userRestController v1.UserRestController) {
	restserver.Router.POST("/register", userRestController.Register)
	restserver.Router.POST("/login", userRestController.Login)
}
