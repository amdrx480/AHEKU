package items

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	CartID        uint
	StockID       uint
	StockName     string
	Quantity      int
	Selling_Price int
	Price         int
}

type DomainCart struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	CustomerID uint
	Items      []Domain
	Total      int
}

type Usecase interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, itemsDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Delete(ctx context.Context, id string) error

	GetByIDCart(ctx context.Context, id string) (DomainCart, error)
	CreateCart(ctx context.Context, cartDomain *DomainCart) (DomainCart, error)
	GetAllCart(ctx context.Context) ([]DomainCart, error)
	DeleteCart(ctx context.Context, id string) error
}

type Repository interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, itemsDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Delete(ctx context.Context, id string) error

	GetByIDCart(ctx context.Context, id string) (DomainCart, error)
	CreateCart(ctx context.Context, cartDomain *DomainCart) (DomainCart, error)
	GetAllCart(ctx context.Context) ([]DomainCart, error)
	DeleteCart(ctx context.Context, id string) error
}

// ToHistory(ctx context.Context, historyDomain *history.Domain, id string) (history.Domain, error)
// ToHistory(ctx context.Context, historyDomain *history.Domain, id string) (history.Domain, error)
