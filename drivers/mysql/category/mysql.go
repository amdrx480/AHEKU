package category

import (
	"backend-golang/businesses/category"
	"context"

	// "errors"

	// "fmt"

	"gorm.io/gorm"
)

type categoryRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) category.Repository {
	return &categoryRepository{
		conn: conn,
	}
}

func (cr *categoryRepository) GetByID(ctx context.Context, id string) (category.Domain, error) {
	var categories Category

	if err := cr.conn.WithContext(ctx).First(&categories, "id = ?", id).Error; err != nil {
		return category.Domain{}, err
	}

	return categories.ToDomain(), nil

}

func (cr *categoryRepository) GetByName(ctx context.Context, name string) (category.Domain, error) {
	var categories Category

	if err := cr.conn.WithContext(ctx).First(&categories, "category_name = ?", name).Error; err != nil {
		return category.Domain{}, err
	}

	return categories.ToDomain(), nil

}

func (cr *categoryRepository) Create(ctx context.Context, categoryDomain *category.Domain) (category.Domain, error) {
	record := FromDomain(categoryDomain)
	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return category.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return category.Domain{}, err
	}

	return record.ToDomain(), nil

}

func (cr *categoryRepository) GetAll(ctx context.Context) ([]category.Domain, error) {
	var records []Category
	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []category.Domain{}

	for _, category := range records {
		domain := category.ToDomain()
		categories = append(categories, domain)
	}

	return categories, nil
}

// func (ur *purchaseRepository) DownloadBarcodeByID(ctx context.Context, id string) (purchase.Domain, error) {
// 	var purchase Purchase

// 	if err := ur.conn.WithContext(ctx).First(&purchase, "id = ?", id).Error; err != nil {
// 		return purchase.Domain{}, err
// 	}

// 	return purchase.ToDomain(), nil

// }

// func (cr *purchaseRepository) PurchaseIn(ctx context.Context, purchaseDomain *purchase.Domain, id string) (purchase.Domain, error) {
// 	purchase, err := cr.DownloadBarcodeByID(ctx, id)

// 	if err != nil {
// 		return purchase.Domain{}, err
// 	}

// 	updatePurchase := FromDomain(&purchase)

// 	updatePurchase.Purchase_Total += purchaseDomain.Purchase_In

// 	if err := cr.conn.WithContext(ctx).Save(&updatePurchase).Error; err != nil {
// 		return purchase.Domain{}, err
// 	}

// 	return updatePurchase.ToDomain(), nil
// }

// func (cr *purchaseRepository) PurchaseOut(ctx context.Context, purchaseDomain *purchase.Domain, id string) (purchase.Domain, error) {
// 	purchase, err := cr.DownloadBarcodeByID(ctx, id)

// 	if err != nil {
// 		return purchase.Domain{}, err
// 	}

// 	updatePurchase := FromDomain(&purchase)

// 	updatePurchase.Purchase_Total -= purchaseDomain.Purchase_Out

// 	if err := cr.conn.WithContext(ctx).Save(&updatePurchase).Error; err != nil {
// 		return purchase.Domain{}, err
// 	}

// 	return updatePurchase.ToDomain(), nil
// }
