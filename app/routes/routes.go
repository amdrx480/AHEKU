package routes

import (
	// "backend-golang/app/middlewares"
	// "backend-golang/businesses/purchases"
	"backend-golang/controllers/admin"
	history "backend-golang/controllers/history"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    echojwt.Config
	AdminController  admin.AuthController

	HistoryController history.HistoryController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	// admin := e.Group("auth")
	admin := e.Group("auth")
	// e.POST("register", cl.AuthController.AdminRegister)
	// e.POST("login", cl.AuthController.AdminLogin)

	admin.POST("/register", cl.AdminController.AdminRegister)
	admin.POST("/login", cl.AdminController.AdminLogin)
	admin.POST("/voucher", cl.AdminController.AdminVoucher)

	customers := e.Group("customers")
	customers.POST("", cl.AdminController.CustomersCreate)
	customers.GET("/:id", cl.AdminController.CustomersGetByID)
	customers.GET("", cl.AdminController.CustomersGetAll)

	// categories := e.Group("categories", echojwt.WithConfig(cl.JWTMiddleware))
	categories := e.Group("categories")
	categories.POST("", cl.AdminController.CategoryCreate)
	categories.GET("/:id", cl.AdminController.CategoryGetByID)
	categories.GET("/category_name/:category_name", cl.AdminController.CategoryGetByName)
	categories.GET("", cl.AdminController.CategoryGetAll)

	// vendors := e.Group("vendors", echojwt.WithConfig(cl.JWTMiddleware))
	vendors := e.Group("vendors")
	vendors.POST("", cl.AdminController.VendorsCreate)
	vendors.GET("/:id", cl.AdminController.VendorsGetByID)
	vendors.GET("", cl.AdminController.VendorsGetAll)

	// units := e.Group("units", echojwt.WithConfig(cl.JWTMiddleware))
	units := e.Group("units")
	units.GET("/:id", cl.AdminController.UnitsGetByID)
	units.POST("", cl.AdminController.UnitsCreate)
	units.GET("", cl.AdminController.UnitsGetAll)

	// stocks := e.Group("stocks", echojwt.WithConfig(cl.JWTMiddleware))
	stocks := e.Group("stocks")
	stocks.GET("/:id", cl.AdminController.StocksGetByID)
	stocks.POST("", cl.AdminController.StocksCreate)
	// admin.PUT("/profiles/picture/:id", cl.ProfilesController.UploadProfileImage)
	stocks.GET("", cl.AdminController.StocksGetAll)

	// purchases := e.Group("purchases", echojwt.WithConfig(cl.JWTMiddleware))
	purchases := e.Group("purchases")
	purchases.GET("/:id", cl.AdminController.PurchasesCreate)
	purchases.POST("", cl.AdminController.PurchasesCreate)
	purchases.GET("", cl.AdminController.PurchasesGetAll)

	items := e.Group("cart_items")
	items.POST("", cl.AdminController.CartItemsCreate)
	items.GET("/:id", cl.AdminController.CartItemsGetByID)
	items.GET("/customer/:customer_id", cl.AdminController.CartItemsGetByCustomerID)
	items.GET("", cl.AdminController.CartItemsGetAll)
	items.DELETE("/:id", cl.AdminController.CartItemsDelete)
	// items.POST("/to_history", cl.ItemsController.ToHistory)

	// carts := e.Group("carts", echojwt.WithConfig(cl.JWTMiddleware))
	// carts := e.Group("carts")
	// carts.GET("/:id", cl.AdminController.CartsGetByID)
	// carts.POST("", cl.AdminController.CartsCreate)
	// carts.GET("", cl.AdminController.CartsGetAll)
	// carts.DELETE("/:id", cl.AdminController.CartsDelete)

	// units := e.Group("units", echojwt.WithConfig(cl.JWTMiddleware))
	history := e.Group("history")
	history.POST("", cl.HistoryController.Create)
	history.GET("", cl.HistoryController.GetAll)

	// vendors := e.Group("vendors", echojwt.WithConfig(cl.JWTMiddleware))

}
