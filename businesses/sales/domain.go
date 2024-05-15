package sales

import (
	"backend-golang/businesses/history"

	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	// VendorName    string
	// VendorID      uint
	StockName     string
	StockID       uint
	Quantity      int
	Selling_Price int
	TotalPrice    int
	TotalAllPrice int
}
type Usecase interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, salesDomain *Domain) (Domain, error)
	ToHistory(ctx context.Context, historyDomain *history.Domain, id string) (history.Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Delete(ctx context.Context, id string) error

	// StockIn(ctx context.Context, salesDomain *Domain, id string) (Domain, error)
	// StockOut(ctx context.Context, salesDomain *Domain, id string) (Domain, error)
}

type Repository interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, salesDomain *Domain) (Domain, error)
	ToHistory(ctx context.Context, historyDomain *history.Domain, id string) (history.Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Delete(ctx context.Context, id string) error
	// DownloadBarcodeByID(ctx context.Context, id string) (Domain, error)
	// StockIn(ctx context.Context, salesDomain *Domain, id string) (Domain, error)
	// StockOut(ctx context.Context, salesDomain *Domain, id string) (Domain, error)

	// Register(ctx context.Context, userDomain *Domain) (Domain, error)
	// GetByEmail(ctx context.Context, userDomain *Domain) (Domain, error)
}
