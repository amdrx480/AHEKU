package admin

import (
	"backend-golang/businesses/admin"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ImagePath string         `json:"image_path"`
	Name      string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Phone     string         `json:"phone" gorm:"unique;not null"`
	Role      Role           `gorm:"foreignKey:RoleID"`
	RoleID    uint           `json:"role_id"`
	RoleName  string         `json:"role_name"`
	Voucher   string         `json:"voucher" gorm:"unique;not null"`
	Password  string         `json:"password" gorm:"not null"`
}

type Role struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	RoleName  string         `json:"role_name" gorm:"unique;not null"`
}

// RoleName  string         `json:"role_name" gorm:"primaryKey;unique;not null"`
// RoleName  string         `json:"role_name" gorm:"primaryKey;unique;not null"`

// type AdminProfile struct {
// 	// Admin
// 	ID        uint           `json:"id" gorm:"primaryKey"`
// 	CreatedAt time.Time      `json:"created_at"`
// 	UpdatedAt time.Time      `json:"updated_at"`
// 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
// 	// AdminID    uint           `json:"admin_id" gorm:"unique"` // Bidang ini menunjukkan hubungan one-to-one dengan Admin
// 	Name       string `json:"name"`
// 	Nip        string `json:"nip"`
// 	Division   string `json:"division"`
// 	Image_Path string `json:"image_path"`
// }

type Customers struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CustomerName    string         `json:"customer_name" gorm:"unique;not null"`
	CustomerEmail   string         `json:"customer_email"`
	CustomerAddress string         `json:"customer_address"`
	CustomerPhone   string         `json:"customer_phone"`
	CartItems       []CartItems    `gorm:"foreignKey:customer_id"` // Menentukan foreign key
}

type PackagingOfficer struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	OfficerName    string         `json:"officer_name"`
	OfficerAddress string         `json:"officer_address"`
	OfficerEmail   string         `json:"officer_email"`
	OfficerPhone   string         `json:"officer_phone"`
}

type Categories struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CategoryName string         `json:"category_name" gorm:"unique;not null"`
}

type Vendors struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	VendorName    string         `json:"vendor_name" gorm:"unique;not null"`
	VendorAddress string         `json:"vendor_address"`
	VendorEmail   string         `json:"vendor_email"`
	VendorPhone   string         `json:"vendor_phone"`
}

type Units struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	UnitName  string         `json:"unit_name" gorm:"unique;not null"`
}

type Stocks struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	StockName    string         `json:"stock_name"`
	StockCode    string         `json:"stock_code"`
	Categories   Categories     `json:"-" gorm:"foreignKey:CategoryID"`
	CategoryID   uint           `json:"category_id"`
	Units        Units          `json:"-" gorm:"foreignKey:UnitID"`
	UnitID       uint           `json:"unit_id"`
	Description  string         `json:"description"`
	Image_Path   string         `json:"image_path"`
	StockTotal   int            `json:"stock_total"`
	SellingPrice int            `json:"selling_price"`
}

type Purchases struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	//Preload Memuat data Dari isi Nama Tabel Struck
	Vendors       Vendors    `json:"-" gorm:"foreignKey:VendorID"`
	VendorID      uint       `json:"vendor_id"`
	StockName     string     `json:"stock_name"`
	StockCode     string     `json:"stock_code"`
	Categories    Categories `json:"-" gorm:"foreignKey:CategoryID"`
	CategoryID    uint       `json:"category_id"`
	Units         Units      `gorm:"foreignKey:UnitID"`
	UnitID        uint       `json:"unit_id"`
	Description   string     `json:"description"`
	Quantity      int        `json:"quantity"`
	PurchasePrice int        `json:"purchase_price"`
	SellingPrice  int        `json:"selling_price"`
}

