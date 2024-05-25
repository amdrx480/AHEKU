package admin

import (
	"backend-golang/businesses/admin"
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type adminRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) admin.Repository {
	return &adminRepository{
		conn: conn,
	}
}

func (ur *adminRepository) AdminRegister(ctx context.Context, adminDomain *admin.AdminsDomain) (admin.AdminsDomain, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(adminDomain.Password), bcrypt.DefaultCost)

	if err != nil {
		return admin.AdminsDomain{}, err
	}

	record := FromAdminsDomain(adminDomain)

	record.Password = string(password)

	result := ur.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return admin.AdminsDomain{}, err
	}

	err = result.Last(&record).Error

	if err != nil {
		return admin.AdminsDomain{}, err
	}

	return record.ToAdminsDomain(), nil
}

func (ur *adminRepository) AdminGetByName(ctx context.Context, adminDomain *admin.AdminsDomain) (admin.AdminsDomain, error) {
	var admins Admins

	err := ur.conn.WithContext(ctx).First(&admins, "name = ?", adminDomain.Name).Error

	if err != nil {
		return admin.AdminsDomain{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admins.Password), []byte(adminDomain.Password))

	if err != nil {
		return admin.AdminsDomain{}, err
	}

	return admins.ToAdminsDomain(), nil
}

func (ur *adminRepository) AdminGetByVoucher(ctx context.Context, adminDomain *admin.AdminsDomain) (admin.AdminsDomain, error) {
	var admins Admins

	err := ur.conn.WithContext(ctx).First(&admins, "voucher = ?", adminDomain.Voucher).Error

	if err != nil {
		return admin.AdminsDomain{}, err
	}

	return admins.ToAdminsDomain(), nil
}

func (cr *adminRepository) CustomersCreate(ctx context.Context, customersDomain *admin.CustomersDomain) (admin.CustomersDomain, error) {
	record := FromCustomersDomain(customersDomain)
	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return admin.CustomersDomain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return admin.CustomersDomain{}, err
	}

	return record.ToCustomersDomain(), nil

}

func (cr *adminRepository) CustomersGetByID(ctx context.Context, id string) (admin.CustomersDomain, error) {
	var customer Customers

	if err := cr.conn.WithContext(ctx).Preload("CartItems.Customers").Preload("CartItems.Stocks").
		First(&customer, "id = ?", id).Error; err != nil {
		return admin.CustomersDomain{}, err
	}

	return customer.ToCustomersDomain(), nil

}

func (sr *adminRepository) CustomersGetAll(ctx context.Context) ([]admin.CustomersDomain, error) {
	var records []Customers
	// Melakukan Preload untuk menampilkan Slice CartItems yang berisi Customers dan Stocks
	if err := sr.conn.WithContext(ctx).Preload("CartItems.Customers").Preload("CartItems.Stocks").
		Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []admin.CustomersDomain{}

	for _, category := range records {
		domain := category.ToCustomersDomain()
		categories = append(categories, domain)
	}

	return categories, nil
}

// Categories
func (cr *adminRepository) CategoryCreate(ctx context.Context, AdminsDomain *admin.CategoriesDomain) (admin.CategoriesDomain, error) {
	record := FromCategoriesDomain(AdminsDomain)
	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return admin.CategoriesDomain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return admin.CategoriesDomain{}, err
	}

	return record.ToCategoriesDomain(), nil

}

func (cr *adminRepository) CategoryGetByID(ctx context.Context, id string) (admin.CategoriesDomain, error) {
	var categories Categories

	if err := cr.conn.WithContext(ctx).First(&categories, "id = ?", id).Error; err != nil {
		return admin.CategoriesDomain{}, err
	}

	return categories.ToCategoriesDomain(), nil

}

func (cr *adminRepository) CategoryGetByName(ctx context.Context, name string) (admin.CategoriesDomain, error) {
	var categories Categories

	if err := cr.conn.WithContext(ctx).First(&categories, "admin_name = ?", name).Error; err != nil {
		return admin.CategoriesDomain{}, err
	}

	return categories.ToCategoriesDomain(), nil

}

