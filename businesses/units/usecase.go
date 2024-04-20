package units

import (
	"backend-golang/app/middlewares"
	"context"
)

type unitsUsecase struct {
	unitsRepository Repository
	jwtAuth         *middlewares.JWTConfig
}

func NewUnitsUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &unitsUsecase{
		unitsRepository: repository,
		jwtAuth:         jwtAuth,
	}
}

func (usecase *unitsUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.unitsRepository.GetByID(ctx, id)
}

func (usecase *unitsUsecase) Create(ctx context.Context, unitsDomain *Domain) (Domain, error) {
	return usecase.unitsRepository.Create(ctx, unitsDomain)
}

func (usecase *unitsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.unitsRepository.GetAll(ctx)
}
