package request

import (
	"backend-golang/businesses/stocks"

	"github.com/go-playground/validator/v10"
)

type Stock struct {
	Stock_Name    string `json:"stock_name"`
	Stock_Code    string `json:"stock_code"`
	CategoryID    uint   `json:"category_id"`
	UnitsID       uint   `json:"units_id"`
	Stock_Total   int    `json:"stock_total"`
	Selling_Price int    `json:"selling_price"`
}

func (req *Stock) ToDomain() *stocks.Domain {
	return &stocks.Domain{
		Stock_Name:    req.Stock_Name,
		Stock_Code:    req.Stock_Code,
		CategoryID:    req.CategoryID,
		UnitsID:       req.UnitsID,
		Stock_Total:   req.Stock_Total,
		Selling_Price: req.Selling_Price,
	}
}

func (req *Stock) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
