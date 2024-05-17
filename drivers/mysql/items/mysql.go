package items

import (
	// "backend-golang/businesses/history"
	"backend-golang/businesses/items"
	"log"

	// _dbHistory "backend-golang/drivers/mysql/history"
	// _dbCart "backend-golang/drivers/mysql/cart"
	_dbStocks "backend-golang/drivers/mysql/stocks"

	"context"
	"fmt"

	"gorm.io/gorm"
)

type itemsRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) items.Repository {
	return &itemsRepository{
		conn: conn,
	}
}

func (sr *itemsRepository) GetByID(ctx context.Context, id string) (items.Domain, error) {
	var item Items

	if err := sr.conn.WithContext(ctx).Preload("Stock").First(&item, "id = ?", id).Error; err != nil {
		return items.Domain{}, err
	}

	return item.ToDomain(), nil

}

func (pr *itemsRepository) Create(ctx context.Context, itemDomain *items.Domain) (items.Domain, error) {
	record := FromDomain(itemDomain)
	// Pertama, periksa apakah stok mencukupi
	var stock _dbStocks.Stock
	err := pr.conn.WithContext(ctx).
		Where("id = ?", record.StockID).
		First(&stock).Error

	if err == gorm.ErrRecordNotFound {
		errMsg := fmt.Sprintf("stok tidak ditemukan dengan stock_id %d", record.StockID)
		log.Println(errMsg)
		return items.Domain{}, fmt.Errorf("stok tidak ditemukan dengan stock_id %d", record.StockID)
	} else if err != nil {
		// Jika ada kesalahan lain, kembalikan error
		return items.Domain{}, err
	}

	if stock.Stock_Total < record.Quantity {
		// Jika stok tidak mencukupi, kembalikan error
		errMsg := fmt.Sprintf("stok tidak cukup untuk penjualan dengan stock_id %d", record.StockID)
		log.Println(errMsg)
		return items.Domain{}, fmt.Errorf(errMsg)
	}

	// Hitung total harga sebelum menyimpan catatan penjualan
	record.Price = record.Quantity * stock.Selling_Price

	// Simpan catatan penjualan
	result := pr.conn.WithContext(ctx).Create(&record)
	if err := result.Error; err != nil {
		return items.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (ir *itemsRepository) CreateCart(ctx context.Context, cartDomain *items.DomainCart) (items.DomainCart, error) {
	record := FromDomainCart(cartDomain)
	// Pertama, periksa apakah stok mencukupi
	// // Hitung total harga sebelum menyimpan catatan penjualan
	// record.Price = record.Quantity * stock.Selling_Price

	// Simpan catatan penjualan
	result := ir.conn.WithContext(ctx).Create(&record)
	if err := result.Error; err != nil {
		return items.DomainCart{}, err
	}
	err := ir.conn.WithContext(ctx).Last(&record).Error

	if err != nil {
		return items.DomainCart{}, err
	}

	return record.ToDomainCart(), nil
}

func (sr *itemsRepository) GetAll(ctx context.Context) ([]items.Domain, error) {
	var records []Items
	if err := sr.conn.WithContext(ctx).Preload("Stock").
		Find(&records).Error; err != nil {
		return nil, err
	}

	itemsDomain := []items.Domain{}

	for _, purchase := range records {
		// Konversi ke domain
		domain := purchase.ToDomain()
		itemsDomain = append(itemsDomain, domain)
	}

	return itemsDomain, nil
}

func (sr *itemsRepository) Delete(ctx context.Context, id string) error {
	items, err := sr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedItems := FromDomain(&items)

	// Ambil semua penjualan lainnya
	var allItems []Items
	err = sr.conn.WithContext(ctx).Find(&allItems).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// // Kurangi TotalPrice dari penjualan yang dihapus dari TotalAllPrice global
	// totalAllPrice := 0
	// for _, item := range allItems {
	// 	if item.ID != deletedItems.ID {
	// 		totalAllPrice += item.TotalPrice
	// 	}
	// }

	// // Perbarui TotalAllPrice untuk semua penjualan lainnya
	// for _, item := range allItems {
	// 	if item.ID != deletedItems.ID {
	// 		item.TotalAllPrice = totalAllPrice
	// 		if err := sr.conn.WithContext(ctx).Save(&item).Error; err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	// Hapus data penjualan
	if err := sr.conn.WithContext(ctx).Unscoped().Delete(&deletedItems).Error; err != nil {
		return err
	}

	return nil
}

func (sr *itemsRepository) GetByIDCart(ctx context.Context, id string) (items.DomainCart, error) {
	var cart Cart

	if err := sr.conn.WithContext(ctx).Preload("Items").First(&cart, "id = ?", id).Error; err != nil {
		return items.DomainCart{}, err
	}

	return cart.ToDomainCart(), nil

}

func (sr *itemsRepository) GetAllCart(ctx context.Context) ([]items.DomainCart, error) {
	var records []Cart
	if err := sr.conn.WithContext(ctx).Preload("Items").
		// if err := sr.conn.WithContext(ctx).
		Find(&records).Error; err != nil {
		return nil, err
	}

	itemsDomain := []items.DomainCart{}

	for _, purchase := range records {
		// Konversi ke domain
		domain := purchase.ToDomainCart()
		itemsDomain = append(itemsDomain, domain)
	}

	return itemsDomain, nil
}

func (sr *itemsRepository) DeleteCart(ctx context.Context, id string) error {
	items, err := sr.GetByIDCart(ctx, id)

	if err != nil {
		return err
	}

	deletedItems := FromDomainCart(&items)

	// Ambil semua penjualan lainnya
	var allItems []Cart
	err = sr.conn.WithContext(ctx).Find(&allItems).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// // Kurangi TotalPrice dari penjualan yang dihapus dari TotalAllPrice global
	// totalAllPrice := 0
	// for _, item := range allItems {
	// 	if item.ID != deletedItems.ID {
	// 		totalAllPrice += item.TotalPrice
	// 	}
	// }

	// // Perbarui TotalAllPrice untuk semua penjualan lainnya
	// for _, item := range allItems {
	// 	if item.ID != deletedItems.ID {
	// 		item.TotalAllPrice = totalAllPrice
	// 		if err := sr.conn.WithContext(ctx).Save(&item).Error; err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	// Hapus data penjualan
	if err := sr.conn.WithContext(ctx).Unscoped().Delete(&deletedItems).Error; err != nil {
		return err
	}

	return nil
}

// func (pr *itemsRepository) Create(ctx context.Context, itemDomain *items.Domain) (items.Domain, error) {
// 	record := FromDomain(itemDomain)
// 	// Pertama, periksa apakah stok mencukupi
// 	var stock _dbStocks.Stock
// 	err := pr.conn.WithContext(ctx).
// 		Where("id = ?", record.StockID).
// 		First(&stock).Error

// 	if err == gorm.ErrRecordNotFound {
// 		errMsg := fmt.Sprintf("stok tidak ditemukan dengan stock_id %d", record.StockID)
// 		log.Println(errMsg)
// 		return items.Domain{}, fmt.Errorf("stok tidak ditemukan dengan stock_id %d", record.StockID)
// 	} else if err != nil {
// 		// Jika ada kesalahan lain, kembalikan error
// 		return items.Domain{}, err
// 	}

// 	if stock.Stock_Total < record.Quantity {
// 		// Jika stok tidak mencukupi, kembalikan error
// 		errMsg := fmt.Sprintf("stok tidak cukup untuk penjualan dengan stock_id %d", record.StockID)
// 		log.Println(errMsg)
// 		return items.Domain{}, fmt.Errorf(errMsg)
// 	}

// 	// Hitung total harga sebelum menyimpan catatan penjualan
// 	record.Price = record.Quantity * stock.Selling_Price

// 	// Hitung total harga sebelum menyimpan catatan penjualan
// 	record.Price = record.Quantity * stock.Selling_Price

// 	// Hitung TotalAllPrice secara global
// 	var allItems []Items
// 	var Carts _dbCart.Cart
// 	err = pr.conn.WithContext(ctx).
// 		Find(&allItems).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return items.Domain{}, err
// 	}

// 	totalPriceItems := record.Price
// 	for _, item := range allItems {
// 		totalPriceItems += item.Price
// 	}
// 	Carts.Total = totalPriceItems

// 	// Simpan catatan penjualan
// 	result := pr.conn.WithContext(ctx).Create(&record)
// 	if err := result.Error; err != nil {
// 		return items.Domain{}, err
// 	}

// 	// Perbarui TotalAllPrice untuk semua penjualan lainnya
// 	for _, item := range allItems {
// 		Carts.Total = totalPriceItems
// 		if err := pr.conn.WithContext(ctx).Save(&item).Error; err != nil {
// 			return items.Domain{}, err
// 		}
// 	}

// 	return record.ToDomain(), nil
// }

// func (pr *itemsRepository) Create(ctx context.Context, itemDomain *items.Domain) (items.Domain, error) {
// 	record := FromDomain(itemDomain)

// 	// Pertama, periksa apakah stok mencukupi
// 	var stock _dbStocks.Stock
// 	err := pr.conn.WithContext(ctx).
// 		Where("id = ?", record.StockID).
// 		First(&stock).Error

// 	if err == gorm.ErrRecordNotFound {
// 		errMsg := fmt.Sprintf("stok tidak ditemukan dengan stock_id %d", record.StockID)
// 		log.Println(errMsg)
// 		return items.Domain{}, fmt.Errorf("stok tidak ditemukan dengan stock_id %d", record.StockID)
// 	} else if err != nil {
// 		// Jika ada kesalahan lain, kembalikan error
// 		return items.Domain{}, err
// 	}

// 	if stock.Stock_Total < record.Quantity {
// 		// Jika stok tidak mencukupi, kembalikan error
// 		errMsg := fmt.Sprintf("stok tidak cukup untuk penjualan dengan stock_id %d", record.StockID)
// 		log.Println(errMsg)
// 		return items.Domain{}, fmt.Errorf(errMsg)
// 	}

// 	// Hitung total harga sebelum menyimpan catatan penjualan
// 	record.Price = record.Quantity * stock.Selling_Price

// 	// Mulai transaksi
// 	tx := pr.conn.Begin()
// 	if tx.Error != nil {
// 		return items.Domain{}, tx.Error
// 	}

// 	// // Periksa apakah cart_id valid
// 	// var cart _dbCart.Cart
// 	// err = tx.WithContext(ctx).
// 	// 	Where("id = ?", record.CartID).
// 	// 	First(&cart).Error

// 	// if err == gorm.ErrRecordNotFound {
// 	// 	errMsg := fmt.Sprintf("keranjang tidak ditemukan dengan cart_id %d", record.CartID)
// 	// 	log.Println(errMsg)
// 	// 	tx.Rollback()
// 	// 	return items.Domain{}, fmt.Errorf("keranjang tidak ditemukan dengan cart_id %d", record.CartID)
// 	// } else if err != nil {
// 	// 	tx.Rollback()
// 	// 	return items.Domain{}, err
// 	// }

// 	// Hitung TotalAllPrice secara global
// 	// var allItems []Items
// 	// err = tx.WithContext(ctx).
// 	// 	Where("cart_id = ?", record.CartID).
// 	// 	Find(&allItems).Error
// 	// if err != nil && err != gorm.ErrRecordNotFound {
// 	// 	tx.Rollback()
// 	// 	return items.Domain{}, err
// 	// }

// 	// totalPriceItems := record.Price
// 	// for _, item := range allItems {
// 	// 	totalPriceItems += item.Price
// 	// }
// 	// cart.Total = totalPriceItems

// 	// Simpan catatan penjualan
// 	result := tx.WithContext(ctx).Create(&record)
// 	if result.Error != nil {
// 		tx.Rollback()
// 		return items.Domain{}, result.Error
// 	}

// 	// Perbarui total harga di keranjang
// 	// if err := tx.WithContext(ctx).Save(&cart).Error; err != nil {
// 	if err := tx.WithContext(ctx).Save(&record).Error; err != nil {

// 		tx.Rollback()
// 		return items.Domain{}, err
// 	}

// 	// Commit transaksi
// 	if err := tx.Commit().Error; err != nil {
// 		tx.Rollback()
// 		return items.Domain{}, err
// 	}

// 	return record.ToDomain(), nil
// }

// func (sr *itemsRepository) ToHistory(ctx context.Context, historyDomain *history.Domain, id string) (history.Domain, error) {
// 	record := _dbHistory.FromDomain(historyDomain)
// 	var itemsData []Items

// 	// Ambil semua data penjualan
// 	if err := sr.conn.WithContext(ctx).Find(&itemsData).Error; err != nil {
// 		return history.Domain{}, err
// 	}

// 	for _, item := range itemsData {
// 		// Periksa stok yang terkait dengan penjualan
// 		var stock _dbStocks.Stock
// 		if err := sr.conn.WithContext(ctx).Where("id = ?", item.StockID).First(&stock).Error; err != nil {
// 			return history.Domain{}, err
// 		}

// 		// Kurangi stok dengan jumlah yang dijual
// 		if stock.Stock_Total < item.Quantity {
// 			errMsg := fmt.Sprintf("stok tidak cukup untuk penjualan dengan stock_id %d", record.StockID)
// 			log.Println(errMsg)
// 			return history.Domain{}, fmt.Errorf(errMsg)
// 		}
// 		stock.Stock_Total -= item.Quantity

// 		// Simpan perubahan stok ke database
// 		if err := sr.conn.WithContext(ctx).Save(&stock).Error; err != nil {
// 			return history.Domain{}, err
// 		}

// 		// Buat catatan history
// 		historyRecord := _dbHistory.History{
// 			StockID:    item.StockID,
// 			StockName:  item.StockName,
// 			Quantity:   item.Quantity,
// 			TotalPrice: item.TotalPrice,
// 			// Sesuaikan dengan kolom lain jika perlu
// 		}

// 		// Simpan catatan history
// 		if err := sr.conn.WithContext(ctx).Create(&historyRecord).Error; err != nil {
// 			return history.Domain{}, err
// 		}
// 	}

// 	// Hapus semua data penjualan sekaligus
// 	if err := sr.conn.WithContext(ctx).Unscoped().Delete(itemsData).Error; err != nil {
// 		return history.Domain{}, err
// 	}

// 	return record.ToDomain(), nil
// }

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

//type itemsRepository struct {
// 	connItems *gorm.DB // Koneksi ke database items
// 	connCarts *gorm.DB // Koneksi ke database carts
// }

// // Buat dua koneksi database: satu untuk database items dan satu lagi untuk database carts.
// func NewMySQLRepository(connItems *gorm.DB, connCarts *gorm.DB) items.Repository {
// 	return &itemsRepository{
// 		connItems: connItems,
// 		connCarts: connCarts,
// 	}
// }

// type itemsRepository struct {
// 	conn *gorm.DB
// }

// func NewMySQLRepository(conn *gorm.DB) items.Repository {
// 	return &itemsRepository{
// 		conn: conn,
// 	}
// }
