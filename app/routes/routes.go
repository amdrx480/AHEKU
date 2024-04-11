package routes

import (
	// "backend-golang/app/middlewares"
	stocks "backend-golang/controllers/stocks"
	"backend-golang/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    echojwt.Config
	AuthController   users.AuthController

	StocksController stocks.StockController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	// users := e.Group("auth")

	// users.POST("/register", cl.AuthController.Register)
	// users.POST("/login", cl.AuthController.Login)
	e.POST("register", cl.AuthController.Register)
	e.POST("login", cl.AuthController.Login)

	stocks := e.Group("stocks", echojwt.WithConfig(cl.JWTMiddleware))
	stocks.GET("/:id", cl.StocksController.GetByID)
	stocks.POST("", cl.StocksController.Create)
	stocks.GET("", cl.StocksController.GetAll)

}
