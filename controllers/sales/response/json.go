package response

import (
	"backend-golang/businesses/sales"

	"time"

	"gorm.io/gorm"
)

type Sales struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	VendorName string         `json:"vendor_name"`
	VendorID   uint           `json:"vendor_id"`
	StockName  string         `json:"stock_name"`
	StockID    uint           `json:"stock_id"`
	Quantity   int            `json:"quantity"`
	// Selling_Price int            `json:"selling_price"`
	Total_Price int `json:"total_price"`
}

func FromDomain(domain sales.Domain) Sales {
	return Sales{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
		VendorName: domain.VendorName,
		VendorID:   domain.VendorID,
		StockName:  domain.StockName,
		StockID:    domain.StockID,
		Quantity:   domain.Quantity,
		// Selling_Price: domain.Selling_Price,
		Total_Price: domain.Total_Price,
	}
}
