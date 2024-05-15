package routes

import (
	// "backend-golang/app/middlewares"
	// "backend-golang/businesses/purchases"
	purchases "backend-golang/controllers/purchases"
	sales "backend-golang/controllers/sales"

	category "backend-golang/controllers/category"
	units "backend-golang/controllers/units"

	stocks "backend-golang/controllers/stocks"
	vendors "backend-golang/controllers/vendors"

	history "backend-golang/controllers/history"

	"backend-golang/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    echojwt.Config
	AuthController   users.AuthController

	StocksController    stocks.StockController
	PurchasesController purchases.PurchasesController
	SalesController     sales.SalesController

	VendorsController  vendors.VendorsController
	CategoryController category.CategoryController
	UnitsController    units.UnitsController
	HistoryController  history.HistoryController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	// users := e.Group("auth")

	// users.POST("/register", cl.AuthController.Register)
	// users.POST("/login", cl.AuthController.Login)
	e.POST("register", cl.AuthController.Register)
	e.POST("login", cl.AuthController.Login)

	// stocks := e.Group("stocks", echojwt.WithConfig(cl.JWTMiddleware))
	stocks := e.Group("stocks")
	stocks.GET("/:id", cl.StocksController.GetByID)
	stocks.POST("", cl.StocksController.Create)
	// users.PUT("/profiles/picture/:id", cl.ProfilesController.UploadProfileImage)
	stocks.GET("", cl.StocksController.GetAll)

	// purchases := e.Group("purchases", echojwt.WithConfig(cl.JWTMiddleware))
	purchases := e.Group("purchases")
	purchases.GET("/:id", cl.PurchasesController.GetByID)
	purchases.POST("", cl.PurchasesController.Create)
	purchases.GET("", cl.PurchasesController.GetAll)

	sales := e.Group("sales")
	sales.GET("/:id", cl.SalesController.GetByID)
	sales.POST("", cl.SalesController.Create)
	sales.POST("/to_history", cl.SalesController.ToHistory)
	sales.GET("", cl.SalesController.GetAll)
	sales.DELETE("/:id", cl.SalesController.Delete)

	// vendors := e.Group("vendors", echojwt.WithConfig(cl.JWTMiddleware))
	vendors := e.Group("vendors")
	vendors.GET("/:id", cl.VendorsController.GetByID)
	vendors.POST("", cl.VendorsController.Create)
	vendors.GET("", cl.VendorsController.GetAll)

	// category := e.Group("category", echojwt.WithConfig(cl.JWTMiddleware))
	category := e.Group("category")
	category.GET("/:id", cl.CategoryController.GetByID)
	category.GET("/category_name/:category_name", cl.CategoryController.GetByName)
	category.POST("", cl.CategoryController.Create)
	category.GET("", cl.CategoryController.GetAll)

	// units := e.Group("units", echojwt.WithConfig(cl.JWTMiddleware))
	units := e.Group("units")
	units.GET("/:id", cl.UnitsController.GetByID)
	units.POST("", cl.UnitsController.Create)
	units.GET("", cl.UnitsController.GetAll)

}
