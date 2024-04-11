package request

import (
	"backend-golang/businesses/stocks"

	"github.com/go-playground/validator/v10"
)

type Stock struct {
	Stock_Location string `json:"stock_location"`
	Stock_Code     string `json:"stock_code"`
	Stock_Category string `json:"stock_category"`
	// Stock_QRCode   string         `json:"stock_qrcode"`
	Stock_Name  string `json:"stock_name"`
	Stock_Pcs   int    `json:"stock_pcs"`
	Stock_Pack  int    `json:"stock_pack"`
	Stock_Roll  int    `json:"stock_roll"`
	Stock_Meter int    `json:"stock_meter"`
	// Stock_Total    int            `json:"stock_total"`
}

func (req *Stock) ToDomain() *stocks.Domain {
	return &stocks.Domain{
		Stock_Location: req.Stock_Location,
		Stock_Code:     req.Stock_Code,
		Stock_Category: req.Stock_Category,
		Stock_Name:     req.Stock_Name,
		Stock_Pcs:      req.Stock_Pcs,
		Stock_Pack:     req.Stock_Pack,
		Stock_Roll:     req.Stock_Roll,
		Stock_Meter:    req.Stock_Meter,
	}
}

func (req *Stock) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
