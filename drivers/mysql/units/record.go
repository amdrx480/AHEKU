package units

import (
	// "backend-golang/businesses/stocks"
	// stockhistory "backend-golang/drivers/mysql/stock_history"
	// stockins "backend-golang/drivers/mysql/stock_ins"

	"backend-golang/businesses/units"
	"time"

	"gorm.io/gorm"
)

type Units struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	UnitsName string         `json:"units_name" gorm:"unique"`
}

func (rec *Units) ToDomain() units.Domain {
	return units.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		UnitsName: rec.UnitsName,
	}
}
func FromDomain(domain *units.Domain) *Units {
	return &Units{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		UnitsName: domain.UnitsName,
	}
}
