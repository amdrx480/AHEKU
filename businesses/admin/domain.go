package admin

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Admins
type AdminsDomain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
	Voucher   string
	Password  string
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
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	VendorID       uint
	VendorName     string
	StockName      string
	StockCode      string
	CategoryID     uint
	CategoryName   string
	UnitID         uint
	UnitName       string
	Quantity       int
	Description    string
	PurchasesPrice int
	SellingPrice   int
}

// Many Items To One Carts
type ItemsDomain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	CartID       uint
	StockID      uint
	StockName    string
	Quantity     int
	SellingPrice int
	Price        int
}

// One Carts To Many Items
type CartsDomain struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	CustomerID uint
	Items      []ItemsDomain
	Total      int
	Status     string
}

type Usecase interface {
	// Admin
	AdminRegister(ctx context.Context, adminDomain *AdminsDomain) (AdminsDomain, error)
	AdminLogin(ctx context.Context, adminDomain *AdminsDomain) (string, error)
	AdminVoucher(ctx context.Context, adminDomain *AdminsDomain) (string, error)

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
	StocksGetAll(ctx context.Context) ([]StocksDomain, error)

	// Purchases
	PurchasesGetByID(ctx context.Context, id string) (PurchasesDomain, error)
	PurchasesCreate(ctx context.Context, purchasesDomain *PurchasesDomain) (PurchasesDomain, error)
	PurchasesGetAll(ctx context.Context) ([]PurchasesDomain, error)

	// Items
	ItemsGetByID(ctx context.Context, id string) (ItemsDomain, error)
	ItemsCreate(ctx context.Context, itemsDomain *ItemsDomain) (ItemsDomain, error)
	ItemsGetAll(ctx context.Context) ([]ItemsDomain, error)
	ItemsDelete(ctx context.Context, id string) error

	// Carts
	CartsGetByID(ctx context.Context, id string) (CartsDomain, error)
	CartsCreate(ctx context.Context, cartDomain *CartsDomain) (CartsDomain, error)
	CartsGetAll(ctx context.Context) ([]CartsDomain, error)
	CartsDelete(ctx context.Context, id string) error
}

type Repository interface {
	// Admin
	AdminRegister(ctx context.Context, adminDomain *AdminsDomain) (AdminsDomain, error)
	AdminGetByName(ctx context.Context, adminDomain *AdminsDomain) (AdminsDomain, error)
	AdminGetByVoucher(ctx context.Context, adminDomain *AdminsDomain) (AdminsDomain, error)

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
	StocksGetAll(ctx context.Context) ([]StocksDomain, error)

	// Purchases
	PurchasesGetByID(ctx context.Context, id string) (PurchasesDomain, error)
	PurchasesCreate(ctx context.Context, purchasesDomain *PurchasesDomain) (PurchasesDomain, error)
	PurchasesGetAll(ctx context.Context) ([]PurchasesDomain, error)

	// Items
	ItemsGetByID(ctx context.Context, id string) (ItemsDomain, error)
	ItemsCreate(ctx context.Context, itemsDomain *ItemsDomain) (ItemsDomain, error)
	ItemsGetAll(ctx context.Context) ([]ItemsDomain, error)
	ItemsDelete(ctx context.Context, id string) error

	// Carts
	CartsGetByID(ctx context.Context, id string) (CartsDomain, error)
	CartsCreate(ctx context.Context, cartDomain *CartsDomain) (CartsDomain, error)
	CartsGetAll(ctx context.Context) ([]CartsDomain, error)
	CartsDelete(ctx context.Context, id string) error
}
