package sales

import (
	// "backend-golang/businesses/stocks"
	// stockhistory "backend-golang/drivers/mysql/stock_history"
	// stockins "backend-golang/drivers/mysql/stock_ins"

	"backend-golang/businesses/sales"
	"backend-golang/drivers/mysql/stocks"

	// "backend-golang/drivers/mysql/units"
	"backend-golang/drivers/mysql/vendors"

	"time"

	"gorm.io/gorm"
)

type Sales struct {
	ID          uint            `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
	Vendor      vendors.Vendors `json:"-" gorm:"foreignKey:vendor_id"`
	VendorID    uint            `json:"vendor_id"`
	Stock       stocks.Stock    `json:"-" gorm:"foreignKey:stock_id"`
	StockID     uint            `json:"stock_id"`
	Quantity    int             `json:"quantity"`
	Total_Price int             `json:"total_price"`
}

func (rec *Sales) ToDomain() sales.Domain {
	return sales.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		VendorID:  rec.VendorID,
		StockID:   rec.StockID,
		Quantity:  rec.Quantity,
		// Tambahkan nama entitas terkait ke domain
		Total_Price: rec.Total_Price,
		VendorName:  rec.Vendor.Vendor_Name, // Nama vendor
		StockName:   rec.Stock.Stock_Name,   // Nama Stock

	}
}
func FromDomain(domain *sales.Domain) *Sales {
	return &Sales{
		ID:          domain.ID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
		VendorID:    domain.VendorID,
		StockID:     domain.StockID,
		Quantity:    domain.Quantity,
		Total_Price: domain.Total_Price,
	}
}
