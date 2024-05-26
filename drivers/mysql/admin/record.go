package admin

import (
	"backend-golang/businesses/admin"
	"time"

	"gorm.io/gorm"
)

type Admins struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name" gorm:"unique;not null"`
	Voucher   string         `json:"voucher" gorm:"unique;not null"`
	Password  string         `json:"password"`
}

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
	Vendors  Vendors `json:"-" gorm:"foreignKey:VendorID"`
	VendorID uint    `json:"vendor_id"`
	// VendorName string     `json:"Vendor_name"`
	StockName  string     `json:"stock_name"`
	StockCode  string     `json:"stock_code"`
	Categories Categories `json:"-" gorm:"foreignKey:CategoryID"`
	CategoryID uint       `json:"category_id"`
	// CategoryName   string     `json:"category_name"`
	Units  Units `gorm:"foreignKey:UnitID"`
	UnitID uint  `json:"unit_id"`
	// UnitName       string `json:"unit_name"`
	Description    string `json:"description"`
	Quantity       int    `json:"quantity"`
	PurchasesPrice int    `json:"purchases_price"`
	SellingPrice   int    `json:"selling_price"`
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
	// Categories Categories     `json:"-" gorm:"foreignKey:category_id"`
	// CategoryID uint           `json:"category_id"`
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
	SubTotal int `json:"sub_total"`
}

type ItemTransactions struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CustomerID uint           `json:"customer_id"`
	StockID    uint           `json:"stock_id"`
	Quantity   int            `json:"quantity"`
	SubTotal   int            `json:"sub_total"`
}

func (record *Admins) ToAdminsDomain() admin.AdminsDomain {
	return admin.AdminsDomain{
		ID:        record.ID,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		DeletedAt: record.DeletedAt,
		Name:      record.Name,
		Voucher:   record.Voucher,
		Password:  record.Password,
	}
}

func FromAdminsDomain(domain *admin.AdminsDomain) *Admins {
	return &Admins{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
		Voucher:   domain.Voucher,
		Password:  domain.Password,
	}
}

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
		ID:        record.ID,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		DeletedAt: record.DeletedAt,
		VendorID:  record.VendorID,
		// VendorName:     record.Vendors.VendorName,
		StockName:  record.StockName,
		StockCode:  record.StockCode,
		CategoryID: record.CategoryID,
		// CategoryName:   record.Categories.CategoryName,
		UnitID: record.UnitID,
		// UnitName:       record.Units.UnitName,
		Description:    record.Description,
		Quantity:       record.Quantity,
		PurchasesPrice: record.PurchasesPrice,
		SellingPrice:   record.SellingPrice,
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
		Description:    domain.Description,
		Quantity:       domain.Quantity,
		PurchasesPrice: domain.PurchasesPrice,
		SellingPrice:   domain.SellingPrice,
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
		// CategoryName: record.Categories.CategoryName,
		Quantity: record.Quantity,
		Price:    record.Price,
		SubTotal: record.SubTotal,
		// CartID:    record.CartID,
	}
}
func FromCartItemsDomain(domain *admin.CartItemsDomain) *CartItems {
	return &CartItems{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
		CustomerID: domain.CustomerID,
		// CustomerName: domain.CustomerName,
		// CartID:    domain.CartID,
		StockID: domain.StockID,
		// StockName: domain.StockName,
		Quantity: domain.Quantity,
		Price:    domain.Price,
		SubTotal: domain.SubTotal,
	}
}

func (record *ItemTransactions) ToItemTransactionsDomain() admin.ItemTransactionsDomain {
	return admin.ItemTransactionsDomain{
		ID:         record.ID,
		CreatedAt:  record.CreatedAt,
		DeletedAt:  record.DeletedAt,
		CustomerID: record.CustomerID,
		StockID:    record.StockID,
		Quantity:   record.Quantity,
		SubTotal:   record.SubTotal,
	}
}

func FromItemTransactionsDomain(domain *admin.ItemTransactionsDomain) *ItemTransactions {
	return &ItemTransactions{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		DeletedAt: domain.DeletedAt,
		StockID:   domain.StockID,
		Quantity:  domain.Quantity,
		SubTotal:  domain.SubTotal,
	}
}

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
