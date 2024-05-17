package purchases

import (
	"backend-golang/businesses/purchases"

	_dbStocks "backend-golang/drivers/mysql/stocks"

	"context"

	// "errors"
	// _dbCategory "backend-golang/drivers/mysql/category"

	// _dbUnits "backend-golang/drivers/mysql/units"
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
	// var category _dbCategory.Category
	// if err := pr.conn.WithContext(ctx).
	// 	Where("category_name = ?", purchaseDomain.CategoryName).
	// 	First(&category).Error; err != nil {
	// 	// Jika Category tidak ditemukan, kembalikan kesalahan
	// 	if err == gorm.ErrRecordNotFound {
	// 		return purchases.Domain{}, fmt.Errorf("Category not found: %w", err)
	// 	}
	// 	return purchases.Domain{}, fmt.Errorf("Failed to fetch category: %w", err)
	// }
	// // Set CategoryID ke stockDomain berdasarkan Category yang ditemukan
	// purchaseDomain.CategoryID = category.ID

	// var units _dbUnits.Units
	// if err := pr.conn.WithContext(ctx).
	// 	Where("units_name = ?", purchaseDomain.UnitsName).
	// 	First(&units).Error; err != nil {
	// 	// Jika Units tidak ditemukan, kembalikan kesalahan
	// 	if err == gorm.ErrRecordNotFound {
	// 		return purchases.Domain{}, fmt.Errorf("Units not found: %w", err)
	// 	}
	// 	return purchases.Domain{}, fmt.Errorf("Failed to fetch Units: %w", err)
	// }
	// // Set UnitsID ke stockDomain berdasarkan Units yang ditemukan
	// purchaseDomain.UnitsID = units.ID

	records := FromDomain(purchaseDomain)
	// preload hasil response saat melakukan post pada purchase untuk isi field Units, Vendor, Category
	result := pr.conn.WithContext(ctx).Preload("Units").Preload("Vendor").Preload("Category").Create(&records)
	// result := pr.conn.WithContext(ctx).Create(&records)

	if err := result.Error; err != nil {
		return purchases.Domain{}, err
	}

	if err := result.Last(&records).Error; err != nil {
		return purchases.Domain{}, err
	}

	// Tambahkan atau perbarui stok terkait setelah membuat Purchase
	var stock _dbStocks.Stock
	// Cari stok berdasarkan kombinasi stock_code, category_namedan stock_unit
	err := pr.conn.WithContext(ctx).
		// Where("stock_code = ? AND category_name = ? AND units_name = ?", records.Stock_Code, records.CategoryName, records.UnitsName).
		// Where("stock_code = ? AND stock_Name = ? AND category_id = ? AND units_id = ?", records.Stock_Code, records.Stock_Name, records.CategoryID, records.UnitsID).

		Where("stock_code = ? AND stock_Name = ? AND units_id = ?", records.Stock_Code, records.Stock_Name, records.UnitsID).
		First(&stock).Error

	if err == gorm.ErrRecordNotFound {
		// Jika stok belum ada, buat stok baru
		newStock := _dbStocks.Stock{
			Stock_Name: records.Stock_Name,
			Stock_Code: records.Stock_Code,
			CategoryID: records.CategoryID,
			// CategoryName: records.CategoryName,

			UnitsID:     records.UnitsID,
			Description: records.Description,
			// UnitsName: records.UnitsName,

			Stock_Total:   records.Quantity, // Jumlah yang dibeli ditambahkan ke stok total
			Selling_Price: records.Selling_Price,
		}
		pr.conn.WithContext(ctx).Create(&newStock)
	} else if err == nil {
		// Jika stok sudah ada, perbarui stok total
		stock.Stock_Total += records.Quantity // Tambahkan jumlah yang dibeli ke stok total
		pr.conn.WithContext(ctx).Save(&stock)
	} else {
		// Jika ada kesalahan lain, kembalikan error
		return purchases.Domain{}, err
	}

	return records.ToDomain(), nil

}

func (sr *purchaseRepository) GetAll(ctx context.Context) ([]purchases.Domain, error) {
	// var records []Purchase
	// if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
	// 	return nil, err
	// }
	// Memuat data Purchase beserta relasi Vendor, Category, dan Units
	var records []Purchase
	if err := sr.conn.WithContext(ctx).
		Preload("Vendor").Preload("Category").Preload("Units").
		Find(&records).Error; err != nil {
		return nil, err
	}

	purchasesDomain := []purchases.Domain{}

	for _, purchase := range records {
		// Konversi ke domain
		domain := purchase.ToDomain()

		// Tambahkan nama Vendor, Category, dan Units ke domain
		// domain.VendorName = purchases.Vendor.Name            // Nama vendor
		// domain.CategoryName = purchase.Category.CategoryName // Nama kategori
		// domain.UnitsName = purchase.Units.Units              // Nama unit

		// Tambahkan ke hasil
		purchasesDomain = append(purchasesDomain, domain)
	}

	return purchasesDomain, nil
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
