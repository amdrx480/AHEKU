package response

import (
	"backend-golang/businesses/admin"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	ImagePath string         `json:"image_path"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	// Role      admin.RoleDomain
	RoleID   uint   `json:"role_id"`
	RoleName string `json:"role_name"`
	Voucher  string `json:"voucher"`
	Password string `json:"password"`
}

// Role           string         `json:"role"`

type Role struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	RoleName  string         `json:"role_name"`
}

// type AdminProfile struct {
// 	ID        uint           `json:"id" gorm:"primaryKey"`
// 	CreatedAt time.Time      `json:"created_at"`
// 	UpdatedAt time.Time      `json:"updated_at"`
// 	DeletedAt gorm.DeletedAt `json:"deleted_at"`
// 	Name      string         `json:"name"`
// 	Nip       string         `json:"nip"`
// 	Division  string         `json:"division"`
// }

type Customers struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
	CustomerName    string         `json:"customer_name"`
	CustomerAddress string         `json:"customer_address"`
	CustomerEmail   string         `json:"customer_email"`
	CustomerPhone   string         `json:"customer_phone"`
	CartItems       []CartItems    `json:"cart_items"`
}

type PackagingOfficer struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	OfficerName    string         `json:"officer_name" validate:"required"`
	OfficerAddress string         `json:"officer_address" validate:"required"`
	OfficerEmail   string         `json:"officer_email" validate:"required"`
	OfficerPhone   string         `json:"officer_phone" validate:"required"`
}

type Category struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	CategoryName string         `json:"category_name"`
}

type Vendors struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	VendorName    string         `json:"vendor_name"`
	VendorAddress string         `json:"vendor_address"`
	VendorEmail   string         `json:"vendor_email"`
	VendorPhone   string         `json:"vendor_phone"`
}

type Units struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	UnitName  string         `json:"unit_name"`
}

type Stocks struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	StockName    string         `json:"stock_name"`
	StockCode    string         `json:"stock_code"`
	CategoryName string         `json:"category_name"`
	CategoryID   uint           `json:"category_id"`
	UnitID       uint           `json:"unit_id"`
	UnitName     string         `json:"unit_name"`
	Description  string         `json:"description"`
	StockTotal   int            `json:"stock_total"`
	SellingPrice int            `json:"selling_price"`
}

type Purchases struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	VendorID       uint           `json:"vendor_id"`
	VendorName     string         `json:"vendor_name"`
	StockName      string         `json:"stock_name"`
	StockCode      string         `json:"stock_code"`
	CategoryID     uint           `json:"category_id"`
	CategoryName   string         `json:"category_name"`
	UnitID         uint           `json:"unit_id"`
	UnitName       string         `json:"unit_name"`
	Quantity       int            `json:"quantity"`
	Description    string         `json:"description"`
	Purchase_Price int            `json:"purchase_price"`
	SellingPrice   int            `json:"selling_price"`
}

type CartItems struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	CustomerID    uint           `json:"customer_id"`
	CustomerName  string         `json:"customer_name"`
	StockID       uint           `json:"stock_id"`
	StockName     string         `json:"stock_name"`
	UnitID        uint           `json:"unit_id"`
	UnitName      string         `json:"unit_name"`
	Quantity      int            `json:"quantity"`
	Selling_Price int            `json:"selling_price"`
	Price         int            `json:"price"`
	SubTotal      int            `json:"sub_total"`
}

type ItemTransactions struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CustomerID   uint           `json:"customer_id"`
	CustomerName string         `json:"customer_name"`
	StockID      uint           `json:"stock_id"`
	StockName    string         `json:"stock_name"`
	StockCode    string         `json:"stock_code"`
	UnitID       uint           `json:"unit_id"`
	UnitName     string         `json:"unit_name"`
	CategoryID   uint           `json:"category_id"`
	CategoryName string         `json:"category_name"`
	Quantity     int            `json:"quantity"`
	SubTotal     int            `json:"sub_total"`
}

// type ReminderPurchaseOrder struct {
// 	ID             uint           `json:"id" gorm:"primaryKey"`
// 	CreatedAt      time.Time      `json:"created_at"`
// 	UpdatedAt      time.Time      `json:"updated_at"`
// 	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
// 	OfficerName    string         `json:"officer_name"`
// 	OfficerAddress string         `json:"officer_address"`
// 	OfficerEmail   string         `json:"officer_email"`
// 	OfficerPhone   string         `json:"officer_phone"`
// 	CustomerName   string         `json:"customer_name"`
// 	StockName      string         `json:"stock_name"`
// 	Quantity       int            `json:"quantity"`
// }

type ReminderPurchaseOrder struct {
	ID           uint           `json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	AdminID      string         `json:"admin_id" `
	CartItem     string         // Asumsikan struktur CartItems sesuai dengan yang Anda definisikan sebelumnya
	ReminderTime time.Time      `json:"reminder_time"`
	Packaging    string         `json:"packaging"` // Asumsikan struktur Packaging sudah didefinisikan
}

// Pagination struct holds pagination information
// type Pagination struct {
// 	TotalItems  int `json:"totalItems"`
// 	TotalPages  int `json:"totalPages"`
// 	CurrentPage int `json:"currentPage"`
// 	PageSize    int `json:"pageSize"`
// }

