package admin

import (
	"backend-golang/businesses/admin"
	"backend-golang/controllers"
	"backend-golang/controllers/admin/request"
	"backend-golang/controllers/admin/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUseCase admin.Usecase
}

func NewAdminController(authUC admin.Usecase) *AuthController {
	return &AuthController{
		authUseCase: authUC,
	}
}

func (ctrl *AuthController) AdminRegister(c echo.Context) error {
	adminInput := request.AdminRegistration{}
	ctx := c.Request().Context()

	if err := c.Bind(&adminInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid register request", "")
	}

	err := adminInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "validation register failed", "")
	}

	admin, err := ctrl.authUseCase.AdminRegister(ctx, adminInput.ToAdminRegistrationDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "error when inserting register data", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "admin registered", response.FromAdminsDomain(admin))
}

func (ctrl *AuthController) AdminLogin(c echo.Context) error {
	adminInput := request.AdminLogin{}
	ctx := c.Request().Context()

	if err := c.Bind(&adminInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid login request", "")
	}

	err := adminInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "validation login failed", "")
	}
	loginResult, err := ctrl.authUseCase.AdminLogin(ctx, adminInput.ToAdminLoginDomain())

	var isFailed bool = err != nil || loginResult == ""

	if isFailed {
		return controllers.NewResponse(c, http.StatusUnauthorized, true, "invalid name or password", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "login success", loginResult)
}

func (ctrl *AuthController) AdminVoucher(c echo.Context) error {
	adminInput := request.AdminVoucher{}
	ctx := c.Request().Context()

	if err := c.Bind(&adminInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid voucher request", "")
	}

	err := adminInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "validation voucher failed", "")
	}
	voucherResult, err := ctrl.authUseCase.AdminVoucher(ctx, adminInput.ToAdminVoucherDomain())

	var isFailed bool = err != nil || voucherResult == ""

	if isFailed {
		return controllers.NewResponse(c, http.StatusUnauthorized, true, "invalid voucher or password", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "voucher success", voucherResult)
}

func (cc *AuthController) CustomersGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	customersID := c.Param("id")

	customers, err := cc.authUseCase.CustomersGetByID(ctx, customersID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "customer not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "customer found", response.FromCustomersDomain(customers))
}

