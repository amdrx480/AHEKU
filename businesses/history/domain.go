package history

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	StockID    uint
	Quantity   int
	TotalPrice int
}
type Usecase interface {
	Create(ctx context.Context, categoryDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Create(ctx context.Context, vendorsDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}