func (cr *adminRepository) CategoryGetAll(ctx context.Context) ([]admin.CategoriesDomain, error) {
	var records []Categories
	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []admin.CategoriesDomain{}

	for _, admin := range records {
		domain := admin.ToCategoriesDomain()
		categories = append(categories, domain)
	}

	return categories, nil
}

func (vr *adminRepository) VendorsCreate(ctx context.Context, purchaseDomain *admin.VendorsDomain) (admin.VendorsDomain, error) {
	record := FromVendorsDomain(purchaseDomain)
	result := vr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return admin.VendorsDomain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return admin.VendorsDomain{}, err
	}

	return record.ToVendorsDomain(), nil

}

func (vr *adminRepository) VendorsGetByID(ctx context.Context, id string) (admin.VendorsDomain, error) {
	var vendor Vendors

	if err := vr.conn.WithContext(ctx).First(&vendor, "id = ?", id).Error; err != nil {
		return admin.VendorsDomain{}, err
	}

	return vendor.ToVendorsDomain(), nil

}

func (sr *adminRepository) VendorsGetAll(ctx context.Context) ([]admin.VendorsDomain, error) {
	var records []Vendors
	if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []admin.VendorsDomain{}

	for _, category := range records {
		domain := category.ToVendorsDomain()
		categories = append(categories, domain)
	}

	return categories, nil
}

func (ur *adminRepository) UnitsCreate(ctx context.Context, unitsDomain *admin.UnitsDomain) (admin.UnitsDomain, error) {
	record := FromUnitsDomain(unitsDomain)
	result := ur.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return admin.UnitsDomain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return admin.UnitsDomain{}, err
	}

	return record.ToUnitsDomain(), nil
}

func (ur *adminRepository) UnitsGetByID(ctx context.Context, id string) (admin.UnitsDomain, error) {
	var unit Units

	if err := ur.conn.WithContext(ctx).First(&unit, "id = ?", id).Error; err != nil {
		return admin.UnitsDomain{}, err
	}

	return unit.ToUnitsDomain(), nil

}

func (ur *adminRepository) UnitsGetAll(ctx context.Context) ([]admin.UnitsDomain, error) {
	var records []Units
	if err := ur.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	units := []admin.UnitsDomain{}

	for _, unit := range records {
		domain := unit.ToUnitsDomain()
		units = append(units, domain)
	}

	return units, nil
}

func (cr *adminRepository) StocksCreate(ctx context.Context, stockDomain *admin.StocksDomain) (admin.StocksDomain, error) {

	// // Cari Categories berdasarkan CategoryName yang diberikan
	// var category _dbCategory.Categories
	// if err := cr.conn.WithContext(ctx).Where("category_name = ?", stockDomain.CategoryName).First(&category).Error; err != nil {
	// 	// Jika Categories tidak ditemukan, kembalikan kesalahan
	// 	if err == gorm.ErrRecordNotFound {
	// 		return admin.StocksDomain{}, fmt.Errorf("Categories not found: %w", err)
	// 	}
	// 	return admin.StocksDomain{}, fmt.Errorf("Failed to fetch category: %w", err)
	// }

	// // Set CategoryID ke stockDomain berdasarkan Categories yang ditemukan
	// stockDomain.CategoryID = category.ID
	// // stockDomain.CategoryName = category.CategoryName

	record := FromStocksDomain(stockDomain)
	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return admin.StocksDomain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return admin.StocksDomain{}, err
	}

	return record.ToStocksDomain(), nil

}

func (ur *adminRepository) StocksGetByID(ctx context.Context, id string) (admin.StocksDomain, error) {
	var stock Stocks

	if err := ur.conn.WithContext(ctx).First(&stock, "id = ?", id).Error; err != nil {
		return admin.StocksDomain{}, err
	}

	return stock.ToStocksDomain(), nil

}

func (sr *adminRepository) StocksGetAll(ctx context.Context) ([]admin.StocksDomain, error) {
	var records []Stocks
	if err := sr.conn.WithContext(ctx).
		Preload("Categories").Preload("Units").
		Find(&records).Error; err != nil {
		return nil, err
	}

	stocksDomain := []admin.StocksDomain{}

	for _, stocks := range records {
		domain := stocks.ToStocksDomain()
		stocksDomain = append(stocksDomain, domain)
	}

	return stocksDomain, nil
}

