package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rosaekapratama/go-starter/log"
	"github.com/rosaekapratama/go-starter/response"
	"github.com/rosaekapratama/go-starter/transport/restserver"
	"github.com/rosaekapratama/mnc-go-test2/models/rest"
	"github.com/rosaekapratama/mnc-go-test2/services"
)

func NewUserRestController(_ context.Context, userService services.UserService) UserRestController {
	return &userRestControllerImpl{
		userService: userService,
	}
}

func (ctrl *userRestControllerImpl) Register(c *gin.Context) {
	ctx := c.Request.Context()
	w := c.Writer

	req := &rest.RegisterRequest{}
	err := c.BindJSON(req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetRawResponse(w, response.GeneralError)
		c.JSON(response.GeneralError.HttpStatusCode(), &rest.BaseResponse[*rest.RegisterResponse]{
			Message: response.GeneralError.Description(),
		})
		return
	}

	res, err := ctrl.userService.Register(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetRawResponse(w, response.GeneralError)
		c.JSON(response.GeneralError.HttpStatusCode(), &rest.BaseResponse[*rest.RegisterResponse]{
			Message: response.GeneralError.Description(),
		})
		return
	}

	restserver.SetRawResponse(w, response.Success)
	c.JSON(response.Success.HttpStatusCode(), res)
}

func (ctrl *userRestControllerImpl) Login(c *gin.Context) {
	ctx := c.Request.Context()
	w := c.Writer

	req := &rest.LoginRequest{}
	err := c.BindJSON(req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetRawResponse(w, response.GeneralError)
		c.JSON(response.GeneralError.HttpStatusCode(), &rest.BaseResponse[*rest.RegisterResponse]{
			Message: response.GeneralError.Description(),
		})
		return
	}

	res, err := ctrl.userService.Login(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetRawResponse(w, response.GeneralError)
		c.JSON(response.GeneralError.HttpStatusCode(), &rest.BaseResponse[*rest.RegisterResponse]{
			Message: response.GeneralError.Description(),
		})
		return
	}

	restserver.SetRawResponse(w, response.Success)
	c.JSON(response.Success.HttpStatusCode(), res)
}
