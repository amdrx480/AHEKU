package stocks

import (
	"backend-golang/businesses/stocks"
	// stockhistory "backend-golang/drivers/mysql/stock_history"
	// stockins "backend-golang/drivers/mysql/stock_ins"

	"time"

	"gorm.io/gorm"
)

// ID             uint
// CreatedAt      time.Time
// UpdatedAt      time.Time
// DeletedAt      gorm.DeletedAt
// Stock_Location string
// Stock_Code     string
// Stock_Category string
// // Stock_QRCode   string
// Stock_Name string
// // Stock_Unit     string
// Stock_Pcs   string
// Stock_Pack  string
// Stock_Roll  string
// Stock_Meter string
// // Stock_Total int

type Stock struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Stock_Location string         `json:"stock_location"`
	Stock_Code     string         `json:"stock_code"`
	Stock_Category string         `json:"stock_category"`
	// Stock_QRCode   string         `json:"stock_qrcode"`
	Stock_Name  string `json:"stock_name"`
	Stock_Pcs   int    `json:"stock_pcs"`
	Stock_Pack  int    `json:"stock_pack"`
	Stock_Roll  int    `json:"stock_roll"`
	Stock_Meter int    `json:"stock_meter"`
	// Stock_Total    int            `json:"stock_total"`
}

func (rec *Stock) ToDomain() stocks.Domain {
	return stocks.Domain{
		ID:             rec.ID,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
		DeletedAt:      rec.DeletedAt,
		Stock_Code:     rec.Stock_Code,
		Stock_Location: rec.Stock_Location,
		Stock_Category: rec.Stock_Category,
		Stock_Name:     rec.Stock_Name,
		Stock_Pcs:      rec.Stock_Pcs,
		Stock_Pack:     rec.Stock_Pack,
		Stock_Roll:     rec.Stock_Roll,
		Stock_Meter:    rec.Stock_Meter,
	}
}
func FromDomain(domain *stocks.Domain) *Stock {
	return &Stock{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		Stock_Location: domain.Stock_Location,
		Stock_Code:     domain.Stock_Code,
		Stock_Category: domain.Stock_Category,
		Stock_Name:     domain.Stock_Name,
		Stock_Pcs:      domain.Stock_Pcs,
		Stock_Pack:     domain.Stock_Pack,
		Stock_Roll:     domain.Stock_Roll,
		Stock_Meter:    domain.Stock_Meter,
	}
}
