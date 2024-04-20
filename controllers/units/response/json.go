package response

import (
	"backend-golang/businesses/units"

	"time"

	"gorm.io/gorm"
)

type Units struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Units     string         `json:"units"`
}

func FromDomain(domain units.Domain) Units {
	return Units{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Units:     domain.Units,
	}
}
