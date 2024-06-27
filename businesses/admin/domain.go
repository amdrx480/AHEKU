package admin

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Admin
type AdminDomain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	ImagePath string // Perubahan nama dari Image_Path ke ImagePath
	Name      string
	Email     string
	Phone     string
	RoleID    uint
	RoleName  string
	Voucher   string
	Password  string
}

type RoleDomain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	RoleName  string
}

// Customers
type CustomersDomain struct {
	ID              uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	CustomerName    string
	CustomerAddress string
	CustomerEmail   string
	CustomerPhone   string
	CartItems       []CartItemsDomain
}

// Categories
type CategoriesDomain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	CategoryName string
}

// Vendors
type VendorsDomain struct {
	ID            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	VendorName    string
	VendorAddress string
	VendorEmail   string
	VendorPhone   string
}

type UnitsDomain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	UnitName  string
}

type StocksDomain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	StockName    string
	StockCode    string
	CategoryName string
	CategoryID   uint
	UnitID       uint
	UnitName     string
	Description  string
	ImagePath    string
	StockTotal   int
	SellingPrice int
}

type PurchasesDomain struct {
	ID            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	VendorID      uint
	VendorName    string
	StockName     string
	StockCode     string
	CategoryID    uint
	CategoryName  string
	UnitID        uint
	UnitName      string
	Quantity      int
	Description   string
	PurchasePrice int
	SellingPrice  int
}

// Many CartItems To One Customer
type CartItemsDomain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	CustomerID   uint
	CustomerName string
	StockID      uint
	StockName    string
	UnitID       uint
	UnitName     string
	Quantity     int
	SellingPrice int
	Price        int
	SubTotal     int
}

type ItemTransactionsDomain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	CustomerID   uint
	CustomerName string
	StockID      uint
	StockName    string
	UnitID       uint
	UnitName     string
	CategoryID   uint
	CategoryName string
	Quantity     int
	Price        int
	SubTotal     int
}

// Role       RoleDomain

// type AdminProfileDomain struct {
// 	ID         uint
// 	CreatedAt  time.Time
// 	UpdatedAt  time.Time
// 	DeletedAt  gorm.DeletedAt
// 	Name       string
// 	Nip        string
// 	Division   string
// 	Image_Path string
// }

// One Carts To Many CartItems
// type CartsDomain struct {
// 	ID         uint
// 	CreatedAt  time.Time
// 	UpdatedAt  time.Time
// 	DeletedAt  gorm.DeletedAt
// 	CustomerID uint
// 	CartItems  []CartItemsDomain
// 	Total      int
// 	Status     string
// }

