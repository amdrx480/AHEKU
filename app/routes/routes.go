package routes

import (
	// "backend-golang/app/middlewares"
	// "backend-golang/businesses/purchases"
	"backend-golang/app/middlewares"
	"backend-golang/controllers/admin"
	history "backend-golang/controllers/history"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// ControllerList holds all controllers and their dependencies
type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    echojwt.Config
	AdminController  admin.AuthController

	HistoryController history.HistoryController
}

// RegisterRoutes registers all routes with the provided Echo instance
func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	// Rute untuk file statis
	// ketika menggunakan local
	e.Static("/images", "D:/Skripsi/AHEKU/AHEKU/images")

	// Public routes
	auth := e.Group("auth")
	auth.POST("/register", cl.AdminController.AdminRegister)
	auth.POST("/login", cl.AdminController.AdminLogin)
	auth.POST("/login_voucher", cl.AdminController.AdminVoucher)

	// Admin routes
	admin := e.Group("/admin")
	admin.Use(echojwt.WithConfig(cl.JWTMiddleware))
	// admin.Use(middlewares.VerifyToken)
	admin.POST("/voucher", cl.AdminController.AdminVoucher)
	admin.PUT("/profile", cl.AdminController.AdminProfileUpdate)
	admin.GET("/profile", cl.AdminController.AdminGetProfile)
	admin.GET("/:id", cl.AdminController.AdminGetByID)

	// Role routes
	role := e.Group("/role")
	role.Use(echojwt.WithConfig(cl.JWTMiddleware))
	role.Use(middlewares.VerifyToken)
	// role.Use(middlewares.RBAC(1))
	// role.Use(middlewares.RBAC("superadmin"))

	role.POST("", cl.AdminController.RoleCreate)
	role.GET("/:id", cl.AdminController.RoleGetByID)
	role.GET("", cl.AdminController.RoleGetAll)

	// Customers routes
	customers := e.Group("/customers")
	// customers.Use(echojwt.WithConfig(cl.JWTMiddleware))
	// customers.Use(middlewares.VerifyToken)
	// customers.Use(middlewares.RBAC(2))
	// customers.Use(middlewares.RBAC("admin"))

	customers.POST("", cl.AdminController.CustomersCreate)
	customers.GET("/:id", cl.AdminController.CustomersGetByID)
	customers.GET("", cl.AdminController.CustomersGetAll)

	// Categories routes
	categories := e.Group("/categories")
	// categories.Use(echojwt.WithConfig(cl.JWTMiddleware))
	// categories.Use(middlewares.VerifyToken)
	// categories.Use(middlewares.RBAC(2))
	// categories.Use(middlewares.RBAC("admin"))

	categories.POST("", cl.AdminController.CategoryCreate)
	categories.GET("/:id", cl.AdminController.CategoryGetByID)
	categories.GET("/category_name/:category_name", cl.AdminController.CategoryGetByName)
	categories.GET("", cl.AdminController.CategoryGetAll)

	// Vendors routes
	vendors := e.Group("/vendors")
	// vendors.Use(echojwt.WithConfig(cl.JWTMiddleware))
	// vendors.Use(middlewares.VerifyToken)
	// vendors.Use(middlewares.RBAC(2))
	// vendors.Use(middlewares.RBAC("admin"))

	vendors.POST("", cl.AdminController.VendorsCreate)
	vendors.GET("/:id", cl.AdminController.VendorsGetByID)
	vendors.GET("", cl.AdminController.VendorsGetAll)

	// Units routes
	units := e.Group("/units")
	// units.Use(echojwt.WithConfig(cl.JWTMiddleware))
	// units.Use(middlewares.VerifyToken)
	// units.Use(middlewares.RBAC(2))
	// units.Use(middlewares.RBAC("admin"))

	units.GET("/:id", cl.AdminController.UnitsGetByID)
	units.POST("", cl.AdminController.UnitsCreate)
	units.GET("", cl.AdminController.UnitsGetAll)

	// Stocks routes
	stocks := e.Group("/stocks")
	// stocks.Use(echojwt.WithConfig(cl.JWTMiddleware))
	// stocks.Use(middlewares.VerifyToken)
	// stocks.Use(middlewares.RBAC(2))
	// stocks.Use(middlewares.RBAC("admin"))

	stocks.GET("/:id", cl.AdminController.StocksGetByID)
	stocks.POST("", cl.AdminController.StocksCreate)
	stocks.GET("", cl.AdminController.StocksGetAll)

	// Purchases routes
	purchases := e.Group("/purchases")
	// purchases.Use(echojwt.WithConfig(cl.JWTMiddleware))
	// purchases.Use(middlewares.VerifyToken)
	// purchases.Use(middlewares.RBAC(2))
	// purchases.Use(middlewares.RBAC("admin"))

	purchases.POST("", cl.AdminController.PurchasesCreate)
	purchases.GET("/:id", cl.AdminController.PurchasesGetByID)
	purchases.GET("", cl.AdminController.PurchasesGetAll)

	// Cart items routes
	items := e.Group("/cart_items")
	// items.Use(echojwt.WithConfig(cl.JWTMiddleware))
	// items.Use(middlewares.VerifyToken)
	// items.Use(middlewares.RBAC(2))
	// items.Use(middlewares.RBAC("admin"))

	items.POST("", cl.AdminController.CartItemsCreate)
	items.GET("/:id", cl.AdminController.CartItemsGetByID)
	items.GET("/customer/:customer_id", cl.AdminController.CartItemsGetAllByCustomerID)
	items.GET("", cl.AdminController.CartItemsGetAll)
	items.DELETE("/:id", cl.AdminController.CartItemsDelete)

	// Item transactions routes
	itemTransactions := e.Group("/item_transactions")
	// itemTransactions.Use(echojwt.WithConfig(cl.JWTMiddleware))
	// itemTransactions.Use(middlewares.VerifyToken)
	// itemTransactions.Use(middlewares.RBAC(2))
	// itemTransactions.Use(middlewares.RBAC("admin"))

	itemTransactions.POST("/:customer_id", cl.AdminController.ItemTransactionsCreate)
	itemTransactions.GET("", cl.AdminController.ItemTransactionsGetAll)
}

