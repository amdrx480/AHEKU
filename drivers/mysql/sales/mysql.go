package sales

import (
	"backend-golang/businesses/sales"
	_dbStocks "backend-golang/drivers/mysql/stocks"
	"context"
	"fmt"

	// "errors"

	// "fmt"

	"gorm.io/gorm"
)

type salesRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) sales.Repository {
	return &salesRepository{
		conn: conn,
	}
}

func (sr *salesRepository) GetByID(ctx context.Context, id string) (sales.Domain, error) {
	var sale Sales

	if err := sr.conn.WithContext(ctx).First(&sale, "id = ?", id).Error; err != nil {
		return sales.Domain{}, err
	}

	return sale.ToDomain(), nil

}

func (pr *salesRepository) Create(ctx context.Context, saleDomain *sales.Domain) (sales.Domain, error) {
	record := FromDomain(saleDomain)
	result := pr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return sales.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return sales.Domain{}, err
	}

	// Tambahkan atau perbarui stok terkait setelah membuat Purchase

	var stock _dbStocks.Stock
	err := pr.conn.WithContext(ctx).
		Where("id = ?", record.StockID).
		First(&stock).Error

	if err == gorm.ErrRecordNotFound {
		return sales.Domain{}, fmt.Errorf("stock not found with stock_id %d", record.StockID)
	} else if err == nil {
		if stock.Stock_Total < record.Quantity {
			// Jika stok tidak mencukupi, kembalikan error
			return sales.Domain{}, fmt.Errorf("not enough stock for sale with stock_id %d", record.StockID)
		}
		// Jika stok sudah ada, perbarui stok total
		stock.Stock_Total -= record.Quantity // Tambahkan jumlah yang dibeli ke stok total
		record.Total_Price = record.Quantity * stock.Selling_Price

		pr.conn.WithContext(ctx).Save(&stock)
		pr.conn.WithContext(ctx).Save(&record)

	} else {
		// Jika ada kesalahan lain, kembalikan error
		return sales.Domain{}, err
	}

	return record.ToDomain(), nil

}

func (sr *salesRepository) GetAll(ctx context.Context) ([]sales.Domain, error) {
	// var records []Purchase
	// if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
	// 	return nil, err
	// }
	// Memuat data Purchase beserta relasi Vendor, Category, dan Units
	var records []Sales
	if err := sr.conn.WithContext(ctx).
		Preload("Vendor").Preload("Stock").
		Find(&records).Error; err != nil {
		return nil, err
	}

	salesDomain := []sales.Domain{}

	for _, purchase := range records {
		// Konversi ke domain
		domain := purchase.ToDomain()

		// Tambahkan nama Vendor, Category, dan Units ke domain
		// domain.VendorName = purchases.Vendor.Name            // Nama vendor
		// domain.CategoryName = purchase.Category.CategoryName // Nama kategori
		// domain.UnitsName = purchase.Units.Units              // Nama unit

		// Tambahkan ke hasil
		salesDomain = append(salesDomain, domain)
	}

	return salesDomain, nil
}

// func (ur *purchaseRepository) DownloadBarcodeByID(ctx context.Context, id string) (purchase.Domain, error) {
// 	var purchase Purchase

// 	if err := ur.conn.WithContext(ctx).First(&purchase, "id = ?", id).Error; err != nil {
// 		return purchase.Domain{}, err
// 	}

// 	return purchase.ToDomain(), nil

// }

// func (cr *purchaseRepository) PurchaseIn(ctx context.Context, purchaseDomain *purchase.Domain, id string) (purchase.Domain, error) {
// 	purchase, err := cr.DownloadBarcodeByID(ctx, id)

// 	if err != nil {
// 		return purchase.Domain{}, err
// 	}

// 	updatePurchase := FromDomain(&purchase)

// 	updatePurchase.Purchase_Total += purchaseDomain.Purchase_In

// 	if err := cr.conn.WithContext(ctx).Save(&updatePurchase).Error; err != nil {
// 		return purchase.Domain{}, err
// 	}

// 	return updatePurchase.ToDomain(), nil
// }

// func (cr *purchaseRepository) PurchaseOut(ctx context.Context, purchaseDomain *purchase.Domain, id string) (purchase.Domain, error) {
// 	purchase, err := cr.DownloadBarcodeByID(ctx, id)

// 	if err != nil {
// 		return purchase.Domain{}, err
// 	}

// 	updatePurchase := FromDomain(&purchase)

// 	updatePurchase.Purchase_Total -= purchaseDomain.Purchase_Out

// 	if err := cr.conn.WithContext(ctx).Save(&updatePurchase).Error; err != nil {
// 		return purchase.Domain{}, err
// 	}

// 	return updatePurchase.ToDomain(), nil
// }