type CartItems struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Customers  Customers      `gorm:"foreignKey:customer_id"` // Menentukan foreign key
	CustomerID uint           `json:"customer_id"`
	Stocks     Stocks         `json:"-" gorm:"foreignKey:stock_id"`
	StockID    uint           `json:"stock_id"`
	Quantity   int            `json:"quantity"`
	Price      int            `json:"price"`
	SubTotal   int            `json:"sub_total"`
}

type ItemTransactions struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Customers  Customers      `gorm:"foreignKey:customer_id"` // Menentukan foreign key
	CustomerID uint           `json:"customer_id"`
	Stocks     Stocks         `json:"-" gorm:"foreignKey:stock_id"`
	StockID    uint           `json:"stock_id"`
	UnitID     uint           `json:"unit_id"`
	Categories Categories     `json:"-" gorm:"foreignKey:CategoryID"`
	CategoryID uint           `json:"category_id"`
	Quantity   int            `json:"quantity"`
	Price      int            `json:"price"`
	SubTotal   int            `json:"sub_total"`
}

type ReminderPurchaseOrder struct {
	ID                 uint           `json:"id" gorm:"primaryKey"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	PackagingOfficerID uint           `json:"packaging_officer_id"`
	ReminderTime       time.Time      `json:"reminder_time"`

	// Admins             Admin            `json:"-" gorm:"foreignKey:AdminID"`
	// AdminID            uint      `json:"admin_id" `
	// Packaging          PackagingOfficer `json:"-" gorm:"foreignKey:PackagingOfficerID"` // Asumsikan struktur Packaging sudah didefinisikan

	// Admins             Admin            `json:"-" gorm:"foreignKey:AdminID"`

}

func (record *Admin) ToAdminDomain() admin.AdminDomain {
	// baseURL := os.Getenv("BASE_URL")

	return admin.AdminDomain{
		ID:        record.ID,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		DeletedAt: record.DeletedAt,
		ImagePath: record.ImagePath,
		// ImagePath: fmt.Sprintf("%s/images/%s", baseURL, record.ImagePath),
		Name:     record.Name,
		Email:    record.Email,
		Phone:    record.Phone,
		RoleID:   record.RoleID,
		RoleName: record.Role.RoleName,
		Voucher:  record.Voucher,
		Password: record.Password,
	}
}

func FromAdminDomain(domain *admin.AdminDomain) *Admin {
	return &Admin{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		ImagePath: domain.ImagePath,
		Name:      domain.Name,
		Email:     domain.Email,
		Phone:     domain.Phone,
		RoleID:    domain.RoleID,
		// RoleName:  domain.RoleName,
		Voucher:  domain.Voucher,
		Password: domain.Password,
	}
}

// Role:      domain.Role, // Menggunakan string untuk peran

func (record *Role) ToRoleDomain() admin.RoleDomain {
	return admin.RoleDomain{
		ID:        record.ID,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		DeletedAt: record.DeletedAt,
		RoleName:  record.RoleName,
	}
}
func FromRoleDomain(domain *admin.RoleDomain) *Role {
	return &Role{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		RoleName:  domain.RoleName,
	}
}

// func (record *AdminProfile) ToAdminProfileDomain() admin.AdminProfileDomain {
// 	return admin.AdminProfileDomain{
// 		ID:         record.ID,
// 		CreatedAt:  record.CreatedAt,
// 		UpdatedAt:  record.UpdatedAt,
// 		DeletedAt:  record.DeletedAt,
// 		Name:       record.Name,
// 		Nip:        record.Nip,
// 		Division:   record.Division,
// 		Image_Path: record.Image_Path,
// 	}
// }

// func FromAdminProfileDomain(domain *admin.AdminProfileDomain) *AdminProfile {
// 	return &AdminProfile{
// 		ID:         domain.ID,
// 		CreatedAt:  domain.CreatedAt,
// 		UpdatedAt:  domain.UpdatedAt,
// 		DeletedAt:  domain.DeletedAt,
// 		Name:       domain.Name,
// 		Nip:        domain.Nip,
// 		Division:   domain.Division,
// 		Image_Path: domain.Image_Path,
// 	}
// }

func (record *Customers) ToCustomersDomain() admin.CustomersDomain {
	cartItemsDomain := []admin.CartItemsDomain{}
	for _, item := range record.CartItems {
		cartItemsDomain = append(cartItemsDomain, item.ToCartItemsDomain())
	}
	return admin.CustomersDomain{
		ID:              record.ID,
		CreatedAt:       record.CreatedAt,
		UpdatedAt:       record.UpdatedAt,
		DeletedAt:       record.DeletedAt,
		CustomerName:    record.CustomerName,
		CustomerAddress: record.CustomerAddress,
		CustomerEmail:   record.CustomerEmail,
		CustomerPhone:   record.CustomerPhone,
		CartItems:       cartItemsDomain,
	}
}

func FromCustomersDomain(domain *admin.CustomersDomain) *Customers {
	cartItems := []CartItems{}
	for _, item := range domain.CartItems {
		cartItems = append(cartItems, *FromCartItemsDomain(&item))
	}

	return &Customers{
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

func (record *PackagingOfficer) ToPackagingOfficerDomain() admin.PackagingOfficerDomain {
	return admin.PackagingOfficerDomain{
		ID:             record.ID,
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
		DeletedAt:      record.DeletedAt,
		OfficerName:    record.OfficerName,
		OfficerAddress: record.OfficerAddress,
		OfficerEmail:   record.OfficerEmail,
		OfficerPhone:   record.OfficerPhone,
	}
}

func FromPackagingOfficerDomain(domain *admin.PackagingOfficerDomain) *PackagingOfficer {
	return &PackagingOfficer{
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

func (record *Categories) ToCategoriesDomain() admin.CategoriesDomain {
	return admin.CategoriesDomain{
		ID:           record.ID,
		CreatedAt:    record.CreatedAt,
		UpdatedAt:    record.UpdatedAt,
		DeletedAt:    record.DeletedAt,
		CategoryName: record.CategoryName,
	}
}

func FromCategoriesDomain(domain *admin.CategoriesDomain) *Categories {
	return &Categories{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		CategoryName: domain.CategoryName,
	}
}

func (record *Vendors) ToVendorsDomain() admin.VendorsDomain {
	return admin.VendorsDomain{
		ID:            record.ID,
		CreatedAt:     record.CreatedAt,
		UpdatedAt:     record.UpdatedAt,
		DeletedAt:     record.DeletedAt,
		VendorName:    record.VendorName,
		VendorAddress: record.VendorAddress,
		VendorEmail:   record.VendorEmail,
		VendorPhone:   record.VendorPhone,
	}
}
func FromVendorsDomain(domain *admin.VendorsDomain) *Vendors {
	return &Vendors{
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

func (record *Units) ToUnitsDomain() admin.UnitsDomain {
	return admin.UnitsDomain{
		ID:        record.ID,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		DeletedAt: record.DeletedAt,
		UnitName:  record.UnitName,
	}
}
func FromUnitsDomain(domain *admin.UnitsDomain) *Units {
	return &Units{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		UnitName:  domain.UnitName,
	}
}
func (record *Stocks) ToStocksDomain() admin.StocksDomain {
	return admin.StocksDomain{
		ID:           record.ID,
		CreatedAt:    record.CreatedAt,
		UpdatedAt:    record.UpdatedAt,
		DeletedAt:    record.DeletedAt,
		StockCode:    record.StockCode,
		StockName:    record.StockName,
		CategoryID:   record.CategoryID,
		CategoryName: record.Categories.CategoryName,
		UnitID:       record.UnitID,
		UnitName:     record.Units.UnitName,
		StockTotal:   record.StockTotal,
		SellingPrice: record.SellingPrice,
		Description:  record.Description,
	}
}

func FromStocksDomain(domain *admin.StocksDomain) *Stocks {
	return &Stocks{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		StockName:    domain.StockName,
		StockCode:    domain.StockCode,
		CategoryID:   domain.CategoryID,
		UnitID:       domain.UnitID,
		Description:  domain.Description,
		StockTotal:   domain.StockTotal,
		SellingPrice: domain.SellingPrice,
	}
}

func (record *Purchases) ToPurchasesDomain() admin.PurchasesDomain {
	return admin.PurchasesDomain{
		ID:            record.ID,
		CreatedAt:     record.CreatedAt,
		UpdatedAt:     record.UpdatedAt,
		DeletedAt:     record.DeletedAt,
		VendorID:      record.VendorID,
		VendorName:    record.Vendors.VendorName,
		StockName:     record.StockName,
		StockCode:     record.StockCode,
		CategoryID:    record.CategoryID,
		CategoryName:  record.Categories.CategoryName,
		UnitID:        record.UnitID,
		UnitName:      record.Units.UnitName,
		Description:   record.Description,
		Quantity:      record.Quantity,
		PurchasePrice: record.PurchasePrice,
		SellingPrice:  record.SellingPrice,
	}
}
func FromPurchasesDomain(domain *admin.PurchasesDomain) *Purchases {
	return &Purchases{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		VendorID:  domain.VendorID,
		// VendorName: domain.VendorName,
		StockName:  domain.StockName,
		StockCode:  domain.StockCode,
		CategoryID: domain.CategoryID,
		// CategoryName:   domain.CategoryName,
		UnitID: domain.UnitID,
		// UnitName:       domain.UnitName,
		Description:   domain.Description,
		Quantity:      domain.Quantity,
		PurchasePrice: domain.PurchasePrice,
		SellingPrice:  domain.SellingPrice,
	}
}

func (record *CartItems) ToCartItemsDomain() admin.CartItemsDomain {
	return admin.CartItemsDomain{
		ID:           record.ID,
		CreatedAt:    record.CreatedAt,
		UpdatedAt:    record.UpdatedAt,
		DeletedAt:    record.DeletedAt,
		CustomerID:   record.CustomerID,
		CustomerName: record.Customers.CustomerName,
		StockID:      record.StockID,
		StockName:    record.Stocks.StockName,
		// Stocks berisi Units yang nantinya akan menampilkan data units menggunakan eager loading (`Preload``)
		// ambil id units dari stocks yang berisi one to one dari untis
		UnitID: record.Stocks.UnitID,
		// ambil UnitName units dari stocks yang berisi one to one dari untis
		UnitName: record.Stocks.Units.UnitName,
		Quantity: record.Quantity,
		Price:    record.Price,
		SubTotal: record.SubTotal,
	}
}
func FromCartItemsDomain(domain *admin.CartItemsDomain) *CartItems {
	return &CartItems{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
		CustomerID: domain.CustomerID,
		StockID:    domain.StockID,

		// UnitsID: domain.UnitsID,

		Quantity: domain.Quantity,
		Price:    domain.Price,
		SubTotal: domain.SubTotal,

		// StockName: domain.StockName,
		// CustomerName: domain.CustomerName,
		// CartID:    domain.CartID,
	}
}

