package request

import (
	"backend-golang/businesses/purchases"

	"github.com/go-playground/validator/v10"
)

type Purchases struct {
	VendorID       uint   `json:"vendor_id"`
	Stock_Name     string `json:"stock_name"`
	Stock_Code     string `json:"stock_code"`
	CategoryName   string `json:"category_name"`
	CategoryID     uint   `json:"category_id"`
	UnitsID        uint   `json:"units_id"`
	UnitsName      string `json:"units_name"`
	Quantity       int    `json:"quantity"`
	Purchase_Price int    `json:"purchase_price"`
	Selling_Price  int    `json:"selling_price"`
}

func (req *Purchases) ToDomain() *purchases.Domain {
	return &purchases.Domain{
		VendorID:       req.VendorID,
		Stock_Name:     req.Stock_Name,
		Stock_Code:     req.Stock_Code,
		CategoryID:     req.CategoryID,
		CategoryName:   req.CategoryName,
		UnitsID:        req.UnitsID,
		UnitsName:      req.UnitsName,
		Quantity:       req.Quantity,
		Purchase_Price: req.Purchase_Price,
		Selling_Price:  req.Selling_Price,
	}
}

func (req *Purchases) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
