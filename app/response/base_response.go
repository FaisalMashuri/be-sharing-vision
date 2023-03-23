package response

import (
	"be-sharing-vision/app/domain"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, status int, data domain.PostDomain) error {
	return c.JSON(status, PostFromDomain(data))
}

func ErrorResponse(c echo.Context, status int, err error) error {
	res := BaseResponse{
		Message: err.Error(),
	}
	return c.JSON(status, res)
}
