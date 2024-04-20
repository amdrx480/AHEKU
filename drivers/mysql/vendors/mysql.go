package vendors

import (
	"backend-golang/businesses/vendors"
	"context"

	// "errors"

	// "fmt"

	"gorm.io/gorm"
)

type vendorsRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) vendors.Repository {
	return &vendorsRepository{
		conn: conn,
	}
}

func (vr *vendorsRepository) GetByID(ctx context.Context, id string) (vendors.Domain, error) {
	var vendor Vendors

	if err := vr.conn.WithContext(ctx).First(&vendor, "id = ?", id).Error; err != nil {
		return vendors.Domain{}, err
	}

	return vendor.ToDomain(), nil

}

func (vr *vendorsRepository) Create(ctx context.Context, purchaseDomain *vendors.Domain) (vendors.Domain, error) {
	record := FromDomain(purchaseDomain)
	result := vr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return vendors.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return vendors.Domain{}, err
	}

	return record.ToDomain(), nil

}

func (sr *vendorsRepository) GetAll(ctx context.Context) ([]vendors.Domain, error) {
	var records []Vendors
	if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []vendors.Domain{}

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
