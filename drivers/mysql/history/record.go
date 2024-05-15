package history

import (
	"backend-golang/businesses/history"
	// "courses-api/drivers/mysql/stocks"
	"time"

	"gorm.io/gorm"
)

type History struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	StockID    uint           `json:"stock_id"`
	Quantity   int            `json:"quantity"`
	TotalPrice int            `json:"total_price"`
}

func (rec *History) ToDomain() history.Domain {
	return history.Domain{
		ID:         rec.ID,
		CreatedAt:  rec.CreatedAt,
		DeletedAt:  rec.DeletedAt,
		StockID:    rec.StockID,
		Quantity:   rec.Quantity,
		TotalPrice: rec.TotalPrice,
	}
}

func FromDomain(domain *history.Domain) *History {
	return &History{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		DeletedAt:  domain.DeletedAt,
		StockID:    domain.StockID,
		Quantity:   domain.Quantity,
		TotalPrice: domain.TotalPrice,
	}
}
