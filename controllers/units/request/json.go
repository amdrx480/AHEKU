package request

import (
	"backend-golang/businesses/units"

	"github.com/go-playground/validator/v10"
)

type Units struct {
	UnitsName string `json:"units_name"`
}

func (req *Units) ToDomain() *units.Domain {
	return &units.Domain{
		UnitsName: req.UnitsName,
	}
}

func (req *Units) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
