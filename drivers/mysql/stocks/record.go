package stocks

import (
	"backend-golang/businesses/stocks"
	"backend-golang/drivers/mysql/category"
	"backend-golang/drivers/mysql/units"

	// stockhistory "backend-golang/drivers/mysql/stock_history"
	// stockins "backend-golang/drivers/mysql/stock_ins"

	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID            uint              `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	DeletedAt     gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
	Stock_Name    string            `json:"stock_name"`
	Stock_Code    string            `json:"stock_code"`
	Category      category.Category `json:"-" gorm:"foreignKey:CategoryID"`
	CategoryName  string            `json:"category_name"`
	CategoryID    uint              `json:"category_id"`
	Units         units.Units       `json:"-" gorm:"foreignKey:UnitsID"`
	UnitsID       uint              `json:"units_id"`
	UnitsName     string            `json:"units_name"`
	Description   string            `json:"description"`
	Image_Path    string            `json:"image_path"`
	Stock_Total   int               `json:"stock_total"`
	Selling_Price int               `json:"selling_price"`
}

func (rec *Stock) ToDomain() stocks.Domain {
	return stocks.Domain{
		ID:            rec.ID,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
		DeletedAt:     rec.DeletedAt,
		Stock_Code:    rec.Stock_Code,
		Stock_Name:    rec.Stock_Name,
		CategoryID:    rec.CategoryID,
		CategoryName:  rec.Category.CategoryName,
		UnitsID:       rec.UnitsID,
		UnitsName:     rec.Units.UnitsName,
		Stock_Total:   rec.Stock_Total,
		Selling_Price: rec.Selling_Price,
		Description:   rec.Description,
	}
}
func FromDomain(domain *stocks.Domain) *Stock {
	return &Stock{
		ID:            domain.ID,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		DeletedAt:     domain.DeletedAt,
		Stock_Name:    domain.Stock_Name,
		Stock_Code:    domain.Stock_Code,
		CategoryID:    domain.CategoryID,
		CategoryName:  domain.CategoryName,
		UnitsID:       domain.UnitsID,
		UnitsName:     domain.UnitsName,
		Description:   domain.Description,
		Stock_Total:   domain.Stock_Total,
		Selling_Price: domain.Selling_Price,
	}
}
