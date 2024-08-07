package admin

import (
	"backend-golang/businesses/admin"
	"backend-golang/utils"
	"context"
	"fmt"
	"log"
	"time"

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

func (ur *adminRepository) AdminRegister(ctx context.Context, adminDomain *admin.AdminDomain) (admin.AdminDomain, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(adminDomain.Password), bcrypt.DefaultCost)
	if err != nil {
		return admin.AdminDomain{}, err
	}

	// Membuat Admin dari AdminDomain
	record := FromAdminDomain(adminDomain)
	record.Password = string(password)

	// Simpan Admin
	result := ur.conn.WithContext(ctx).Create(&record)
	if result.Error != nil {
		return admin.AdminDomain{}, result.Error
	}

	// Mengambil data Admin terbaru dengan Role terkait
	err = ur.conn.WithContext(ctx).Preload("Role").Last(&record).Error
	if err != nil {
		return admin.AdminDomain{}, err
	}

	return record.ToAdminDomain(), nil
}

// func (ur *adminRepository) AdminRegister(ctx context.Context, adminDomain *admin.AdminDomain) (admin.AdminDomain, error) {
// 	// Membuat AdminProfile dari AdminDomain
// 	adminProfile := FromAdminProfileDomain(&adminDomain.AdminProfile)

// 	// Set nama admin di profil admin
// 	adminProfile.Name = adminDomain.Name

// 	// Simpan AdminProfile terlebih dahulu
// 	profileResult := ur.conn.WithContext(ctx).Create(&adminProfile)
// 	if profileResult.Error != nil {
// 		return admin.AdminDomain{}, profileResult.Error
// 	}

// 	password, err := bcrypt.GenerateFromPassword([]byte(adminDomain.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return admin.AdminDomain{}, err
// 	}

// 	// Set AdminProfileID dengan ID AdminProfile yang baru disimpan
// 	adminDomain.AdminProfileID = adminProfile.ID

// 	// Membuat Admin dari AdminDomain
// 	record := FromAdminDomain(adminDomain)
// 	record.Password = string(password)

// 	// Simpan Admin
// 	result := ur.conn.WithContext(ctx).Create(&record)
// 	if result.Error != nil {
// 		return admin.AdminDomain{}, result.Error
// 	}

// 	// Mengambil data Admin terbaru dengan AdminProfile terkait
// 	err = ur.conn.WithContext(ctx).Preload("AdminProfile").Last(&record).Error
// 	if err != nil {
// 		return admin.AdminDomain{}, err
// 	}

// 	return record.ToAdminDomain(), nil
// }

// func (ur *adminRepository) AdminRegister(ctx context.Context, adminDomain *admin.AdminDomain) (admin.AdminDomain, error) {
// 	var adminProfile AdminProfile
// 	// adminProfile.Name = adminDomain.AdminProfile.Name
// 	// adminProfile.Nip = adminDomain.AdminProfile.Nip
// 	// adminProfile.Division = adminDomain.AdminProfile.Division
// 	// adminProfile.Image_Path = adminDomain.AdminProfile.Image_Path

// 	// // Membuat AdminProfile
// 	// adminProfile := AdminProfile{
// 	// 	Name:       adminDomain.AdminProfile.Name,
// 	// 	Nip:        adminDomain.AdminProfile.Nip,
// 	// 	Division:   adminDomain.AdminProfile.Division,
// 	// 	Image_Path: adminDomain.AdminProfile.Image_Path,
// 	// }

// 	// Simpan AdminProfile terlebih dahulu
// 	profileResult := ur.conn.WithContext(ctx).Create(&adminProfile)
// 	if profileResult.Error != nil {
// 		return admin.AdminDomain{}, profileResult.Error
// 	}

// 	password, err := bcrypt.GenerateFromPassword([]byte(adminDomain.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return admin.AdminDomain{}, err
// 	}

// 	record := FromAdminDomain(adminDomain)
// 	record.Password = string(password)
// 	record.AdminProfileID = adminProfile.ID // Set ID dari AdminProfile yang baru disimpan
// 	// record.Name = adminProfile.Name

// 	result := ur.conn.WithContext(ctx).Create(&record)
// 	if result.Error != nil {
// 		return admin.AdminDomain{}, result.Error
// 	}

// 	err = ur.conn.WithContext(ctx).Preload("AdminProfile").Last(&record).Error
// 	if err != nil {
// 		return admin.AdminDomain{}, err
// 	}

// 	return record.ToAdminDomain(), nil
// }

// func (ur *adminRepository) AdminRegister(ctx context.Context, adminDomain *admin.AdminDomain) (admin.AdminDomain, error) {
// 	var adminProfile AdminProfile
// 	adminProfile.Name = adminDomain.Name

// 	password, err := bcrypt.GenerateFromPassword([]byte(adminDomain.Password), bcrypt.DefaultCost)

// 	if err != nil {
// 		return admin.AdminDomain{}, err
// 	}

// 	record := FromAdminDomain(adminDomain)
// 	record.Password = string(password)
// 	record.AdminProfileID = record.AdminProfile.ID

// 	result := ur.conn.WithContext(ctx).Preload("AdminProfile").Create(&record)

// 	if err := result.Error; err != nil {
// 		return admin.AdminDomain{}, err
// 	}

// 	err = result.Last(&record).Error

// 	if err != nil {
// 		return admin.AdminDomain{}, err
// 	}

// 	return record.ToAdminDomain(), nil
// }

func (ur *adminRepository) AdminGetByEmail(ctx context.Context, adminDomain *admin.AdminDomain) (admin.AdminDomain, error) {
	var admins Admin

	err := ur.conn.WithContext(ctx).First(&admins, "email = ?", adminDomain.Name).Error

	if err != nil {
		return admin.AdminDomain{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admins.Password), []byte(adminDomain.Password))

	if err != nil {
		return admin.AdminDomain{}, err
	}

	return admins.ToAdminDomain(), nil
}

func (ur *adminRepository) AdminGetByVoucher(ctx context.Context, adminDomain *admin.AdminDomain) (admin.AdminDomain, error) {
	var admins Admin

	err := ur.conn.WithContext(ctx).First(&admins, "voucher = ?", adminDomain.Voucher).Error

	if err != nil {
		return admin.AdminDomain{}, err
	}

	return admins.ToAdminDomain(), nil
}

// // Ambil data admin berdasarkan ID
// admins, err := ur.AdminGetByID(ctx, id)
// if err != nil {
// 	return admin.AdminDomain{}, "", err
// }

// // Perbarui field-field yang ada pada adminDomain
// admins.Name = adminDomain.Name
// admins.Email = adminDomain.Email
// admins.Phone = adminDomain.Phone

// // Jika imagePath baru diberikan, perbarui ImagePath
// var prevURL string
// if imagePath != "" {
// 	prevURL = admins.ImagePath
// 	admins.ImagePath = imagePath
// }

// // Simpan perubahan ke database
// if err := ur.conn.WithContext(ctx).Save(&admins).Error; err != nil {
// 	return admin.AdminDomain{}, "", err
// }

func (ur *adminRepository) AdminProfileUpdate(ctx context.Context, adminDomain *admin.AdminDomain, imagePath string, id string) (admin.AdminDomain, string, error) {
	admins, err := ur.AdminGetByID(ctx, id)
	if err != nil {
		return admin.AdminDomain{}, "", err
	}

	// Ambil nama file menggunakan utilitas
	fileName := utils.GetFileName(imagePath)

	updateAdmin := FromAdminDomain(&admins)

	if adminDomain.Name != "" {
		updateAdmin.Name = adminDomain.Name
	}
	if adminDomain.Email != "" {
		updateAdmin.Email = adminDomain.Email
	}
	if adminDomain.Phone != "" {
		updateAdmin.Phone = adminDomain.Phone
	}

	// Simpan nama file jika ada perubahan
	if fileName != "" && fileName != "." {
		updateAdmin.ImagePath = fileName // Hanya menyimpan nama file
	}

	if err := ur.conn.WithContext(ctx).Save(&updateAdmin).Error; err != nil {
		return admin.AdminDomain{}, "", err
	}

	return updateAdmin.ToAdminDomain(), fileName, nil
}

// func (ur *adminRepository) AdminProfileUpdate(ctx context.Context, profileDomain *admin.AdminDomain, avatarPath string, id string) (admin.AdminDomain, string, error) {
// 	var admins Admin

