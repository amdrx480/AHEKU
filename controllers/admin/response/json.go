package response

import (
	"backend-golang/businesses/admin"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `json:"name"`
	Voucher   string         `json:"voucher"`
	Password  string         `json:"password"`
}

type Customers struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
	CustomerName    string         `json:"customer_name"`
	CustomerAddress string         `json:"customer_address"`
	CustomerEmail   string         `json:"customer_email"`
	CustomerPhone   string         `json:"customer_phone"`
	CartItems       []CartItems    `json:"cart_items"`
}

type Category struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	CategoryName string         `json:"category_name"`
}

type Vendors struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	VendorName    string         `json:"vendor_name"`
	VendorAddress string         `json:"vendor_address"`
	VendorEmail   string         `json:"vendor_email"`
	VendorPhone   string         `json:"vendor_phone"`
}

type Units struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	UnitName  string         `json:"units_name"`
}

type Stocks struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	StockName    string         `json:"stock_name"`
	StockCode    string         `json:"stock_code"`
	CategoryName string         `json:"category_name"`
	CategoryID   uint           `json:"category_id"`
	UnitID       uint           `json:"units_id"`
	UnitName     string         `json:"units_name"`
	Description  string         `json:"description"`
	StockTotal   int            `json:"stock_total"`
	SellingPrice int            `json:"selling_price"`
}

type Purchases struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	VendorID       uint           `json:"vendor_id"`
	StockName      string         `json:"stock_name"`
	StockCode      string         `json:"stock_code"`
	CategoryID     uint           `json:"category_id"`
	UnitID         uint           `json:"units_id"`
	Quantity       int            `json:"quantity"`
	Description    string         `json:"description"`
	Purchase_Price int            `json:"purchase_price"`
	SellingPrice   int            `json:"selling_price"`
}

type CartItems struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	// CartID        uint           `json:"cart_id"`
	CustomerID   uint   `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	StockID      uint   `json:"stock_id"`
	StockName    string `json:"stock_name"`
	// CategoryID    uint   `json:"category_id"`
	// CategoryName  string `json:"category_name"`
	Quantity      int `json:"quantity"`
	Selling_Price int `json:"selling_price"`
	Price         int `json:"price"`
	SubTotal      int `json:"sub_total"`
}

type ItemTransactions struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	StockID   uint           `json:"stock_id"`
	Quantity  int            `json:"quantity"`
	SubTotal  int            `json:"sub_total"`
}

func FromAdminsDomain(domain admin.AdminsDomain) Admin {
	return Admin{
		ID:        domain.ID,
		Name:      domain.Name,
		Voucher:   domain.Voucher,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func FromCustomersDomain(domain admin.CustomersDomain) Customers {
	cartItems := []CartItems{}
	for _, item := range domain.CartItems {
		cartItems = append(cartItems, FromCartItemsDomain(item))
	}
	return Customers{
		ID:              domain.ID,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
		DeletedAt:       domain.DeletedAt,
		CustomerName:    domain.CustomerName,
		CustomerAddress: domain.CustomerAddress,
		CustomerEmail:   domain.CustomerEmail,
		CustomerPhone:   domain.CustomerPhone,
		CartItems:       cartItems,
	}
}

func FromCategoryDomain(domain admin.CategoriesDomain) Category {
	return Category{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		CategoryName: domain.CategoryName,
	}
}

func FromVendorsDomain(domain admin.VendorsDomain) Vendors {
	return Vendors{
		ID:            domain.ID,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		DeletedAt:     domain.DeletedAt,
		VendorName:    domain.VendorName,
		VendorAddress: domain.VendorAddress,
		VendorEmail:   domain.VendorEmail,
		VendorPhone:   domain.VendorPhone,
	}
}

func FromUnitsDomain(domain admin.UnitsDomain) Units {
	return Units{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		UnitName:  domain.UnitName,
	}
}

func FromStocksDomain(domain admin.StocksDomain) Stocks {
	return Stocks{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		StockName:    domain.StockName,
		StockCode:    domain.StockCode,
		CategoryName: domain.CategoryName,
		CategoryID:   domain.CategoryID,
		UnitName:     domain.UnitName,
		UnitID:       domain.UnitID,
		Description:  domain.Description,
		StockTotal:   domain.StockTotal,
		SellingPrice: domain.SellingPrice,
	}
}

func FromPurchasesDomain(domain admin.PurchasesDomain) Purchases {
	return Purchases{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		VendorID:  domain.VendorID,
		// VendorName:     domain.VendorName,
		StockName:  domain.StockName,
		StockCode:  domain.StockCode,
		CategoryID: domain.CategoryID,
		// CategoryName:   domain.CategoryName,
		UnitID: domain.UnitID,
		// UnitName:       domain.UnitName,
		Quantity:       domain.Quantity,
		Description:    domain.Description,
		Purchase_Price: domain.PurchasesPrice,
		SellingPrice:   domain.SellingPrice,
	}
}

func FromCartItemsDomain(domain admin.CartItemsDomain) CartItems {
	return CartItems{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		CustomerID:   domain.CustomerID,
		CustomerName: domain.CustomerName,
		StockID:      domain.StockID,
		StockName:    domain.StockName,
		// CategoryID:   domain.CategoryID,
		// CategoryName: domain.CategoryName,
		Quantity: domain.Quantity,
		Price:    domain.Price,
		SubTotal: domain.SubTotal,
	}
}

func FromItemTransactionsDomain(domain admin.ItemTransactionsDomain) ItemTransactions {
	return ItemTransactions{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		DeletedAt: domain.DeletedAt,
		StockID:   domain.StockID,
		Quantity:  domain.Quantity,
		SubTotal:  domain.SubTotal,
	}
}