func FromAdminsDomain(domain admin.AdminDomain) Admin {
	return Admin{
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		ID:        domain.ID,
		ImagePath: domain.ImagePath,
		Name:      domain.Name,
		Email:     domain.Email,
		Phone:     domain.Phone,
		RoleID:    domain.RoleID,
		RoleName:  domain.RoleName,
		Voucher:   domain.Voucher,
		Password:  domain.Password,
	}
}

func FromAdminUpdateDomain(domain admin.AdminDomain) Admin {
	return Admin{
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		ID:        domain.ID,
		ImagePath: domain.ImagePath,
		Name:      domain.Name,
		Email:     domain.Email,
		Phone:     domain.Phone,
		// Role: FromAdminUpdateDomain(),
		RoleID:   domain.RoleID,
		RoleName: domain.RoleName,
		Voucher:  domain.Voucher,
	}
}

func FromRoleDomain(domain admin.RoleDomain) Role {
	return Role{
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		ID:        domain.ID,
		RoleName:  domain.RoleName,
	}
}

// func FromAdminProfileDomain(domain admin.AdminProfileDomain) AdminProfile {
// 	return AdminProfile{
// 		ID:        domain.ID,
// 		Name:      domain.Name,
// 		Nip:       domain.Nip,
// 		Division:  domain.Division,
// 		CreatedAt: domain.CreatedAt,
// 		UpdatedAt: domain.UpdatedAt,
// 		DeletedAt: domain.DeletedAt,
// 	}
// }

func FromCustomersDomain(domain admin.CustomersDomain) Customers {
	cartItems := []CartItems{}
	for _, item := range domain.CartItems {
		cartItems = append(cartItems, FromCartItemsDomain(item))
	}
	return Customers{
		ID:              domain.ID,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
		DeletedAt:       domain.DeletedAt,
		CustomerName:    domain.CustomerName,
		CustomerAddress: domain.CustomerAddress,
		CustomerEmail:   domain.CustomerEmail,
		CustomerPhone:   domain.CustomerPhone,
		CartItems:       cartItems,
	}
}

func FromPackagingOfficerDomain(domain admin.PackagingOfficerDomain) PackagingOfficer {
	return PackagingOfficer{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		OfficerName:    domain.OfficerName,
		OfficerAddress: domain.OfficerAddress,
		OfficerEmail:   domain.OfficerEmail,
		OfficerPhone:   domain.OfficerPhone,
	}
}

func FromCategoryDomain(domain admin.CategoriesDomain) Category {
	return Category{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		CategoryName: domain.CategoryName,
	}
}

func FromVendorsDomain(domain admin.VendorsDomain) Vendors {
	return Vendors{
		ID:            domain.ID,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		DeletedAt:     domain.DeletedAt,
		VendorName:    domain.VendorName,
		VendorAddress: domain.VendorAddress,
		VendorEmail:   domain.VendorEmail,
		VendorPhone:   domain.VendorPhone,
	}
}

func FromUnitsDomain(domain admin.UnitsDomain) Units {
	return Units{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		UnitName:  domain.UnitName,
	}
}

func FromStocksDomain(domain admin.StocksDomain) Stocks {
	return Stocks{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		StockName:    domain.StockName,
		StockCode:    domain.StockCode,
		CategoryName: domain.CategoryName,
		CategoryID:   domain.CategoryID,
		UnitName:     domain.UnitName,
		UnitID:       domain.UnitID,
		Description:  domain.Description,
		StockTotal:   domain.StockTotal,
		SellingPrice: domain.SellingPrice,
	}
}

func FromPurchasesDomain(domain admin.PurchasesDomain) Purchases {
	return Purchases{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		VendorID:       domain.VendorID,
		VendorName:     domain.VendorName,
		StockName:      domain.StockName,
		StockCode:      domain.StockCode,
		CategoryID:     domain.CategoryID,
		CategoryName:   domain.CategoryName,
		UnitID:         domain.UnitID,
		UnitName:       domain.UnitName,
		Quantity:       domain.Quantity,
		Description:    domain.Description,
		Purchase_Price: domain.PurchasePrice,
		SellingPrice:   domain.SellingPrice,
	}
}

func FromCartItemsDomain(domain admin.CartItemsDomain) CartItems {
	return CartItems{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		CustomerID:   domain.CustomerID,
		CustomerName: domain.CustomerName,
		StockID:      domain.StockID,
		StockName:    domain.StockName,
		UnitID:       domain.UnitID,
		UnitName:     domain.UnitName,
		Quantity:     domain.Quantity,
		Price:        domain.Price,
		SubTotal:     domain.SubTotal,
	}
}

func FromItemTransactionsDomain(domain admin.ItemTransactionsDomain) ItemTransactions {
	return ItemTransactions{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		DeletedAt:    domain.DeletedAt,
		CustomerID:   domain.CustomerID,
		CustomerName: domain.CustomerName,
		StockID:      domain.StockID,
		StockName:    domain.StockName,
		StockCode:    domain.StockCode,
		UnitID:       domain.UnitID,
		UnitName:     domain.UnitName,
		CategoryID:   domain.CategoryID,
		CategoryName: domain.CategoryName,
		Quantity:     domain.Quantity,
		SubTotal:     domain.SubTotal,
	}
}

func FromReminderPurchaseOrderDomain(domain admin.ReminderPurchaseOrderDomain) ReminderPurchaseOrder {
	return ReminderPurchaseOrder{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		// AdminID:   domain.AdminID,
		CartItem: domain.CartItem.CustomerName,
	}
}