// RegisterRoutes registers all routes with the provided Echo instance
// func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
// 	e.Use(cl.LoggerMiddleware)

// 	// admin := e.Group("auth")
// 	admin := e.Group("auth")
// 	// e.POST("register", cl.AuthController.AdminRegister)
// 	// e.POST("login", cl.AuthController.AdminLogin)

// 	admin.POST("/register", cl.AdminController.AdminRegister)
// 	admin.POST("/login", cl.AdminController.AdminLogin)
// 	admin.POST("/voucher", cl.AdminController.AdminVoucher)
// 	admin.GET("/:id", cl.AdminController.AdminGetByID)

// role := e.Group("role", echojwt.WithConfig(cl.JWTMiddleware))
// 	role := e.Group("role")
// 	role.POST("", cl.AdminController.RoleCreate)
// 	role.GET("/:id", cl.AdminController.RoleGetByID)
// 	role.GET("", cl.AdminController.RoleGetAll)

// 	customers := e.Group("customers")
// 	customers.POST("", cl.AdminController.CustomersCreate)
// 	customers.GET("/:id", cl.AdminController.CustomersGetByID)
// 	customers.GET("", cl.AdminController.CustomersGetAll)

// 	// categories := e.Group("categories", echojwt.WithConfig(cl.JWTMiddleware))
// 	categories := e.Group("categories")
// 	categories.POST("", cl.AdminController.CategoryCreate)
// 	categories.GET("/:id", cl.AdminController.CategoryGetByID)
// 	categories.GET("/category_name/:category_name", cl.AdminController.CategoryGetByName)
// 	categories.GET("", cl.AdminController.CategoryGetAll)

// 	// vendors := e.Group("vendors", echojwt.WithConfig(cl.JWTMiddleware))
// 	vendors := e.Group("vendors")
// 	vendors.POST("", cl.AdminController.VendorsCreate)
// 	vendors.GET("/:id", cl.AdminController.VendorsGetByID)
// 	vendors.GET("", cl.AdminController.VendorsGetAll)

// 	// units := e.Group("units", echojwt.WithConfig(cl.JWTMiddleware))
// 	units := e.Group("units")
// 	units.GET("/:id", cl.AdminController.UnitsGetByID)
// 	units.POST("", cl.AdminController.UnitsCreate)
// 	units.GET("", cl.AdminController.UnitsGetAll)

// 	// stocks := e.Group("stocks", echojwt.WithConfig(cl.JWTMiddleware))
// 	stocks := e.Group("stocks")
// 	stocks.GET("/:id", cl.AdminController.StocksGetByID)
// 	stocks.POST("", cl.AdminController.StocksCreate)
// 	// admin.PUT("/profiles/picture/:id", cl.AdminController.UploadProfileImage)
// 	stocks.GET("", cl.AdminController.StocksGetAll)

// 	// purchases := e.Group("purchases", echojwt.WithConfig(cl.JWTMiddleware))
// 	purchases := e.Group("purchases")
// 	purchases.POST("", cl.AdminController.PurchasesCreate)
// 	purchases.GET("/:id", cl.AdminController.PurchasesGetByID)
// 	purchases.GET("", cl.AdminController.PurchasesGetAll)

// 	items := e.Group("cart_items")
// 	items.POST("", cl.AdminController.CartItemsCreate)
// 	items.GET("/:id", cl.AdminController.CartItemsGetByID)
// 	items.GET("/customer/:customer_id", cl.AdminController.CartItemsGetAllByCustomerID)
// 	items.GET("", cl.AdminController.CartItemsGetAll)
// 	items.DELETE("/:id", cl.AdminController.CartItemsDelete)

// 	itemTransactions := e.Group("item_transactions")
// 	itemTransactions.POST("/:customer_id", cl.AdminController.ItemTransactionsCreate)
// 	itemTransactions.GET("", cl.AdminController.ItemTransactionsGetAll)
// }

// admin.GET("/admin_profile/:id", cl.AdminController.AdminProfileGetByID)
// admin.PUT("/admin_profile/admin_update/:id", cl.AdminController.AdminProfileUpdate)
// admin.PUT("/admin_profile/admin_picture/:id", cl.AdminController.AdminProfileUploadImage)

// carts := e.Group("carts", echojwt.WithConfig(cl.JWTMiddleware))
// carts := e.Group("carts")
// carts.GET("/:id", cl.AdminController.CartsGetByID)
// carts.POST("", cl.AdminController.CartsCreate)
// carts.GET("", cl.AdminController.CartsGetAll)
// carts.DELETE("/:id", cl.AdminController.CartsDelete)

// history := e.Group("history")
// history.POST("", cl.HistoryController.Create)
// history.GET("", cl.HistoryController.GetAll)
