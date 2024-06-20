package request

import (
	"backend-golang/businesses/admin"
	"strings"

	"github.com/go-playground/validator/v10"
)

type AdminRegistration struct {
	ImagePath string `json:"image_path"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
	RoleID    uint   `json:"role_id" validate:"required"`

	// Role       Role   `json:"role" validate:"required"`
	// Role       string `json:"role" validate:"required"`
	Voucher  string `json:"voucher" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AdminVoucher struct {
	Voucher string `json:"voucher" validate:"required"`
}

type AdminUpdate struct {
	ImagePath string `json:"image_path"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone" validate:"required"`
	Voucher   string `json:"voucher"`
}

type Role struct {
	RoleName string `json:"role_name"`
}

// type AdminProfile struct {
// 	Name       string `json:"name" validate:"NotEmpty"`
// 	Nip        string `json:"nip" validate:"required,NotEmpty"`
// 	Division   string `json:"division" validate:"required,NotEmpty"`
// 	Image_Path string `json:"image_path"`
// }

type Customers struct {
	CustomerName    string `json:"customer_name" validate:"required"`
	CustomerAddress string `json:"customer_address" validate:"required"`
	CustomerEmail   string `json:"customer_email" validate:"required"`
	CustomerPhone   string `json:"customer_phone" validate:"required"`
}

type Category struct {
	CategoryName string `json:"category_name" validate:"required"`
}

type Vendors struct {
	VendorName    string `json:"vendor_name" validate:"required"`
	VendorAddress string `json:"vendor_address" validate:"required"`
	VendorEmail   string `json:"vendor_email" validate:"required"`
	VendorPhone   string `json:"vendor_phone" validate:"required"`
}

type Units struct {
	UnitName string `json:"unit_name" validate:"required"`
}

type Stocks struct {
	StockName    string `json:"stock_name" form:"stock_name"`
	StockCode    string `json:"stock_code" form:"stock_code"`
	CategoryName string `json:"category_name" form:"category_name"`
	CategoryID   uint   `json:"category_id" form:"category_id"`
	UnitID       uint   `json:"unit_id" form:"unit_id"`
	UnitName     string `json:"unit_name" form:"unit_name"`
	Description  string `json:"description" form:"description"`
	ImagePath    string `json:"image_path" form:"image_path"`
	StockTotal   int    `json:"stock_total" form:"stock_total"`
	SellingPrice int    `json:"selling_price" form:"selling_price"`
}

type Purchases struct {
	VendorID      uint   `json:"vendor_id"`
	StockName     string `json:"stock_name"`
	StockCode     string `json:"stock_code"`
	CategoryID    uint   `json:"category_id"`
	UnitID        uint   `json:"unit_id"`
	Quantity      int    `json:"quantity"`
	Description   string `json:"description"`
	PurchasePrice int    `json:"purchase_price"`
	SellingPrice  int    `json:"selling_price"`
}

type CartItems struct {
	CustomerID uint `json:"customer_id" validate:"required"`
	StockID    uint `json:"stock_id" validate:"required"`
	// UnitsID    uint `json:"unit_id"`
	Quantity int `json:"quantity" validate:"required"`
}

// type ItemTransactions struct {
// 	CustomerID uint `json:"customer_id" validate:"required"`
// }

func (request *AdminRegistration) ToAdminRegistrationDomain() *admin.AdminDomain {
	return &admin.AdminDomain{
		ImagePath: request.ImagePath,
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		RoleID:    request.RoleID,
		// Role:       admin.RoleDomain{RoleName: request.Role.RoleName},
		Voucher:  request.Voucher,
		Password: request.Password,
	}
}

func (request *AdminLogin) ToAdminLoginDomain() *admin.AdminDomain {
	return &admin.AdminDomain{
		Name:     request.Email,
		Password: request.Password,
	}
}

func (request *AdminVoucher) ToAdminVoucherDomain() *admin.AdminDomain {
	return &admin.AdminDomain{
		Voucher: request.Voucher,
	}
}

func (request *AdminUpdate) ToAdminUpdateDomain() *admin.AdminDomain {
	return &admin.AdminDomain{
		ImagePath: request.ImagePath,
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		Voucher:   request.Voucher,
	}
}

func (req *Role) ToRoleDomain() *admin.RoleDomain {
	return &admin.RoleDomain{
		RoleName: req.RoleName,
	}
}

// func (req *AdminProfile) ToDomain() *admin.AdminProfileDomain {
// 	return &admin.AdminProfileDomain{
// 		Name:       req.Name,
// 		Nip:        req.Nip,
// 		Division:   req.Division,
// 		Image_Path: req.Image_Path,
// 	}
// }

func (request *Customers) ToCustomersDomain() *admin.CustomersDomain {
	return &admin.CustomersDomain{
		CustomerName:    request.CustomerName,
		CustomerAddress: request.CustomerAddress,
		CustomerEmail:   request.CustomerEmail,
		CustomerPhone:   request.CustomerPhone,
	}
}

func (request *Category) ToCategoriesDomain() *admin.CategoriesDomain {
	return &admin.CategoriesDomain{
		CategoryName: request.CategoryName,
	}
}

func (request *Vendors) ToVendorsDomain() *admin.VendorsDomain {
	return &admin.VendorsDomain{
		VendorName:    request.VendorName,
		VendorAddress: request.VendorAddress,
		VendorEmail:   request.VendorEmail,
		VendorPhone:   request.VendorPhone,
	}
}

func (request *Units) ToUnitsDomain() *admin.UnitsDomain {
	return &admin.UnitsDomain{
		UnitName: request.UnitName,
	}
}

func (request *Stocks) ToStocksDomain() *admin.StocksDomain {
	return &admin.StocksDomain{
		StockName:    request.StockName,
		StockCode:    request.StockCode,
		CategoryID:   request.CategoryID,
		CategoryName: request.CategoryName,
		UnitID:       request.UnitID,
		Description:  request.Description,
		ImagePath:    request.ImagePath,
		StockTotal:   request.StockTotal,
		SellingPrice: request.SellingPrice,
	}
}

func (request *Purchases) ToPurchasesDomain() *admin.PurchasesDomain {
	return &admin.PurchasesDomain{
		VendorID:      request.VendorID,
		StockName:     request.StockName,
		StockCode:     request.StockCode,
		CategoryID:    request.CategoryID,
		UnitID:        request.UnitID,
		Quantity:      request.Quantity,
		Description:   request.Description,
		PurchasePrice: request.PurchasePrice,
		SellingPrice:  request.SellingPrice,
	}
}

func (request *CartItems) ToCartItemsDomain() *admin.CartItemsDomain {
	return &admin.CartItemsDomain{
		CustomerID: request.CustomerID,
		StockID:    request.StockID,
		// UnitsID:    request.UnitsID,
		Quantity: request.Quantity,
	}
}

// func (request *ItemTransactions) ToItemTransactionsDomain() *admin.ItemTransactionsDomain {
// 	return &admin.ItemTransactionsDomain{
// 		CustomerID: request.CustomerID,
// 		// StockID:  request.StockID,
// 		// Quantity: request.Quantity,
// 		// SubTotal: request.SubTotal,
// 	}
// }

func validateRequest(request interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("NotEmpty", NotEmpty)

	err := validate.Struct(request)

	return err
}

func NotEmpty(fl validator.FieldLevel) bool {
	inputData := fl.Field().String()
	inputData = strings.TrimSpace(inputData)

	return inputData != ""
}

func (request *AdminRegistration) Validate() error {
	return validateRequest(request)
}

func (request *AdminLogin) Validate() error {
	return validateRequest(request)
}

func (request *AdminVoucher) Validate() error {
	validate := validator.New()

	err := validate.Struct(request)

	return err
}

func (request *Role) Validate() error {
	return validateRequest(request)
}

// func (request *AdminProfile) Validate() error {
// 	return validateRequest(request)
// }

func (request *Customers) Validate() error {
	validate := validator.New()

	err := validate.Struct(request)

	return err
}

func (request *Category) Validate() error {
	validate := validator.New()

	err := validate.Struct(request)

	return err
}

func (request *Vendors) Validate() error {
	validate := validator.New()

	err := validate.Struct(request)

	return err
}

func (request *Units) Validate() error {
	validate := validator.New()

	err := validate.Struct(request)

	return err
}

func (request *Stocks) Validate() error {
	validate := validator.New()

	err := validate.Struct(request)

	return err
}

func (request *Purchases) Validate() error {
	validate := validator.New()

	err := validate.Struct(request)

	return err
}

func (request *CartItems) Validate() error {
	validate := validator.New()

	err := validate.Struct(request)

	return err
}

// func (request *ItemTransactions) Validate() error {
// 	validate := validator.New()

// 	err := validate.Struct(request)

// 	return err
// }