type Usecase interface {
	// Admin
	AdminRegister(ctx context.Context, adminDomain *AdminDomain) (AdminDomain, error)
	AdminLogin(ctx context.Context, adminDomain *AdminDomain) (string, error)
	AdminVoucher(ctx context.Context, adminDomain *AdminDomain) (string, error)
	AdminProfileUpdate(ctx context.Context, adminDomain *AdminDomain, imangePath string, id string) (AdminDomain, string, error)
	AdminGetProfile(ctx context.Context, id string) (AdminDomain, error)
	AdminGetByID(ctx context.Context, id string) (AdminDomain, error)

	RoleCreate(ctx context.Context, roleDomain *RoleDomain) (RoleDomain, error)
	RoleGetByID(ctx context.Context, id string) (RoleDomain, error)
	RoleGetAll(ctx context.Context) ([]RoleDomain, error)

	// Admin Profile
	// AdminProfileUpdate(ctx context.Context, profileDomain *AdminProfileDomain, id string) (AdminProfileDomain, error)
	// AdminProfileUploadImage(ctx context.Context, profileDomain *AdminProfileDomain, avatarPath string, id string) (AdminProfileDomain, string, error)
	// AdminProfileGetByID(ctx context.Context, id string) (AdminProfileDomain, error)

	// Customers
	CustomersGetByID(ctx context.Context, id string) (CustomersDomain, error)
	CustomersCreate(ctx context.Context, customersDomain *CustomersDomain) (CustomersDomain, error)
	CustomersGetAll(ctx context.Context) ([]CustomersDomain, error)

	// Category
	CategoryGetByID(ctx context.Context, id string) (CategoriesDomain, error)
	CategoryGetByName(ctx context.Context, name string) (CategoriesDomain, error)
	CategoryCreate(ctx context.Context, categoryDomain *CategoriesDomain) (CategoriesDomain, error)
	CategoryGetAll(ctx context.Context) ([]CategoriesDomain, error)

	// Vendors
	VendorsCreate(ctx context.Context, vendorsDomain *VendorsDomain) (VendorsDomain, error)
	VendorsGetByID(ctx context.Context, id string) (VendorsDomain, error)
	VendorsGetAll(ctx context.Context) ([]VendorsDomain, error)

	// Units
	UnitsGetByID(ctx context.Context, id string) (UnitsDomain, error)
	UnitsCreate(ctx context.Context, unitsDomain *UnitsDomain) (UnitsDomain, error)
	UnitsGetAll(ctx context.Context) ([]UnitsDomain, error)

	// Stocks
	StocksCreate(ctx context.Context, stocksDomain *StocksDomain) (StocksDomain, error)
	StocksGetByID(ctx context.Context, id string) (StocksDomain, error)
	StocksGetAll(ctx context.Context, page int, limit int, sort string, order string, search string, filters map[string]interface{}) ([]StocksDomain, int, error)
	// StocksGetAll(ctx context.Context, page int, limit int, sort string, order string, search string) ([]StocksDomain, int, error)
	// StocksGetAll(ctx context.Context) ([]StocksDomain, error)

	// Purchases
	PurchasesGetByID(ctx context.Context, id string) (PurchasesDomain, error)
	PurchasesCreate(ctx context.Context, purchasesDomain *PurchasesDomain) (PurchasesDomain, error)
	PurchasesGetAll(ctx context.Context, page int, limit int, sort string, order string, search string, filters map[string]interface{}) ([]PurchasesDomain, int, error)
	// PurchasesGetAll(ctx context.Context, page int, limit int, sort string, order string, search string) ([]PurchasesDomain, int, error)

	// CartItems
	CartItemsCreate(ctx context.Context, cartItemsDomain *CartItemsDomain) (CartItemsDomain, error)
	CartItemsGetByID(ctx context.Context, id string) (CartItemsDomain, error)
	// CartItemsGetByCustomerID(ctx context.Context, cartItemsDomain *CartItemsDomain) (CartItemsDomain, error)
	// CartItemsGetByCustomerID(ctx context.Context, customerId string) ([]CartItemsDomain, error)
	CartItemsGetAllByCustomerID(ctx context.Context, customerId string) ([]CartItemsDomain, error)
	CartItemsGetAll(ctx context.Context) ([]CartItemsDomain, error)
	CartItemsDelete(ctx context.Context, id string) error

	// ItemTransactionsCreate(ctx context.Context, itemTransactionsDomain *ItemTransactionsDomain, id string) (ItemTransactionsDomain, error)
	ItemTransactionsCreate(ctx context.Context, customerId string) (ItemTransactionsDomain, error)
	ItemTransactionsGetAll(ctx context.Context) ([]ItemTransactionsDomain, error)

	// Carts
	// CartsGetByID(ctx context.Context, id string) (CartsDomain, error)
	// CartsCreate(ctx context.Context, cartDomain *CartsDomain) (CartsDomain, error)
	// CartsGetAll(ctx context.Context) ([]CartsDomain, error)
	// CartsDelete(ctx context.Context, id string) error
}

