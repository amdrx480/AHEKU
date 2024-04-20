package drivers

import (
	userDomain "backend-golang/businesses/users"
	userDB "backend-golang/drivers/mysql/users"

	stockDomain "backend-golang/businesses/stocks"
	stockDB "backend-golang/drivers/mysql/stocks"

	purchasesDomain "backend-golang/businesses/purchases"
	purchasesDB "backend-golang/drivers/mysql/purchases"

	vendorsDomain "backend-golang/businesses/vendors"
	vendorsDB "backend-golang/drivers/mysql/vendors"

	categoryDomain "backend-golang/businesses/category"
	categoryDB "backend-golang/drivers/mysql/category"

	unitsDomain "backend-golang/businesses/units"
	unitsDB "backend-golang/drivers/mysql/units"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewStockRepository(conn *gorm.DB) stockDomain.Repository {
	return stockDB.NewMySQLRepository(conn)
}

func NewPurchasesRepository(conn *gorm.DB) purchasesDomain.Repository {
	return purchasesDB.NewMySQLRepository(conn)
}

func NewVendorsRepository(conn *gorm.DB) vendorsDomain.Repository {
	return vendorsDB.NewMySQLRepository(conn)
}

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewUnitsRepository(conn *gorm.DB) unitsDomain.Repository {
	return unitsDB.NewMySQLRepository(conn)
}
