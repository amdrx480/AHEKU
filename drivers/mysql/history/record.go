package history

import (
	"backend-golang/businesses/history"
	// "courses-api/drivers/mysql/stocks"
	"time"

	"gorm.io/gorm"
)

type History struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CustomerID   uint           `json:"customer_id"`
	CustomerName string         `json:"customer_name"`
	StockID      uint           `json:"stock_id"`
	StockName    string         `json:"stock_name"`
	Quantity     int            `json:"quantity"`
	TotalPrice   int            `json:"total_price"`
}

func (rec *History) ToDomain() history.Domain {
	return history.Domain{
		ID:           rec.ID,
		CreatedAt:    rec.CreatedAt,
		DeletedAt:    rec.DeletedAt,
		CustomerID:   rec.CustomerID,
		CustomerName: rec.CustomerName,
		StockID:      rec.StockID,
		StockName:    rec.CustomerName,
		Quantity:     rec.Quantity,
		TotalPrice:   rec.TotalPrice,
	}
}

func FromDomain(domain *history.Domain) *History {
	return &History{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		DeletedAt:    domain.DeletedAt,
		CustomerID:   domain.CustomerID,
		CustomerName: domain.CustomerName,
		StockID:      domain.StockID,
		StockName:    domain.CustomerName,
		Quantity:     domain.Quantity,
		TotalPrice:   domain.TotalPrice,
	}
}
