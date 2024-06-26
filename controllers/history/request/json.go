package request

import (
	"backend-golang/businesses/history"

	"github.com/go-playground/validator/v10"
)

type History struct {
	CustomerID uint `json:"customer_id" `
	StockID    uint `json:"stock_id" `
	Quantity   int  `json:"quantity" `
	TotalPrice int  `json:"total_price" `
}

func (req *History) ToDomain() *history.Domain {
	return &history.Domain{
		CustomerID: req.CustomerID,
		StockID:    req.StockID,
		Quantity:   req.Quantity,
		TotalPrice: req.TotalPrice,
	}
}

func (req *History) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
