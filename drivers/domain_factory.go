package drivers

import (
	adminDomain "backend-golang/businesses/admin"
	adminDB "backend-golang/drivers/mysql/admin"

	historyDomain "backend-golang/businesses/history"
	historyDB "backend-golang/drivers/mysql/history"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMySQLRepository(conn)
}

func NewHistoryRepository(conn *gorm.DB) historyDomain.Repository {
	return historyDB.NewMySQLRepository(conn)
}
