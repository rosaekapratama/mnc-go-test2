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

func NewFinanceRestController(_ context.Context, financeService services.FinanceService) FinanceRestController {
	return &financeRestControllerImpl{
		financeService: financeService,
	}
}

func (ctrl *financeRestControllerImpl) Topup(c *gin.Context) {
	ctx := c.Request.Context()
	w := c.Writer

	req := &rest.TopupRequest{}
	err := c.BindJSON(req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetRawResponse(w, response.GeneralError)
		c.JSON(response.GeneralError.HttpStatusCode(), &rest.BaseResponse[*rest.TopUpResponse]{
			Message: response.GeneralError.Description(),
		})
		return
	}

	res, err := ctrl.financeService.Topup(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetRawResponse(w, response.GeneralError)
		c.JSON(response.GeneralError.HttpStatusCode(), &rest.BaseResponse[*rest.TopUpResponse]{
			Message: response.GeneralError.Description(),
		})
		return
	}

	restserver.SetRawResponse(w, response.Success)
	c.JSON(response.Success.HttpStatusCode(), res)
}

func (ctrl *financeRestControllerImpl) Payment(c *gin.Context) {
	ctx := c.Request.Context()
	w := c.Writer

	req := &rest.PaymentRequest{}
	err := c.BindJSON(req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetRawResponse(w, response.GeneralError)
		c.JSON(response.GeneralError.HttpStatusCode(), &rest.BaseResponse[*rest.PaymentResponse]{
			Message: response.GeneralError.Description(),
		})
		return
	}

	res, err := ctrl.financeService.Payment(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetRawResponse(w, response.GeneralError)
		c.JSON(response.GeneralError.HttpStatusCode(), &rest.BaseResponse[*rest.PaymentResponse]{
			Message: response.GeneralError.Description(),
		})
		return
	}

	restserver.SetRawResponse(w, response.Success)
	c.JSON(response.Success.HttpStatusCode(), res)
}

func (ctrl *financeRestControllerImpl) Transfer(c *gin.Context) {
	ctx := c.Request.Context()
	w := c.Writer

	req := &rest.TransferRequest{}
	err := c.BindJSON(req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetRawResponse(w, response.GeneralError)
		c.JSON(response.GeneralError.HttpStatusCode(), &rest.BaseResponse[*rest.TransferResponse]{
			Message: response.GeneralError.Description(),
		})
		return
	}

	res, err := ctrl.financeService.Transfer(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetRawResponse(w, response.GeneralError)
		c.JSON(response.GeneralError.HttpStatusCode(), &rest.BaseResponse[*rest.TransferResponse]{
			Message: response.GeneralError.Description(),
		})
		return
	}

	restserver.SetRawResponse(w, response.Success)
	c.JSON(response.Success.HttpStatusCode(), res)
}

func (ctrl *financeRestControllerImpl) Transactions(c *gin.Context) {
	ctx := c.Request.Context()
	w := c.Writer

	res, err := ctrl.financeService.FindAllTransaction(ctx)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetRawResponse(w, response.GeneralError)
		c.JSON(response.GeneralError.HttpStatusCode(), &rest.BaseResponse[[]*rest.TransactionDetailResponse]{
			Message: response.GeneralError.Description(),
		})
		return
	}

	restserver.SetRawResponse(w, response.Success)
	c.JSON(response.Success.HttpStatusCode(), res)
}
