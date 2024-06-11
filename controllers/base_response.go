package controllers

import (
	"github.com/labstack/echo/v4"
	// "backend-golang/controllers/users/response"
)

type LoginVoucherResult struct {
	Token string `json:"token"`
}

type LoginVoucherResponse struct {
	Error              bool               `json:"error"`
	Message            string             `json:"message"`
	LoginVoucherResult LoginVoucherResult `json:"login_result"`
}

type Response[T any] struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ResponseWithoutData struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func NewResponseLoginVoucher(c echo.Context, statusCode int, statusMessage bool, message string, loginVoucherResult string) error {
	return c.JSON(statusCode, LoginVoucherResponse{
		Error:   statusMessage,
		Message: message,
		LoginVoucherResult: LoginVoucherResult{
			Token: loginVoucherResult,
		},
	})
}

func NewResponse[T any](c echo.Context, statusCode int, statusError bool, message string, data T) error {
	return c.JSON(statusCode, Response[T]{
		Error:   statusError,
		Message: message,
		Data:    data,
	})
}

func NewResponseWithoutData(c echo.Context, statusCode int, statusError bool, message string) error {
	return c.JSON(statusCode, ResponseWithoutData{
		Error:   statusError,
		Message: message,
	})
}
