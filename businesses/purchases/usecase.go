package purchases

import (
	"backend-golang/app/middlewares"
	"context"
)

type purchaseUsecase struct {
	purchaseRepository Repository
	jwtAuth            *middlewares.JWTConfig
}

func NewPurchasesUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &purchaseUsecase{
		purchaseRepository: repository,
		jwtAuth:            jwtAuth,
	}
}

func (usecase *purchaseUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.purchaseRepository.GetByID(ctx, id)
}

func (usecase *purchaseUsecase) Create(ctx context.Context, purchasesDomain *Domain) (Domain, error) {
	return usecase.purchaseRepository.Create(ctx, purchasesDomain)
}

func (usecase *purchaseUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.purchaseRepository.GetAll(ctx)
}