type Repository interface {
	// Admin
	AdminRegister(ctx context.Context, adminDomain *AdminDomain) (AdminDomain, error)
	AdminGetByEmail(ctx context.Context, adminDomain *AdminDomain) (AdminDomain, error)
	AdminGetByVoucher(ctx context.Context, adminDomain *AdminDomain) (AdminDomain, error)
	AdminProfileUpdate(ctx context.Context, adminDomain *AdminDomain, imangePath string, id string) (AdminDomain, string, error)
	// AdminGetInfo(ctx context.Context, id string) (AdminDomain, error)

	AdminGetByID(ctx context.Context, id string) (AdminDomain, error)

	RoleCreate(ctx context.Context, roleDomain *RoleDomain) (RoleDomain, error)
	RoleGetByID(ctx context.Context, id string) (RoleDomain, error)
	RoleGetAll(ctx context.Context) ([]RoleDomain, error)

	// Admin Profile
	// AdminProfileUpdate(ctx context.Context, profileDomain *AdminProfileDomain, id string) (AdminProfileDomain, error)
	// AdminProfileUploadImage(ctx context.Context, profileDomain *AdminProfileDomain, avatarPath string, id string) (AdminProfileDomain, string, error)
	// AdminProfileGetByID(ctx context.Context, id string) (AdminProfileDomain, error)

	// Customers
	CustomersGetByID(ctx context.Context, id string) (CustomersDomain, error)
	CustomersCreate(ctx context.Context, customersDomain *CustomersDomain) (CustomersDomain, error)
	CustomersGetAll(ctx context.Context) ([]CustomersDomain, error)

	// Category
	CategoryCreate(ctx context.Context, categoryDomain *CategoriesDomain) (CategoriesDomain, error)
	CategoryGetByID(ctx context.Context, id string) (CategoriesDomain, error)
	CategoryGetByName(ctx context.Context, categoryName string) (CategoriesDomain, error)
	CategoryGetAll(ctx context.Context) ([]CategoriesDomain, error)

	// Vendors
	VendorsCreate(ctx context.Context, vendorsDomain *VendorsDomain) (VendorsDomain, error)
	VendorsGetByID(ctx context.Context, id string) (VendorsDomain, error)
	VendorsGetAll(ctx context.Context) ([]VendorsDomain, error)

	// Units
	UnitsCreate(ctx context.Context, unitsDomain *UnitsDomain) (UnitsDomain, error)
	UnitsGetByID(ctx context.Context, id string) (UnitsDomain, error)
	UnitsGetAll(ctx context.Context) ([]UnitsDomain, error)

	// Stocks
	StocksCreate(ctx context.Context, stocksDomain *StocksDomain) (StocksDomain, error)
	StocksGetByID(ctx context.Context, id string) (StocksDomain, error)
	StocksGetAll(ctx context.Context, page int, limit int, sort string, order string, search string, filters map[string]interface{}) ([]StocksDomain, int, error)
	// StocksGetAll(ctx context.Context, page int, limit int, sort string, order string, search string) ([]StocksDomain, int, error)
	// StocksGetAll(ctx context.Context) ([]StocksDomain, error)

	// Purchases
	PurchasesGetByID(ctx context.Context, id string) (PurchasesDomain, error)
	PurchasesCreate(ctx context.Context, purchasesDomain *PurchasesDomain) (PurchasesDomain, error)
	PurchasesGetAll(ctx context.Context, page int, limit int, sort string, order string, search string, filters map[string]interface{}) ([]PurchasesDomain, int, error)
	// PurchasesGetAll(ctx context.Context, page int, limit int, sort string, order string, search string) ([]PurchasesDomain, int, error)

	// CartItems
	CartItemsCreate(ctx context.Context, cartItemsDomain *CartItemsDomain) (CartItemsDomain, error)
	CartItemsGetByID(ctx context.Context, id string) (CartItemsDomain, error)
	// CartItemsGetByCustomerID(ctx context.Context, cartItemsDomain *CartItemsDomain) (CartItemsDomain, error)
	CartItemsGetAllByCustomerID(ctx context.Context, customerId string) ([]CartItemsDomain, error)
	CartItemsGetAll(ctx context.Context) ([]CartItemsDomain, error)
	CartItemsDelete(ctx context.Context, id string) error

	// ItemTransactionsCreate(ctx context.Context, itemTransactionsDomain *ItemTransactionsDomain, id string) (ItemTransactionsDomain, error)
	ItemTransactionsCreate(ctx context.Context, customerId string) (ItemTransactionsDomain, error)
	ItemTransactionsGetAll(ctx context.Context) ([]ItemTransactionsDomain, error)

	// Carts
	// CartsGetByID(ctx context.Context, id string) (CartsDomain, error)
	// CartsCreate(ctx context.Context, cartDomain *CartsDomain) (CartsDomain, error)
	// CartsGetAll(ctx context.Context) ([]CartsDomain, error)
	// CartsDelete(ctx context.Context, id string) error
}
