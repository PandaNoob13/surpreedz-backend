package response

import (
	"errors"
	"net/http"
	"surpreedz-backend/utils"
)

type Status struct {
	ResponsCode    string `json:"responsCode"`
	ResponsMessage string `json:"responsMessage"`
}

type Response struct {
	Status
	Data interface{} `json:"data,omitempty"`
}

func NewSuccessMessage(data interface{}) (httpStatusCode int, apiRespons Response) {
	status := Status{
		ResponsCode:    SuccessCode,
		ResponsMessage: SuccessMessage,
	}
	httpStatusCode = http.StatusOK
	apiRespons = Response{
		Status: status,
		Data:   data,
	}
	return
}

func NewErrorMessage(err error) (httpStatusCode int, apiResponse Response) {
	var userError utils.AppError
	var status Status
	if errors.As(err, &userError) {
		status = Status{
			ResponsCode:    userError.ErrorCode,
			ResponsMessage: userError.ErrorMessage,
		}
		httpStatusCode = userError.ErrorType
	} else {
		status = Status{
			ResponsCode:    DefaultErrorCode,
			ResponsMessage: DefaultErrorMessage,
		}
	}
	apiResponse = Response{
		Status: status,
		Data:   nil,
	}
	return
}
