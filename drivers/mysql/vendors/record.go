package vendors

import (
	// "backend-golang/businesses/stocks"
	// stockhistory "backend-golang/drivers/mysql/stock_history"
	// stockins "backend-golang/drivers/mysql/stock_ins"

	"backend-golang/businesses/vendors"
	"time"

	"gorm.io/gorm"
)

type Vendors struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Vendor_Name    string         `json:"vendor_name"`
	Vendor_Address string         `json:"vendor_address"`
	Vendor_Email   string         `json:"vendor_email"`
	Vendor_Phone   string         `json:"vendor_phone"`
}

func (rec *Vendors) ToDomain() vendors.Domain {
	return vendors.Domain{
		ID:             rec.ID,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
		DeletedAt:      rec.DeletedAt,
		Vendor_Name:    rec.Vendor_Name,
		Vendor_Address: rec.Vendor_Address,
		Vendor_Email:   rec.Vendor_Email,
		Vendor_Phone:   rec.Vendor_Phone,
	}
}
func FromDomain(domain *vendors.Domain) *Vendors {
	return &Vendors{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		Vendor_Name:    domain.Vendor_Name,
		Vendor_Address: domain.Vendor_Address,
		Vendor_Email:   domain.Vendor_Email,
		Vendor_Phone:   domain.Vendor_Phone,
	}
}
