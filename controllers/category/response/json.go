package response

import (
	"backend-golang/businesses/category"

	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	CategoryName string         `json:"category_name"`
}

func FromDomain(domain category.Domain) Category {
	return Category{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		CategoryName: domain.CategoryName,
	}
}