func (cc *AuthController) CustomersCreate(c echo.Context) error {
	input := request.Customers{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	customers, err := cc.authUseCase.CustomersCreate(ctx, input.ToCustomersDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a customer", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "customer registered", response.FromCustomersDomain(customers))
}

func (cc *AuthController) CustomersGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	categoriesData, err := cc.authUseCase.CustomersGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	categories := []response.Customers{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromCustomersDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all customers", categories)
}

func (cc *AuthController) CategoryGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	categoryID := c.Param("id")

	category, err := cc.authUseCase.CategoryGetByID(ctx, categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "category found", response.FromCategoryDomain(category))
}

func (cc *AuthController) CategoryGetByName(c echo.Context) error {
	ctx := c.Request().Context()

	categoryName := c.Param("category_name")

	category, err := cc.authUseCase.CategoryGetByName(ctx, categoryName)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "categoryName not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "categoryName found", response.FromCategoryDomain(category))
}

func (pc *AuthController) CategoryCreate(c echo.Context) error {
	input := request.Category{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	category, err := pc.authUseCase.CategoryCreate(ctx, input.ToCategoriesDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a category", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "category registered", response.FromCategoryDomain(category))
}

func (pc *AuthController) CategoryGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	categoryData, err := pc.authUseCase.CategoryGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	categories := []response.Category{}

	for _, category := range categoryData {
		categories = append(categories, response.FromCategoryDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all category", categories)
}

func (pc *AuthController) VendorsCreate(c echo.Context) error {
	input := request.Vendors{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	vendors, err := pc.authUseCase.VendorsCreate(ctx, input.ToVendorsDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a vendor", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "vendor registered", response.FromVendorsDomain(vendors))
}

func (cc *AuthController) VendorsGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	vendorsID := c.Param("id")

	vendors, err := cc.authUseCase.VendorsGetByID(ctx, vendorsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "vendor not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "vendor found", response.FromVendorsDomain(vendors))
}

func (pc *AuthController) VendorsGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	categoriesData, err := pc.authUseCase.VendorsGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	categories := []response.Vendors{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromVendorsDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all vendors", categories)
}

func (pc *AuthController) UnitsCreate(c echo.Context) error {
	input := request.Units{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	unit, err := pc.authUseCase.UnitsCreate(ctx, input.ToUnitsDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a unit", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "unit registered", response.FromUnitsDomain(unit))
}

func (cc *AuthController) UnitsGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	unitsID := c.Param("id")

	unit, err := cc.authUseCase.UnitsGetByID(ctx, unitsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "unit not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "unit found", response.FromUnitsDomain(unit))
}

func (pc *AuthController) UnitsGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	unitsData, err := pc.authUseCase.UnitsGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	units := []response.Units{}

	for _, unit := range unitsData {
		units = append(units, response.FromUnitsDomain(unit))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all units", units)
}

func (pc *AuthController) PurchasesCreate(c echo.Context) error {
	input := request.Purchases{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	purchases, err := pc.authUseCase.PurchasesCreate(ctx, input.ToPurchasesDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a purchases", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "purchases registered", response.FromPurchasesDomain(purchases))
}

func (cc *AuthController) PurchasesGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	purchasesID := c.Param("id")

	purchases, err := cc.authUseCase.PurchasesGetByID(ctx, purchasesID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "category found", response.FromPurchasesDomain(purchases))
}

func (pc *AuthController) PurchasesGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	categoriesData, err := pc.authUseCase.PurchasesGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	categories := []response.Purchases{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromPurchasesDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all categories", categories)
}

func (sc *AuthController) StocksCreate(c echo.Context) error {
	input := request.Stocks{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	stock, err := sc.authUseCase.StocksCreate(ctx, input.ToStocksDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a stock", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "stock registered", response.FromStocksDomain(stock))
}

func (cc *AuthController) StocksGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	stockID := c.Param("id")

	stock, err := cc.authUseCase.StocksGetByID(ctx, stockID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "category found", response.FromStocksDomain(stock))
}

func (sc *AuthController) StocksGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	categoriesData, err := sc.authUseCase.StocksGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	categories := []response.Stocks{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromStocksDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all categories", categories)
}

func (sc *AuthController) CartItemsCreate(c echo.Context) error {
	input := request.CartItems{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	// err := input.Validate()

	// if err != nil {
	// 	return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	// }

	items, err := sc.authUseCase.CartItemsCreate(ctx, input.ToCartItemsDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a items", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "items registered", response.FromCartItemsDomain(items))
}

func (sc *AuthController) CartItemsGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	itemsID := c.Param("id")

	items, err := sc.authUseCase.CartItemsGetByID(ctx, itemsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "items not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "items found", response.FromCartItemsDomain(items))
}

func (sc *AuthController) CartItemsGetByCustomerID(c echo.Context) error {
	ctx := c.Request().Context()

	customerID := c.Param("customer_id")

	cartitems, err := sc.authUseCase.CartItemsGetByCustomerID(ctx, customerID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "items not found", "")
	}

	cartItemsResponse := []response.CartItems{}

	for _, items := range cartitems {
		cartItemsResponse = append(cartItemsResponse, response.FromCartItemsDomain(items))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all cartItems", cartItemsResponse)
}

func (sc *AuthController) CartItemsGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	itemsData, err := sc.authUseCase.CartItemsGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	itemsResponse := []response.CartItems{}

	for _, items := range itemsData {
		itemsResponse = append(itemsResponse, response.FromCartItemsDomain(items))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all items", itemsResponse)
}

func (sc *AuthController) CartItemsDelete(c echo.Context) error {
	categoryID := c.Param("id")
	ctx := c.Request().Context()

	err := sc.authUseCase.CartItemsDelete(ctx, categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, false, "failed to delete a items", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "items deleted", "")
}

// func (sc *AuthController) CartsCreate(c echo.Context) error {
// 	input := request.Carts{}
// 	ctx := c.Request().Context()

// 	if err := c.Bind(&input); err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
// 	}

// 	// err := input.Validate()

// 	// if err != nil {
// 	// 	return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
// 	// }

// 	items, err := sc.authUseCase.CartsCreate(ctx, input.ToCartsDomain())

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a cart", "")
// 	}

// 	return controllers.NewResponse(c, http.StatusCreated, false, "cart registered", response.FromCartsDomain(items))
// }

// func (sc *AuthController) CartsGetByID(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	itemsID := c.Param("id")

// 	items, err := sc.authUseCase.CartsGetByID(ctx, itemsID)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusNotFound, true, "cart not found", "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "cart found", response.FromCartsDomain(items))
// }

// func (sc *AuthController) CartsGetAll(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	itemsData, err := sc.authUseCase.CartsGetAll(ctx)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
// 	}

// 	itemsResponse := []response.Carts{}

// 	for _, items := range itemsData {
// 		itemsResponse = append(itemsResponse, response.FromCartsDomain(items))
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "all cart", itemsResponse)
// }

// func (sc *AuthController) CartsDelete(c echo.Context) error {
// 	categoryID := c.Param("id")
// 	ctx := c.Request().Context()

// 	err := sc.authUseCase.CartsDelete(ctx, categoryID)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, false, "failed to delete a cart", "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "cart deleted", "")
// }
