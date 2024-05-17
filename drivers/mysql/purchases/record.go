package purchases

import (
	// "backend-golang/businesses/stocks"
	// stockhistory "backend-golang/drivers/mysql/stock_history"
	// stockins "backend-golang/drivers/mysql/stock_ins"

	"backend-golang/businesses/purchases"
	// unitsDomain "backend-golang/businesses/units"

	"backend-golang/drivers/mysql/category"
	"backend-golang/drivers/mysql/units"
	"backend-golang/drivers/mysql/vendors"

	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	ID             uint              `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
	DeletedAt      gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
	Vendor         vendors.Vendors   `json:"-" gorm:"foreignKey:VendorID"`
	VendorID       uint              `json:"vendor_id"`
	VendorName     string            `json:"Vendor_name"`
	Stock_Name     string            `json:"stock_name"`
	Stock_Code     string            `json:"stock_code"`
	Category       category.Category `json:"-" gorm:"foreignKey:CategoryID"`
	CategoryID     uint              `json:"category_id"`
	CategoryName   string            `json:"category_name"`
	Units          units.Units       `gorm:"foreignKey:UnitsID"`
	UnitsID        uint              `json:"units_id"`
	UnitsName      string            `json:"units_name"`
	Description    string            `json:"description"`
	Quantity       int               `json:"quantity"`
	Purchase_Price int               `json:"purchase_price"`
	Selling_Price  int               `json:"selling_price"`
}

func (record *Purchase) ToDomain() purchases.Domain {
	return purchases.Domain{
		ID:             record.ID,
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
		DeletedAt:      record.DeletedAt,
		VendorID:       record.VendorID,
		VendorName:     record.Vendor.Vendor_Name,
		Stock_Name:     record.Stock_Name,
		Stock_Code:     record.Stock_Code,
		CategoryID:     record.CategoryID,
		CategoryName:   record.Category.CategoryName,
		UnitsID:        record.UnitsID,
		UnitsName:      record.Units.UnitsName,
		Description:    record.Description,
		Quantity:       record.Quantity,
		Purchase_Price: record.Purchase_Price,
		Selling_Price:  record.Selling_Price,
	}
}
func FromDomain(domain *purchases.Domain) *Purchase {
	return &Purchase{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		VendorID:       domain.VendorID,
		VendorName:     domain.VendorName,
		Stock_Name:     domain.Stock_Name,
		Stock_Code:     domain.Stock_Code,
		CategoryID:     domain.CategoryID,
		CategoryName:   domain.CategoryName,
		UnitsID:        domain.UnitsID,
		UnitsName:      domain.UnitsName,
		Description:    domain.Description,
		Quantity:       domain.Quantity,
		Purchase_Price: domain.Purchase_Price,
		Selling_Price:  domain.Selling_Price,
	}
}

// Tambahkan nama entitas terkait ke domain
// VendorName: record.Vendor.Vendor_Name, // Nama vendor
// CategoryName: record.Category.CategoryName, // Nama kategori
// UnitsName: record.Units.UnitsName, // Nama unit
