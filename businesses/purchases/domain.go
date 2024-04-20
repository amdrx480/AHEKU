package purchases

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
	VendorID       uint
	Stock_Name     string
	Stock_Code     string
	CategoryName   string
	CategoryID     uint
	UnitsID        uint
	Quantity       int
	Purchase_Price int
	Selling_Price  int
	// StockID        uint
}
type Usecase interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, categoryDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)

	// StockIn(ctx context.Context, categoryDomain *Domain, id string) (Domain, error)
	// StockOut(ctx context.Context, categoryDomain *Domain, id string) (Domain, error)
}

type Repository interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, categoryDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)

	// DownloadBarcodeByID(ctx context.Context, id string) (Domain, error)
	// StockIn(ctx context.Context, categoryDomain *Domain, id string) (Domain, error)
	// StockOut(ctx context.Context, categoryDomain *Domain, id string) (Domain, error)

	// Register(ctx context.Context, userDomain *Domain) (Domain, error)
	// GetByEmail(ctx context.Context, userDomain *Domain) (Domain, error)
	// DeleteCustomer(ctx context.Context, id string) (error)
}
