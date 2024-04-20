package request

import (
	"backend-golang/businesses/units"

	"github.com/go-playground/validator/v10"
)

type Units struct {
	Units string `json:"units"`
}

func (req *Units) ToDomain() *units.Domain {
	return &units.Domain{
		Units: req.Units,
	}
}

func (req *Units) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
