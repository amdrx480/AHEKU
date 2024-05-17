package customers

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID               uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	Customer_Name    string
	Customer_Address string
	Customer_Email   string
	Customer_Phone   string
}
type Usecase interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, customersDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	// Delete(ctx context.Context, id string) error
}

type Repository interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, customersDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	// Delete(ctx context.Context, id string) error
}
