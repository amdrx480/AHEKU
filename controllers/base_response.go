package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	// "backend-golang/controllers/users/response"
	// "backend-golang/controllers/admin/response"
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

// ////////////////////
// PaginatedResponse is a generic structure for paginated responses
type Pagination struct {
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
	TotalPages  int `json:"total_pages"`
	TotalItems  int `json:"total_items"`
}

type PaginationResponse[T any] struct {
	Error      bool       `json:"error"`
	Message    string     `json:"message"`
	Data       T          `json:"data"`
	Pagination Pagination `json:"pagination,omitempty"`
}

////////////////////////

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

// ////////////////////////
// NewPaginatedResponse creates a new paginated response
func NewPaginatedResponse[T any](c echo.Context, statusCode int, message string, data T, currentPage int, pageSize int, totalItems int) error {
	var pagination Pagination

	if currentPage < 1 {
		currentPage = 1
	}

	if pageSize < 1 {
		pageSize = 5 // Default page size
	}

	totalPages := (totalItems + pageSize - 1) / pageSize

	if currentPage > totalPages {
		errorMessage := fmt.Sprintf("Invalid page number. Current page: %d, Total pages: %d", currentPage, totalPages)
		return c.JSON(http.StatusBadRequest, PaginationResponse[T]{
			Error:   true,
			Message: errorMessage,
			Data:    data,
			Pagination: Pagination{
				CurrentPage: currentPage,
				PageSize:    pageSize,
				TotalPages:  totalPages,
				TotalItems:  totalItems,
			},
		})
	}

	pagination = Pagination{
		CurrentPage: currentPage,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}

	return c.JSON(statusCode, PaginationResponse[T]{
		Error:      false,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}

// func NewPaginatedResponse[T any](c echo.Context, statusCode int, message string, data T, currentPage int, pageSize int, totalItems int) error {
// 	totalPages := (totalItems + pageSize - 1) / pageSize // Calculate total pages

// 	// Check if currentPage is valid
// 	if currentPage < 1 || currentPage > totalPages {
// 		errorMessage := fmt.Sprintf("Invalid page number. Current page: %d, Total pages: %d", currentPage, totalPages)
// 		return c.JSON(http.StatusBadRequest, PaginationResponse[T]{
// 			Error:      true,
// 			Message:    errorMessage,
// 			Data:       data,
// 			Pagination: Pagination{},
// 		})
// 	}

// 	pagination := Pagination{
// 		CurrentPage: currentPage,
// 		PageSize:    pageSize,
// 		TotalPages:  totalPages,
// 		TotalItems:  totalItems,
// 	}

// 	return c.JSON(statusCode, PaginationResponse[T]{
// 		Error:      false,
// 		Message:    message,
// 		Data:       data,
// 		Pagination: pagination,
// 	})
// }

// func NewPaginatedResponse[T any](c echo.Context, statusCode int, statusError bool, message string, data T, currentPage int, pageSize int, totalItems int) error {
// 	totalPages := (totalItems + pageSize - 1) / pageSize // Calculate total pages
// 	pagination := Pagination{
// 		CurrentPage: currentPage,
// 		PageSize:    pageSize,
// 		TotalPages:  totalPages,
// 		TotalItems:  totalItems,
// 	}
// 	return c.JSON(statusCode, PaginationResponse[T]{
// 		Error:      statusError,
// 		Message:    message,
// 		Data:       data,
// 		Pagination: pagination,
// 	})
// }
