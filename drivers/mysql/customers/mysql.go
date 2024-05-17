package customers

import (
	"backend-golang/businesses/customers"
	"context"

	// "errors"

	// "fmt"

	"gorm.io/gorm"
)

type customersRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) customers.Repository {
	return &customersRepository{
		conn: conn,
	}
}

func (cr *customersRepository) GetByID(ctx context.Context, id string) (customers.Domain, error) {
	var customer Customers

	if err := cr.conn.WithContext(ctx).First(&customer, "id = ?", id).Error; err != nil {
		return customers.Domain{}, err
	}

	return customer.ToDomain(), nil

}

func (cr *customersRepository) Create(ctx context.Context, customersDomain *customers.Domain) (customers.Domain, error) {
	record := FromDomain(customersDomain)
	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return customers.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return customers.Domain{}, err
	}

	return record.ToDomain(), nil

}

func (sr *customersRepository) GetAll(ctx context.Context) ([]customers.Domain, error) {
	var records []Customers
	if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []customers.Domain{}

	for _, category := range records {
		domain := category.ToDomain()
		categories = append(categories, domain)
	}

	return categories, nil
}

// func (ur *customersRepository) DownloadBarcodeByID(ctx context.Context, id string) (customers.Domain, error) {
// 	var customers customers

// 	if err := ur.conn.WithContext(ctx).First(&customers, "id = ?", id).Error; err != nil {
// 		return customers.Domain{}, err
// 	}

// 	return customers.ToDomain(), nil

// }

// func (cr *customersRepository) customersIn(ctx context.Context, customersDomain *customers.Domain, id string) (customers.Domain, error) {
// 	customers, err := cr.DownloadBarcodeByID(ctx, id)

// 	if err != nil {
// 		return customers.Domain{}, err
// 	}

// 	updatecustomers := FromDomain(&customers)

// 	updatecustomers.customers_Total += customersDomain.customers_In

// 	if err := cr.conn.WithContext(ctx).Save(&updatecustomers).Error; err != nil {
// 		return customers.Domain{}, err
// 	}

// 	return updatecustomers.ToDomain(), nil
// }

// func (cr *customersRepository) customersOut(ctx context.Context, customersDomain *customers.Domain, id string) (customers.Domain, error) {
// 	customers, err := cr.DownloadBarcodeByID(ctx, id)

// 	if err != nil {
// 		return customers.Domain{}, err
// 	}

// 	updatecustomers := FromDomain(&customers)

// 	updatecustomers.customers_Total -= customersDomain.customers_Out

// 	if err := cr.conn.WithContext(ctx).Save(&updatecustomers).Error; err != nil {
// 		return customers.Domain{}, err
// 	}

// 	return updatecustomers.ToDomain(), nil
// }
