package vendors

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Vendor_Name    string
	Vendor_Address string
	Vendor_Email   string
	Vendor_Phone   string
}
type Usecase interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, vendorsDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	// Delete(ctx context.Context, id string) error
}

type Repository interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, vendorsDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	// Delete(ctx context.Context, id string) error
}
