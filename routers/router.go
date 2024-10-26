package routers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rosaekapratama/go-starter/config"
	"github.com/rosaekapratama/go-starter/constant/headers"
	"github.com/rosaekapratama/go-starter/constant/str"
	"github.com/rosaekapratama/go-starter/log"
	"github.com/rosaekapratama/go-starter/response"
	"github.com/rosaekapratama/go-starter/slices"
	"github.com/rosaekapratama/go-starter/transport/restserver"
	v1 "github.com/rosaekapratama/mnc-go-test2/controllers/rest/v1"
	"github.com/rosaekapratama/mnc-go-test2/crypto"
	"github.com/rosaekapratama/mnc-go-test2/models/rest"
	"strings"
)

var (
	protectedPath = []string{
		"/topup",
		"/pay",
		"/transfer",
		"/transactions",
	}
)

func Init(userRestController v1.UserRestController, financeRestController v1.FinanceRestController) {
	restserver.Router.Use(validateToken)
	restserver.Router.POST("/register", userRestController.Register)
	restserver.Router.POST("/login", userRestController.Login)
	restserver.Router.POST("/topup", financeRestController.Topup)
	restserver.Router.POST("/pay", financeRestController.Payment)
	restserver.Router.POST("/transfer", financeRestController.Transfer)
	restserver.Router.GET("/transactions", financeRestController.Transactions)
}

func validateToken(c *gin.Context) {
	ctx := c.Request.Context()
	w := c.Writer

	res := &rest.BaseResponse[interface{}]{}
	path := c.Request.URL.Path
	if slices.ContainStringCaseSensitive(protectedPath, path) {
		auth := c.Request.Header.Get(headers.Authorization)

		// Reject if auth header doesn't have Bearer prefix
		if !strings.HasPrefix(auth, headers.BearerTokenPrefix) {
			res.Message = "Invalid token"
			restserver.SetRawResponse(w, response.InvalidArgument)
			c.JSON(response.InvalidArgument.HttpStatusCode(), res)
			c.Abort()
			return
		}

		secret, err := config.Instance.GetString("app.secret")
		if err != nil {
			log.Error(ctx, err)
			res.Message = response.GeneralError.Description()
			restserver.SetRawResponse(w, response.GeneralError)
			c.JSON(response.GeneralError.HttpStatusCode(), res)
			c.Abort()
			return
		}

		tokenStr := auth[len(headers.BearerTokenPrefix):]
		claim, err := crypto.ExtractClaim(ctx, secret, tokenStr)
		if err != nil {
			log.Error(ctx, err)
			res.Message = response.UnauthorizedAccess.Description()
			restserver.SetRawResponse(w, response.UnauthorizedAccess)
			c.JSON(response.UnauthorizedAccess.HttpStatusCode(), res)
			c.Abort()
			return
		}

		if claim.Sub == str.Empty || claim.PhoneNumber == str.Empty {
			log.Error(ctx, response.UnauthorizedAccess, "sub and phoneNo claim must not be empty")
			res.Message = response.UnauthorizedAccess.Description()
			restserver.SetRawResponse(w, response.UnauthorizedAccess)
			c.JSON(response.UnauthorizedAccess.HttpStatusCode(), res)
			c.Abort()
			return
		}

		ctx = context.WithValue(ctx, "userId", claim.Sub)
		ctx = context.WithValue(ctx, "phoneNo", claim.PhoneNumber)
		c.Request = c.Request.WithContext(ctx)
	}
}
