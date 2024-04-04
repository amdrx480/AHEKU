package drivers

import (
	userDomain "backend-golang/businesses/users"
	userDB "backend-golang/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
