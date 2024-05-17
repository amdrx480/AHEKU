package customers

import (
	// "backend-golang/businesses/stocks"
	// stockhistory "backend-golang/drivers/mysql/stock_history"
	// stockins "backend-golang/drivers/mysql/stock_ins"

	"backend-golang/businesses/customers"
	"time"

	"gorm.io/gorm"
)

type Customers struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Customer_Name    string         `json:"customer_name"`
	Customer_Address string         `json:"customer_address"`
	Customer_Email   string         `json:"customer_email"`
	Customer_Phone   string         `json:"customer_phone"`
}

func (rec *Customers) ToDomain() customers.Domain {
	return customers.Domain{
		ID:               rec.ID,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt,
		DeletedAt:        rec.DeletedAt,
		Customer_Name:    rec.Customer_Name,
		Customer_Address: rec.Customer_Address,
		Customer_Email:   rec.Customer_Email,
		Customer_Phone:   rec.Customer_Phone,
	}
}
func FromDomain(domain *customers.Domain) *Customers {
	return &Customers{
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
