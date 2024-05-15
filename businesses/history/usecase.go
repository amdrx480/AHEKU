package history

import (
	"backend-golang/app/middlewares"
	"context"
)

type historyUsecase struct {
	historyRepository Repository
	jwtAuth           *middlewares.JWTConfig
}

func NewHistoryUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &historyUsecase{
		historyRepository: repository,
		jwtAuth:           jwtAuth,
	}
}

// func (usecase *historyUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
// 	return usecase.historyRepository.GetByID(ctx, id)
// }

// func (usecase *historyUsecase) GetByName(ctx context.Context, name string) (Domain, error) {
// 	return usecase.historyRepository.GetByName(ctx, name)
// }

func (usecase *historyUsecase) Create(ctx context.Context, historyDomain *Domain) (Domain, error) {
	return usecase.historyRepository.Create(ctx, historyDomain)
}

func (usecase *historyUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.historyRepository.GetAll(ctx)
}