func (record *ItemTransactions) ToItemTransactionsDomain() admin.ItemTransactionsDomain {
	return admin.ItemTransactionsDomain{
		ID:           record.ID,
		CreatedAt:    record.CreatedAt,
		DeletedAt:    record.DeletedAt,
		CustomerID:   record.CustomerID,
		CustomerName: record.Customers.CustomerName,
		StockID:      record.StockID,
		StockName:    record.Stocks.StockName,
		StockCode:    record.Stocks.StockCode,
		UnitID:       record.Stocks.UnitID,
		UnitName:     record.Stocks.Units.UnitName,
		CategoryID:   record.Stocks.CategoryID,
		CategoryName: record.Stocks.Categories.CategoryName,
		Quantity:     record.Quantity,
		Price:        record.Price,
		SubTotal:     record.SubTotal,
	}
}

func FromItemTransactionsDomain(domain *admin.ItemTransactionsDomain) *ItemTransactions {
	return &ItemTransactions{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		DeletedAt:  domain.DeletedAt,
		CustomerID: domain.CustomerID,
		StockID:    domain.StockID,
		UnitID:     domain.UnitID,
		CategoryID: domain.CategoryID,
		Quantity:   domain.Quantity,
		Price:      domain.Price,
		SubTotal:   domain.SubTotal,
	}
}

