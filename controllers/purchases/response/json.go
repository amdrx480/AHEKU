package response

import (
	"backend-golang/businesses/purchases"

	"time"

	"gorm.io/gorm"
)

type Purchases struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	VendorName     string         `json:"vendor_name"`
	VendorID       uint           `json:"vendor_id"`
	Stock_Name     string         `json:"stock_name"`
	Stock_Code     string         `json:"stock_code"`
	CategoryName   string         `json:"category_name"`
	CategoryID     uint           `json:"category_id"`
	UnitsName      string         `json:"units_name"`
	UnitsID        uint           `json:"units"`
	Quantity       int            `json:"quantity"`
	Purchase_Price int            `json:"purchase_price"`
	Selling_Price  int            `json:"selling_price"`
}

func FromDomain(domain purchases.Domain) Purchases {
	return Purchases{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		VendorName:     domain.VendorName,
		VendorID:       domain.VendorID,
		Stock_Name:     domain.Stock_Name,
		Stock_Code:     domain.Stock_Code,
		CategoryName:   domain.CategoryName,
		CategoryID:     domain.CategoryID,
		UnitsName:      domain.UnitsName,
		UnitsID:        domain.UnitsID,
		Quantity:       domain.Quantity,
		Purchase_Price: domain.Purchase_Price,
		Selling_Price:  domain.Selling_Price,
	}
}
