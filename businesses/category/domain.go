package category

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	CategoryName string
}
type Usecase interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	GetByName(ctx context.Context, name string) (Domain, error)

	Create(ctx context.Context, categoryDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	// Delete(ctx context.Context, id string) error
}

type Repository interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	GetByName(ctx context.Context, name string) (Domain, error)

	Create(ctx context.Context, vendorsDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	// Delete(ctx context.Context, id string) error
}