func (pr *adminRepository) PurchasesCreate(ctx context.Context, purchaseDomain *admin.PurchasesDomain) (admin.PurchasesDomain, error) {
	// var category _dbCategory.Categories
	// if err := pr.conn.WithContext(ctx).
	// 	Where("category_name = ?", purchaseDomain.CategoryName).
	// 	First(&category).Error; err != nil {
	// 	// Jika Categories tidak ditemukan, kembalikan kesalahan
	// 	if err == gorm.ErrRecordNotFound {
	// 		return admin.PurchasesDomain{}, fmt.Errorf("Categories not found: %w", err)
	// 	}
	// 	return admin.PurchasesDomain{}, fmt.Errorf("Failed to fetch category: %w", err)
	// }
	// // Set CategoryID ke stockDomain berdasarkan Categories yang ditemukan
	// purchaseDomain.CategoryID = category.ID

	// var units _dbUnits.Units
	// if err := pr.conn.WithContext(ctx).
	// 	Where("units_name = ?", purchaseDomain.UnitsName).
	// 	First(&units).Error; err != nil {
	// 	// Jika Units tidak ditemukan, kembalikan kesalahan
	// 	if err == gorm.ErrRecordNotFound {
	// 		return admin.PurchasesDomain{}, fmt.Errorf("Units not found: %w", err)
	// 	}
	// 	return admin.PurchasesDomain{}, fmt.Errorf("Failed to fetch Units: %w", err)
	// }
	// // Set UnitID ke stockDomain berdasarkan Units yang ditemukan
	// purchaseDomain.UnitID = units.ID

	records := FromPurchasesDomain(purchaseDomain)
	// preload hasil response saat melakukan post pada purchase untuk isi field Units, Vendors, Categories
	// result := pr.conn.WithContext(ctx).Preload("Units").Preload("Vendors").Preload("Categories").Create(&records)
	result := pr.conn.WithContext(ctx).Create(&records)

	if err := result.Error; err != nil {
		return admin.PurchasesDomain{}, err
	}

	if err := result.Last(&records).Error; err != nil {
		return admin.PurchasesDomain{}, err
	}

	// Tambahkan atau perbarui stok terkait setelah membuat Purchases
	var stock Stocks
	// Cari stok berdasarkan kombinasi stock_code, category_namedan stock_unit
	err := pr.conn.WithContext(ctx).
		// Where("stock_code = ? AND category_name = ? AND units_name = ?", records.StockCode, records.CategoryName, records.UnitsName).
		// Where("stock_code = ? AND stock_Name = ? AND category_id = ? AND units_id = ?", records.StockCode, records.StockName, records.CategoryID, records.UnitID).

		Where("stock_code = ? AND stock_Name = ? AND unit_id = ?", records.StockCode, records.StockName, records.UnitID).
		First(&stock).Error

	if err == gorm.ErrRecordNotFound {
		// Jika stok belum ada, buat stok baru
		newStock := Stocks{
			StockName:  records.StockName,
			StockCode:  records.StockCode,
			CategoryID: records.CategoryID,
			// CategoryName: records.CategoryName,
			UnitID:       records.UnitID,
			Description:  records.Description,
			StockTotal:   records.Quantity, // Jumlah yang dibeli ditambahkan ke stok total
			SellingPrice: records.SellingPrice,
		}
		pr.conn.WithContext(ctx).Create(&newStock)
	} else if err == nil {
		// Jika stok sudah ada, perbarui stok total
		stock.StockTotal += records.Quantity // Tambahkan jumlah yang dibeli ke stok total
		pr.conn.WithContext(ctx).Save(&stock)
	} else {
		// Jika ada kesalahan lain, kembalikan error
		return admin.PurchasesDomain{}, err
	}

	return records.ToPurchasesDomain(), nil

}

