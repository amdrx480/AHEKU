package response

import (
	"backend-golang/businesses/history"
	"time"

	"gorm.io/gorm"
)

type History struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	StockID    uint           `json:"stock_id" `
	Quantity   int            `json:"quantity" `
	TotalPrice int            `json:"total_price" `
}

func FromDomain(domain history.Domain) History {
	return History{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		DeletedAt:  domain.DeletedAt,
		StockID:    domain.StockID,
		Quantity:   domain.Quantity,
		TotalPrice: domain.TotalPrice,
	}
}
