package controllers

import (
	"github.com/labstack/echo/v4"
	// "backend-golang/controllers/users/response"
)

// type LoginResult struct {
// 	// Name  string `json:"name"`
// 	Token string `json:"token"`
// }

type Response[T any] struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    T      `json:"token"`
}

// type ResponseLogin struct {
// 	Error       bool        `json:"error"`
// 	Message     string      `json:"message"`
// 	LoginResult LoginResult `json:"loginResult"`
// }

func NewResponse[T any](c echo.Context, statusCode int, statusError bool, message string, data T) error {
	return c.JSON(statusCode, Response[T]{
		Error:   statusError,
		Message: message,
		Data:    data,
	})
}

// func NewResponseLogin(c echo.Context, statusCode int, statusMessage bool, message string, loginResult string) error {
// 	return c.JSON(statusCode, ResponseLogin{
// 		Error:   statusMessage,
// 		Message: message,
// 		LoginResult: LoginResult{
// 			Token: loginResult,
// 		},
// 	})
// }
