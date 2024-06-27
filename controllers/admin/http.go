package admin

import (
	"backend-golang/app/middlewares"
	"backend-golang/businesses/admin"
	"backend-golang/controllers"
	"backend-golang/controllers/admin/request"
	"backend-golang/controllers/admin/response"
	"backend-golang/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

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
		return controllers.NewResponseLoginVoucher(c, http.StatusBadRequest, true, "invalid voucher request", "")
	}

	err := adminInput.Validate()

	if err != nil {
		return controllers.NewResponseLoginVoucher(c, http.StatusBadRequest, true, "validation voucher failed", "")
	}
	voucherResult, err := ctrl.authUseCase.AdminVoucher(ctx, adminInput.ToAdminVoucherDomain())

	var isFailed bool = err != nil || voucherResult == ""

	if isFailed {
		return controllers.NewResponseLoginVoucher(c, http.StatusUnauthorized, true, "invalid voucher or password", "")
	}

	return controllers.NewResponseLoginVoucher(c, http.StatusOK, false, "voucher success", voucherResult)
}

func (ctrl *AuthController) AdminProfileUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	userData, err := middlewares.GetUser(c)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnauthorized, true, err.Error(), "")
	}

	// Handling Upload File (Gambar)
	file, err := c.FormFile("image")
	var fileName string
	if err == nil {
		// Jika ada file yang diunggah, simpan ke direktori yang sesuai
		uploadDir := "images"
		filePath, err := utils.SaveUploadedFile(file, uploadDir)
		if err != nil {
			return controllers.NewResponse(c, http.StatusInternalServerError, true, "unable to save file", "")
		}
		fileName = utils.GetFileName(filePath)
	}

	// Bind JSON data dari form-data
	jsonData := c.FormValue("data")
	var adminInput request.AdminUpdate
	if jsonData != "" {
		if err := json.Unmarshal([]byte(jsonData), &adminInput); err != nil {
			return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid JSON data", "")
		}
	}

	// Jika tidak ada perubahan yang dilakukan pada kedua data admin dan gambar
	if fileName == "" && reflect.DeepEqual(adminInput, request.AdminUpdate{}) {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "no updates provided", "")
	}

	// Update Profil Admin jika ada perubahan pada data admin atau gambar
	if fileName != "" || !reflect.DeepEqual(adminInput, request.AdminUpdate{}) {
		_, _, err := ctrl.authUseCase.AdminProfileUpdate(ctx, adminInput.ToAdminUpdateDomain(), fileName, strconv.Itoa(userData.ID))
		if err != nil {
			return controllers.NewResponse(c, http.StatusNotFound, true, err.Error(), "")
		}
	}

	// Mendapatkan data admin setelah pembaruan
	user, err := ctrl.authUseCase.AdminGetByID(ctx, strconv.Itoa(userData.ID))
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, err.Error(), "")
	}

	// Menyiapkan dan Mengembalikan Respons
	return controllers.NewResponse(c, http.StatusOK, false, "admin profile updated", response.FromAdminsDomain(user))
}

// func (ctrl *AuthController) AdminProfileUpdate(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	userData, err := middlewares.GetUser(c)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusUnauthorized, true, err.Error(), "")
// 	}

// 	// adminInput := request.AdminUpdate{}
// 	// if err := c.Bind(&adminInput); err != nil {
// 	// 	return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
// 	// }

// 	// Bind JSON data from form-data
// 	// jsonData := c.FormValue("data")
// 	// adminInput := request.AdminUpdate{}
// 	// if err := json.Unmarshal([]byte(jsonData), &adminInput); err != nil {
// 	// 	return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid JSON data", "")
// 	// }

// 	file, err := c.FormFile("image")
// 	var fileName string
// 	if err == nil {
// 		// uploadDir := "D:/Skripsi/AHEKU/AHEKU/images"
// 		uploadDir := "images"
// 		filePath, err := utils.SaveUploadedFile(file, uploadDir)
// 		if err != nil {
// 			return controllers.NewResponse(c, http.StatusInternalServerError, true, "unable to save file", "")
// 		}
// 		fileName = utils.GetFileName(filePath)
// 	}

