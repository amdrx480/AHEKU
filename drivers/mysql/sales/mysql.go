package sales

import (
	"backend-golang/businesses/history"
	"backend-golang/businesses/sales"
	"log"

	_dbHistory "backend-golang/drivers/mysql/history"
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

	if err := sr.conn.WithContext(ctx).Preload("Stock").First(&sale, "id = ?", id).Error; err != nil {
		return sales.Domain{}, err
	}

	return sale.ToDomain(), nil

}

func (pr *salesRepository) Create(ctx context.Context, saleDomain *sales.Domain) (sales.Domain, error) {
	record := FromDomain(saleDomain)
	// Pertama, periksa apakah stok mencukupi
	var stock _dbStocks.Stock
	err := pr.conn.WithContext(ctx).
		Where("id = ?", record.StockID).
		First(&stock).Error

	if err == gorm.ErrRecordNotFound {
		errMsg := fmt.Sprintf("stok tidak ditemukan dengan stock_id %d", record.StockID)
		log.Println(errMsg)
		return sales.Domain{}, fmt.Errorf("stok tidak ditemukan dengan stock_id %d", record.StockID)
	} else if err != nil {
		// Jika ada kesalahan lain, kembalikan error
		return sales.Domain{}, err
	}

	if stock.Stock_Total < record.Quantity {
		// Jika stok tidak mencukupi, kembalikan error
		errMsg := fmt.Sprintf("stok tidak cukup untuk penjualan dengan stock_id %d", record.StockID)
		log.Println(errMsg)
		return sales.Domain{}, fmt.Errorf(errMsg)
	}

	// Hitung total harga sebelum menyimpan catatan penjualan
	record.TotalPrice = record.Quantity * stock.Selling_Price

	// Hitung total harga sebelum menyimpan catatan penjualan
	record.TotalPrice = record.Quantity * stock.Selling_Price

	// Hitung TotalAllPrice secara global
	var allSales []Sales
	err = pr.conn.WithContext(ctx).
		Find(&allSales).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return sales.Domain{}, err
	}

	totalAllPrice := record.TotalPrice
	for _, sale := range allSales {
		totalAllPrice += sale.TotalPrice
	}
	record.TotalAllPrice = totalAllPrice

	// Simpan catatan penjualan
	result := pr.conn.WithContext(ctx).Create(&record)
	if err := result.Error; err != nil {
		return sales.Domain{}, err
	}

	// Perbarui TotalAllPrice untuk semua penjualan lainnya
	for _, sale := range allSales {
		sale.TotalAllPrice = totalAllPrice
		if err := pr.conn.WithContext(ctx).Save(&sale).Error; err != nil {
			return sales.Domain{}, err
		}
	}

	return record.ToDomain(), nil
}

func (sr *salesRepository) ToHistory(ctx context.Context, historyDomain *history.Domain, id string) (history.Domain, error) {
	record := _dbHistory.FromDomain(historyDomain)
	var salesData []Sales

	// Ambil semua data penjualan
	if err := sr.conn.WithContext(ctx).Find(&salesData).Error; err != nil {
		return history.Domain{}, err
	}

	for _, sale := range salesData {
		// Periksa stok yang terkait dengan penjualan
		var stock _dbStocks.Stock
		if err := sr.conn.WithContext(ctx).Where("id = ?", sale.StockID).First(&stock).Error; err != nil {
			return history.Domain{}, err
		}

		// Kurangi stok dengan jumlah yang dijual
		if stock.Stock_Total < sale.Quantity {
			errMsg := fmt.Sprintf("stok tidak cukup untuk penjualan dengan stock_id %d", record.StockID)
			log.Println(errMsg)
			return history.Domain{}, fmt.Errorf(errMsg)
		}
		stock.Stock_Total -= sale.Quantity

		// Simpan perubahan stok ke database
		if err := sr.conn.WithContext(ctx).Save(&stock).Error; err != nil {
			return history.Domain{}, err
		}

		// Buat catatan history
		historyRecord := _dbHistory.History{
			StockID:    sale.StockID,
			Quantity:   sale.Quantity,
			TotalPrice: sale.TotalPrice,
			// Sesuaikan dengan kolom lain jika perlu
		}

		// Simpan catatan history
		if err := sr.conn.WithContext(ctx).Create(&historyRecord).Error; err != nil {
			return history.Domain{}, err
		}
	}

	// Hapus semua data penjualan sekaligus
	if err := sr.conn.WithContext(ctx).Unscoped().Delete(salesData).Error; err != nil {
		return history.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (sr *salesRepository) GetAll(ctx context.Context) ([]sales.Domain, error) {
	var records []Sales
	if err := sr.conn.WithContext(ctx).Preload("Stock").
		Find(&records).Error; err != nil {
		return nil, err
	}

	salesDomain := []sales.Domain{}

	for _, purchase := range records {
		// Konversi ke domain
		domain := purchase.ToDomain()
		salesDomain = append(salesDomain, domain)
	}

	return salesDomain, nil
}

func (sr *salesRepository) Delete(ctx context.Context, id string) error {
	sales, err := sr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedSales := FromDomain(&sales)

	// Ambil semua penjualan lainnya
	var allSales []Sales
	err = sr.conn.WithContext(ctx).Find(&allSales).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// Kurangi TotalPrice dari penjualan yang dihapus dari TotalAllPrice global
	totalAllPrice := 0
	for _, sale := range allSales {
		if sale.ID != deletedSales.ID {
			totalAllPrice += sale.TotalPrice
		}
	}

	// Perbarui TotalAllPrice untuk semua penjualan lainnya
	for _, sale := range allSales {
		if sale.ID != deletedSales.ID {
			sale.TotalAllPrice = totalAllPrice
			if err := sr.conn.WithContext(ctx).Save(&sale).Error; err != nil {
				return err
			}
		}
	}

	// Hapus data penjualan
	if err := sr.conn.WithContext(ctx).Unscoped().Delete(&deletedSales).Error; err != nil {
		return err
	}

	return nil
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
