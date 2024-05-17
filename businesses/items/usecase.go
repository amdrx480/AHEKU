package items

import (
	// "backend-golang/businesses/history"

	"backend-golang/app/middlewares"
	"context"
)

type itemsUsecase struct {
	itemsRepository Repository
	jwtAuth         *middlewares.JWTConfig
}

func NewItemsUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &itemsUsecase{
		itemsRepository: repository,
		jwtAuth:         jwtAuth,
	}
}

func (usecase *itemsUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.itemsRepository.GetByID(ctx, id)
}

func (usecase *itemsUsecase) Create(ctx context.Context, itemsDomain *Domain) (Domain, error) {
	return usecase.itemsRepository.Create(ctx, itemsDomain)
}

func (usecase *itemsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.itemsRepository.GetAll(ctx)
}

func (usecase *itemsUsecase) Delete(ctx context.Context, id string) error {
	return usecase.itemsRepository.Delete(ctx, id)
}

func (usecase *itemsUsecase) GetByIDCart(ctx context.Context, id string) (DomainCart, error) {
	return usecase.itemsRepository.GetByIDCart(ctx, id)
}

func (usecase *itemsUsecase) CreateCart(ctx context.Context, itemsDomain *DomainCart) (DomainCart, error) {
	return usecase.itemsRepository.CreateCart(ctx, itemsDomain)
}

func (usecase *itemsUsecase) GetAllCart(ctx context.Context) ([]DomainCart, error) {
	return usecase.itemsRepository.GetAllCart(ctx)
}

func (usecase *itemsUsecase) DeleteCart(ctx context.Context, id string) error {
	return usecase.itemsRepository.DeleteCart(ctx, id)
}

// func (usecase *itemsUsecase) ToHistory(ctx context.Context, historyDomain *history.Domain, id string) (history.Domain, error) {
// 	return usecase.itemsRepository.ToHistory(ctx, historyDomain, id)
// }
