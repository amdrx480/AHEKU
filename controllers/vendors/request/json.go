package request

import (
	"backend-golang/businesses/vendors"

	"github.com/go-playground/validator/v10"
)

type Vendors struct {
	Vendor_Name    string `json:"vendor_name"`
	Vendor_Address string `json:"vendor_address"`
	Vendor_Email   string `json:"vendor_email"`
	Vendor_Phone   string `json:"vendor_phone"`
}

func (req *Vendors) ToDomain() *vendors.Domain {
	return &vendors.Domain{
		Vendor_Name:    req.Vendor_Name,
		Vendor_Address: req.Vendor_Address,
		Vendor_Email:   req.Vendor_Email,
		Vendor_Phone:   req.Vendor_Phone,
	}
}

func (req *Vendors) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