func (record *ReminderPurchaseOrder) ToReminderPurchaseOrderDomain() admin.ReminderPurchaseOrderDomain {
	return admin.ReminderPurchaseOrderDomain{
		ID:        record.ID,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		DeletedAt: record.DeletedAt,
		// AdminID:            record.AdminID,
		PackagingOfficerID: record.PackagingOfficerID,
		ReminderTime:       record.ReminderTime,
	}
}

func FromReminderPurchaseOrderDomain(domain admin.ReminderPurchaseOrderDomain) ReminderPurchaseOrder {
	return ReminderPurchaseOrder{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		// AdminID:            domain.AdminID,
		PackagingOfficerID: domain.PackagingOfficerID,
		ReminderTime:       domain.ReminderTime,
	}
}

// Stocks berisi Units yang nantinya akan menampilkan data units menggunakan eager loading (`Preload``)
// Units      Units          `gorm:"foreignKey:unit_id"`
//akan terjadi constraint foreginKey dari unit_id
// UnitID    uint           `json:"unit_id"`

// AdminProfileID uint           `json:"admin_profile_id" gorm:"not null"`
// AdminProfile   AdminProfile   `json:"admin_profile" gorm:"foreignKey:AdminProfileID"`
// func (record *Carts) ToCartsDomain() admin.CartsDomain {
// 	cartItemsDomain := []admin.CartItemsDomain{}
// 	for _, item := range record.CartItems {
// 		cartItemsDomain = append(cartItemsDomain, item.ToItemsDomain())
// 	}
// 	return admin.CartsDomain{
// 		ID:         record.ID,
// 		CreatedAt:  record.CreatedAt,
// 		UpdatedAt:  record.UpdatedAt,
// 		DeletedAt:  record.DeletedAt,
// 		CustomerID: record.CustomerID,
// 		CartItems:  cartItemsDomain,
// 		Total:      record.Total,
// 	}
// }

