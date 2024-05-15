package stocks

import (
	"backend-golang/app/middlewares"
	"context"
)

type stockUsecase struct {
	stockRepository Repository
	jwtAuth         *middlewares.JWTConfig
}

func NewStockUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &stockUsecase{
		stockRepository: repository,
		jwtAuth:         jwtAuth,
	}
}

func (usecase *stockUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.stockRepository.GetByID(ctx, id)
}

// func (usecase *stockUsecase) DownloadBarcodeByID(ctx context.Context, id string) (Domain, error) {
// 	return usecase.stockRepository.DownloadBarcodeByID(ctx, id)
// }

func (usecase *stockUsecase) Create(ctx context.Context, stockDomain *Domain) (Domain, error) {
	return usecase.stockRepository.Create(ctx, stockDomain)
}

// func (usecase *stockUsecase) Create(ctx context.Context, stockDomain *Domain, imagePath string, id string) (Domain, string, error) {
// 	return usecase.stockRepository.Create(ctx, stockDomain, imagePath, id)
// }

func (usecase *stockUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.stockRepository.GetAll(ctx)
}

// func (usecase *stockUsecase) StockIn(ctx context.Context, stockDomain *Domain, id string) (Domain, error) {
// 	return usecase.stockRepository.StockIn(ctx, stockDomain, id)
// }

// func (usecase *stockUsecase) StockOut(ctx context.Context, stockDomain *Domain, id string) (Domain, error) {
// 	return usecase.stockRepository.StockOut(ctx, stockDomain, id)
// }
