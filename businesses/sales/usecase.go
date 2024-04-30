package sales

import (
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

func (usecase *salesUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.salesRepository.GetAll(ctx)
}
