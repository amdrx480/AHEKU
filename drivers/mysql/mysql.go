package mysql_driver

import (
	"backend-golang/drivers/mysql/admin"
	"errors"
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
	// err := db.AutoMigrate(&admin.Admin{}, &admin.Customers{}, &admin.Categories{}, &admin.Vendors{}, &admin.Units{}, &admin.Stocks{}, &admin.Purchases{}, &admin.CartItems{}, &admin.Carts{})
	err := db.AutoMigrate(&admin.Admin{}, &admin.Role{}, &admin.Customers{}, &admin.PackagingOfficer{}, &admin.Categories{}, &admin.Vendors{}, &admin.Units{}, &admin.Stocks{}, &admin.Purchases{}, &admin.CartItems{}, &admin.ItemTransactions{}, &admin.ReminderPurchaseOrder{})
	// err := db.AutoMigrate(&admin.Admin{}, &admin.Role{}, &admin.Customers{}, &admin.PackagingOfficer{}, &admin.Categories{}, &admin.Vendors{}, &admin.Units{}, &admin.Stocks{}, &admin.Purchases{}, &admin.CartItems{}, &admin.ItemTransactions{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}

}

func SeedAdminData(db *gorm.DB) error {
	// Periksa apakah tabel 'admins' sudah ada
	if !db.Migrator().HasTable("admins") {
		// Jika belum ada, buat tabel 'admins' terlebih dahulu
		if err := db.AutoMigrate(&admin.Admin{}); err != nil {
			return fmt.Errorf("failed to perform migration for admins table: %w", err)
		}
	}

	// Memeriksa apakah admin sudah ada dalam database
	var existingAdmin []admin.Admin
	result := db.Find(&existingAdmin)
	if result.Error != nil {
		return fmt.Errorf("failed to query admin: %v", result.Error)
	}

	// Jika admin sudah ada, log pesan dan keluar
	if len(existingAdmin) > 0 {
		log.Println("admin already exist")
		return nil
	}

	// Data admin yang ingin dibuat
	admins := []admin.Admin{
		{Name: "superadmin", Email: "superadmin@example.com", Phone: "081382815860", RoleID: 1, Voucher: "superadmin123", Password: "superadmin12345"},
		{Name: "admin1", Email: "admin1@example.com", Phone: "987654321", RoleID: 2, Voucher: "admin1123", Password: "admin112345"},
		{Name: "admin2", Email: "admin2@example.com", Phone: "987654322", RoleID: 2, Voucher: "admin2123", Password: "admin212345"},
	}

	// Loop melalui setiap admin yang ingin ditambahkan
	for _, adminData := range admins {
		// Menghasilkan hash password admin
		password, err := bcrypt.GenerateFromPassword([]byte(adminData.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash admin password: %w", err)
		}

		// Mengubah password menjadi hash
		adminData.Password = string(password)

		// Membuat admin baru
		createResult := db.Create(&adminData)
		if createResult.Error != nil {
			return fmt.Errorf("failed to create admin: %w", createResult.Error)
		}

		// Log pesan sukses
		log.Printf("admin %s created\n", adminData.Name)
	}

	return nil
}

func SeedRolesData(db *gorm.DB) error {
	// Periksa apakah tabel 'roles' sudah ada
	if !db.Migrator().HasTable("roles") {
		// Jika belum ada, buat tabel 'roles' terlebih dahulu
		if err := db.AutoMigrate(&admin.Role{}); err != nil {
			return fmt.Errorf("failed to perform migration for roles table: %v", err)
		}
	}

	// Memeriksa apakah roles sudah ada dalam database
	var existingRoles []admin.Role
	result := db.Find(&existingRoles)
	if result.Error != nil {
		return fmt.Errorf("failed to query roles: %v", result.Error)
	}

	// Jika roles sudah ada, log pesan dan keluar
	if len(existingRoles) > 0 {
		log.Println("roles already exist")
		return nil
	}

	// Data role yang ingin dibuat
	roleData := []admin.Role{
		{RoleName: "Super Admin"},
		{RoleName: "Admin"},
		{RoleName: "User"},
	}

	// Loop melalui setiap role yang ingin ditambahkan
	for _, role := range roleData {
		// Membuat role baru
		createResult := db.Create(&role)
		if createResult.Error != nil {
			return fmt.Errorf("failed to create role: %v", createResult.Error)
		}

		// Log pesan sukses
		log.Printf("role %s created\n", role.RoleName)
	}

	return nil
}

func SeedCustomersData(db *gorm.DB) error {
	// Periksa apakah tabel 'customers' sudah ada
	if !db.Migrator().HasTable("customers") {
		// Jika belum ada, buat tabel 'customers' terlebih dahulu
		if err := db.AutoMigrate(&admin.Customers{}); err != nil {
			return fmt.Errorf("failed to perform migration for customers table: %w", err)
		}
	}

	// Memeriksa apakah customers sudah ada dalam database
	var existingCustomers []admin.Customers
	result := db.Find(&existingCustomers)
	if result.Error != nil {
		return fmt.Errorf("failed to query customers: %v", result.Error)
	}

	// Jika customers sudah ada, log pesan dan keluar
	if len(existingCustomers) > 0 {
		log.Println("customers already exist")
		return nil
	}

	// Data customers yang ingin dibuat
	customersData := []admin.Customers{
		{CustomerName: "Ajax", CustomerEmail: "Ajax@Gmail.com", CustomerAddress: "PT Farm Land", CustomerPhone: "+5412345678901"},
		{CustomerName: "Doss", CustomerEmail: "Doss@Gmail.com", CustomerAddress: "PT Valley", CustomerPhone: "+5212345678901"},
		{CustomerName: "Fred", CustomerEmail: "Fred@Gmail.com", CustomerAddress: "PT Northbridge", CustomerPhone: "+5512345678901"},
		{CustomerName: "Renoir", CustomerEmail: "Renoir@Gmail.com", CustomerAddress: "PT Armory", CustomerPhone: "+5712345678901"},
	}

	// Loop melalui setiap customer yang ingin ditambahkan
	for _, customer := range customersData {
		// Membuat customer baru
		createResult := db.Create(&customer)
		if createResult.Error != nil {
			return fmt.Errorf("failed to create customer: %v", createResult.Error)
		}

		// Log pesan sukses
		log.Printf("customer %s created\n", customer.CustomerName)
	}

	return nil
}

func SeedCategoryData(db *gorm.DB) error {
	// Periksa apakah tabel 'categories' sudah ada
	if !db.Migrator().HasTable("categories") {
		// Jika belum ada, buat tabel 'categories' terlebih dahulu
		if err := db.AutoMigrate(&admin.Categories{}); err != nil {
			return fmt.Errorf("failed to perform migration for categories table: %v", err)
		}
	}

	// Memeriksa apakah categories sudah ada dalam database
	var existingCategories []admin.Categories
	result := db.Find(&existingCategories)
	if result.Error != nil {
		return fmt.Errorf("failed to query categories: %v", result.Error)
	}

	// Jika categories sudah ada, log pesan dan keluar
	if len(existingCategories) > 0 {
		log.Println("categories already exist")
		return nil
	}

	// Data kategori yang ingin dibuat
	categoryData := []admin.Categories{
		{CategoryName: "Kabel"},
		{CategoryName: "Lampu"},
		{CategoryName: "Contactor"},
		{CategoryName: "MCB"},
		{CategoryName: "Inverter"},
	}

	// Loop melalui setiap kategori yang ingin ditambahkan
	for _, category := range categoryData {
		// Membuat kategori baru
		createResult := db.Create(&category)
		if createResult.Error != nil {
			return fmt.Errorf("failed to create category: %v", createResult.Error)
		}

		// Log pesan sukses
		log.Printf("category %s created\n", category.CategoryName)
	}

	return nil
}

func SeedVendorsData(db *gorm.DB) error {
	// Periksa apakah tabel 'vendors' sudah ada
	if !db.Migrator().HasTable("vendors") {
		// Jika belum ada, buat tabel 'vendors' terlebih dahulu
		if err := db.AutoMigrate(&admin.Vendors{}); err != nil {
			return fmt.Errorf("failed to perform migration for vendors table: %v", err)
		}
	}

	// Memeriksa apakah vendors sudah ada dalam database
	var existingVendors []admin.Vendors
	result := db.Find(&existingVendors)
	if result.Error != nil {
		return fmt.Errorf("failed to query vendors: %v", result.Error)
	}

	// Jika vendors sudah ada, log pesan dan keluar
	if len(existingVendors) > 0 {
		log.Println("vendors already exist")
		return nil
	}

	// Data vendor yang ingin dibuat
	vendorsData := []admin.Vendors{
		{VendorName: "PT Skuy Makmur", VendorAddress: "Tangerang, Jalan Makmur No 22", VendorEmail: "SkuyMakmur@Gmail.com", VendorPhone: "081381814040"},
		{VendorName: "PT Guanzho", VendorAddress: "Wuhan, Covid No 19", VendorEmail: "Guanzho@Gmail.com", VendorPhone: "01230987896"},
		{VendorName: "PT Kuat Perkasa", VendorAddress: "Konoha, JL Ninjaku No 90", VendorEmail: "KuatPerkasa@Gmail.com", VendorPhone: "081567834908"},
	}

	// Loop melalui setiap vendor yang ingin ditambahkan
	for _, vendor := range vendorsData {
		// Membuat vendor baru
		createResult := db.Create(&vendor)
		if createResult.Error != nil {
			return fmt.Errorf("failed to create vendor: %v", createResult.Error)
		}

		// Log pesan sukses
		log.Printf("vendor %s created\n", vendor.VendorName)
	}

	return nil
}

func SeedUnitsData(db *gorm.DB) error {
	// Periksa apakah tabel 'units' sudah ada
	if !db.Migrator().HasTable("units") {
		// Jika belum ada, buat tabel 'units' terlebih dahulu
		if err := db.AutoMigrate(&admin.Units{}); err != nil {
			return fmt.Errorf("failed to perform migration for units table: %v", err)
		}
	}

	// Memeriksa apakah units sudah ada dalam database
	var existingUnits []admin.Units
	result := db.Find(&existingUnits)
	if result.Error != nil {
		return fmt.Errorf("failed to query units: %v", result.Error)
	}

	// Jika units sudah ada, log pesan dan keluar
	if len(existingUnits) > 0 {
		log.Println("units already exist")
		return nil
	}

	// Data unit yang ingin dibuat
	unitsData := []admin.Units{
		{UnitName: "Pcs"},
		{UnitName: "Pack"},
		{UnitName: "Roll"},
		{UnitName: "Meter"},
	}

	// Loop melalui setiap unit yang ingin ditambahkan
	for _, unit := range unitsData {
		// Membuat unit baru
		createResult := db.Create(&unit)
		if createResult.Error != nil {
			return fmt.Errorf("failed to create unit: %v", createResult.Error)
		}

		// Log pesan sukses
		log.Printf("unit %s created\n", unit.UnitName)
	}

	return nil
}

func SeedPurchasesData(db *gorm.DB) error {
	// Periksa apakah tabel 'purchases' dan 'stocks' sudah ada
	purchasesTableExists := db.Migrator().HasTable(&admin.Purchases{})
	stocksTableExists := db.Migrator().HasTable(&admin.Stocks{})

	// Jika tabel tidak ada, buat tabel 'purchases' dan 'stocks' terlebih dahulu
	if !purchasesTableExists {
		if err := db.AutoMigrate(&admin.Purchases{}); err != nil {
			return fmt.Errorf("failed to perform migration for purchases table: %w", err)
		}
	}
	if !stocksTableExists {
		if err := db.AutoMigrate(&admin.Stocks{}); err != nil {
			return fmt.Errorf("failed to perform migration for stocks table: %w", err)
		}
	}

	// Memeriksa apakah data purchases dan stocks sudah ada dalam database
	var existingPurchases []admin.Purchases
	var existingStocks []admin.Stocks
	resultPurchases := db.Find(&existingPurchases)
	if resultPurchases.Error != nil {
		return fmt.Errorf("failed to query purchases: %v", resultPurchases.Error)
	}
	resultStocks := db.Find(&existingStocks)
	if resultStocks.Error != nil {
		return fmt.Errorf("failed to query stocks: %v", resultStocks.Error)
	}

	// Jika tabel purchases dan stocks sudah ada tetapi kosong, maka isi SeedPurchasesData
	if len(existingPurchases) == 0 && len(existingStocks) == 0 {
		purchasesData := []admin.Purchases{
			{
				VendorID:    1, // Pastikan vendor ini ada di tabel vendors
				StockName:   "Produk A",
				StockCode:   "A001",
				CategoryID:  1, // Pastikan kategori ini ada di tabel categories
				UnitID:      1, // Pastikan unit ini ada di tabel units
				Description: "Lorem ipsum dolor sit amet.",
				Quantity:    50, // Jumlah yang dibeli
				// Quantity:      100000000000000, // Jumlah yang dibeli
				PurchasePrice: 500,  // Harga beli
				SellingPrice:  1000, // Harga jual
			},
			{
				VendorID:    2, // Pastikan vendor ini ada di tabel vendors
				StockName:   "Produk B",
				StockCode:   "B001",
				CategoryID:  2, // Pastikan kategori ini ada di tabel categories
				UnitID:      2, // Pastikan unit ini ada di tabel units
				Description: "Lorem ipsum dolor sit amet.",
				Quantity:    75, // Jumlah yang dibeli
				// Quantity:      100000000000000, // Jumlah yang dibeli
				PurchasePrice: 750,  // Harga beli
				SellingPrice:  2000, // Harga jual
			},
			// Tambahkan data lainnya sesuai kebutuhan...
			{
				VendorID:      3,
				StockName:     "Produk C",
				StockCode:     "C001",
				CategoryID:    3,
				UnitID:        3,
				Description:   "Lorem ipsum dolor sit amet.",
				Quantity:      60,
				PurchasePrice: 600,
				SellingPrice:  1500,
			},
			{
				VendorID:      1,
				StockName:     "Produk D",
				StockCode:     "D001",
				CategoryID:    4,
				UnitID:        4,
				Description:   "Lorem ipsum dolor sit amet.",
				Quantity:      80,
				PurchasePrice: 800,
				SellingPrice:  3000,
			},
			{
				VendorID:      2,
				StockName:     "Produk E",
				StockCode:     "E001",
				CategoryID:    5,
				UnitID:        1,
				Description:   "Lorem ipsum dolor sit amet.",
				Quantity:      70,
				PurchasePrice: 700,
				SellingPrice:  2500,
			},
		}

		// Loop melalui setiap purchase yang ingin ditambahkan
		for _, purchase := range purchasesData {
			// Membuat purchase baru
			createResult := db.Create(&purchase)
			if createResult.Error != nil {
				return fmt.Errorf("failed to create purchase: %v", createResult.Error)
			}

			// Periksa stok berdasarkan StockID
			var stock admin.Stocks
			stockResult := db.First(&stock, "stock_code = ?", purchase.StockCode)
			if stockResult.Error != nil {
				if errors.Is(stockResult.Error, gorm.ErrRecordNotFound) {
					// Jika stok belum ada, buat stok baru
					newStock := admin.Stocks{
						StockName:    purchase.StockName,
						StockCode:    purchase.StockCode,
						CategoryID:   purchase.CategoryID,
						UnitID:       purchase.UnitID,
						StockTotal:   purchase.Quantity,
						SellingPrice: purchase.SellingPrice,
					}
					createStockResult := db.Create(&newStock)
					if createStockResult.Error != nil {
						return fmt.Errorf("failed to create new stock: %v", createStockResult.Error)
					}
				} else {
					return fmt.Errorf("failed to query stock: %v", stockResult.Error)
				}
			}

			// Log pesan sukses
			log.Printf("purchase %s created and stock processed\n", purchase.StockName)
		}
		return nil
	}

	// Jika tabel purchases dan stocks sudah ada dan data tidak kosong, log pesan dan keluar
	log.Println("purchases and stocks already exist, skipping SeedPurchasesData")
	return nil
}

func SeedCartItemsData(db *gorm.DB) error {
	// Periksa apakah tabel 'cart_items' sudah ada
	if !db.Migrator().HasTable("cart_items") {
		// Jika belum ada, buat tabel 'cart_items' terlebih dahulu
		if err := db.AutoMigrate(&admin.CartItems{}); err != nil {
			return fmt.Errorf("failed to perform migration for cart_items table: %v", err)
		}
	}

	// Memeriksa apakah item keranjang sudah ada dalam database
	var existingCartItems []admin.CartItems
	result := db.Find(&existingCartItems)
	if result.Error != nil {
		return fmt.Errorf("failed to query cart items: %v", result.Error)
	}

	// Jika item keranjang sudah ada, log pesan dan keluar
	if len(existingCartItems) > 0 {
		log.Println("cart items already exist")
		return nil
	}

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

	// Loop melalui setiap item keranjang yang ingin ditambahkan
	for _, cartItem := range cartItemsData {
		// Membuat item keranjang baru
		createResult := db.Create(&cartItem)
		if createResult.Error != nil {
			return fmt.Errorf("failed to create cart item: %v", createResult.Error)
		}

		// Log pesan sukses
		log.Printf("cart item for stock %d created\n", cartItem.StockID)
	}

	return nil
}

func SeedStocksData(db *gorm.DB) error {
	// Periksa apakah tabel 'stocks' sudah ada
	if !db.Migrator().HasTable("stocks") {
		// Jika belum ada, buat tabel 'stocks' terlebih dahulu
		if err := db.AutoMigrate(&admin.Stocks{}); err != nil {
			return fmt.Errorf("failed to perform migration for stocks table: %v", err)
		}
	}

	// Memeriksa apakah stok sudah ada dalam database
	var existingStocks []admin.Stocks
	result := db.Find(&existingStocks)
	if result.Error != nil {
		return fmt.Errorf("failed to query stocks: %v", result.Error)
	}

	// Jika stok sudah ada, log pesan dan keluar
	if len(existingStocks) > 0 {
		log.Println("stocks already exist")
		return nil
	}

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

	// Loop melalui setiap stok yang ingin ditambahkan
	for _, stock := range stocksData {
		// Membuat stok baru
		createResult := db.Create(&stock)
		if createResult.Error != nil {
			return fmt.Errorf("failed to create stock: %v", createResult.Error)
		}

		// Log pesan sukses
		log.Printf("stock %s created\n", stock.StockName)
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