// func FromCartsDomain(domain *admin.CartsDomain) *Carts {
// 	cartItems := []CartItems{}
// 	for _, item := range domain.CartItems {
// 		cartItems = append(cartItems, *FromCartItemsDomain(&item))
// 	}

// 	return &Carts{
// 		ID:         domain.ID,
// 		CreatedAt:  domain.CreatedAt,
// 		UpdatedAt:  domain.UpdatedAt,
// 		DeletedAt:  domain.DeletedAt,
// 		CustomerID: domain.CustomerID,
// 		CartItems:  cartItems,
// 		Total:      domain.Total,
// 	}
// }

//	type Admin struct {
//		ID             uint           `json:"id" gorm:"primaryKey"`
//		CreatedAt      time.Time      `json:"created_at"`
//		UpdatedAt      time.Time      `json:"updated_at"`
//		DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
//		Name           string         `json:"name" gorm:"not null"`
//		Email          string         `json:"email" gorm:"unique;not null"`
//		Phone          string         `json:"phone" gorm:"unique;not null"`
//		Role           string         `json:"role"`
//		Voucher        string         `json:"voucher" gorm:"unique;not null"`
//		Password       string         `json:"password"`
//		AdminProfileID uint           `json:"admin_profile_id"`
//		AdminProfile   AdminProfile   `gorm:"foreignKey:AdminProfileID"`
//	}
