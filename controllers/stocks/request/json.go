package request

import (
	"backend-golang/businesses/stocks"

	"github.com/go-playground/validator/v10"
)

type Stock struct {
	Stock_Name    string `json:"stock_name" form:"stock_name"`
	Stock_Code    string `json:"stock_code" form:"stock_code"`
	CategoryName  string `json:"category_name" form:"category_name"`
	CategoryID    uint   `json:"category_id" form:"category_id"`
	UnitsID       uint   `json:"units_id" form:"units_id"`
	UnitsName     string `json:"units_name" form:"units_name"`
	Description   string `json:"description" form:"description"`
	Image_Path    string `json:"image_path" form:"image_path"`
	Stock_Total   int    `json:"stock_total" form:"stock_total"`
	Selling_Price int    `json:"selling_price" form:"selling_price"`
}

func (req *Stock) ToDomain() *stocks.Domain {
	return &stocks.Domain{
		Stock_Name:    req.Stock_Name,
		Stock_Code:    req.Stock_Code,
		CategoryID:    req.CategoryID,
		CategoryName:  req.CategoryName,
		UnitsID:       req.UnitsID,
		Description:   req.Description,
		Image_Path:    req.Image_Path,
		Stock_Total:   req.Stock_Total,
		Selling_Price: req.Selling_Price,
	}
}

func (req *Stock) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