func (pr *adminRepository) PurchasesGetByID(ctx context.Context, id string) (admin.PurchasesDomain, error) {
	var purchase Purchases

	if err := pr.conn.WithContext(ctx).First(&purchase, "id = ?", id).Error; err != nil {
		return admin.PurchasesDomain{}, err
	}

	return purchase.ToPurchasesDomain(), nil

}

func (sr *adminRepository) PurchasesGetAll(ctx context.Context) ([]admin.PurchasesDomain, error) {
	// Memuat data Purchases beserta relasi Vendors, Categories, dan Units
	var records []Purchases
	if err := sr.conn.WithContext(ctx).
		// Preload("Vendors").Preload("Categories").Preload("Units").
		Preload("Vendors").Preload("Categories").Preload("Units").
		Find(&records).Error; err != nil {
		return nil, err
	}

	purchasesDomain := []admin.PurchasesDomain{}

	for _, purchase := range records {
		// Konversi ke domain
		domain := purchase.ToPurchasesDomain()
		// Tambahkan ke hasil
		purchasesDomain = append(purchasesDomain, domain)
	}

	return purchasesDomain, nil
}

func (pr *adminRepository) CartItemsCreate(ctx context.Context, itemDomain *admin.CartItemsDomain) (admin.CartItemsDomain, error) {
	record := FromCartItemsDomain(itemDomain)
	// Pertama, periksa apakah stok mencukupi
	var stock Stocks
	err := pr.conn.WithContext(ctx).
		Where("id = ?", record.StockID).
		First(&stock).Error

	if err == gorm.ErrRecordNotFound {
		errMsg := fmt.Sprintf("stok tidak ditemukan dengan stock_id %d", record.StockID)
		log.Println(errMsg)
		return admin.CartItemsDomain{}, fmt.Errorf("stok tidak ditemukan dengan stock_id %d", record.StockID)
	} else if err != nil {
		// Jika ada kesalahan lain, kembalikan error
		return admin.CartItemsDomain{}, err
	}

	if stock.StockTotal < record.Quantity {
		// Jika stok tidak mencukupi, kembalikan error
		errMsg := fmt.Sprintf("stok tidak cukup untuk penjualan dengan stock_id %d", record.StockID)
		log.Println(errMsg)
		return admin.CartItemsDomain{}, fmt.Errorf(errMsg)
	}

	// Hitung total harga sebelum menyimpan catatan penjualan
	record.Price = record.Quantity * stock.SellingPrice

	// Hitung SubTotal berdasarkan barang yang dibeli customer yang berbeda
	var customerCartItems []CartItems
	err = pr.conn.WithContext(ctx).
		Where("customer_id = ?", record.CustomerID).
		Find(&customerCartItems).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return admin.CartItemsDomain{}, err
	}

	subTotal := record.Price
	for _, sale := range customerCartItems {
		subTotal += sale.Price
	}
	record.SubTotal = subTotal

	// Simpan catatan penjualan
	result := pr.conn.WithContext(ctx).Create(&record)
	if err := result.Error; err != nil {
		return admin.CartItemsDomain{}, err
	}

	// Perbarui SubTotal untuk semua penjualan lainnya dari pelanggan yang sama
	for _, sale := range customerCartItems {
		sale.SubTotal = subTotal
		if err := pr.conn.WithContext(ctx).Save(&sale).Error; err != nil {
			return admin.CartItemsDomain{}, err
		}
	}

	return record.ToCartItemsDomain(), nil

	// // Hitung TotalAllPrice secara global
	// var allCartItems []CartItems
	// err = pr.conn.WithContext(ctx).
	// 	Find(&allCartItems).Error
	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return admin.CartItemsDomain{}, err
	// }

	// totalAllPrice := record.SubTotal
	// for _, sale := range allCartItems {
	// 	totalAllPrice += sale.Price
	// }
	// record.SubTotal = totalAllPrice

	// // Simpan catatan penjualan
	// result := pr.conn.WithContext(ctx).Create(&record)
	// if err := result.Error; err != nil {
	// 	return admin.CartItemsDomain{}, err
	// }

	// // Perbarui TotalAllPrice untuk semua penjualan lainnya
	// for _, sale := range allCartItems {
	// 	sale.SubTotal = totalAllPrice
	// 	if err := pr.conn.WithContext(ctx).Save(&sale).Error; err != nil {
	// 		return admin.CartItemsDomain{}, err
	// 	}
	// }

}

