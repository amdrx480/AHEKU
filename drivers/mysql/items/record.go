package items

import (
	"backend-golang/businesses/items"
	"backend-golang/drivers/mysql/customers"
	"backend-golang/drivers/mysql/stocks"

	"time"

	"gorm.io/gorm"
)

type Items struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Cart      Cart           `gorm:"foreignKey:cart_id"` // Menentukan foreign key
	CartID    uint           `json:"cart_id"`
	Stock     stocks.Stock   `json:"-" gorm:"foreignKey:stock_id"`
	StockID   uint           `json:"stock_id"`
	StockName string         `json:"stock_name"`
	Quantity  int            `json:"quantity"`
	Price     int            `json:"price"`
}

type Cart struct {
	ID         uint                `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	DeletedAt  gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
	Customer   customers.Customers `gorm:"foreignKey:customer_id"`
	CustomerID uint                `json:"customer_id"`
	Items      []Items             `gorm:"foreignKey:cart_id"` // Menentukan foreign key
	Total      int                 `json:"total"`
}

func (rec *Items) ToDomain() items.Domain {
	return items.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		CartID:    rec.CartID,
		StockID:   rec.StockID,
		StockName: rec.Stock.Stock_Name,
		Quantity:  rec.Quantity,
		Price:     rec.Price,
	}
}
func FromDomain(domain *items.Domain) *Items {
	return &Items{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		CartID:    domain.CartID,
		StockID:   domain.StockID,
		StockName: domain.StockName,
		Quantity:  domain.Quantity,
		Price:     domain.Price,
	}
}

func (record *Cart) ToDomainCart() items.DomainCart {
	itemsDomain := []items.Domain{}
	for _, item := range record.Items {
		itemsDomain = append(itemsDomain, item.ToDomain())
	}
	return items.DomainCart{
		ID:         record.ID,
		CreatedAt:  record.CreatedAt,
		UpdatedAt:  record.UpdatedAt,
		DeletedAt:  record.DeletedAt,
		CustomerID: record.CustomerID,
		Items:      itemsDomain,
		Total:      record.Total,
	}
}

func FromDomainCart(domain *items.DomainCart) *Cart {
	itemsModel := []Items{}
	for _, item := range domain.Items {
		itemsModel = append(itemsModel, *FromDomain(&item))
	}

	return &Cart{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
		CustomerID: domain.CustomerID,
		Items:      itemsModel,
		Total:      domain.Total,
	}
}
