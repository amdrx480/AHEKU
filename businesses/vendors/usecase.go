package vendors

import (
	"backend-golang/app/middlewares"
	"context"
)

type vendorsUsecase struct {
	vendorsRepository Repository
	jwtAuth           *middlewares.JWTConfig
}

func NewVendorsUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &vendorsUsecase{
		vendorsRepository: repository,
		jwtAuth:           jwtAuth,
	}
}

func (usecase *vendorsUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.vendorsRepository.GetByID(ctx, id)
}

func (usecase *vendorsUsecase) Create(ctx context.Context, vendorsDomain *Domain) (Domain, error) {
	return usecase.vendorsRepository.Create(ctx, vendorsDomain)
}

func (usecase *vendorsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.vendorsRepository.GetAll(ctx)
}