func (sr *adminRepository) CartItemsGetByID(ctx context.Context, id string) (admin.CartItemsDomain, error) {
	var item CartItems

	if err := sr.conn.WithContext(ctx).Preload("Customers").Preload("Stocks").First(&item, "id = ?", id).Error; err != nil {
		return admin.CartItemsDomain{}, err
	}

	return item.ToCartItemsDomain(), nil

}

// func (sr *adminRepository) CartItemsGetByCustomerID(ctx context.Context, customerId string) (admin.CustomersDomain, error) {
// 	var customer Customers

// 	// if err := sr.conn.WithContext(ctx).Preload("Customers").Preload("Stocks").First(&item, "customer_id = ?", cartItemsDomain.CustomerID).Error; err != nil {
// 	if err := sr.conn.WithContext(ctx).Preload("CartItems").Where(" id = ?", customerId).First(&customer).Error; err != nil {
// 		return admin.CustomersDomain{}, err
// 	}
// 	// if err != nil {
// 	// 	return admin.CartItemsDomain{}, err
// 	// }

// 	return customer.ToCustomersDomain(), nil

// }

// Sulit-Sulit :v
func (sr *adminRepository) CartItemsGetByCustomerID(ctx context.Context, customerId string) ([]admin.CartItemsDomain, error) {
	var cartItems []CartItems

	// if err := sr.conn.WithContext(ctx).Preload("Customers").Preload("Stocks").First(&item, "customer_id = ?", cartItemsDomain.CustomerID).Error; err != nil {
	if err := sr.conn.WithContext(ctx).
		Preload("Customers").
		Preload("Stocks").
		Where("customer_id  = ?", customerId).
		Find(&cartItems).Error; err != nil {
		return nil, err
	}

	cartItemsDomain := make([]admin.CartItemsDomain, len(cartItems))

	for i, purchase := range cartItems {
		// Konversi ke domain
		// domain := purchase.ToCartItemsDomain()
		cartItemsDomain[i] = purchase.ToCartItemsDomain()
	}

	return cartItemsDomain, nil

}

func (sr *adminRepository) CartItemsGetAll(ctx context.Context) ([]admin.CartItemsDomain, error) {
	var records []CartItems
	if err := sr.conn.WithContext(ctx).
		Preload("Customers").Preload("Stocks").Preload("Categories").
		Find(&records).Error; err != nil {
		return nil, err
	}

	cartItemsDomain := []admin.CartItemsDomain{}

	for _, purchase := range records {
		// Konversi ke domain
		domain := purchase.ToCartItemsDomain()
		cartItemsDomain = append(cartItemsDomain, domain)
	}

	return cartItemsDomain, nil
}

func (sr *adminRepository) CartItemsDelete(ctx context.Context, id string) error {
	items, err := sr.CartItemsGetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedItems := FromCartItemsDomain(&items)

	// Ambil semua item keranjang dari pelanggan yang sama
	var customerItems []CartItems
	err = sr.conn.WithContext(ctx).
		Where("customer_id = ?", deletedItems.CustomerID).
		Find(&customerItems).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// Kurangi harga item yang dihapus dari subtotal
	subTotal := 0
	for _, item := range customerItems {
		if item.ID != deletedItems.ID {
			subTotal += item.Price
		}
	}

	// Perbarui subtotal untuk semua item keranjang lainnya dari pelanggan yang sama
	for _, item := range customerItems {
		if item.ID != deletedItems.ID {
			item.SubTotal = subTotal
			if err := sr.conn.WithContext(ctx).Save(&item).Error; err != nil {
				return err
			}
		}
	}

	// Hapus data item keranjang
	if err := sr.conn.WithContext(ctx).Unscoped().Delete(&deletedItems).Error; err != nil {
		return err
	}

	return nil

	// // Ambil semua penjualan lainnya
	// var allItems []CartItems
	// err = sr.conn.WithContext(ctx).Find(&allItems).Error
	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return err
	// }

	// // Kurangi TotalPrice dari penjualan yang dihapus dari TotalAllPrice global
	// totalAllPrice := 0
	// for _, item := range allItems {
	// 	if item.ID != deletedItems.ID {
	// 		totalAllPrice += item.SubTotal
	// 	}
	// }

	// // Perbarui TotalAllPrice untuk semua penjualan lainnya
	// for _, item := range allItems {
	// 	if item.ID != deletedItems.ID {
	// 		item.SubTotal = totalAllPrice
	// 		if err := sr.conn.WithContext(ctx).Save(&item).Error; err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	// // Hapus data penjualan
	// if err := sr.conn.WithContext(ctx).Unscoped().Delete(&deletedItems).Error; err != nil {
	// 	return err
	// }

	// return nil
}