// 	// Lakukan bind JSON data jika tidak ada file yang diunggah
// 	var adminInput request.AdminUpdate
// 	if fileName == "" {
// 		jsonData := c.FormValue("data")
// 		if err := json.Unmarshal([]byte(jsonData), &adminInput); err != nil {
// 			return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid JSON data", "")
// 		}
// 	}

// 	// Update admin data
// 	user, _, err := ctrl.authUseCase.AdminProfileUpdate(ctx, adminInput.ToAdminUpdateDomain(), fileName, strconv.Itoa(userData.ID))
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusNotFound, true, err.Error(), "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "admin profile updated", response.FromAdminUpdateDomain(user))

// 	//agar dapat menampilkan gambar jika dibuat url
// 	// fileURL := fmt.Sprintf("http://localhost:8080/images/%s", fileName)
// 	// fileURL := fmt.Sprintf(fileName)

// 	// response := response.FromAdminUpdateDomain(user)
// 	// response.ImagePath = fileURL

// 	// return controllers.NewResponse(c, http.StatusOK, false, "admin profile updated", response)
// }

// func (ctrl *AuthController) AdminProfileUpdate(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	// Mendapatkan data user dari token
// 	userData, err := middlewares.GetUser(c)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusUnauthorized, true, err.Error(), "")
// 	}

// 	adminInput := request.AdminUpdate{}

// 	// userID := c.Param("id")
// 	file, err := c.FormFile("image")

// 	var filePath string
// 	if err != nil {

// 		fmt.Println("Unable to open file:", err.Error())

// 		return controllers.NewResponse(c, http.StatusBadRequest, true, "error handling file upload", "")
// 	}

// 	src, err := file.Open()
// 	if err != nil {
// 		fmt.Println("Unable to open file:", err.Error())

// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "Unable to open file", "")
// 	}
// 	defer src.Close()

// 	avatarPath := file.Filename
// 	user, _, err := ctrl.authUseCase.AdminProfileUpdate(ctx, adminInput.ToAdminUpdateDomain(), avatarPath, strconv.Itoa(userData.ID))

// 	if err != nil {
// 		fmt.Println("Unable to open file:", err.Error())

// 		return controllers.NewResponse(c, http.StatusNotFound, true, err.Error(), "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "customer profile image updated", response.FromAdminUpdateDomain(user))
// }

func (ctrl *AuthController) AdminGetProfile(c echo.Context) error {
	ctx := c.Request().Context()

	// Mendapatkan data user dari token
	userData, err := middlewares.GetUser(c)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnauthorized, true, err.Error(), "")
	}

	// Menggunakan ID dari token untuk mengambil data admin
	user, err := ctrl.authUseCase.AdminGetByID(ctx, strconv.Itoa(userData.ID))
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, err.Error(), "")
	}

	// Konversi admin domain ke respons dengan URL gambar lengkap
	resp := response.FromAdminsDomain(user)
	if user.ImagePath != "" {
		resp.ImagePath = fmt.Sprintf("http://192.168.253.91:8080/images/%s", user.ImagePath)
	}
	return controllers.NewResponse(c, http.StatusOK, false, "admin info found", resp)

	// return controllers.NewResponse(c, http.StatusOK, false, "admin info found", response.FromAdminsDomain(user))
}

// func (ctrl *AuthController) AdminGetInfo(c echo.Context) error {
// 	// Ambil informasi pengguna dari token JWT
// 	user, err := middlewares.GetUser(c)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusUnauthorized, true, "invalid token", "")
// 	}

// 	ctx := c.Request().Context()

// 	// Ambil informasi admin berdasarkan ID yang terdapat dalam token JWT
// 	admin, err := ctrl.authUseCase.AdminGetByID(ctx, strconv.Itoa(user.ID))
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusNotFound, true, err.Error(), "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "admin get info found", response.FromAdminsDomain(admin))
// }

