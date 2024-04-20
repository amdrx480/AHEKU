package units

import (
	"backend-golang/businesses/units"
	"context"

	// "errors"

	// "fmt"

	"gorm.io/gorm"
)

type unitsRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) units.Repository {
	return &unitsRepository{
		conn: conn,
	}
}

func (ur *unitsRepository) GetByID(ctx context.Context, id string) (units.Domain, error) {
	var unit Units

	if err := ur.conn.WithContext(ctx).First(&unit, "id = ?", id).Error; err != nil {
		return units.Domain{}, err
	}

	return unit.ToDomain(), nil

}

func (ur *unitsRepository) Create(ctx context.Context, unitsDomain *units.Domain) (units.Domain, error) {
	record := FromDomain(unitsDomain)
	result := ur.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return units.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return units.Domain{}, err
	}

	return record.ToDomain(), nil

}

func (ur *unitsRepository) GetAll(ctx context.Context) ([]units.Domain, error) {
	var records []Units
	if err := ur.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	units := []units.Domain{}

	for _, unit := range records {
		domain := unit.ToDomain()
		units = append(units, domain)
	}

	return units, nil
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
