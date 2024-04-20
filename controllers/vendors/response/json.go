package response

import (
	"backend-golang/businesses/vendors"

	"time"

	"gorm.io/gorm"
)

type Vendors struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	Vendor_Name    string         `json:"vendor_name"`
	Vendor_Address string         `json:"vendor_address"`
	Vendor_Email   string         `json:"vendor_email"`
	Vendor_Phone   string         `json:"vendor_phone"`
}

func FromDomain(domain vendors.Domain) Vendors {
	return Vendors{
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
