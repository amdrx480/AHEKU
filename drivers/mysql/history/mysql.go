package history

import (
	"backend-golang/businesses/history"
	"context"

	"gorm.io/gorm"
)

type historyRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) history.Repository {
	return &historyRepository{
		conn: conn,
	}
}

func (cr *historyRepository) GetAll(ctx context.Context) ([]history.Domain, error) {
	var records []History

	err := cr.conn.WithContext(ctx).Find(&records).Error

	if err != nil {
		return nil, err
	}

	stockhistory := []history.Domain{}

	for _, stockHistory := range records {
		stockhistory = append(stockhistory, stockHistory.ToDomain())
	}

	return stockhistory, nil
}

func (cr *historyRepository) Create(ctx context.Context, courseDomain *history.Domain) (history.Domain, error) {
	record := FromDomain(courseDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return history.Domain{}, err
	}

	err := cr.conn.WithContext(ctx).Last(&record).Error

	if err != nil {
		return history.Domain{}, err
	}

	return record.ToDomain(), nil
}
