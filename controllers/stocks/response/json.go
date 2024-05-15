package response

import (
	"backend-golang/businesses/stocks"

	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	Stock_Name    string         `json:"stock_name"`
	Stock_Code    string         `json:"stock_code"`
	CategoryName  string         `json:"category_name"`
	CategoryID    uint           `json:"category_id"`
	UnitsID       uint           `json:"units_id"`
	UnitsName     string         `json:"units_name"`
	Description   string         `json:"description"`
	Stock_Total   int            `json:"stock_total"`
	Selling_Price int            `json:"selling_price"`
}

func FromDomain(domain stocks.Domain) Stock {
	return Stock{
		ID:            domain.ID,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		DeletedAt:     domain.DeletedAt,
		Stock_Name:    domain.Stock_Name,
		Stock_Code:    domain.Stock_Code,
		CategoryName:  domain.CategoryName,
		CategoryID:    domain.CategoryID,
		UnitsName:     domain.UnitsName,
		UnitsID:       domain.UnitsID,
		Description:   domain.Description,
		Stock_Total:   domain.Stock_Total,
		Selling_Price: domain.Selling_Price,
	}
}
