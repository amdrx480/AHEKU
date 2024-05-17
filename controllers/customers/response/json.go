package response

import (
	"backend-golang/businesses/customers"

	"time"

	"gorm.io/gorm"
)

type Customers struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
	Customer_Name    string         `json:"customer_name"`
	Customer_Address string         `json:"customer_address"`
	Customer_Email   string         `json:"customer_email"`
	Customer_Phone   string         `json:"customer_phone"`
}

func FromDomain(domain customers.Domain) Customers {
	return Customers{
		ID:               domain.ID,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
		DeletedAt:        domain.DeletedAt,
		Customer_Name:    domain.Customer_Name,
		Customer_Address: domain.Customer_Address,
		Customer_Email:   domain.Customer_Email,
		Customer_Phone:   domain.Customer_Phone,
	}
}
