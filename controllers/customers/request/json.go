package request

import (
	"backend-golang/businesses/customers"

	"github.com/go-playground/validator/v10"
)

type Customers struct {
	Customer_Name    string `json:"customer_name"`
	Customer_Address string `json:"customer_address"`
	Customer_Email   string `json:"customer_email"`
	Customer_Phone   string `json:"customer_phone"`
}

func (req *Customers) ToDomain() *customers.Domain {
	return &customers.Domain{
		Customer_Name:    req.Customer_Name,
		Customer_Address: req.Customer_Address,
		Customer_Email:   req.Customer_Email,
		Customer_Phone:   req.Customer_Phone,
	}
}

func (req *Customers) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
