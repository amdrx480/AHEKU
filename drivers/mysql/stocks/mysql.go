package stocks

import (
	"backend-golang/businesses/stocks"
	"context"

	// "fmt"

	// _dbCategory "backend-golang/drivers/mysql/category"

	"gorm.io/gorm"
)

type stockRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) stocks.Repository {
	return &stockRepository{
		conn: conn,
	}
}

func (ur *stockRepository) GetByID(ctx context.Context, id string) (stocks.Domain, error) {
	var stock Stock

	if err := ur.conn.WithContext(ctx).First(&stock, "id = ?", id).Error; err != nil {
		return stocks.Domain{}, err
	}

	return stock.ToDomain(), nil

}

func (cr *stockRepository) Create(ctx context.Context, stockDomain *stocks.Domain) (stocks.Domain, error) {

	// // Cari Category berdasarkan CategoryName yang diberikan
	// var category _dbCategory.Category
	// if err := cr.conn.WithContext(ctx).Where("category_name = ?", stockDomain.CategoryName).First(&category).Error; err != nil {
	// 	// Jika Category tidak ditemukan, kembalikan kesalahan
	// 	if err == gorm.ErrRecordNotFound {
	// 		return stocks.Domain{}, fmt.Errorf("Category not found: %w", err)
	// 	}
	// 	return stocks.Domain{}, fmt.Errorf("Failed to fetch category: %w", err)
	// }

	// // Set CategoryID ke stockDomain berdasarkan Category yang ditemukan
	// stockDomain.CategoryID = category.ID
	// // stockDomain.CategoryName = category.CategoryName

	record := FromDomain(stockDomain)
	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return stocks.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return stocks.Domain{}, err
	}

	return record.ToDomain(), nil

}

func (sr *stockRepository) GetAll(ctx context.Context) ([]stocks.Domain, error) {
	// var records []Stock
	// if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
	// 	return nil, err
	// }
	var records []Stock
	if err := sr.conn.WithContext(ctx).
		Preload("Category").Preload("Units").
		Find(&records).Error; err != nil {
		return nil, err
	}

	stocksDomain := []stocks.Domain{}

	for _, stocks := range records {
		domain := stocks.ToDomain()
		stocksDomain = append(stocksDomain, domain)
	}

	return stocksDomain, nil
}

// func (ur *stockRepository) DownloadBarcodeByID(ctx context.Context, id string) (stocks.Domain, error) {
// 	var stock Stock

// 	if err := ur.conn.WithContext(ctx).First(&stock, "id = ?", id).Error; err != nil {
// 		return stocks.Domain{}, err
// 	}

// 	return stock.ToDomain(), nil

// }

// func (cr *stockRepository) StockIn(ctx context.Context, stockDomain *stocks.Domain, id string) (stocks.Domain, error) {
// 	stock, err := cr.DownloadBarcodeByID(ctx, id)

// 	if err != nil {
// 		return stocks.Domain{}, err
// 	}

// 	updateStock := FromDomain(&stock)

// 	updateStock.Stock_Total += stockDomain.Stock_In

// 	if err := cr.conn.WithContext(ctx).Save(&updateStock).Error; err != nil {
// 		return stocks.Domain{}, err
// 	}

// 	return updateStock.ToDomain(), nil
// }

// func (cr *stockRepository) StockOut(ctx context.Context, stockDomain *stocks.Domain, id string) (stocks.Domain, error) {
// 	stock, err := cr.DownloadBarcodeByID(ctx, id)

// 	if err != nil {
// 		return stocks.Domain{}, err
// 	}

// 	updateStock := FromDomain(&stock)

// 	updateStock.Stock_Total -= stockDomain.Stock_Out

// 	if err := cr.conn.WithContext(ctx).Save(&updateStock).Error; err != nil {
// 		return stocks.Domain{}, err
// 	}

// 	return updateStock.ToDomain(), nil
// }
