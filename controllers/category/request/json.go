package request

import (
	"backend-golang/businesses/category"

	"github.com/go-playground/validator/v10"
)

type Category struct {
	CategoryName string `json:"category_name"`
}

func (req *Category) ToDomain() *category.Domain {
	return &category.Domain{
		CategoryName: req.CategoryName,
	}
}

func (req *Category) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
