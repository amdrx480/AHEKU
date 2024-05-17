package response

import (
	"backend-golang/businesses/history"
	"time"

	"gorm.io/gorm"
)

type History struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	CustomerID   uint           `json:"customer_id"`
	CustomerName string         `json:"customer_name"`
	StockID      uint           `json:"stock_id" `
	Quantity     int            `json:"quantity" `
	TotalPrice   int            `json:"total_price"`
}

func FromDomain(domain history.Domain) History {
	return History{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		DeletedAt:    domain.DeletedAt,
		CustomerID:   domain.CustomerID,
		CustomerName: domain.CustomerName,
		StockID:      domain.StockID,
		Quantity:     domain.Quantity,
		TotalPrice:   domain.TotalPrice,
	}
}
