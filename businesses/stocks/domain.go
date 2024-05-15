package stocks

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
	Stock_Name   string
	Stock_Code   string
	CategoryName string
	CategoryID   uint
	UnitsID      uint
	UnitsName    string
	Description  string

	Image_Path string

	Stock_Total   int
	Selling_Price int
}
type Usecase interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	// Create(ctx context.Context, stocksDomain *Domain, imagePath string, id string) (Domain, string, error)
	Create(ctx context.Context, stocksDomain *Domain) (Domain, error)
	// UploadStocksImage(ctx context.Context, stocksDomain *Domain, imagePath string, id string) (Domain, string, error)
	GetAll(ctx context.Context) ([]Domain, error)

	// DownloadBarcodeByID(ctx context.Context, id string) (Domain, error)
	// StockIn(ctx context.Context, categoryDomain *Domain, id string) (Domain, error)
	// StockOut(ctx context.Context, categoryDomain *Domain, id string) (Domain, error)

	// Register(ctx context.Context, userDomain *Domain) (Domain, error)
	// Login(ctx context.Context, userDomain *Domain) (string, error)
	// DeleteUser(ctx context.Context, id string) (error)
}

type Repository interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, categoryDomain *Domain) (Domain, error)
	// Create(ctx context.Context, categoryDomain *Domain, imagePath string, id string) (Domain, string, error)

	// UploadStocksImage(ctx context.Context, stocksDomain *Domain, imagePath string, id string) (Domain, string, error)
	GetAll(ctx context.Context) ([]Domain, error)

	// DownloadBarcodeByID(ctx context.Context, id string) (Domain, error)
	// StockIn(ctx context.Context, categoryDomain *Domain, id string) (Domain, error)
	// StockOut(ctx context.Context, categoryDomain *Domain, id string) (Domain, error)

	// Register(ctx context.Context, userDomain *Domain) (Domain, error)
	// GetByEmail(ctx context.Context, userDomain *Domain) (Domain, error)
	// DeleteCustomer(ctx context.Context, id string) (error)
}