// func (ir *adminRepository) CartsCreate(ctx context.Context, cartDomain *admin.CartsDomain) (admin.CartsDomain, error) {
// 	record := FromCartsDomain(cartDomain)
// 	// Pertama, periksa apakah stok mencukupi
// 	// // Hitung total harga sebelum menyimpan catatan penjualan
// 	// record.Price = record.Quantity * stock.SellingPrice

// 	// Simpan catatan penjualan
// 	result := ir.conn.WithContext(ctx).Create(&record)
// 	if err := result.Error; err != nil {
// 		return admin.CartsDomain{}, err
// 	}
// 	err := ir.conn.WithContext(ctx).Last(&record).Error

// 	if err != nil {
// 		return admin.CartsDomain{}, err
// 	}

// 	return record.ToCartsDomain(), nil
// }

// func (sr *adminRepository) CartsGetByID(ctx context.Context, id string) (admin.CartsDomain, error) {
// 	var cart Carts

// 	if err := sr.conn.WithContext(ctx).Preload("CartItems").First(&cart, "id = ?", id).Error; err != nil {
// 		return admin.CartsDomain{}, err
// 	}

// 	return cart.ToCartsDomain(), nil

// }

// func (sr *adminRepository) CartsGetAll(ctx context.Context) ([]admin.CartsDomain, error) {
// 	var records []Carts
// 	if err := sr.conn.WithContext(ctx).Preload("CartItems").
// 		// if err := sr.conn.WithContext(ctx).
// 		Find(&records).Error; err != nil {
// 		return nil, err
// 	}

// 	cartItemsDomain := []admin.CartsDomain{}

// 	for _, purchase := range records {
// 		// Konversi ke domain
// 		domain := purchase.ToCartsDomain()
// 		cartItemsDomain = append(cartItemsDomain, domain)
// 	}

// 	return cartItemsDomain, nil
// }

// func (sr *adminRepository) CartsDelete(ctx context.Context, id string) error {
// 	items, err := sr.CartsGetByID(ctx, id)

// 	if err != nil {
// 		return err
// 	}

// 	deletedItems := FromCartsDomain(&items)

// 	// Ambil semua penjualan lainnya
// 	var allItems []Carts
// 	err = sr.conn.WithContext(ctx).Find(&allItems).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return err
// 	}

// 	// // Kurangi TotalPrice dari penjualan yang dihapus dari TotalAllPrice global
// 	// totalAllPrice := 0
// 	// for _, item := range allItems {
// 	// 	if item.ID != deletedItems.ID {
// 	// 		totalAllPrice += item.TotalPrice
// 	// 	}
// 	// }

// 	// // Perbarui TotalAllPrice untuk semua penjualan lainnya
// 	// for _, item := range allItems {
// 	// 	if item.ID != deletedItems.ID {
// 	// 		item.TotalAllPrice = totalAllPrice
// 	// 		if err := sr.conn.WithContext(ctx).Save(&item).Error; err != nil {
// 	// 			return err
// 	// 		}
// 	// 	}
// 	// }

// 	// Hapus data penjualan
// 	if err := sr.conn.WithContext(ctx).Unscoped().Delete(&deletedItems).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
