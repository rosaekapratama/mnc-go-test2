package response

import (
	"github.com/rosaekapratama/go-starter/response"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

const (
	AccountNotFound response.Response = iota + 100
)

var (
	descriptions = map[response.Response]string{
		AccountNotFound: "Account not found",
	}

	httpStatusCodes = map[response.Response]int{
		AccountNotFound: http.StatusOK,
	}

	otelCodes = map[response.Response]codes.Code{
		AccountNotFound: codes.Ok,
	}
)

func init() {
	response.AppendDescription(descriptions)
	response.AppendHttpStatusCode(httpStatusCodes)
	response.AppendOtelCode(otelCodes)
}
