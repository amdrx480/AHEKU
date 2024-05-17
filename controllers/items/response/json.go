package response

import (
	// "backend-golang/businesses/cart"
	"backend-golang/businesses/items"

	"time"

	"gorm.io/gorm"
)

type Items struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	CartID        uint           `json:"cart_id"`
	StockID       uint           `json:"stock_id"`
	StockName     string         `json:"stock_name"`
	Quantity      int            `json:"quantity"`
	Selling_Price int            `json:"selling_price"`
	Price         int            `json:"price"`
}

type Cart struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	CustomerID uint           `json:"customer_id"`
	Items      []Items        `json:"items"`
	Total      int            `json:"total"`
}

func FromDomain(domain items.Domain) Items {
	return Items{
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

// Fungsi untuk mengubah dari domain ke response struct Cart
func FromDomainCart(domain items.DomainCart) Cart {
	items := []Items{}
	for _, item := range domain.Items {
		items = append(items, FromDomain(item))
	}
	return Cart{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
		CustomerID: domain.CustomerID,
		Items:      items,
		Total:      domain.Total,
	}
}
