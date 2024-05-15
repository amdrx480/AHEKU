package request

import (
	"backend-golang/businesses/sales"

	"github.com/go-playground/validator/v10"
)

type Sales struct {
	// VendorID uint `json:"vendor_id"`
	StockID  uint `json:"stock_id"`
	Quantity int  `json:"quantity"`
	// Selling_Price int  `json:"selling_price"`
}

func (req *Sales) ToDomain() *sales.Domain {
	return &sales.Domain{
		// VendorID: req.VendorID,
		StockID:  req.StockID,
		Quantity: req.Quantity,
		// Selling_Price: req.Selling_Price,
	}
}

func (req *Sales) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
