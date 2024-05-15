package sales

import (
	"backend-golang/businesses/history"

	"backend-golang/app/middlewares"
	"context"
)

type salesUsecase struct {
	salesRepository Repository
	jwtAuth         *middlewares.JWTConfig
}

func NewSalesUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &salesUsecase{
		salesRepository: repository,
		jwtAuth:         jwtAuth,
	}
}

func (usecase *salesUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.salesRepository.GetByID(ctx, id)
}

func (usecase *salesUsecase) Create(ctx context.Context, salesDomain *Domain) (Domain, error) {
	return usecase.salesRepository.Create(ctx, salesDomain)
}

func (usecase *salesUsecase) ToHistory(ctx context.Context, historyDomain *history.Domain, id string) (history.Domain, error) {
	return usecase.salesRepository.ToHistory(ctx, historyDomain, id)
}

func (usecase *salesUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.salesRepository.GetAll(ctx)
}

func (usecase *salesUsecase) Delete(ctx context.Context, id string) error {
	return usecase.salesRepository.Delete(ctx, id)
}
