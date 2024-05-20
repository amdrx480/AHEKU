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
	VendorName     string         `json:"vendor_name"`
	StockName      string         `json:"stock_name"`
	StockCode      string         `json:"stock_code"`
	CategoryID     uint           `json:"category_id"`
	CategoryName   string         `json:"category_name"`
	UnitID         uint           `json:"units_id"`
	UnitName       string         `json:"units_name"`
	Quantity       int            `json:"quantity"`
	Description    string         `json:"description"`
	Purchase_Price int            `json:"purchase_price"`
	SellingPrice   int            `json:"selling_price"`
}

type Items struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	CartID        uint           `json:"cart_id"`
	StockID       uint           `json:"stock_id"`
	StockName     string         `json:"stock_name"`
	Quantity      int            `json:"quantity"`
	Selling_Price int            `json:"selling_price"`
	Price         int            `json:"price"`
}

type Carts struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	CustomerID uint           `json:"customer_id"`
	Items      []Items        `json:"items"`
	Total      int            `json:"total"`
	Status     string         `json:"status"`
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
	return Customers{
		ID:              domain.ID,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
		DeletedAt:       domain.DeletedAt,
		CustomerName:    domain.CustomerName,
		CustomerAddress: domain.CustomerAddress,
		CustomerEmail:   domain.CustomerEmail,
		CustomerPhone:   domain.CustomerPhone,
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
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		VendorID:       domain.VendorID,
		VendorName:     domain.VendorName,
		StockName:      domain.StockName,
		StockCode:      domain.StockCode,
		CategoryID:     domain.CategoryID,
		CategoryName:   domain.CategoryName,
		UnitID:         domain.UnitID,
		UnitName:       domain.UnitName,
		Quantity:       domain.Quantity,
		Description:    domain.Description,
		Purchase_Price: domain.PurchasesPrice,
		SellingPrice:   domain.SellingPrice,
	}
}

func FromItemsDomain(domain admin.ItemsDomain) Items {
	return Items{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		CartID:    domain.CartID,
		StockID:   domain.StockID,
		StockName: domain.StockName,
		Quantity:  domain.Quantity,
		Price:     domain.Price,
	}
}

// Fungsi untuk mengubah dari domain ke response struct Carts
func FromCartsDomain(domain admin.CartsDomain) Carts {
	items := []Items{}
	for _, item := range domain.Items {
		items = append(items, FromItemsDomain(item))
	}
	return Carts{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
		CustomerID: domain.CustomerID,
		Items:      items,
		Total:      domain.Total,
		Status:     domain.Status,
	}
}