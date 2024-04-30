package category

import (
	"backend-golang/app/middlewares"
	"context"
)

type categoryUsecase struct {
	categoryRepository Repository
	jwtAuth            *middlewares.JWTConfig
}

func NewCategoryUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &categoryUsecase{
		categoryRepository: repository,
		jwtAuth:            jwtAuth,
	}
}

func (usecase *categoryUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.categoryRepository.GetByID(ctx, id)
}

func (usecase *categoryUsecase) GetByName(ctx context.Context, name string) (Domain, error) {
	return usecase.categoryRepository.GetByName(ctx, name)
}

func (usecase *categoryUsecase) Create(ctx context.Context, categoryDomain *Domain) (Domain, error) {
	return usecase.categoryRepository.Create(ctx, categoryDomain)
}

func (usecase *categoryUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.categoryRepository.GetAll(ctx)
}
