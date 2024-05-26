package mysql_driver

import (
	"backend-golang/drivers/mysql/admin"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func (config *DBConfig) InitDB() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
	}

	log.Println("connected to the database")

	return db
}

func MigrateDB(db *gorm.DB) {
	// err := db.AutoMigrate(&admin.Admins{}, &admin.Customers{}, &admin.Categories{}, &admin.Vendors{}, &admin.Units{}, &admin.Stocks{}, &admin.Purchases{}, &admin.CartItems{}, &admin.Carts{})
	err := db.AutoMigrate(&admin.Admins{}, &admin.Customers{}, &admin.Categories{}, &admin.Vendors{}, &admin.Units{}, &admin.Stocks{}, &admin.Purchases{}, &admin.CartItems{}, &admin.ItemTransactions{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
}

// func ModifyTableCollation(db *gorm.DB) error {
// 	// Jalankan perintah ALTER TABLE
// 	// MODIFY COLUMN stock_name VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin,

// 	err := db.Exec(`
//         ALTER TABLE purchases
//         MODIFY COLUMN category_name VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
//     `).Error
// 	return err
// }

func SeedAdminData(db *gorm.DB) error {
	// Data admin yang ingin dibuat
	adminData := admin.Admins{
		Name:     "admin",
		Voucher:  "admin123",
		Password: "admin12345",
	}

	// Menghasilkan hash password admin
	password, err := bcrypt.GenerateFromPassword([]byte(adminData.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash admin password: %w", err)
	}

	// Mengubah password menjadi hash
	adminData.Password = string(password)

	// Memeriksa apakah admin sudah ada berdasarkan email
	var record admin.Admins
	result := db.First(&record)

	// Jika admin sudah ada, log pesan dan keluar
	if result.Error == nil && record.ID != 0 {
		log.Println("admin already exists")
		return nil
	}

	// Membuat admin baru jika belum ada
	createResult := db.Create(&adminData)
	if createResult.Error != nil {
		return fmt.Errorf("failed to create admin: %w", createResult.Error)
	}

	// Log pesan sukses
	log.Println("admin created")
	return nil
}

func SeedCustomersData(db *gorm.DB) error {
	customersData := []admin.Customers{
		{CustomerName: "Ajax", CustomerEmail: "Ajax@Gmail.com", CustomerAddress: "PT Farm Land", CustomerPhone: "+5412345678901"},
		{CustomerName: "Doss", CustomerEmail: "Doss@Gmail.com", CustomerAddress: "PT Valley", CustomerPhone: "+5212345678901"},
		{CustomerName: "Fred", CustomerEmail: "Fred@Gmail.com", CustomerAddress: "PT Northbridge", CustomerPhone: "+5512345678901"},
		{CustomerName: "Renoir", CustomerEmail: "Renoir@Gmail.com", CustomerAddress: "PT Armory", CustomerPhone: "+5712345678901"},
	}

	var record admin.Customers
	_ = db.First(&record)

	if record.ID != 0 {
		// if record.CategoryName != "" {
		log.Printf("customers detail already exists\n")
	} else {
		for _, customers := range customersData {
			result := db.Create(&customers)
			if result.Error != nil {
				return result.Error
			}
		}
		log.Printf("%d customers detail created\n", len(customersData))
	}

	return nil
}

func SeedCategoryData(db *gorm.DB) error {
	categoryData := []admin.Categories{
		{CategoryName: "Kabel"},
		{CategoryName: "Lampu"},
		{CategoryName: "Contactor"},
		{CategoryName: "MCB"},
		{CategoryName: "Inverter"},
	}

	var record admin.Categories
	_ = db.First(&record)

	if record.ID != 0 {
		// if record.CategoryName != "" {
		log.Printf("category detail already exists\n")
	} else {
		for _, category := range categoryData {
			result := db.Create(&category)
			if result.Error != nil {
				return result.Error
			}
		}
		log.Printf("%d category detail created\n", len(categoryData))
	}

	return nil
}

func SeedVendorsData(db *gorm.DB) error {
	vendorsData := []admin.Vendors{
		{VendorName: "PT Skuy Makmur", VendorAddress: "Tangerang, Jalan Makmur No 22", VendorEmail: "SkuyMakmur@Gmail.com", VendorPhone: "081381814040"},
		{VendorName: "PT Guanzho", VendorAddress: "Wuhan, Covid No 19", VendorEmail: "Guanzho@Gmail.com", VendorPhone: "01230987896"},
		{VendorName: "PT Kuat Perkasa", VendorAddress: "Konoha, JL Ninjaku No 90", VendorEmail: "KuatPerkasa@Gmail.com", VendorPhone: "081567834908"},
	}

	var record admin.Vendors
	_ = db.First(&record)

	if record.ID != 0 {
		log.Printf("vendors detail already exists\n")
	} else {
		for _, vendors := range vendorsData {
			result := db.Create(&vendors)
			if result.Error != nil {
				return result.Error
			}
		}
		log.Printf("%d vendors detail created\n", len(vendorsData))
	}

	return nil
}

func SeedUnitsData(db *gorm.DB) error {
	unitsData := []admin.Units{
		{UnitName: "Pcs"},
		{UnitName: "Pack"},
		{UnitName: "Roll"},
		{UnitName: "Meter"},
	}

	var record admin.Units
	_ = db.First(&record)

	if record.ID != 0 {
		log.Printf("units detail already exists\n")
	} else {
		for _, units_name := range unitsData {
			result := db.Create(&units_name)
			if result.Error != nil {
				return result.Error
			}
		}
		log.Printf("%d units detail created\n", len(unitsData))
	}

	return nil
}

func SeedPurchasesData(db *gorm.DB) error {
	purchasesData := []admin.Purchases{
		{
			VendorID:   1, // Pastikan vendor ini ada di tabel vendors
			StockName:  "Produk A",
			StockCode:  "A001",
			CategoryID: 1, // Pastikan kategori ini ada di tabel categories
			// CategoryName:   "Kabel",
			UnitID: 1, // Pastikan unit ini ada di tabel units
			// UnitName:      "Pcs",
			Description:    "Lorem ipsum dolor sit amet.",
			Quantity:       50,   // Jumlah yang dibeli
			PurchasesPrice: 500,  // Harga beli
			SellingPrice:   1000, // Harga jual
		},
		{
			VendorID:   2, // Pastikan vendor ini ada di tabel vendors
			StockName:  "Produk B",
			StockCode:  "B001",
			CategoryID: 2, // Pastikan kategori ini ada di tabel categories
			// CategoryName:   "Lampu",
			UnitID: 2, // Pastikan unit ini ada di tabel units
			// UnitName:   "Pack",
			Description: "Lorem ipsum dolor sit amet.",

			Quantity:       75,   // Jumlah yang dibeli
			PurchasesPrice: 750,  // Harga beli
			SellingPrice:   2000, // Harga jual
		},
		{
			VendorID:   3, // Pastikan vendor ini ada di tabel vendors
			StockName:  "Produk C",
			StockCode:  "C001",
			CategoryID: 3, // Pastikan kategori ini ada di tabel categories
			// CategoryName:   "Contactor",
			UnitID: 3, // Pastikan unit ini ada di tabel units
			// UnitName:   "Roll",
			Description: "Lorem ipsum dolor sit amet.",

			Quantity:       60,   // Jumlah yang dibeli
			PurchasesPrice: 600,  // Harga beli
			SellingPrice:   1500, // Harga jual
		},
		{
			VendorID:   1, // Pastikan vendor ini ada di tabel vendors
			StockName:  "Produk D",
			StockCode:  "D001",
			CategoryID: 4, // Pastikan kategori ini ada di tabel categories
			// CategoryName:   "MCB",
			UnitID: 4, // Pastikan unit ini ada di tabel units
			// UnitName:   "Meter",
			Description: "Lorem ipsum dolor sit amet.",

			Quantity:       80,   // Jumlah yang dibeli
			PurchasesPrice: 800,  // Harga beli
			SellingPrice:   3000, // Harga jual
		},
		{
			VendorID:   2, // Pastikan vendor ini ada di tabel vendors
			StockName:  "Produk E",
			StockCode:  "E001",
			CategoryID: 5, // Pastikan kategori ini ada di tabel categories
			// CategoryName: "Inverter",
			UnitID: 1, // Pastikan unit ini ada di tabel units
			// UnitName:   "Pcs",
			Description: "Lorem ipsum dolor sit amet.",

			Quantity:       70,   // Jumlah yang dibeli
			PurchasesPrice: 700,  // Harga beli
			SellingPrice:   2500, // Harga jual
		},
	}

	// Masukkan data ke dalam tabel purchases
	for _, purchase := range purchasesData {
		result := db.Create(&purchase)
		if result.Error != nil {
			return result.Error
		}

		// Setelah berhasil membuat pembelian, perbarui atau tambahkan stok
		var stock admin.Stocks
		// err := db.Where("stock_code = ? AND units_name = ?", purchase.StockCode, purchase.UnitName).
		err := db.Where("stock_code = ?", purchase.StockCode).
			First(&stock).Error
		if err == gorm.ErrRecordNotFound {
			// Jika stok belum ada, buat stok baru
			newStock := admin.Stocks{
				StockName:  purchase.StockName,
				StockCode:  purchase.StockCode,
				CategoryID: purchase.CategoryID,
				// CategoryName: purchase.Categories.CategoryName,
				UnitID: purchase.UnitID,
				// UnitName:     purchase.UnitName,
				Description:  purchase.Description,
				StockTotal:   purchase.Quantity, // Jumlah yang dibeli ditambahkan ke stok total
				SellingPrice: purchase.SellingPrice,
			}
			db.Create(&newStock)
		} else {
			// Jika ada kesalahan lain, kembalikan error
			return err
		}
	}
	// else if err == nil {
	// 	// Jika stok sudah ada, perbarui stok total
	// 	stock.StockTotal += purchase.Quantity // Tambahkan jumlah yang dibeli ke stok total
	// 	db.Save(&stock)
	// }

	// Log pesan sukses
	log.Println("Purchases data seeded")
	return nil
}

func SeedCartItemsData(db *gorm.DB) error {
	cartItemsData := []admin.CartItems{
		{
			CustomerID: 1, // Pastikan vendor ini ada di tabel vendors
			StockID:    1,
			Quantity:   1,
			Price:      1000,
			SubTotal:   21000,
		},
		{
			CustomerID: 1, // Pastikan vendor ini ada di tabel vendors
			StockID:    1,
			Quantity:   2,
			Price:      2000,
			SubTotal:   21000,
		},
		{
			CustomerID: 1, // Pastikan vendor ini ada di tabel vendors
			StockID:    1,
			Quantity:   3,
			Price:      3000,
			SubTotal:   21000,
		},
		{
			CustomerID: 1, // Pastikan vendor ini ada di tabel vendors
			StockID:    1,
			Quantity:   4,
			Price:      4000,
			SubTotal:   21000,
		},
		{
			CustomerID: 1, // Pastikan vendor ini ada di tabel vendors
			StockID:    1,
			Quantity:   5,
			Price:      5000,
			SubTotal:   21000,
		},
		{
			CustomerID: 1, // Pastikan vendor ini ada di tabel vendors
			StockID:    1,
			Quantity:   6,
			Price:      6000,
			SubTotal:   21000,
		},

		//

		{
			CustomerID: 2,
			StockID:    2,
			Quantity:   1,
			Price:      2000,
			SubTotal:   42000,
		},

		{
			CustomerID: 2,
			StockID:    2,
			Quantity:   2,
			Price:      4000,
			SubTotal:   42000,
		},

		{
			CustomerID: 2,
			StockID:    2,
			Quantity:   3,
			Price:      6000,
			SubTotal:   42000,
		},
		{
			CustomerID: 2,
			StockID:    2,
			Quantity:   4,
			Price:      8000,
			SubTotal:   42000,
		},
		{
			CustomerID: 2,
			StockID:    2,
			Quantity:   5,
			Price:      10000,
			SubTotal:   42000,
		},
		{
			CustomerID: 2,
			StockID:    2,
			Quantity:   6,
			Price:      12000,
			SubTotal:   42000,
		},
	}

	var record admin.CartItems
	_ = db.First(&record)

	if record.ID != 0 {
		log.Printf("category detail already exists\n")
	} else {
		for _, cartItems := range cartItemsData {
			result := db.Create(&cartItems)
			if result.Error != nil {
				return result.Error
			}
		}
		log.Printf("%d category detail created\n", len(cartItemsData))
	}

	return nil

	// // Masukkan data ke dalam tabel purchases
	// for _, cartItems := range cartItemsData {
	// 	result := db.Create(&cartItems)
	// 	if result.Error != nil {
	// 		return result.Error
	// 	}

	// 	// Setelah berhasil membuat pembelian, perbarui atau tambahkan stok

	// 	// err := db.Where("stock_code = ? AND units_name = ?", purchase.StockCode, purchase.UnitName).
	// 	err := db.Where(" stock_id = ? AND customer_id", cartItems.CustomerID, cartItems.StockID).
	// 		First(&stock).Error
	// 	if err == gorm.ErrRecordNotFound {
	// 		// Jika stok belum ada, buat stok baru
	// 		newStock := admin.Stocks{
	// 			CustomerID:  purchase.StockName,
	// 			StockID:  purchase.StockCode,
	// 			Quantity: purchase.CategoryID,
	// 			// CategoryName: purchase.Categories.CategoryName,
	// 			UnitID: purchase.UnitID,
	// 			// UnitName:     purchase.UnitName,
	// 			Description:  purchase.Description,
	// 			StockTotal:   purchase.Quantity, // Jumlah yang dibeli ditambahkan ke stok total
	// 			SellingPrice: purchase.SellingPrice,
	// 		}
	// 		db.Create(&newStock)
	// 	} else {
	// 		// Jika ada kesalahan lain, kembalikan error
	// 		return err
	// 	}
	// }
	// // else if err == nil {
	// // 	// Jika stok sudah ada, perbarui stok total
	// // 	stock.StockTotal += purchase.Quantity // Tambahkan jumlah yang dibeli ke stok total
	// // 	db.Save(&stock)
	// // }

	// // Log pesan sukses
	// log.Println("Purchases data seeded")
	// return nil
}

func SeedStocksData(db *gorm.DB) error {
	// {CategoryName: "Kabel"},
	// {CategoryName: "Lampu"},
	// {CategoryName: "Contactor"},
	// {CategoryName: "MCB"},
	// {CategoryName: "Inverter"},
	stocksData := []admin.Stocks{
		{
			StockName: "Produk A",
			StockCode: "A001",
			// CategoryID:    1,    // Pastikan kategori ini ada di tabel categories
			// CategoryName:  "Kabel",
			UnitID:       1,    // Pastikan unit ini ada di tabel units
			StockTotal:   100,  // Jumlah stok total
			SellingPrice: 1000, // Harga jual
		},
		{
			StockName: "Produk B",
			StockCode: "B001",
			// CategoryID:    2,    // Pastikan kategori ini ada di tabel categories
			// CategoryName:  "Lampu",
			UnitID:       2,    // Pastikan unit ini ada di tabel units
			StockTotal:   200,  // Jumlah stok total
			SellingPrice: 2000, // Harga jual
		},
		{
			StockName: "Produk C",
			StockCode: "C001",
			// CategoryID:    3,    // Pastikan kategori ini ada di tabel categories
			// CategoryName:  "Contactor",
			UnitID:       3,    // Pastikan unit ini ada di tabel units
			StockTotal:   150,  // Jumlah stok total
			SellingPrice: 1500, // Harga jual
		},
		{
			StockName: "Produk D",
			StockCode: "D001",
			// CategoryID:    4,    // Pastikan kategori ini ada di tabel categories
			// CategoryName:  "MCB",
			UnitID:       4,    // Pastikan unit ini ada di tabel units
			StockTotal:   300,  // Jumlah stok total
			SellingPrice: 3000, // Harga jual
		},
		{
			StockName: "Produk E",
			StockCode: "E001",
			// CategoryID:    5,    // Pastikan kategori ini ada di tabel categories
			// CategoryName:  "Inverter",
			UnitID:       4,    // Pastikan unit ini ada di tabel units
			StockTotal:   250,  // Jumlah stok total
			SellingPrice: 2500, // Harga jual
		},
	}

	var record admin.Stocks
	_ = db.First(&record)

	if record.ID != 0 {
		log.Printf("units detail already exists\n")
	} else {
		for _, stocks := range stocksData {
			result := db.Create(&stocks)
			if result.Error != nil {
				return result.Error
			}
		}
		log.Printf("%d units detail created\n", len(stocksData))
	}

	return nil
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()

	if err != nil {
		log.Printf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection: %v", err)
		return err
	}

	log.Println("database connection is closed")

	return nil
}