// 	if err := ur.conn.WithContext(ctx).First(&admins, "id = ?", id).Error; err != nil {
// 		return admin.AdminDomain{}, "", err
// 	}

// 	prev_url := admins.ImagePath
// 	admins.ImagePath = avatarPath

// 	if err := ur.conn.WithContext(ctx).Save(&admins).Error; err != nil {
// 		return admin.AdminDomain{}, "", err
// 	}

// 	return admins.ToAdminDomain(), prev_url, nil
// }

// func (ur *adminRepository) AdminGetInfo(ctx context.Context, id string) (admin.AdminDomain, error) {
// 	var admins Admin

// 	if err := ur.conn.WithContext(ctx).First(&admins, "id = ?", id).Error; err != nil {
// 		return admin.AdminDomain{}, err
// 	}

// 	return admins.ToAdminDomain(), nil

// }

func (ur *adminRepository) AdminGetByID(ctx context.Context, id string) (admin.AdminDomain, error) {
	var admins Admin

	if err := ur.conn.WithContext(ctx).Preload("Role").First(&admins, "id = ?", id).Error; err != nil {
		return admin.AdminDomain{}, err
	}

	return admins.ToAdminDomain(), nil

}

func (vr *adminRepository) RoleCreate(ctx context.Context, roleDomain *admin.RoleDomain) (admin.RoleDomain, error) {
	record := FromRoleDomain(roleDomain)
	result := vr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return admin.RoleDomain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return admin.RoleDomain{}, err
	}

	return record.ToRoleDomain(), nil

}

func (ar *adminRepository) RoleGetByID(ctx context.Context, id string) (admin.RoleDomain, error) {
	var role Role

	if err := ar.conn.WithContext(ctx).First(&role, "id = ?", id).Error; err != nil {
		return admin.RoleDomain{}, err
	}

	return role.ToRoleDomain(), nil

}

func (ar *adminRepository) RoleGetAll(ctx context.Context) ([]admin.RoleDomain, error) {
	var records []Role
	if err := ar.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	roles := []admin.RoleDomain{}

	for _, role := range records {
		domain := role.ToRoleDomain()
		roles = append(roles, domain)
	}

	return roles, nil
}

// func (ur *adminRepository) AdminProfileUpdate(ctx context.Context, profileDomain *admin.AdminProfileDomain, id string) (admin.AdminProfileDomain, error) {
// 	var profile AdminProfile

// 	// Preload "Admin" untuk memastikan data User ter-load
// 	if err := ur.conn.WithContext(ctx).First(&profile, "id = ?", id).Error; err != nil {
// 		return admin.AdminProfileDomain{}, err
// 	}

// 	// Update hanya jika nilai berbeda
// 	if profile.Name != profileDomain.Name {
// 		profile.Name = profileDomain.Name
// 	}

// 	if profile.Nip != profileDomain.Nip {
// 		profile.Nip = profileDomain.Nip
// 	}

// 	if profile.Division != profileDomain.Division {
// 		profile.Division = profileDomain.Division
// 	}

// 	// Simpan perubahan ke database
// 	if err := ur.conn.WithContext(ctx).Save(&profile).Error; err != nil {
// 		return admin.AdminProfileDomain{}, err
// 	}

// 	// Mengembalikan profil yang telah diperbarui
// 	return profile.ToAdminProfileDomain(), nil
// }

// func (ur *adminRepository) AdminProfileUploadImage(ctx context.Context, profileDomain *admin.AdminProfileDomain, avatarPath string, id string) (admin.AdminProfileDomain, string, error) {
// 	var profile AdminProfile

// 	if err := ur.conn.WithContext(ctx).First(&profile, "id = ?", id).Error; err != nil {
// 		return admin.AdminProfileDomain{}, "", err
// 	}

// 	prev_url := profile.Image_Path
// 	profile.Image_Path = avatarPath

// 	if err := ur.conn.WithContext(ctx).Save(&profile).Error; err != nil {
// 		return admin.AdminProfileDomain{}, "", err
// 	}

// 	return profile.ToAdminProfileDomain(), prev_url, nil
// }

// func (ur *adminRepository) AdminProfileGetByID(ctx context.Context, id string) (admin.AdminProfileDomain, error) {
// 	var profile AdminProfile

// 	if err := ur.conn.WithContext(ctx).First(&profile, "id = ?", id).Error; err != nil {
// 		return admin.AdminProfileDomain{}, err
// 	}

// 	return profile.ToAdminProfileDomain(), nil

// }

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

func (ar *adminRepository) CustomersGetAll(ctx context.Context) ([]admin.CustomersDomain, error) {
	var records []Customers
	// Melakukan Preload untuk menampilkan Slice CartItems yang berisi Customers dan Stocks
	if err := ar.conn.WithContext(ctx).Preload("CartItems.Customers").Preload("CartItems.Stocks").
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

func (ar *adminRepository) CustomerDelete(ctx context.Context, id string) error {
	customer, err := ar.CustomersGetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedCustomer := FromCustomersDomain(&customer)

	err = ar.conn.WithContext(ctx).Delete(&deletedCustomer).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *adminRepository) PackagingOfficerCreate(ctx context.Context, packagingOfficerDomain *admin.PackagingOfficerDomain) (admin.PackagingOfficerDomain, error) {
	record := FromPackagingOfficerDomain(packagingOfficerDomain)
	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return admin.PackagingOfficerDomain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return admin.PackagingOfficerDomain{}, err
	}

	return record.ToPackagingOfficerDomain(), nil

}

func (cr *adminRepository) PackagingOfficerGetByID(ctx context.Context, id string) (admin.PackagingOfficerDomain, error) {
	var packagingOfficer PackagingOfficer

	if err := cr.conn.WithContext(ctx).First(&packagingOfficer, "id = ?", id).Error; err != nil {
		return admin.PackagingOfficerDomain{}, err
	}

	return packagingOfficer.ToPackagingOfficerDomain(), nil

}

func (ar *adminRepository) PackagingOfficerGetAll(ctx context.Context) ([]admin.PackagingOfficerDomain, error) {
	var records []PackagingOfficer
	// Melakukan Preload untuk menampilkan Slice CartItems yang berisi Customers dan Stocks
	if err := ar.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []admin.PackagingOfficerDomain{}

	for _, category := range records {
		domain := category.ToPackagingOfficerDomain()
		categories = append(categories, domain)
	}

	return categories, nil
}

