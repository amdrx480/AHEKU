package purchases

import (
	// "backend-golang/businesses/stocks"
	// stockhistory "backend-golang/drivers/mysql/stock_history"
	// stockins "backend-golang/drivers/mysql/stock_ins"

	"backend-golang/businesses/purchases"
	"backend-golang/drivers/mysql/category"
	"backend-golang/drivers/mysql/units"
	"backend-golang/drivers/mysql/vendors"

	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	ID         uint            `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	DeletedAt  gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
	Vendor     vendors.Vendors `json:"-" gorm:"foreignKey:vendor_id"`
	VendorID   uint            `json:"vendor_id"`
	Stock_Name string          `json:"stock_name"`
	Stock_Code string          `json:"stock_code"`

	Category   category.Category `json:"-" gorm:"foreignKey:CategoryID"`
	CategoryID uint              `json:"category_id"`
	// CategoryName string            `json:"category_name"`

	Units     units.Units `json:"-" gorm:"foreignKey:UnitsID"`
	UnitsID   uint        `json:"units_id"`
	UnitsName string      `json:"units_name"`

	Quantity       int `json:"quantity"`
	Purchase_Price int `json:"purchase_price"`
	Selling_Price  int `json:"selling_price"`
}

func (rec *Purchase) ToDomain() purchases.Domain {
	return purchases.Domain{
		ID:         rec.ID,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
		DeletedAt:  rec.DeletedAt,
		VendorID:   rec.VendorID,
		Stock_Name: rec.Stock_Name,
		Stock_Code: rec.Stock_Code,

		CategoryID: rec.CategoryID,

		UnitsID:        rec.UnitsID,
		Quantity:       rec.Quantity,
		Purchase_Price: rec.Purchase_Price,
		Selling_Price:  rec.Selling_Price,

		// Tambahkan nama entitas terkait ke domain
		VendorName: rec.Vendor.Vendor_Name, // Nama vendor
		// CategoryName: rec.Category.CategoryName, // Nama kategori
		UnitsName: rec.Units.UnitsName, // Nama unit
	}
}
func FromDomain(domain *purchases.Domain) *Purchase {
	return &Purchase{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
		VendorID:   domain.VendorID,
		Stock_Name: domain.Stock_Name,
		Stock_Code: domain.Stock_Code,
		CategoryID: domain.CategoryID,
		// CategoryName:   domain.CategoryName,
		UnitsID:        domain.UnitsID,
		UnitsName:      domain.UnitsName,
		Quantity:       domain.Quantity,
		Purchase_Price: domain.Purchase_Price,
		Selling_Price:  domain.Selling_Price,
	}
}
