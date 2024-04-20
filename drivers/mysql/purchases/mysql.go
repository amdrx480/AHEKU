package purchases

import (
	"backend-golang/businesses/purchases"
	_dbStocks "backend-golang/drivers/mysql/stocks"
	"context"

	// "errors"

	// "fmt"

	"gorm.io/gorm"
)

type purchaseRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) purchases.Repository {
	return &purchaseRepository{
		conn: conn,
	}
}

func (pr *purchaseRepository) GetByID(ctx context.Context, id string) (purchases.Domain, error) {
	var purchase Purchase

	if err := pr.conn.WithContext(ctx).First(&purchase, "id = ?", id).Error; err != nil {
		return purchases.Domain{}, err
	}

	return purchase.ToDomain(), nil

}

func (pr *purchaseRepository) Create(ctx context.Context, purchaseDomain *purchases.Domain) (purchases.Domain, error) {
	record := FromDomain(purchaseDomain)
	result := pr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return purchases.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return purchases.Domain{}, err
	}

	// Tambahkan atau perbarui stok terkait setelah membuat Purchase

	var stock _dbStocks.Stock
	// Cari stok berdasarkan kombinasi stock_code dan stock_unit
	err := pr.conn.WithContext(ctx).
		Where("stock_code = ? AND units_id = ?", record.Stock_Code, record.UnitsID).
		First(&stock).Error

	if err == gorm.ErrRecordNotFound {
		// Jika stok belum ada, buat stok baru
		newStock := _dbStocks.Stock{
			Stock_Name:    record.Stock_Name,
			Stock_Code:    record.Stock_Code,
			CategoryID:    record.CategoryID,
			UnitsID:       record.UnitsID,
			Stock_Total:   record.Quantity, // Jumlah yang dibeli ditambahkan ke stok total
			Selling_Price: record.Selling_Price,
		}
		pr.conn.WithContext(ctx).Create(&newStock)
	} else if err == nil {
		// Jika stok sudah ada, perbarui stok total
		stock.Stock_Total += record.Quantity // Tambahkan jumlah yang dibeli ke stok total
		pr.conn.WithContext(ctx).Save(&stock)
	} else {
		// Jika ada kesalahan lain, kembalikan error
		return purchases.Domain{}, err
	}

	return record.ToDomain(), nil

}

func (sr *purchaseRepository) GetAll(ctx context.Context) ([]purchases.Domain, error) {
	var records []Purchase
	if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []purchases.Domain{}

	for _, category := range records {
		domain := category.ToDomain()
		categories = append(categories, domain)
	}

	return categories, nil
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