func (ctrl *AuthController) AdminGetByID(c echo.Context) error {
	var adminsID string = c.Param("id")

	ctx := c.Request().Context()

	user, err := ctrl.authUseCase.AdminGetByID(ctx, adminsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "admin get by id found", response.FromAdminsDomain(user))
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (ac *AuthController) RoleCreate(c echo.Context) error {
	input := request.Role{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	role, err := ac.authUseCase.RoleCreate(ctx, input.ToRoleDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a role", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "role registered", response.FromRoleDomain(role))
}

func (ac *AuthController) RoleGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	roleID := c.Param("id")

	role, err := ac.authUseCase.RoleGetByID(ctx, roleID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "role not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "role found", response.FromRoleDomain(role))
}

func (ac *AuthController) RoleGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	roleData, err := ac.authUseCase.RoleGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	roles := []response.Role{}

	for _, role := range roleData {
		roles = append(roles, response.FromRoleDomain(role))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all vendors", roles)
}

// func (ctrl *AuthController) AdminProfileUpdate(c echo.Context) error {
// 	profileInput := request.AdminProfile{}
// 	ctx := c.Request().Context()

// 	var userID string = c.Param("id")

// 	if err := c.Bind(&profileInput); err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
// 	}

// 	err := profileInput.Validate()
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
// 	}
// 	//percobaan
// 	if userID == "" {
// 		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid user ID", "")
// 	}

// 	// nil
// 	profile, err := ctrl.authUseCase.AdminProfileUpdate(ctx, profileInput.ToDomain(), userID)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusNotFound, true, "No data found", "")
// 	}

// 	// nil
// 	return controllers.NewResponse(c, http.StatusOK, false, "customer updated", response.FromAdminProfileDomain(profile))
// }

// func (ctrl *AuthController) AdminProfileUploadImage(c echo.Context) error {
// 	profileInput := request.AdminProfile{}
// 	ctx := c.Request().Context()

// 	userID := c.Param("id")
// 	file, err := c.FormFile("image")

// 	if err != nil {

// 		fmt.Println("Unable to open file:", err.Error())

// 		return controllers.NewResponse(c, http.StatusBadRequest, true, "error handling file upload", "")
// 	}

// 	src, err := file.Open()
// 	if err != nil {
// 		fmt.Println("Unable to open file:", err.Error())

// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "Unable to open file", "")
// 	}
// 	defer src.Close()

// 	avatarPath := file.Filename
// 	user, _, err := ctrl.authUseCase.AdminProfileUploadImage(ctx, profileInput.ToDomain(), avatarPath, userID)

// 	if err != nil {
// 		fmt.Println("Unable to open file:", err.Error())

// 		return controllers.NewResponse(c, http.StatusNotFound, true, err.Error(), "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "customer profile image updated", user)
// }

// func (ctrl *AuthController) AdminProfileGetByID(c echo.Context) error {
// 	var userID string = c.Param("id")

// 	ctx := c.Request().Context()

// 	user, err := ctrl.authUseCase.AdminProfileGetByID(ctx, userID)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusNotFound, true, err.Error(), "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "admin profile found", response.FromAdminProfileDomain(user))
// }

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (ac *AuthController) CustomersGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	customersID := c.Param("id")

	customers, err := ac.authUseCase.CustomersGetByID(ctx, customersID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "customer not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "customer found", response.FromCustomersDomain(customers))
}

func (ac *AuthController) CustomersCreate(c echo.Context) error {
	input := request.Customers{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	customers, err := ac.authUseCase.CustomersCreate(ctx, input.ToCustomersDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a customer", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "customer registered", response.FromCustomersDomain(customers))
}

func (ac *AuthController) CustomersGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	categoriesData, err := ac.authUseCase.CustomersGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	categories := []response.Customers{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromCustomersDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all customers", categories)
}

func (ac *AuthController) CategoryGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	categoryID := c.Param("id")

	category, err := ac.authUseCase.CategoryGetByID(ctx, categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "category found", response.FromCategoryDomain(category))
}

func (ac *AuthController) CategoryGetByName(c echo.Context) error {
	ctx := c.Request().Context()

	categoryName := c.Param("category_name")

	category, err := ac.authUseCase.CategoryGetByName(ctx, categoryName)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "categoryName not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "categoryName found", response.FromCategoryDomain(category))
}

func (ac *AuthController) CategoryCreate(c echo.Context) error {
	input := request.Category{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	category, err := ac.authUseCase.CategoryCreate(ctx, input.ToCategoriesDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a category", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "category registered", response.FromCategoryDomain(category))
}

func (ac *AuthController) CategoryGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	categoryData, err := ac.authUseCase.CategoryGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	categories := []response.Category{}

	for _, category := range categoryData {
		categories = append(categories, response.FromCategoryDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all category", categories)
}

func (ac *AuthController) VendorsCreate(c echo.Context) error {
	input := request.Vendors{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	vendors, err := ac.authUseCase.VendorsCreate(ctx, input.ToVendorsDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a vendor", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "vendor registered", response.FromVendorsDomain(vendors))
}

func (ac *AuthController) VendorsGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	vendorsID := c.Param("id")

	vendors, err := ac.authUseCase.VendorsGetByID(ctx, vendorsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "vendor not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "vendor found", response.FromVendorsDomain(vendors))
}

func (ac *AuthController) VendorsGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	categoriesData, err := ac.authUseCase.VendorsGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	categories := []response.Vendors{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromVendorsDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all vendors", categories)
}

func (ac *AuthController) UnitsCreate(c echo.Context) error {
	input := request.Units{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	unit, err := ac.authUseCase.UnitsCreate(ctx, input.ToUnitsDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a unit", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "unit registered", response.FromUnitsDomain(unit))
}

func (ac *AuthController) UnitsGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	unitsID := c.Param("id")

	unit, err := ac.authUseCase.UnitsGetByID(ctx, unitsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "unit not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "unit found", response.FromUnitsDomain(unit))
}

func (ac *AuthController) UnitsGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	unitsData, err := ac.authUseCase.UnitsGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	units := []response.Units{}

	for _, unit := range unitsData {
		units = append(units, response.FromUnitsDomain(unit))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all units", units)
}

func (ac *AuthController) PurchasesCreate(c echo.Context) error {
	input := request.Purchases{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	purchases, err := ac.authUseCase.PurchasesCreate(ctx, input.ToPurchasesDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a purchases", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "purchases registered", response.FromPurchasesDomain(purchases))
}

func (ac *AuthController) PurchasesGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	purchasesID := c.Param("id")

	purchases, err := ac.authUseCase.PurchasesGetByID(ctx, purchasesID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "category found", response.FromPurchasesDomain(purchases))
}

func (ac *AuthController) PurchasesGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit <= 0 {
		limit = 5
	}

	sort := c.QueryParam("sort")
	if sort == "" {
		sort = "id"
	}

	order := c.QueryParam("order")
	if order == "" {
		order = "asc"
	}

	search := c.QueryParam("search")

	filters := make(map[string]interface{})

	vendorID, _ := strconv.Atoi(c.QueryParam("vendor_id"))
	if vendorID != 0 {
		filters["vendor_id"] = vendorID
	}

	categoryID, _ := strconv.Atoi(c.QueryParam("category_id"))
	if categoryID != 0 {
		filters["category_id"] = categoryID
	}

	unitID, _ := strconv.Atoi(c.QueryParam("unit_id"))
	if unitID != 0 {
		filters["unit_id"] = unitID
	}

	quantityMin, _ := strconv.Atoi(c.QueryParam("quantity_min"))
	if quantityMin != 0 {
		filters["quantity_min"] = quantityMin
	}

	quantityMax, _ := strconv.Atoi(c.QueryParam("quantity_max"))
	if quantityMax != 0 {
		filters["quantity_max"] = quantityMax
	}

	purchasePriceMin, _ := strconv.Atoi(c.QueryParam("purchase_price_min"))
	if purchasePriceMin != 0 {
		filters["purchase_price_min"] = purchasePriceMin
	}

	purchasePriceMax, _ := strconv.Atoi(c.QueryParam("purchase_price_max"))
	if purchasePriceMax != 0 {
		filters["purchase_price_max"] = purchasePriceMax
	}

	sellingPriceMin, _ := strconv.Atoi(c.QueryParam("selling_price_min"))
	if sellingPriceMin != 0 {
		filters["selling_price_min"] = sellingPriceMin
	}

	sellingPriceMax, _ := strconv.Atoi(c.QueryParam("selling_price_max"))
	if sellingPriceMax != 0 {
		filters["selling_price_max"] = sellingPriceMax
	}

	purchasesData, totalItems, err := ac.authUseCase.PurchasesGetAll(ctx, page, limit, sort, order, search, filters)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "Failed to fetch data", "")
	}

	purchases := []response.Purchases{}
	for _, purchase := range purchasesData {
		purchases = append(purchases, response.FromPurchasesDomain(purchase))
	}

	return controllers.NewPaginatedResponse(c, http.StatusOK, "All purchases", purchases, page, limit, totalItems)
}

// func (ac *AuthController) PurchasesGetAll(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	// Get query parameters for pagination, sorting, and search
// 	page, err := strconv.Atoi(c.QueryParam("page"))
// 	if err != nil || page < 1 {
// 		page = 1
// 	}

// 	limit, err := strconv.Atoi(c.QueryParam("limit"))
// 	if err != nil || limit < 1 {
// 		limit = 10
// 	}

// 	sort := c.QueryParam("sort")
// 	if sort == "" {
// 		sort = "id"
// 	}

// 	order := c.QueryParam("order")
// 	if order == "" {
// 		order = "asc"
// 	}

// 	search := c.QueryParam("search")

// 	// Call usecase to get purchases data
// 	purchases, totalItems, err := ac.authUseCase.PurchasesGetAll(ctx, page, limit, sort, order, search)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "Failed to fetch data", "")
// 	}

// 	// Use NewPaginatedResponse to create paginated response
// 	return controllers.NewPaginatedResponse(c, http.StatusOK, "All purchases", purchases, page, limit, totalItems)
// }

//////
// func (ac *AuthController) PurchasesGetAll(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	// Fetch purchases data from use case
// 	purchasesData, err := ac.authUseCase.PurchasesGetAll(ctx)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "Failed to fetch data", "")
// 	}

// 	// Convert domain purchases to response purchases
// 	purchases := []response.Purchases{}
// 	for _, purchase := range purchasesData {
// 		purchases = append(purchases, response.FromPurchasesDomain(purchase))
// 	}

// 	// Paginate the purchases data based on query parameters
// 	page, _ := strconv.Atoi(c.QueryParam("page"))
// 	if page == 0 {
// 		page = 1 // Default to page 1 if not specified or invalid
// 	}

// 	// Menentukan default item per page yang di tampilkan
// 	size, _ := strconv.Atoi(c.QueryParam("size"))
// 	if size == 0 {
// 		size = 5 // Default to 10 items per page if not specified or invalid
// 	}
// 	totalItems := len(purchases)

// 	// Slice data based on pagination parameters
// 	paginatedData := utils.Paginate(purchases, page, size)

// 	// paginatedData, err := utils.Paginate(purchases, page, size)
// 	// if err != nil {
// 	// 	return controllers.NewResponse(c, http.StatusBadRequest, true, err.Error(), "")
// 	// }

// 	// Prepare paginated response using NewPaginatedResponse
// 	return controllers.NewPaginatedResponse(c, http.StatusOK, "All categories", paginatedData, page, size, totalItems)
// }

// func (ac *AuthController) PurchasesGetAll(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	purchasesData, err := ac.authUseCase.PurchasesGetAll(ctx)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
// 	}

// 	purchases := []response.Purchases{}

// 	for _, purchase := range purchasesData {
// 		purchases = append(purchases, response.FromPurchasesDomain(purchase))
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "all purchases", purchases)
// }

func (ac *AuthController) StocksCreate(c echo.Context) error {
	input := request.Stocks{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	stock, err := ac.authUseCase.StocksCreate(ctx, input.ToStocksDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a stock", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "stock registered", response.FromStocksDomain(stock))
}

func (ac *AuthController) StocksGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	stockID := c.Param("id")

	stock, err := ac.authUseCase.StocksGetByID(ctx, stockID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "category found", response.FromStocksDomain(stock))
}

func (ac *AuthController) StocksGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	// Ambil parameter query untuk pagination, sorting, dan search
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1 // Default to page 1 if not specified or invalid
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10 // Default to 10 items per page if not specified or invalid
	}
	sort := c.QueryParam("sort")
	if sort == "" {
		sort = "id" // Default sorting field
	}
	order := c.QueryParam("order")
	if order == "" {
		order = "asc" // Default order
	}
	search := c.QueryParam("search")

	// Example of extracting filters from query params
	filters := make(map[string]interface{})
	categoryID, _ := strconv.Atoi(c.QueryParam("category_id"))
	if categoryID != 0 {
		filters["category_id"] = categoryID
	}
	unitID, _ := strconv.Atoi(c.QueryParam("unit_id"))
	if unitID != 0 {
		filters["unit_id"] = unitID
	}
	stockTotalMin, _ := strconv.Atoi(c.QueryParam("stock_total_min"))
	if stockTotalMin != 0 {
		filters["stock_total_min"] = stockTotalMin
	}
	stockTotalMax, _ := strconv.Atoi(c.QueryParam("stock_total_max"))
	if stockTotalMax != 0 {
		filters["stock_total_max"] = stockTotalMax
	}
	sellingPriceMin, _ := strconv.Atoi(c.QueryParam("selling_price_min"))
	if sellingPriceMin != 0 {
		filters["selling_price_min"] = sellingPriceMin
	}
	sellingPriceMax, _ := strconv.Atoi(c.QueryParam("selling_price_max"))
	if sellingPriceMax != 0 {
		filters["selling_price_max"] = sellingPriceMax
	}
	sellingPriceOrder := c.QueryParam("selling_price_order")
	if sellingPriceOrder != "" {
		filters["selling_price_order"] = sellingPriceOrder
	}

	// Fetch stocks data from use case with pagination, sorting, search, and filters
	stocksData, totalItems, err := ac.authUseCase.StocksGetAll(ctx, page, limit, sort, order, search, filters)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "Failed to fetch data", "")
	}

	// Convert domain stocks to response stocks
	stocks := make([]response.Stocks, len(stocksData))
	for i, stock := range stocksData {
		stocks[i] = response.FromStocksDomain(stock)
	}

	// Prepare paginated response using NewPaginatedResponse
	return controllers.NewPaginatedResponse(c, http.StatusOK, "Stocks retrieved successfully", stocks, page, limit, totalItems)
}

// func (ac *AuthController) StocksGetAll(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	// Ambil parameter query untuk pagination, sorting, dan search
// 	page, _ := strconv.Atoi(c.QueryParam("page"))
// 	if page == 0 {
// 		page = 1 // Default to page 1 if not specified or invalid
// 	}
// 	limit, _ := strconv.Atoi(c.QueryParam("limit"))
// 	if limit == 0 {
// 		limit = 5 // Default to 5 items per page if not specified or invalid
// 	}
// 	sort := c.QueryParam("sort")
// 	if sort == "" {
// 		sort = "id" // Default sorting field
// 	}
// 	order := c.QueryParam("order")
// 	if order == "" {
// 		order = "asc" // Default order
// 	}
// 	search := c.QueryParam("search")

// 	// Fetch stocks data from use case with pagination, sorting, and search
// 	stocksData, totalItems, err := ac.authUseCase.StocksGetAll(ctx, page, limit, sort, order, search)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "Failed to fetch data", "")
// 	}

// 	// Convert domain stocks to response stocks
// 	stocks := []response.Stocks{}
// 	for _, stock := range stocksData {
// 		stocks = append(stocks, response.FromStocksDomain(stock))
// 	}

// 	// Calculate total pages
// 	// totalPages := (totalItems + limit - 1) / limit

// 	// Prepare paginated response using NewPaginatedResponse
// 	return controllers.NewPaginatedResponse(c, http.StatusOK, "All stocks", stocks, page, limit, totalItems)
// 	// return controllers.NewPaginatedResponse(c, http.StatusOK, "All stocks", stocks, page, limit, totalPages, totalItems)
// }

// func (ac *AuthController) StocksGetAll(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	// Fetch stocks data from use case
// 	stocksData, err := ac.authUseCase.StocksGetAll(ctx)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "Failed to fetch data", "")
// 		// return controllers.NewPaginatedResponse(c, http.StatusInternalServerError, "failed to fetch data", nil, 1, 10, 0)
// 	}

// 	// Convert domain stocks to response stocks
// 	stocks := []response.Stocks{}
// 	for _, stock := range stocksData {
// 		stocks = append(stocks, response.FromStocksDomain(stock))
// 	}

// 	// Paginate the stocks data based on query parameters
// 	page, _ := strconv.Atoi(c.QueryParam("page"))
// 	if page == 0 {
// 		page = 1 // Default to page 1 if not specified or invalid
// 	}
// 	limit, _ := strconv.Atoi(c.QueryParam("limit"))
// 	if limit == 0 {
// 		limit = 5 // Default to 10 items per page if not specified or invalid
// 	}
// 	totalItems := len(stocks)

// 	// Slice data based on pagination parameters
// 	paginatedData := utils.Paginate(stocks, page, limit)

// 	// Prepare paginated response using NewPaginatedResponse
// 	return controllers.NewPaginatedResponse(c, http.StatusOK, "All stocks", paginatedData, page, limit, totalItems)
// }

// func (ac *AuthController) StocksGetAll(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	stocksData, err := ac.authUseCase.StocksGetAll(ctx)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
// 	}

// 	categories := []response.Stocks{}

// 	for _, category := range stocksData {
// 		categories = append(categories, response.FromStocksDomain(category))
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "all categories", categories)
// }

func (ac *AuthController) CartItemsCreate(c echo.Context) error {
	input := request.CartItems{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	// err := input.Validate()

	// if err != nil {
	// 	return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	// }

	items, err := ac.authUseCase.CartItemsCreate(ctx, input.ToCartItemsDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a items", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "items registered", response.FromCartItemsDomain(items))
}

func (ac *AuthController) CartItemsGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	itemsID := c.Param("id")

	items, err := ac.authUseCase.CartItemsGetByID(ctx, itemsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "items not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "items found", response.FromCartItemsDomain(items))
}

func (ac *AuthController) CartItemsGetAllByCustomerID(c echo.Context) error {
	ctx := c.Request().Context()

	customerID := c.Param("customer_id")

	cartitems, err := ac.authUseCase.CartItemsGetAllByCustomerID(ctx, customerID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "items not found", "")
	}

	cartItemsResponse := []response.CartItems{}

	for _, items := range cartitems {
		cartItemsResponse = append(cartItemsResponse, response.FromCartItemsDomain(items))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all cartItems", cartItemsResponse)
}

// func (ac *AuthController) CartItemsGetByCustomerID(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	customerID := c.Param("customer_id")

// 	cartitems, err := ac.authUseCase.CartItemsGetByCustomerID(ctx, customerID)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusNotFound, true, "items not found", "")
// 	}

// 	cartItemsResponse := []response.CartItems{}

// 	for _, items := range cartitems {
// 		cartItemsResponse = append(cartItemsResponse, response.FromCartItemsDomain(items))
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "all cartItems", cartItemsResponse)
// }

func (ac *AuthController) CartItemsGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	itemsData, err := ac.authUseCase.CartItemsGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	itemsResponse := []response.CartItems{}

	for _, items := range itemsData {
		itemsResponse = append(itemsResponse, response.FromCartItemsDomain(items))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all items", itemsResponse)
}

func (ac *AuthController) CartItemsDelete(c echo.Context) error {
	categoryID := c.Param("id")
	ctx := c.Request().Context()

	err := ac.authUseCase.CartItemsDelete(ctx, categoryID)

	if err != nil {
		return controllers.NewResponseWithoutData(c, http.StatusInternalServerError, false, "failed to delete a items")
	}

	return controllers.NewResponseWithoutData(c, http.StatusOK, false, "items deleted")
}

func (ac *AuthController) ItemTransactionsCreate(c echo.Context) error {
	customerID := c.Param("customer_id") // Ambil customer_id dari parameter URL
	ctx := c.Request().Context()

	// Panggil use case untuk memproses transaksi berdasarkan customerID
	_, err := ac.authUseCase.ItemTransactionsCreate(ctx, customerID)

	if err != nil {
		log.Printf("Error creating item transactions for customer ID %s: %v\n", customerID, err)
		// return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to create item transactions", "")
		return controllers.NewResponse(c, http.StatusInternalServerError, true, err.Error(), "")
	}

	return controllers.NewResponseWithoutData(c, http.StatusCreated, false, "successfully created item transactions")
}

// func (ac *AuthController) ItemTransactionsCreate(c echo.Context) error {
// 	historyID := c.Param("customer_id")
// 	input := request.ItemTransactions{}
// 	ctx := c.Request().Context()

// 	if err := c.Bind(&input); err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
// 	}

// 	err := input.Validate()

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
// 	}

// 	// history, err := ac.salesUseCase.ToHistory(ctx, input.ToDomain(), historyID)
// 	_, err = ac.authUseCase.ItemTransactionsCreate(ctx, input.ToItemTransactionsDomain(), historyID)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "Missing item transactions data", "")
// 	}

// 	// return controllers.NewResponse(c, http.StatusCreated, false, "Success ToHistory Data", _resHistory.FromAdminProfileDomain(history))
// 	// return controllers.NewResponse(c, http.StatusCreated, false, "Success ToHistory Data", " ")
// 	return controllers.NewResponseWithoutData(c, http.StatusCreated, false, "Success item transactions Data")

// }

func (hc *AuthController) ItemTransactionsGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	historiesData, err := hc.authUseCase.ItemTransactionsGetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch item transactions data", "")
	}

	histories := []response.ItemTransactions{}

	for _, course := range historiesData {
		histories = append(histories, response.FromItemTransactionsDomain(course))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all item transactions", histories)
}

// func (ac *AuthController) CartsCreate(c echo.Context) error {
// 	input := request.Carts{}
// 	ctx := c.Request().Context()

// 	if err := c.Bind(&input); err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
// 	}

// 	// err := input.Validate()

// 	// if err != nil {
// 	// 	return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
// 	// }

// 	items, err := ac.authUseCase.CartsCreate(ctx, input.ToCartsDomain())

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a cart", "")
// 	}

// 	return controllers.NewResponse(c, http.StatusCreated, false, "cart registered", response.FromCartsDomain(items))
// }

// func (ac *AuthController) CartsGetByID(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	itemsID := c.Param("id")

// 	items, err := ac.authUseCase.CartsGetByID(ctx, itemsID)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusNotFound, true, "cart not found", "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "cart found", response.FromCartsDomain(items))
// }

// func (ac *AuthController) CartsGetAll(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	itemsData, err := ac.authUseCase.CartsGetAll(ctx)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
// 	}

// 	itemsResponse := []response.Carts{}

// 	for _, items := range itemsData {
// 		itemsResponse = append(itemsResponse, response.FromCartsDomain(items))
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "all cart", itemsResponse)
// }

// func (ac *AuthController) CartsDelete(c echo.Context) error {
// 	categoryID := c.Param("id")
// 	ctx := c.Request().Context()

// 	err := ac.authUseCase.CartsDelete(ctx, categoryID)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, false, "failed to delete a cart", "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, false, "cart deleted", "")
// }
