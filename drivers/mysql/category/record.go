package category

import (
	// "backend-golang/businesses/stocks"
	// stockhistory "backend-golang/drivers/mysql/stock_history"
	// stockins "backend-golang/drivers/mysql/stock_ins"

	"backend-golang/businesses/category"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CategoryName string         `json:"category_name" gorm:"unique"`
}

func (rec *Category) ToDomain() category.Domain {
	return category.Domain{
		ID:           rec.ID,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
		CategoryName: rec.CategoryName,
	}
}
func FromDomain(domain *category.Domain) *Category {
	return &Category{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		CategoryName: domain.CategoryName,
	}
}
