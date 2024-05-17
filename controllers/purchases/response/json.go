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
	VendorID       uint           `json:"vendor_id"`
	VendorName     string         `json:"vendor_name"`
	Stock_Name     string         `json:"stock_name"`
	Stock_Code     string         `json:"stock_code"`
	CategoryID     uint           `json:"category_id"`
	CategoryName   string         `json:"category_name"`
	UnitsID        uint           `json:"units_id"`
	UnitsName      string         `json:"units_name"`
	Quantity       int            `json:"quantity"`
	Description    string         `json:"description"`
	Purchase_Price int            `json:"purchase_price"`
	Selling_Price  int            `json:"selling_price"`
}

func FromDomain(domain purchases.Domain) Purchases {
	return Purchases{
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
		Quantity:       domain.Quantity,
		Description:    domain.Description,
		Purchase_Price: domain.Purchase_Price,
		Selling_Price:  domain.Selling_Price,
	}
}
