package response

import (
	"github.com/rosaekapratama/go-starter/response"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

const (
	PaymentAmountMustBeGreaterThanZero response.Response = iota + 100
)

var (
	descriptions = map[response.Response]string{
		PaymentAmountMustBeGreaterThanZero: "Payment amount must be greater than zero",
	}

	httpStatusCodes = map[response.Response]int{
		PaymentAmountMustBeGreaterThanZero: http.StatusBadRequest,
	}

	otelCodes = map[response.Response]codes.Code{
		PaymentAmountMustBeGreaterThanZero: codes.Ok,
	}
)

func init() {
	response.AppendDescription(descriptions)
	response.AppendHttpStatusCode(httpStatusCodes)
	response.AppendOtelCode(otelCodes)
}
