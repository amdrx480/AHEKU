package request

import (
	"backend-golang/businesses/admin"
	"strings"

	"github.com/go-playground/validator/v10"
)

type AdminRegistration struct {
	Name     string `json:"name" validate:"required"`
	Voucher  string `json:"voucher" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminLogin struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminVoucher struct {
	Voucher string `json:"voucher" validate:"required"`
}

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
	UnitName string `json:"units_name" validate:"required"`
}

type Stocks struct {
	StockName    string `json:"stock_name" form:"stock_name"`
	StockCode    string `json:"stock_code" form:"stock_code"`
	CategoryName string `json:"category_name" form:"category_name"`
	CategoryID   uint   `json:"category_id" form:"category_id"`
	UnitID       uint   `json:"units_id" form:"units_id"`
	UnitName     string `json:"units_name" form:"units_name"`
	Description  string `json:"description" form:"description"`
	ImagePath    string `json:"image_path" form:"image_path"`
	StockTotal   int    `json:"stock_total" form:"stock_total"`
	SellingPrice int    `json:"selling_price" form:"selling_price"`
}

type Purchases struct {
	VendorID       uint   `json:"vendor_id"`
	StockName      string `json:"stock_name"`
	StockCode      string `json:"stock_code"`
	CategoryID     uint   `json:"category_id"`
	UnitID         uint   `json:"units_id"`
	Quantity       int    `json:"quantity"`
	Description    string `json:"description"`
	PurchasesPrice int    `json:"purchase_price"`
	SellingPrice   int    `json:"selling_price"`
}

type CartItems struct {
	CustomerID uint `json:"customer_id" validate:"required"`
	StockID    uint `json:"stock_id" validate:"required"`
	Quantity   int  `json:"quantity" validate:"required"`
}

type ItemTransactions struct {
	StockID  uint `json:"stock_id" `
	Quantity int  `json:"quantity" `
	SubTotal int  `json:"sub_total" `
}

func (req *AdminRegistration) ToAdminRegistrationDomain() *admin.AdminsDomain {
	return &admin.AdminsDomain{
		Name:     req.Name,
		Voucher:  req.Voucher,
		Password: req.Password,
	}
}

func (req *AdminLogin) ToAdminLoginDomain() *admin.AdminsDomain {
	return &admin.AdminsDomain{
		Name:     req.Name,
		Password: req.Password,
	}
}

func (req *AdminVoucher) ToAdminVoucherDomain() *admin.AdminsDomain {
	return &admin.AdminsDomain{
		Voucher: req.Voucher,
	}
}

func (req *Customers) ToCustomersDomain() *admin.CustomersDomain {
	return &admin.CustomersDomain{
		CustomerName:    req.CustomerName,
		CustomerAddress: req.CustomerAddress,
		CustomerEmail:   req.CustomerEmail,
		CustomerPhone:   req.CustomerPhone,
	}
}

func (req *Category) ToCategoriesDomain() *admin.CategoriesDomain {
	return &admin.CategoriesDomain{
		CategoryName: req.CategoryName,
	}
}

func (req *Vendors) ToVendorsDomain() *admin.VendorsDomain {
	return &admin.VendorsDomain{
		VendorName:    req.VendorName,
		VendorAddress: req.VendorAddress,
		VendorEmail:   req.VendorEmail,
		VendorPhone:   req.VendorPhone,
	}
}

func (req *Units) ToUnitsDomain() *admin.UnitsDomain {
	return &admin.UnitsDomain{
		UnitName: req.UnitName,
	}
}

func (req *Stocks) ToStocksDomain() *admin.StocksDomain {
	return &admin.StocksDomain{
		StockName:    req.StockName,
		StockCode:    req.StockCode,
		CategoryID:   req.CategoryID,
		CategoryName: req.CategoryName,
		UnitID:       req.UnitID,
		Description:  req.Description,
		ImagePath:    req.ImagePath,
		StockTotal:   req.StockTotal,
		SellingPrice: req.SellingPrice,
	}
}

func (req *Purchases) ToPurchasesDomain() *admin.PurchasesDomain {
	return &admin.PurchasesDomain{
		VendorID:       req.VendorID,
		StockName:      req.StockName,
		StockCode:      req.StockCode,
		CategoryID:     req.CategoryID,
		UnitID:         req.UnitID,
		Quantity:       req.Quantity,
		Description:    req.Description,
		PurchasesPrice: req.PurchasesPrice,
		SellingPrice:   req.SellingPrice,
	}
}

func (req *CartItems) ToCartItemsDomain() *admin.CartItemsDomain {
	return &admin.CartItemsDomain{
		CustomerID: req.CustomerID,
		StockID:    req.StockID,
		Quantity:   req.Quantity,
	}
}

func (req *ItemTransactions) ToItemTransactionsDomain() *admin.ItemTransactionsDomain {
	return &admin.ItemTransactionsDomain{
		StockID:  req.StockID,
		Quantity: req.Quantity,
		SubTotal: req.SubTotal,
	}
}

func validateRequest(req interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("NotEmpty", NotEmpty)

	err := validate.Struct(req)

	return err
}

func NotEmpty(fl validator.FieldLevel) bool {
	inputData := fl.Field().String()
	inputData = strings.TrimSpace(inputData)

	return inputData != ""
}

func (req *AdminRegistration) Validate() error {
	return validateRequest(req)
}

func (req *AdminLogin) Validate() error {
	return validateRequest(req)
}

func (req *AdminVoucher) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *Customers) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *Category) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *Vendors) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *Units) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *Stocks) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *Purchases) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *CartItems) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *ItemTransactions) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