// Categories
func (cr *adminRepository) CategoryCreate(ctx context.Context, AdminDomain *admin.CategoriesDomain) (admin.CategoriesDomain, error) {
	record := FromCategoriesDomain(AdminDomain)
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

	if err := cr.conn.WithContext(ctx).First(&categories, "name = ?", name).Error; err != nil {
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

func (ar *adminRepository) VendorsGetAll(ctx context.Context) ([]admin.VendorsDomain, error) {
	var records []Vendors
	if err := ar.conn.WithContext(ctx).Find(&records).Error; err != nil {
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

func (ar *adminRepository) StocksGetAll(ctx context.Context, page int, limit int, sort string, order string, search string, filters map[string]interface{}) ([]admin.StocksDomain, int, error) {
	var records []Stocks
	offset := (page - 1) * limit

	// Build the base query with pagination, sorting, and preload
	query := ar.conn.WithContext(ctx).
		Preload("Categories").
		Preload("Units").
		Offset(offset).
		Limit(limit).
		Order(fmt.Sprintf("%s %s", sort, order))

	// Add search condition if search keyword is provided
	if search != "" {
		query = query.
			Joins("LEFT JOIN categories ON categories.id = stocks.category_id").
			Joins("LEFT JOIN units ON units.id = stocks.unit_id").
			Where("stocks.stock_name LIKE ?", "%"+search+"%").
			Or("stocks.stock_code LIKE ?", "%"+search+"%").
			Or("stocks.description LIKE ?", "%"+search+"%").
			Or("stocks.selling_price LIKE ?", "%"+search+"%").
			Or("categories.category_name LIKE ?", "%"+search+"%").
			Or("units.unit_name LIKE ?", "%"+search+"%")
	}

	// Add filter conditions
	for key, value := range filters {
		switch key {
		case "category_id":
			query = query.Where("category_id = ?", value)
		case "unit_id":
			query = query.Where("unit_id = ?", value)
		case "stock_total_min":
			query = query.Where("stock_total >= ?", value)
		case "stock_total_max":
			query = query.Where("stock_total <= ?", value)
		case "selling_price_min":
			query = query.Where("selling_price >= ?", value)
		case "selling_price_max":
			query = query.Where("selling_price <= ?", value)
		case "category_name":
			// Filter for multiple category names
			categoryName := value.([]string)
			query = query.Joins("LEFT JOIN categories ON categories.id = stocks.category_id").
				Where("categories.category_name IN (?)", categoryName)
		case "unit_name":
			// Filter for multiple unit names
			unitName := value.([]string)
			query = query.Joins("LEFT JOIN units ON units.id = stocks.unit_id").
				Where("units.unit_name IN (?)", unitName)
		}
	}

	// Execute the query
	if err := query.Find(&records).Error; err != nil {
		return nil, 0, err
	}

	// Convert to domain models
	stocksDomain := make([]admin.StocksDomain, len(records))
	for i, stock := range records {
		stocksDomain[i] = stock.ToStocksDomain()
	}

	// Get the total count of items without pagination
	var totalItems int64
	countQuery := ar.conn.Model(&Stocks{}).
		Joins("LEFT JOIN categories ON categories.id = stocks.category_id").
		Joins("LEFT JOIN units ON units.id = stocks.unit_id")

	// Apply search and filter conditions for count query
	if search != "" {
		countQuery = countQuery.
			Where("stocks.stock_name LIKE ?", "%"+search+"%").
			Or("stocks.stock_code LIKE ?", "%"+search+"%").
			Or("stocks.description LIKE ?", "%"+search+"%").
			Or("stocks.selling_price LIKE ?", "%"+search+"%").
			Or("categories.category_name LIKE ?", "%"+search+"%").
			Or("units.unit_name LIKE ?", "%"+search+"%")
	}

	for key, value := range filters {
		switch key {
		case "category_id":
			countQuery = countQuery.Where("category_id = ?", value)
		case "unit_id":
			countQuery = countQuery.Where("unit_id = ?", value)
		case "stock_total_min":
			countQuery = countQuery.Where("stock_total >= ?", value)
		case "stock_total_max":
			countQuery = countQuery.Where("stock_total <= ?", value)
		case "selling_price_min":
			countQuery = countQuery.Where("selling_price >= ?", value)
		case "selling_price_max":
			countQuery = countQuery.Where("selling_price <= ?", value)
		case "category_name":
			// Filter for multiple category names
			categoryName := value.([]string)
			countQuery = countQuery.Joins("LEFT JOIN categories ON categories.id = stocks.category_id").
				Where("categories.category_name IN (?)", categoryName)
		case "unit_name":
			// Filter for multiple unit names
			unitName := value.([]string)
			countQuery = countQuery.Joins("LEFT JOIN units ON units.id = stocks.unit_id").
				Where("units.unit_name IN (?)", unitName)
		}
	}

	countQuery.Count(&totalItems)

	return stocksDomain, int(totalItems), nil
}

// search by name, search by category gagal, filter, asc ok
// func (ar *adminRepository) StocksGetAll(ctx context.Context, page int, limit int, sort string, order string, search string, filters map[string]interface{}) ([]admin.StocksDomain, int, error) {
// 	var records []Stocks
// 	offset := (page - 1) * limit

// 	// Build the base query with pagination and sorting
// 	query := ar.conn.WithContext(ctx).
// 		Model(&Stocks{}).
// 		Preload("Categories").
// 		Preload("Units").
// 		Offset(offset).
// 		Limit(limit).
// 		Order(fmt.Sprintf("%s %s", sort, order))

// 	// Add search condition if search keyword is provided
// 	if search != "" {
// 		query = query.
// 			Where("stocks.stock_name LIKE ?", "%"+search+"%").
// 			Or("stocks.stock_code LIKE ?", "%"+search+"%").
// 			Or("stocks.description LIKE ?", "%"+search+"%").
// 			Or("stocks.selling_price LIKE ?", "%"+search+"%").
// 			// Join dengan Categories untuk mencari category_name
// 			Joins("LEFT JOIN categories ON categories.id = stocks.category_id").
// 			Where("categories.category_name LIKE ?", "%"+search+"%").
// 			// Join dengan Units untuk mencari unit_name
// 			Joins("LEFT JOIN units ON units.id = stocks.unit_id").
// 			Where("units.unit_name LIKE ?", "%"+search+"%")
// 	}

// 	// Add filter conditions
// 	for key, value := range filters {
// 		switch key {
// 		case "category_id":
// 			query = query.Where("category_id = ?", value)
// 		case "unit_id":
// 			query = query.Where("unit_id = ?", value)
// 		case "stock_total_min":
// 			query = query.Where("stock_total >= ?", value)
// 		case "stock_total_max":
// 			query = query.Where("stock_total <= ?", value)
// 		case "selling_price_min":
// 			query = query.Where("selling_price >= ?", value)
// 		case "selling_price_max":
// 			query = query.Where("selling_price <= ?", value)
// 		}
// 	}

// 	// Execute the query
// 	if err := query.Find(&records).Error; err != nil {
// 		return nil, 0, err
// 	}

// 	// Convert to domain models
// 	stocksDomain := make([]admin.StocksDomain, len(records))
// 	for i, stock := range records {
// 		stocksDomain[i] = stock.ToStocksDomain()
// 	}

// 	// Get the total count of items without pagination
// 	var totalItems int64
// 	countQuery := ar.conn.Model(&Stocks{})

// 	// Apply search and filter conditions for count query
// 	if search != "" {
// 		countQuery = countQuery.
// 			Where("stock_name LIKE ?", "%"+search+"%").
// 			Or("stock_code LIKE ?", "%"+search+"%").
// 			Or("description LIKE ?", "%"+search+"%").
// 			Or("selling_price LIKE ?", "%"+search+"%").
// 			Joins("LEFT JOIN categories ON categories.id = stocks.category_id").
// 			Where("categories.category_name LIKE ?", "%"+search+"%").
// 			Joins("LEFT JOIN units ON units.id = stocks.unit_id").
// 			Where("units.unit_name LIKE ?", "%"+search+"%")
// 	}

// 	for key, value := range filters {
// 		switch key {
// 		case "category_id":
// 			countQuery = countQuery.Where("category_id = ?", value)
// 		case "unit_id":
// 			countQuery = countQuery.Where("unit_id = ?", value)
// 		case "stock_total_min":
// 			countQuery = countQuery.Where("stock_total >= ?", value)
// 		case "stock_total_max":
// 			countQuery = countQuery.Where("stock_total <= ?", value)
// 		case "selling_price_min":
// 			countQuery = countQuery.Where("selling_price >= ?", value)
// 		case "selling_price_max":
// 			countQuery = countQuery.Where("selling_price <= ?", value)
// 		}
// 	}

// 	countQuery.Count(&totalItems)

// 	return stocksDomain, int(totalItems), nil
// }

// func (ar *adminRepository) StocksGetAll(ctx context.Context, page int, limit int, sort string, order string, search string) ([]admin.StocksDomain, int, error) {
// 	var records []Stocks
// 	offset := (page - 1) * limit

// 	// Build the query with pagination, sorting, and search
// 	query := ar.conn.WithContext(ctx).Preload("Categories").Preload("Units").
// 		Offset(offset).Limit(limit).Order(fmt.Sprintf("%s %s", sort, order))

// 	if search != "" {
// 		query = query.Where("stock_name LIKE ?", "%"+search+"%")
// 	}

// 	// Execute the query
// 	if err := query.Find(&records).Error; err != nil {
// 		return nil, 0, err
// 	}

// 	// Convert to domain models
// 	stocksDomain := []admin.StocksDomain{}
// 	for _, stock := range records {
// 		domain := stock.ToStocksDomain()
// 		stocksDomain = append(stocksDomain, domain)
// 	}

// 	// Get the total count of items without pagination
// 	var totalItems int64
// 	countQuery := ar.conn.Model(&Stocks{})
// 	if search != "" {
// 		countQuery = countQuery.Where("stock_name LIKE ?", "%"+search+"%")
// 	}
// 	countQuery.Count(&totalItems)

// 	return stocksDomain, int(totalItems), nil
// }

// func (ar *adminRepository) StocksGetAll(ctx context.Context) ([]admin.StocksDomain, error) {
// 	var records []Stocks
// 	if err := ar.conn.WithContext(ctx).
// 		Preload("Categories").Preload("Units").
// 		Find(&records).Error; err != nil {
// 		return nil, err
// 	}

// 	stocksDomain := []admin.StocksDomain{}

// 	for _, stocks := range records {
// 		domain := stocks.ToStocksDomain()
// 		stocksDomain = append(stocksDomain, domain)
// 	}

// 	return stocksDomain, nil
// }

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

func (ar *adminRepository) PurchasesGetAll(ctx context.Context, page int, limit int, sort string, order string, search string, filters map[string]interface{}) ([]admin.PurchasesDomain, int, error) {
	var records []Purchases
	offset := (page - 1) * limit

	// Build the base query with pagination, sorting, and preload
	query := ar.conn.WithContext(ctx).
		Preload("Vendors").
		Preload("Categories").
		Preload("Units").
		Offset(offset).
		Limit(limit).
		Order(fmt.Sprintf("%s %s", sort, order))

	// Add search condition if search keyword is provided
	if search != "" {
		query = query.
			Joins("LEFT JOIN vendors ON vendors.id = purchases.vendor_id").
			Joins("LEFT JOIN categories ON categories.id = purchases.category_id").
			Joins("LEFT JOIN units ON units.id = purchases.unit_id").
			Where("purchases.stock_name LIKE ?", "%"+search+"%").
			Or("purchases.stock_code LIKE ?", "%"+search+"%").
			Or("purchases.description LIKE ?", "%"+search+"%").
			Or("purchases.purchase_price LIKE ?", "%"+search+"%").
			Or("vendors.vendor_name LIKE ?", "%"+search+"%").
			Or("categories.category_name LIKE ?", "%"+search+"%").
			Or("units.unit_name LIKE ?", "%"+search+"%")
	}

	// Add filter conditions
	for key, value := range filters {
		switch key {
		case "vendor_id":
			query = query.Where("vendor_id = ?", value)
		case "category_id":
			query = query.Where("category_id = ?", value)
		case "unit_id":
			query = query.Where("unit_id = ?", value)
		case "purchase_price_min":
			query = query.Where("purchase_price >= ?", value)
		case "purchase_price_max":
			query = query.Where("purchase_price <= ?", value)
		case "category_name":
			categoryNames := value.([]string)
			query = query.Joins("LEFT JOIN categories ON categories.id = purchases.category_id").
				Where("categories.category_name IN (?)", categoryNames)
		case "unit_name":
			unitNames := value.([]string)
			query = query.Joins("LEFT JOIN units ON units.id = purchases.unit_id").
				Where("units.unit_name IN (?)", unitNames)
		case "vendor_name":
			vendorNames := value.([]string)
			query = query.Joins("LEFT JOIN vendors ON vendors.id = purchases.vendor_id").
				Where("vendors.vendor_name IN (?)", vendorNames)
		}
	}

	// Execute the query
	if err := query.Find(&records).Error; err != nil {
		return nil, 0, err
	}

	// Convert to domain models
	purchasesDomain := make([]admin.PurchasesDomain, len(records))
	for i, purchase := range records {
		purchasesDomain[i] = purchase.ToPurchasesDomain()
	}

	// Get the total count of items without pagination
	var totalItems int64
	countQuery := ar.conn.Model(&Purchases{}).
		Joins("LEFT JOIN vendors ON vendors.id = purchases.vendor_id").
		Joins("LEFT JOIN categories ON categories.id = purchases.category_id").
		Joins("LEFT JOIN units ON units.id = purchases.unit_id")

	// Apply search and filter conditions for count query
	if search != "" {
		countQuery = countQuery.
			Where("purchases.stock_name LIKE ?", "%"+search+"%").
			Or("purchases.stock_code LIKE ?", "%"+search+"%").
			Or("purchases.description LIKE ?", "%"+search+"%").
			Or("purchases.purchase_price LIKE ?", "%"+search+"%").
			Or("vendors.vendor_name LIKE ?", "%"+search+"%").
			Or("categories.category_name LIKE ?", "%"+search+"%").
			Or("units.unit_name LIKE ?", "%"+search+"%")
	}

	for key, value := range filters {
		switch key {
		case "vendor_id":
			countQuery = countQuery.Where("vendor_id = ?", value)
		case "category_id":
			countQuery = countQuery.Where("category_id = ?", value)
		case "unit_id":
			countQuery = countQuery.Where("unit_id = ?", value)
		case "purchase_price_min":
			countQuery = countQuery.Where("purchase_price >= ?", value)
		case "purchase_price_max":
			countQuery = countQuery.Where("purchase_price <= ?", value)
		case "category_name":
			categoryNames := value.([]string)
			countQuery = countQuery.Joins("LEFT JOIN categories ON categories.id = purchases.category_id").
				Where("categories.category_name IN (?)", categoryNames)
		case "unit_name":
			unitNames := value.([]string)
			countQuery = countQuery.Joins("LEFT JOIN units ON units.id = purchases.unit_id").
				Where("units.unit_name IN (?)", unitNames)
		case "vendor_name":
			vendorNames := value.([]string)
			countQuery = countQuery.Joins("LEFT JOIN vendors ON vendors.id = purchases.vendor_id").
				Where("vendors.vendor_name IN (?)", vendorNames)
		}
	}

	countQuery.Count(&totalItems)

	return purchasesDomain, int(totalItems), nil
}

// func (ar *adminRepository) PurchasesGetAll(ctx context.Context, page int, limit int, sort string, order string, search string) ([]admin.PurchasesDomain, int, error) {
// 	var records []Purchases
// 	offset := (page - 1) * limit

// 	// Bangun query dengan paginasi, sorting, dan pencarian
// 	query := ar.conn.WithContext(ctx).
// 		Preload("Vendors").Preload("Categories").Preload("Units").
// 		Offset(offset).Limit(limit).Order(fmt.Sprintf("%s %s", sort, order))

// 	if search != "" {
// 		query = query.Where("stock_name LIKE ?", "%"+search+"%")
// 	}

// 	// Eksekusi query
// 	if err := query.Find(&records).Error; err != nil {
// 		return nil, 0, err
// 	}

// 	// Konversi ke domain models
// 	purchasesDomain := []admin.PurchasesDomain{}
// 	for _, purchase := range records {
// 		domain := purchase.ToPurchasesDomain()
// 		purchasesDomain = append(purchasesDomain, domain)
// 	}

// 	// Dapatkan total item tanpa paginasi
// 	var totalItems int64
// 	countQuery := ar.conn.Model(&Purchases{})
// 	if search != "" {
// 		countQuery = countQuery.Where("stock_name LIKE ?", "%"+search+"%")
// 	}
// 	countQuery.Count(&totalItems)

// 	return purchasesDomain, int(totalItems), nil
// }

// func (ar *adminRepository) PurchasesGetAll(ctx context.Context) ([]admin.PurchasesDomain, error) {
// 	// Memuat data Purchases beserta relasi Vendors, Categories, dan Units
// 	var records []Purchases
// 	if err := ar.conn.WithContext(ctx).
// 		// Preload("Vendors").Preload("Categories").Preload("Units").
// 		Preload("Vendors").Preload("Categories").Preload("Units").
// 		Find(&records).Error; err != nil {
// 		return nil, err
// 	}

// 	purchasesDomain := []admin.PurchasesDomain{}

// 	for _, purchase := range records {
// 		// Konversi ke domain
// 		domain := purchase.ToPurchasesDomain()
// 		// Tambahkan ke hasil
// 		purchasesDomain = append(purchasesDomain, domain)
// 	}

// 	return purchasesDomain, nil
// }

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
	for _, cartItems := range customerCartItems {
		subTotal += cartItems.Price
	}
	record.SubTotal = subTotal
	// record.UnitsID = itemDomain.UnitsID
	// record.Stocks.UnitID = record.UnitsID
	// record.UnitsID = record.UnitsID

	// Simpan catatan penjualan
	result := pr.conn.WithContext(ctx).Create(&record)
	if err := result.Error; err != nil {
		return admin.CartItemsDomain{}, err
	}

	// Perbarui SubTotal untuk semua penjualan lainnya dari pelanggan yang sama
	for _, cartItems := range customerCartItems {
		cartItems.SubTotal = subTotal
		if err := pr.conn.WithContext(ctx).Save(&cartItems).Error; err != nil {
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
	// for _, cartItems := range allCartItems {
	// 	totalAllPrice += cartItems.Price
	// }
	// record.SubTotal = totalAllPrice

	// // Simpan catatan penjualan
	// result := pr.conn.WithContext(ctx).Create(&record)
	// if err := result.Error; err != nil {
	// 	return admin.CartItemsDomain{}, err
	// }

	// // Perbarui TotalAllPrice untuk semua penjualan lainnya
	// for _, cartItems := range allCartItems {
	// 	cartItems.SubTotal = totalAllPrice
	// 	if err := pr.conn.WithContext(ctx).Save(&cartItems).Error; err != nil {
	// 		return admin.CartItemsDomain{}, err
	// 	}
	// }

}

func (ar *adminRepository) CartItemsGetByID(ctx context.Context, id string) (admin.CartItemsDomain, error) {
	var item CartItems

	if err := ar.conn.WithContext(ctx).
		Preload("Customers").
		Preload("Stocks").
		Preload("Stocks.Units").
		First(&item, "id = ?", id).Error; err != nil {
		return admin.CartItemsDomain{}, err
	}

	return item.ToCartItemsDomain(), nil

}

// func (ar *adminRepository) CartItemsGetByCustomerID(ctx context.Context, customerId string) (admin.CustomersDomain, error) {
// 	var customer Customers

// 	// if err := ar.conn.WithContext(ctx).Preload("Customers").Preload("Stocks").First(&item, "customer_id = ?", cartItemsDomain.CustomerID).Error; err != nil {
// 	if err := ar.conn.WithContext(ctx).Preload("CartItems").Where(" id = ?", customerId).First(&customer).Error; err != nil {
// 		return admin.CustomersDomain{}, err
// 	}
// 	// if err != nil {
// 	// 	return admin.CartItemsDomain{}, err
// 	// }

// 	return customer.ToCustomersDomain(), nil

// }

// Sulit-Sulit :v
func (ar *adminRepository) CartItemsGetAllByCustomerID(ctx context.Context, customerId string) ([]admin.CartItemsDomain, error) {
	var cartItems []CartItems

	if err := ar.conn.WithContext(ctx).
		Preload("Customers").
		Preload("Stocks").
		Preload("Stocks.Units").
		Where("customer_id = ?", customerId).
		Find(&cartItems).Error; err != nil {
		// return nil, err
		return nil, err
	}

	// Komentar Ini Jangan Di Hapus
	// Jika customer tidak ada data cartItems yang ditemukan, kembalikan error
	// if len(cartItems) == 0 {
	// 	return nil, fmt.Errorf("no cart items found for customer with ID %s", customerId)
	// }

	cartItemsDomain := []admin.CartItemsDomain{}

	for _, purchase := range cartItems {
		// Konversi ke domain
		domain := purchase.ToCartItemsDomain()
		cartItemsDomain = append(cartItemsDomain, domain)

	}
	return cartItemsDomain, nil

	// cartItemsDomain := make([]admin.CartItemsDomain, len(cartItems))

	// for i, purchase := range cartItems {
	// 	// Konversi ke domain
	// 	// domain := purchase.ToCartItemsDomain()
	// 	cartItemsDomain[i] = purchase.ToCartItemsDomain()
	// }

	// return cartItemsDomain, nil

}

func (ar *adminRepository) CartItemsGetAll(ctx context.Context) ([]admin.CartItemsDomain, error) {
	var records []CartItems
	if err := ar.conn.WithContext(ctx).
		Preload("Customers").
		Preload("Stocks").
		Preload("Stocks.Units").
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

func (ar *adminRepository) CartItemsDelete(ctx context.Context, id string) error {
	items, err := ar.CartItemsGetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedItems := FromCartItemsDomain(&items)

	// Ambil semua item keranjang dari pelanggan yang sama
	var customerItems []CartItems
	err = ar.conn.WithContext(ctx).
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
			if err := ar.conn.WithContext(ctx).Save(&item).Error; err != nil {
				return err
			}
		}
	}

	// Hapus data item keranjang
	// Gunakan Unscope untuk menghapus data secara permanent
	if err := ar.conn.WithContext(ctx).Unscoped().Delete(&deletedItems).Error; err != nil {
		return err
	}

	return nil

	// // Ambil semua penjualan lainnya
	// var allItems []CartItems
	// err = ar.conn.WithContext(ctx).Find(&allItems).Error
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
	// 		if err := ar.conn.WithContext(ctx).Save(&item).Error; err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	// // Hapus data penjualan
	// if err := ar.conn.WithContext(ctx).Unscoped().Delete(&deletedItems).Error; err != nil {
	// 	return err
	// }

	// return nil
}

func (ar *adminRepository) ItemTransactionsCreate(ctx context.Context, customerId string) (admin.ItemTransactionsDomain, error) {

	var cartItemsData []CartItems

	// Ambil semua data penjualan berdasarkan customerID
	if err := ar.conn.WithContext(ctx).
		Where("customer_id = ?", customerId).
		Find(&cartItemsData).Error; err != nil {
		return admin.ItemTransactionsDomain{}, err
	}

	// Jika tidak ada item dalam keranjang, kembalikan error
	if len(cartItemsData) == 0 {
		return admin.ItemTransactionsDomain{}, fmt.Errorf("no cart items found for customer with ID %s", customerId)
	}

	for _, cartItems := range cartItemsData {
		// Periksa stok yang terkait dengan penjualan
		var stock Stocks
		if err := ar.conn.WithContext(ctx).Where("id = ?", cartItems.StockID).First(&stock).Error; err != nil {
			return admin.ItemTransactionsDomain{}, err
		}

		// Kurangi stok dengan jumlah yang dijual
		if stock.StockTotal < cartItems.Quantity {
			errMsg := fmt.Sprintf("stok tidak cukup untuk penjualan dengan stock_id %d", cartItems.StockID)
			log.Println(errMsg)
			return admin.ItemTransactionsDomain{}, fmt.Errorf(errMsg)
		}
		stock.StockTotal -= cartItems.Quantity

		// Simpan perubahan stok ke database
		if err := ar.conn.WithContext(ctx).Save(&stock).Error; err != nil {
			return admin.ItemTransactionsDomain{}, err
		}

		// Buat catatan itemTransactions
		itemTransactionsRecord := ItemTransactions{
			CustomerID: cartItems.CustomerID,
			StockID:    cartItems.StockID,
			CategoryID: stock.CategoryID, // Gunakan CategoryID dari Stocks
			UnitID:     stock.UnitID,     // Gunakan UnitID dari Stocks
			Quantity:   cartItems.Quantity,
			Price:      cartItems.Price,
			SubTotal:   cartItems.SubTotal,
			// Sesuaikan dengan kolom lain jika perlu
		}

		// Simpan catatan itemTransactions
		if err := ar.conn.WithContext(ctx).Create(&itemTransactionsRecord).Error; err != nil {
			return admin.ItemTransactionsDomain{}, err
		}
	}

	// Hapus semua data penjualan yang terkait dengan customerID
	if err := ar.conn.WithContext(ctx).Where("customer_id = ?", customerId).Unscoped().Delete(&CartItems{}).Error; err != nil {
		return admin.ItemTransactionsDomain{}, err
	}

	return admin.ItemTransactionsDomain{}, nil
}

func (cr *adminRepository) ItemTransactionsGetAll(ctx context.Context, page int, limit int, sort string, order string, search string, filters map[string]interface{}) ([]admin.ItemTransactionsDomain, int, error) {
	var records []ItemTransactions
	offset := (page - 1) * limit

	// Build the base query with pagination, sorting, and preload
	query := cr.conn.WithContext(ctx).
		Preload("Customers").
		Preload("Stocks").
		Preload("Stocks.Units").
		Preload("Stocks.Categories").
		Offset(offset).
		Limit(limit).
		Order(fmt.Sprintf("%s %s", sort, order))

	// Handle sorting and joins for sorting
	if sort == "stock_name" {
		query = query.Joins("LEFT JOIN stocks ON stocks.id = item_transactions.stock_id").
			Order(fmt.Sprintf("stocks.%s %s", sort, order))
	} else if sort == "customer_name" {
		query = query.Joins("LEFT JOIN customers ON customers.id = item_transactions.customer_id").
			Order(fmt.Sprintf("customers.%s %s", sort, order))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", sort, order))
	}

	// Add search condition if search keyword is provided
	if search != "" {
		query = query.
			Joins("LEFT JOIN customers ON customers.id = item_transactions.customer_id").
			Joins("LEFT JOIN stocks ON stocks.id = item_transactions.stock_id").
			Joins("LEFT JOIN units ON units.id = item_transactions.unit_id").
			Joins("LEFT JOIN categories ON categories.id = item_transactions.category_id").
			Where("stocks.stock_name LIKE ?", "%"+search+"%").
			Or("stocks.stock_code LIKE ?", "%"+search+"%").
			Or("customers.customer_name LIKE ?", "%"+search+"%").
			Or("categories.category_name LIKE ?", "%"+search+"%").
			Or("units.unit_name LIKE ?", "%"+search+"%")
	}

	// Add filter conditions
	for key, value := range filters {
		switch key {
		case "customer_id":
			query = query.Where("customer_id = ?", value)
		case "stock_id":
			query = query.Where("stock_id = ?", value)
		case "unit_id":
			query = query.Where("unit_id = ?", value)
		case "category_id":
			query = query.Where("category_id = ?", value)
		case "quantity_min":
			query = query.Where("quantity >= ?", value)
		case "quantity_max":
			query = query.Where("quantity <= ?", value)
		case "price_min":
			query = query.Where("price >= ?", value)
		case "price_max":
			query = query.Where("price <= ?", value)
		case "sub_total_min":
			query = query.Where("sub_total >= ?", value)
		case "sub_total_max":
			query = query.Where("sub_total <= ?", value)
		case "customer_name":
			customerNames := value.([]string)
			query = query.Joins("LEFT JOIN customers ON customers.id = item_transactions.customer_id").
				Where("customers.customer_name IN (?)", customerNames)
		case "stock_name":
			stockNames := value.([]string)
			query = query.Joins("LEFT JOIN stocks ON stocks.id = item_transactions.stock_id").
				Where("stocks.stock_name IN (?)", stockNames)
		case "unit_name":
			unitNames := value.([]string)
			query = query.Joins("LEFT JOIN units ON units.id = item_transactions.unit_id").
				Where("units.unit_name IN (?)", unitNames)
		case "category_name":
			categoryNames := value.([]string)
			query = query.Joins("LEFT JOIN categories ON categories.id = item_transactions.category_id").
				Where("categories.category_name IN (?)", categoryNames)
		}
	}

	// Execute the query
	if err := query.Find(&records).Error; err != nil {
		return nil, 0, err
	}

	// Convert to domain models
	itemTransactionsDomain := make([]admin.ItemTransactionsDomain, len(records))
	for i, itemTransaction := range records {
		itemTransactionsDomain[i] = itemTransaction.ToItemTransactionsDomain()
	}

	// Get the total count of items without pagination
	var totalItems int64
	countQuery := cr.conn.Model(&ItemTransactions{}).
		Joins("LEFT JOIN customers ON customers.id = item_transactions.customer_id").
		Joins("LEFT JOIN stocks ON stocks.id = item_transactions.stock_id").
		Joins("LEFT JOIN units ON units.id = item_transactions.unit_id").
		Joins("LEFT JOIN categories ON categories.id = item_transactions.category_id")

	// Apply search and filter conditions for count query
	if search != "" {
		countQuery = countQuery.
			Where("stocks.stock_name LIKE ?", "%"+search+"%").
			Or("stocks.stock_code LIKE ?", "%"+search+"%").
			Or("customers.customer_name LIKE ?", "%"+search+"%").
			Or("categories.category_name LIKE ?", "%"+search+"%").
			Or("units.unit_name LIKE ?", "%"+search+"%")
	}

	for key, value := range filters {
		switch key {
		case "customer_id":
			countQuery = countQuery.Where("customer_id = ?", value)
		case "stock_id":
			countQuery = countQuery.Where("stock_id = ?", value)
		case "unit_id":
			countQuery = countQuery.Where("unit_id = ?", value)
		case "category_id":
			countQuery = countQuery.Where("category_id = ?", value)
		case "quantity_min":
			countQuery = countQuery.Where("quantity >= ?", value)
		case "quantity_max":
			countQuery = countQuery.Where("quantity <= ?", value)
		case "price_min":
			countQuery = countQuery.Where("price >= ?", value)
		case "price_max":
			countQuery = countQuery.Where("price <= ?", value)
		case "sub_total_min":
			countQuery = countQuery.Where("sub_total >= ?", value)
		case "sub_total_max":
			countQuery = countQuery.Where("sub_total <= ?", value)
		case "customer_name":
			customerNames := value.([]string)
			countQuery = countQuery.Joins("LEFT JOIN customers ON customers.id = item_transactions.customer_id").
				Where("customers.customer_name IN (?)", customerNames)
		case "stock_name":
			stockNames := value.([]string)
			countQuery = countQuery.Joins("LEFT JOIN stocks ON stocks.id = item_transactions.stock_id").
				Where("stocks.stock_name IN (?)", stockNames)
		case "unit_name":
			unitNames := value.([]string)
			countQuery = countQuery.Joins("LEFT JOIN units ON units.id = item_transactions.unit_id").
				Where("units.unit_name IN (?)", unitNames)
		case "category_name":
			categoryNames := value.([]string)
			countQuery = countQuery.Joins("LEFT JOIN categories ON categories.id = item_transactions.category_id").
				Where("categories.category_name IN (?)", categoryNames)
		}
	}

	countQuery.Count(&totalItems)

	return itemTransactionsDomain, int(totalItems), nil
}

func (ar *adminRepository) ReminderPurchaseOrderCreate(ctx context.Context, reminderPurchaseOrderDomain *admin.ReminderPurchaseOrderDomain) (admin.ItemTransactionsDomain, error) {
	// Ambil semua data item keranjang berdasarkan customerID
	var cartItemsData []CartItems
	if err := ar.conn.WithContext(ctx).
		Where("customer_id = ?", reminderPurchaseOrderDomain.CartItem.CustomerID).
		Preload("Customers").
		Find(&cartItemsData).Error; err != nil {
		return admin.ItemTransactionsDomain{}, err
	}

	// Jika tidak ada item dalam keranjang, kembalikan error
	if len(cartItemsData) == 0 {
		return admin.ItemTransactionsDomain{}, fmt.Errorf("no cart items found for customer with ID %d", reminderPurchaseOrderDomain.CartItem.CustomerID)
	}

	// Buat daftar item untuk pesan
	var itemList string
	customerInfo := struct {
		Name    string
		Email   string
		Address string
		Phone   string
	}{}

	for _, cartItems := range cartItemsData {
		// Ambil customer info dari item keranjang
		if customerInfo.Name == "" {
			customerInfo.Name = cartItems.Customers.CustomerName
			customerInfo.Email = cartItems.Customers.CustomerEmail
			customerInfo.Address = cartItems.Customers.CustomerAddress
			customerInfo.Phone = cartItems.Customers.CustomerPhone
		}

		// Periksa stok yang terkait dengan penjualan
		var stock Stocks
		if err := ar.conn.WithContext(ctx).Where("id = ?", cartItems.StockID).First(&stock).Error; err != nil {
			return admin.ItemTransactionsDomain{}, err
		}

		// Kurangi stok dengan jumlah yang dijual
		if stock.StockTotal < cartItems.Quantity {
			errMsg := fmt.Sprintf("stok tidak cukup untuk penjualan dengan stock_id %d", cartItems.StockID)
			log.Println(errMsg)
			return admin.ItemTransactionsDomain{}, fmt.Errorf(errMsg)
		}
		stock.StockTotal -= cartItems.Quantity

		// Simpan perubahan stok ke database
		if err := ar.conn.WithContext(ctx).Save(&stock).Error; err != nil {
			return admin.ItemTransactionsDomain{}, err
		}

		// Buat catatan itemTransactions
		itemTransactionsRecord := ItemTransactions{
			CustomerID: cartItems.CustomerID,
			StockID:    cartItems.StockID,
			UnitID:     stock.UnitID,
			CategoryID: stock.CategoryID,
			Quantity:   cartItems.Quantity,
			Price:      cartItems.Price,
			SubTotal:   cartItems.SubTotal,
		}

		// Simpan catatan itemTransactions
		if err := ar.conn.WithContext(ctx).Create(&itemTransactionsRecord).Error; err != nil {
			return admin.ItemTransactionsDomain{}, err
		}

		// Tambahkan detail item ke itemList
		itemDetail := fmt.Sprintf("Product: %s, Quantity: %d", stock.StockName, cartItems.Quantity)
		itemList += itemDetail + "\n"
	}

	// Ambil semua data Packaging Officer dan log informasinya
	var packagingOfficers []PackagingOfficer
	if err := ar.conn.WithContext(ctx).Find(&packagingOfficers).Error; err != nil {
		return admin.ItemTransactionsDomain{}, err
	}

	// Format pesan yang akan dikirim
	message := fmt.Sprintf("Reminder for Packaging: Purchase-Order needs to be prepared for customer %s. Please check the item list:\n%s\n\nCustomer Details:\nName: %s\nEmail: %s\nAddress: %s\nPhone: %s",
		customerInfo.Name, itemList, customerInfo.Name, customerInfo.Email, customerInfo.Address, customerInfo.Phone)

	// Hitung durasi sampai ReminderTime
	duration := time.Until(reminderPurchaseOrderDomain.ReminderTime)

	// Jalankan goroutine untuk menunggu sampai ReminderTime dan kirim pesan
	for _, officer := range packagingOfficers {
		log.Printf("Packaging Officer - ID: %d, Name: %s, Phone: %s\n", officer.ID, officer.OfficerName, officer.OfficerPhone)

		go func(officer PackagingOfficer, message string) {
			time.Sleep(duration)

			// Kirim pesan WhatsApp ke Packaging Officer
			err := utils.SendWhatsAppMessage(officer.OfficerPhone, message, "default")
			if err != nil {
				log.Printf("Failed to send WhatsApp message to Packaging Officer: %v\n", err)
			} else {
				log.Printf("WhatsApp message sent to Packaging Officer %s: %s\n", officer.OfficerPhone, message)
			}
		}(officer, message)

		// Simpan reminder ke database
		reminderRecord := ReminderPurchaseOrder{
			PackagingOfficerID: officer.ID,
			ReminderTime:       reminderPurchaseOrderDomain.ReminderTime,
		}

		if err := ar.conn.WithContext(ctx).Create(&reminderRecord).Error; err != nil {
			return admin.ItemTransactionsDomain{}, err
		}
	}

	// Hapus semua data item keranjang yang terkait dengan customerID
	if err := ar.conn.WithContext(ctx).Where("customer_id = ?", reminderPurchaseOrderDomain.CartItem.CustomerID).Unscoped().Delete(&CartItems{}).Error; err != nil {
		return admin.ItemTransactionsDomain{}, err
	}

	return admin.ItemTransactionsDomain{}, nil
}

// // Pastikan AdminVoucher tidak kosong
// if reminderPurchaseOrderDomain.Admins.Voucher == "" {
// 	return admin.ItemTransactionsDomain{}, fmt.Errorf("admin voucher is required")
// }

// // Dapatkan AdminID dari record Admin berdasarkan voucher
// var admins Admin
// if err := ar.conn.WithContext(ctx).
// 	Where("admin_id = ?", reminderPurchaseOrderDomain.Admins.ID).
// 	// First(&admins, "id = ?", reminderPurchaseOrderDomain.AdminID).Error; err != nil {
// 	First(&admins).Error; err != nil {
// 	return admin.ItemTransactionsDomain{}, err
// }

// // Tampilkan log informasi admin
// log.Printf("Admin - ID: %d, Name: %s, Email: %s, Phone: %s, Voucher: %s\n", admins.ID, admins.Name, admins.Email, admins.Phone, admins.Voucher)

// // Ambil semua data item keranjang dan log informasinya
// var cartItems []CartItems
// if err := ar.conn.WithContext(ctx).Find(&cartItems).Error; err != nil {
// 	return admin.ItemTransactionsDomain{}, err
// }

// // Log informasi tentang setiap item dalam keranjang
// for _, items := range cartItems {
// 	log.Printf("Cart Item - ID: %d, CustomerID: %d, StockID: %d, Quantity: %d, Price: %d, SubTotal: %d\n",
// 		items.ID, items.CustomerID, items.StockID, items.Quantity, items.Price, items.SubTotal)
// }

// Ambil semua data item keranjang berdasarkan customerID dan log informasinya
// var cartItems []CartItems
// if err := ar.conn.WithContext(ctx).
// 	Where("customer_id = ?", reminderPurchaseOrderDomain.CartItem.CustomerID).
// 	Find(&cartItems).Error; err != nil {
// 	return admin.ItemTransactionsDomain{}, err
// }

// // Log informasi tentang setiap item dalam keranjang
// for _, items := range cartItems {
// 	log.Printf("Cart Item - ID: %d, CustomerID: %d, StockID: %d, Quantity: %d, Price: %d, SubTotal: %d\n",
// 		items.ID, items.CustomerID, items.StockID, items.Quantity, items.Price, items.SubTotal)
// }

// // Konversi record ke domain
// packagingOfficerDomain := packagingOfficerRecord.ToPackagingOfficerDomain()

// utils.ScheduleNotification(*reminderPurchaseOrderDomain, packagingOfficerDomain) // Menggunakan nomor telepon Packaging Officer untuk mengirim pesan

// Panggil fungsi untuk mengirimkan notifikasi WhatsApp ke Packaging Officer
// admin, err := ar.AdminGetByID(ctx, reminderPurchaseOrderDomain.AdminID)
// if err != nil {
// 	return admin.ItemTransactionsDomain{}, err
// }
// func (ar *adminRepository) ReminderPurchaseOrderCreate(ctx context.Context, reminderPurchaseOrderDomain *admin.ReminderPurchaseOrderDomain) (admin.ItemTransactionsDomain, error) {
// 	// Ambil semua data item keranjang berdasarkan customerID
// 	var cartItemsData []CartItems
// 	if err := ar.conn.WithContext(ctx).
// 		Where("customer_id = ?", reminderPurchaseOrderDomain.CartItem.CustomerID).
// 		Find(&cartItemsData).Error; err != nil {
// 		return admin.ItemTransactionsDomain{}, err
// 	}

// 	// Jika tidak ada item dalam keranjang, kembalikan error
// 	if len(cartItemsData) == 0 {
// 		return admin.ItemTransactionsDomain{}, fmt.Errorf("no cart items found for customer with ID %d", reminderPurchaseOrderDomain.CartItem.CustomerID)
// 	}

// 	for _, cartItems := range cartItemsData {
// 		// Periksa stok yang terkait dengan penjualan
// 		var stock Stocks
// 		if err := ar.conn.WithContext(ctx).Where("id = ?", cartItems.StockID).First(&stock).Error; err != nil {
// 			return admin.ItemTransactionsDomain{}, err
// 		}

// 		// Kurangi stok dengan jumlah yang dijual
// 		if stock.StockTotal < cartItems.Quantity {
// 			errMsg := fmt.Sprintf("stok tidak cukup untuk penjualan dengan stock_id %d", cartItems.StockID)
// 			log.Println(errMsg)
// 			return admin.ItemTransactionsDomain{}, fmt.Errorf(errMsg)
// 		}
// 		stock.StockTotal -= cartItems.Quantity

// 		// Simpan perubahan stok ke database
// 		if err := ar.conn.WithContext(ctx).Save(&stock).Error; err != nil {
// 			return admin.ItemTransactionsDomain{}, err
// 		}

// 		// Buat catatan itemTransactions
// 		itemTransactionsRecord := ItemTransactions{
// 			CustomerID: cartItems.CustomerID,
// 			StockID:    cartItems.StockID,
// 			UnitID:     stock.UnitID,
// 			CategoryID: stock.CategoryID,
// 			Quantity:   cartItems.Quantity,
// 			Price:      cartItems.Price,
// 			SubTotal:   cartItems.SubTotal,
// 			// Sesuaikan dengan kolom lain jika perlu
// 		}

// 		// Simpan catatan itemTransactions
// 		if err := ar.conn.WithContext(ctx).Create(&itemTransactionsRecord).Error; err != nil {
// 			return admin.ItemTransactionsDomain{}, err
// 		}
// 	}

// 	// Hapus semua data item keranjang yang terkait dengan customerID
// 	if err := ar.conn.WithContext(ctx).Where("customer_id = ?", reminderPurchaseOrderDomain.CartItem.CustomerID).Unscoped().Delete(&CartItems{}).Error; err != nil {
// 		return admin.ItemTransactionsDomain{}, err
// 	}

// 	// Panggil fungsi untuk mengirimkan notifikasi WhatsApp ke Packaging Officer
// 	admins, err := ar.AdminGetByID(ctx, reminderPurchaseOrderDomain.AdminID)
// 	if err != nil {
// 		return admin.ItemTransactionsDomain{}, err
// 	}

// 	utils.ScheduleNotification(*reminderPurchaseOrderDomain, admins) // Menggunakan nomor telepon Admin untuk mengirim pesan ke Packaging Officer

// 	return admin.ItemTransactionsDomain{}, nil
// }

/////////////
// func (ar *adminRepository) ReminderPurchaseOrderCreate(ctx context.Context, reminderPurchaseOrderDomain *admin.ReminderPurchaseOrderDomain) (admin.ItemTransactionsDomain, error) {
// 	// Ambil semua data item keranjang berdasarkan customerID
// 	var cartItemsData []CartItems
// 	if err := ar.conn.WithContext(ctx).
// 		Where("customer_id = ?", reminderPurchaseOrderDomain.CartItem.CustomerID).
// 		Find(&cartItemsData).Error; err != nil {
// 		return admin.ItemTransactionsDomain{}, err
// 	}

// 	// Jika tidak ada item dalam keranjang, kembalikan error
// 	if len(cartItemsData) == 0 {
// 		return admin.ItemTransactionsDomain{}, fmt.Errorf("no cart items found for customer with ID %d", reminderPurchaseOrderDomain.CartItem.CustomerID)
// 	}

// 	for _, cartItems := range cartItemsData {
// 		// Periksa stok yang terkait dengan penjualan
// 		var stock Stocks
// 		if err := ar.conn.WithContext(ctx).Where("id = ?", cartItems.StockID).First(&stock).Error; err != nil {
// 			return admin.ItemTransactionsDomain{}, err
// 		}

// 		// Kurangi stok dengan jumlah yang dijual
// 		if stock.StockTotal < cartItems.Quantity {
// 			errMsg := fmt.Sprintf("stok tidak cukup untuk penjualan dengan stock_id %d", cartItems.StockID)
// 			log.Println(errMsg)
// 			return admin.ItemTransactionsDomain{}, fmt.Errorf(errMsg)
// 		}
// 		stock.StockTotal -= cartItems.Quantity

// 		// Simpan perubahan stok ke database
// 		if err := ar.conn.WithContext(ctx).Save(&stock).Error; err != nil {
// 			return admin.ItemTransactionsDomain{}, err
// 		}

// 		// Buat catatan itemTransactions
// 		itemTransactionsRecord := ItemTransactions{
// 			CustomerID: cartItems.CustomerID,
// 			StockID:    cartItems.StockID,
// 			UnitID:     stock.UnitID,
// 			CategoryID: stock.CategoryID,
// 			Quantity:   cartItems.Quantity,
// 			Price:      cartItems.Price,
// 			SubTotal:   cartItems.SubTotal,
// 			// Sesuaikan dengan kolom lain jika perlu
// 		}

// 		// Simpan catatan itemTransactions
// 		if err := ar.conn.WithContext(ctx).Create(&itemTransactionsRecord).Error; err != nil {
// 			return admin.ItemTransactionsDomain{}, err
// 		}
// 	}

// 	// Hapus semua data item keranjang yang terkait dengan customerID
// 	if err := ar.conn.WithContext(ctx).Where("customer_id = ?", reminderPurchaseOrderDomain.CartItem.CustomerID).Unscoped().Delete(&CartItems{}).Error; err != nil {
// 		return admin.ItemTransactionsDomain{}, err
// 	}

// 	// Panggil fungsi untuk mengirimkan notifikasi WhatsApp ke Packaging Officer
// 	admin, err := ar.AdminGetByID(ctx, reminderPurchaseOrderDomain.AdminID)
// 	if err != nil {
// 		return admin.ItemTransactionsDomain{}, err
// 	}

// 	reminderPreOrder := ReminderPreOrder{
// 		ProductID:     reminderPurchaseOrderDomain.CartItem.StockID,
// 		Quantity:      reminderPurchaseOrderDomain.CartItem.Quantity,
// 		CustomerName:  reminderPurchaseOrderDomain.CartItem.Customers.CustomerName,
// 		Packaging:     PackagingOfficer{},
// 		ReminderTime:  time.Now().Add(24 * time.Hour), // Misalnya dijadwalkan untuk 1 hari ke depan
// 		AdminPhone:    admin.Phone,
// 		PackagingType: reminderPurchaseOrderDomain.CartItem.Stocks.CategoryName,
// 	}
// 	scheduleNotification(reminderPreOrder, admin)

// 	return admin.ItemTransactionsDomain{}, nil
// }

// Panggil fungsi untuk mengirimkan notifikasi WhatsApp ke Packaging Officer
// admin, err := ar.AdminGetByID(ctx, reminderPurchaseOrderDomain.AdminID)
// if err != nil {
// 	return admin.ItemTransactionsDomain{}, err
// }

// reminderPreOrder := ReminderPreOrder{
// 	ProductID:    reminderPurchaseOrderDomain.CartItem.StockID,
// 	Quantity:     reminderPurchaseOrderDomain.CartItem.Quantity,
// 	CustomerName: reminderPurchaseOrderDomain.CartItem.Customers.CustomerName,
// 	Packaging: PackagingOfficer{
// 		OfficerPhone: admin.Phone, // Gunakan nomor telepon Admin untuk mengirim pesan ke Packaging Officer
// 	},
// 	ReminderTime: time.Now(), // Atur waktu pengingat, bisa disesuaikan
// }

// scheduleNotification(reminderPreOrder, admin)

// func (cr *adminRepository) ItemTransactionsGetAll(ctx context.Context) ([]admin.ItemTransactionsDomain, error) {
// 	var records []ItemTransactions
// 	if err := cr.conn.WithContext(ctx).
// 		Preload("Customers").
// 		Preload("Stocks").
// 		Preload("Stocks.Units").
// 		Preload("Stocks.Categories").
// 		Find(&records).Error; err != nil {
// 		return nil, err
// 	}

// 	itemTransactionsDomain := []admin.ItemTransactionsDomain{}

// 	for _, stockHistory := range records {
// 		itemTransactionsDomain = append(itemTransactionsDomain, stockHistory.ToItemTransactionsDomain())
// 	}

// 	return itemTransactionsDomain, nil
// }

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

// func (ar *adminRepository) CartsGetByID(ctx context.Context, id string) (admin.CartsDomain, error) {
// 	var cart Carts

// 	if err := ar.conn.WithContext(ctx).Preload("CartItems").First(&cart, "id = ?", id).Error; err != nil {
// 		return admin.CartsDomain{}, err
// 	}

// 	return cart.ToCartsDomain(), nil

// }

// func (ar *adminRepository) CartsGetAll(ctx context.Context) ([]admin.CartsDomain, error) {
// 	var records []Carts
// 	if err := ar.conn.WithContext(ctx).Preload("CartItems").
// 		// if err := ar.conn.WithContext(ctx).
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

// func (ar *adminRepository) CartsDelete(ctx context.Context, id string) error {
// 	items, err := ar.CartsGetByID(ctx, id)

// 	if err != nil {
// 		return err
// 	}

// 	deletedItems := FromCartsDomain(&items)

// 	// Ambil semua penjualan lainnya
// 	var allItems []Carts
// 	err = ar.conn.WithContext(ctx).Find(&allItems).Error
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
// 	// 		if err := ar.conn.WithContext(ctx).Save(&item).Error; err != nil {
// 	// 			return err
// 	// 		}
// 	// 	}
// 	// }

// 	// Hapus data penjualan
// 	if err := ar.conn.WithContext(ctx).Unscoped().Delete(&deletedItems).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
