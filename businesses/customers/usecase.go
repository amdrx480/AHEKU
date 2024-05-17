package customers

import (
	"backend-golang/app/middlewares"
	"context"
)

type customersUsecase struct {
	customersRepository Repository
	jwtAuth             *middlewares.JWTConfig
}

func NewCustomersUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &customersUsecase{
		customersRepository: repository,
		jwtAuth:             jwtAuth,
	}
}

func (usecase *customersUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.customersRepository.GetByID(ctx, id)
}

func (usecase *customersUsecase) Create(ctx context.Context, customersDomain *Domain) (Domain, error) {
	return usecase.customersRepository.Create(ctx, customersDomain)
}

func (usecase *customersUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.customersRepository.GetAll(ctx)
}
