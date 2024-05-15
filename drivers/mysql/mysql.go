package mysql_driver

import (
	"backend-golang/drivers/mysql/category"
	"backend-golang/drivers/mysql/history"
	"backend-golang/drivers/mysql/purchases"
	"backend-golang/drivers/mysql/sales"
	"backend-golang/drivers/mysql/units"

	"backend-golang/drivers/mysql/stocks"
	"backend-golang/drivers/mysql/users"
	"backend-golang/drivers/mysql/vendors"
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
	err := db.AutoMigrate(&users.User{}, &stocks.Stock{}, &purchases.Purchase{}, &sales.Sales{}, &vendors.Vendors{}, &category.Category{}, &units.Units{}, &history.History{})
	// err := db.AutoMigrate(&users.User{}, &stocks.Stock{}, &sales.Sales{}, &vendors.Vendors{}, &category.Category{}, &units.Units{})

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
	adminData := users.User{
		Name:     "irwan",
		Password: "admin",
	}

	// Menghasilkan hash password admin
	password, err := bcrypt.GenerateFromPassword([]byte(adminData.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash admin password: %w", err)
	}

	// Mengubah password menjadi hash
	adminData.Password = string(password)

	// Memeriksa apakah admin sudah ada berdasarkan email
	var record users.User
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

func SeedVendorsData(db *gorm.DB) error {
	vendorsData := []vendors.Vendors{
		{Vendor_Name: "PT Skuy Makmur", Vendor_Address: "Tangerang, Jalan Makmur No 22", Vendor_Email: "SkuyMakmur@Gmail.com", Vendor_Phone: "081381814040"},
		{Vendor_Name: "PT Guanzho", Vendor_Address: "Wuhan, Covid No 19", Vendor_Email: "Guanzho@Gmail.com", Vendor_Phone: "01230987896"},
		{Vendor_Name: "PT Kuat Perkasa", Vendor_Address: "Konoha, JL Ninjaku No 90", Vendor_Email: "KuatPerkasa@Gmail.com", Vendor_Phone: "081567834908"},
	}

	var record vendors.Vendors
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

func SeedCategoryData(db *gorm.DB) error {
	categoryData := []category.Category{
		{CategoryName: "Kabel"},
		{CategoryName: "Lampu"},
		{CategoryName: "Contactor"},
		{CategoryName: "MCB"},
		{CategoryName: "Inverter"},
	}

	var record category.Category
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

func SeedUnitsData(db *gorm.DB) error {
	unitsData := []units.Units{
		{UnitsName: "Pcs"},
		{UnitsName: "Pack"},
		{UnitsName: "Roll"},
		{UnitsName: "Meter"},
	}

	var record units.Units
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
	purchasesData := []purchases.Purchase{
		{
			VendorID:   1, // Pastikan vendor ini ada di tabel vendors
			Stock_Name: "Produk A",
			Stock_Code: "A001",
			CategoryID: 1, // Pastikan kategori ini ada di tabel categories
			// CategoryName:   "Kabel",
			UnitsID:        1, // Pastikan unit ini ada di tabel units
			UnitsName:      "Pcs",
			Description:    "Lorem ipsum dolor sit amet, consectetur adipiscing eli.",
			Quantity:       50,   // Jumlah yang dibeli
			Purchase_Price: 500,  // Harga beli
			Selling_Price:  1000, // Harga jual
		},
		{
			VendorID:   2, // Pastikan vendor ini ada di tabel vendors
			Stock_Name: "Produk B",
			Stock_Code: "B001",
			CategoryID: 2, // Pastikan kategori ini ada di tabel categories
			// CategoryName:   "Lampu",
			UnitsID:     2, // Pastikan unit ini ada di tabel units
			UnitsName:   "Pack",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing eli.",

			Quantity:       75,   // Jumlah yang dibeli
			Purchase_Price: 750,  // Harga beli
			Selling_Price:  2000, // Harga jual
		},
		{
			VendorID:   3, // Pastikan vendor ini ada di tabel vendors
			Stock_Name: "Produk C",
			Stock_Code: "C001",
			CategoryID: 3, // Pastikan kategori ini ada di tabel categories
			// CategoryName:   "Contactor",
			UnitsID:     3, // Pastikan unit ini ada di tabel units
			UnitsName:   "Roll",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing eli.",

			Quantity:       60,   // Jumlah yang dibeli
			Purchase_Price: 600,  // Harga beli
			Selling_Price:  1500, // Harga jual
		},
		{
			VendorID:   1, // Pastikan vendor ini ada di tabel vendors
			Stock_Name: "Produk D",
			Stock_Code: "D001",
			CategoryID: 4, // Pastikan kategori ini ada di tabel categories
			// CategoryName:   "MCB",
			UnitsID:     4, // Pastikan unit ini ada di tabel units
			UnitsName:   "Meter",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing eli.",

			Quantity:       80,   // Jumlah yang dibeli
			Purchase_Price: 800,  // Harga beli
			Selling_Price:  3000, // Harga jual
		},
		{
			VendorID:   2, // Pastikan vendor ini ada di tabel vendors
			Stock_Name: "Produk E",
			Stock_Code: "E001",
			CategoryID: 5, // Pastikan kategori ini ada di tabel categories
			// CategoryName: "Inverter",
			UnitsID:     1, // Pastikan unit ini ada di tabel units
			UnitsName:   "Pcs",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing eli.",

			Quantity:       70,   // Jumlah yang dibeli
			Purchase_Price: 700,  // Harga beli
			Selling_Price:  2500, // Harga jual
		},
	}

	// Masukkan data ke dalam tabel purchases
	for _, purchase := range purchasesData {
		result := db.Create(&purchase)
		if result.Error != nil {
			return result.Error
		}

		// Setelah berhasil membuat pembelian, perbarui atau tambahkan stok
		var stock stocks.Stock
		// err := db.Where("stock_code = ? AND units_name = ?", purchase.Stock_Code, purchase.UnitsName).
		err := db.Where("stock_code = ?", purchase.Stock_Code).
			First(&stock).Error
		if err == gorm.ErrRecordNotFound {
			// Jika stok belum ada, buat stok baru
			newStock := stocks.Stock{
				Stock_Name: purchase.Stock_Name,
				Stock_Code: purchase.Stock_Code,
				// CategoryName:  purchase.CategoryName,
				CategoryID:    purchase.CategoryID,
				UnitsID:       purchase.UnitsID,
				UnitsName:     purchase.UnitsName,
				Description:   purchase.Description,
				Stock_Total:   purchase.Quantity, // Jumlah yang dibeli ditambahkan ke stok total
				Selling_Price: purchase.Selling_Price,
			}
			db.Create(&newStock)
		} else {
			// Jika ada kesalahan lain, kembalikan error
			return err
		}
	}
	// else if err == nil {
	// 	// Jika stok sudah ada, perbarui stok total
	// 	stock.Stock_Total += purchase.Quantity // Tambahkan jumlah yang dibeli ke stok total
	// 	db.Save(&stock)
	// }

	// Log pesan sukses
	log.Println("Purchases data seeded")
	return nil
}

func SeedStocksData(db *gorm.DB) error {
	// {CategoryName: "Kabel"},
	// {CategoryName: "Lampu"},
	// {CategoryName: "Contactor"},
	// {CategoryName: "MCB"},
	// {CategoryName: "Inverter"},
	stocksData := []stocks.Stock{
		{
			Stock_Name: "Produk A",
			Stock_Code: "A001",
			// CategoryID:    1,    // Pastikan kategori ini ada di tabel categories
			// CategoryName:  "Kabel",
			UnitsID:       1,    // Pastikan unit ini ada di tabel units
			Stock_Total:   100,  // Jumlah stok total
			Selling_Price: 1000, // Harga jual
		},
		{
			Stock_Name: "Produk B",
			Stock_Code: "B001",
			// CategoryID:    2,    // Pastikan kategori ini ada di tabel categories
			// CategoryName:  "Lampu",
			UnitsID:       2,    // Pastikan unit ini ada di tabel units
			Stock_Total:   200,  // Jumlah stok total
			Selling_Price: 2000, // Harga jual
		},
		{
			Stock_Name: "Produk C",
			Stock_Code: "C001",
			// CategoryID:    3,    // Pastikan kategori ini ada di tabel categories
			// CategoryName:  "Contactor",
			UnitsID:       3,    // Pastikan unit ini ada di tabel units
			Stock_Total:   150,  // Jumlah stok total
			Selling_Price: 1500, // Harga jual
		},
		{
			Stock_Name: "Produk D",
			Stock_Code: "D001",
			// CategoryID:    4,    // Pastikan kategori ini ada di tabel categories
			// CategoryName:  "MCB",
			UnitsID:       4,    // Pastikan unit ini ada di tabel units
			Stock_Total:   300,  // Jumlah stok total
			Selling_Price: 3000, // Harga jual
		},
		{
			Stock_Name: "Produk E",
			Stock_Code: "E001",
			// CategoryID:    5,    // Pastikan kategori ini ada di tabel categories
			// CategoryName:  "Inverter",
			UnitsID:       4,    // Pastikan unit ini ada di tabel units
			Stock_Total:   250,  // Jumlah stok total
			Selling_Price: 2500, // Harga jual
		},
	}

	var record stocks.Stock
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
