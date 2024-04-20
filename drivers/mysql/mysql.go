package mysql_driver

import (
	"backend-golang/drivers/mysql/category"
	"backend-golang/drivers/mysql/purchases"
	"backend-golang/drivers/mysql/units"

	"backend-golang/drivers/mysql/stocks"
	"backend-golang/drivers/mysql/users"
	"backend-golang/drivers/mysql/vendors"
	"fmt"
	"log"

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
	err := db.AutoMigrate(&users.User{}, &stocks.Stock{}, &purchases.Purchase{}, &vendors.Vendors{}, &category.Category{}, &units.Units{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
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
		{Units: "Pcs"},
		{Units: "Pack"},
		{Units: "Roll"},
		{Units: "Meter"},
	}

	var record units.Units
	_ = db.First(&record)

	if record.ID != 0 {
		log.Printf("units detail already exists\n")
	} else {
		for _, units := range unitsData {
			result := db.Create(&units)
			if result.Error != nil {
				return result.Error
			}
		}
		log.Printf("%d units detail created\n", len(unitsData))
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
