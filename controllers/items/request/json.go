package request

import (
	"backend-golang/businesses/items"

	"github.com/go-playground/validator/v10"
)

type Items struct {
	CartID   uint `json:"cart_id" validate:"required"`
	StockID  uint `json:"stock_id" validate:"required"`
	Quantity int  `json:"quantity" validate:"required"`
}

type Cart struct {
	CustomerID uint `json:"customer_id" validate:"required"`
}

func (req *Items) ToDomain() *items.Domain {
	return &items.Domain{
		CartID:   req.CartID,
		StockID:  req.StockID,
		Quantity: req.Quantity,
	}
}

func (req *Cart) ToDomainCart() *items.DomainCart {
	return &items.DomainCart{
		CustomerID: req.CustomerID,
	}
}

func (req *Items) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *Cart) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
