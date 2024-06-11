package admin

import (
	"backend-golang/app/middlewares"
	"context"
)

// Service
type adminUsecase struct {
	adminRepository Repository
	jwtAuth         *middlewares.JWTConfig
}

func NewAdminUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &adminUsecase{
		adminRepository: repository,
		jwtAuth:         jwtAuth,
	}
}

func (usecase *adminUsecase) AdminRegister(ctx context.Context, AdminsDomain *AdminsDomain) (AdminsDomain, error) {
	return usecase.adminRepository.AdminRegister(ctx, AdminsDomain)
}

func (usecase *adminUsecase) AdminLogin(ctx context.Context, adminDomain *AdminsDomain) (string, error) {
	admin, err := usecase.adminRepository.AdminGetByName(ctx, adminDomain)

	if err != nil {
		return "", err
	}

	token, err := usecase.jwtAuth.GenerateToken(int(admin.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (usecase *adminUsecase) AdminVoucher(ctx context.Context, adminDomain *AdminsDomain) (string, error) {
	admin, err := usecase.adminRepository.AdminGetByVoucher(ctx, adminDomain)

	if err != nil {
		return "", err
	}

	return admin.Voucher, nil
}

func (usecase *adminUsecase) CustomersCreate(ctx context.Context, customersDomain *CustomersDomain) (CustomersDomain, error) {
	return usecase.adminRepository.CustomersCreate(ctx, customersDomain)
}

func (usecase *adminUsecase) CustomersGetByID(ctx context.Context, id string) (CustomersDomain, error) {
	return usecase.adminRepository.CustomersGetByID(ctx, id)
}

func (usecase *adminUsecase) CustomersGetAll(ctx context.Context) ([]CustomersDomain, error) {
	return usecase.adminRepository.CustomersGetAll(ctx)
}

func (usecase *adminUsecase) CategoryCreate(ctx context.Context, CategoriesDomain *CategoriesDomain) (CategoriesDomain, error) {
	return usecase.adminRepository.CategoryCreate(ctx, CategoriesDomain)
}

func (usecase *adminUsecase) CategoryGetByID(ctx context.Context, id string) (CategoriesDomain, error) {
	return usecase.adminRepository.CategoryGetByID(ctx, id)
}

func (usecase *adminUsecase) CategoryGetByName(ctx context.Context, name string) (CategoriesDomain, error) {
	return usecase.adminRepository.CategoryGetByName(ctx, name)
}

func (usecase *adminUsecase) CategoryGetAll(ctx context.Context) ([]CategoriesDomain, error) {
	return usecase.adminRepository.CategoryGetAll(ctx)
}

func (usecase *adminUsecase) VendorsCreate(ctx context.Context, vendorsDomain *VendorsDomain) (VendorsDomain, error) {
	return usecase.adminRepository.VendorsCreate(ctx, vendorsDomain)
}

func (usecase *adminUsecase) VendorsGetByID(ctx context.Context, id string) (VendorsDomain, error) {
	return usecase.adminRepository.VendorsGetByID(ctx, id)
}

func (usecase *adminUsecase) VendorsGetAll(ctx context.Context) ([]VendorsDomain, error) {
	return usecase.adminRepository.VendorsGetAll(ctx)
}

func (usecase *adminUsecase) UnitsCreate(ctx context.Context, unitsDomain *UnitsDomain) (UnitsDomain, error) {
	return usecase.adminRepository.UnitsCreate(ctx, unitsDomain)
}

func (usecase *adminUsecase) UnitsGetByID(ctx context.Context, id string) (UnitsDomain, error) {
	return usecase.adminRepository.UnitsGetByID(ctx, id)
}

func (usecase *adminUsecase) UnitsGetAll(ctx context.Context) ([]UnitsDomain, error) {
	return usecase.adminRepository.UnitsGetAll(ctx)
}

func (usecase *adminUsecase) StocksCreate(ctx context.Context, stocksDomain *StocksDomain) (StocksDomain, error) {
	return usecase.adminRepository.StocksCreate(ctx, stocksDomain)
}

func (usecase *adminUsecase) StocksGetByID(ctx context.Context, id string) (StocksDomain, error) {
	return usecase.adminRepository.StocksGetByID(ctx, id)
}

func (usecase *adminUsecase) StocksGetAll(ctx context.Context) ([]StocksDomain, error) {
	return usecase.adminRepository.StocksGetAll(ctx)
}

func (usecase *adminUsecase) PurchasesCreate(ctx context.Context, purchasesDomain *PurchasesDomain) (PurchasesDomain, error) {
	return usecase.adminRepository.PurchasesCreate(ctx, purchasesDomain)
}

func (usecase *adminUsecase) PurchasesGetByID(ctx context.Context, id string) (PurchasesDomain, error) {
	return usecase.adminRepository.PurchasesGetByID(ctx, id)
}

func (usecase *adminUsecase) PurchasesGetAll(ctx context.Context) ([]PurchasesDomain, error) {
	return usecase.adminRepository.PurchasesGetAll(ctx)
}

func (usecase *adminUsecase) CartItemsGetByID(ctx context.Context, id string) (CartItemsDomain, error) {
	return usecase.adminRepository.CartItemsGetByID(ctx, id)
}

func (usecase *adminUsecase) CartItemsGetAllByCustomerID(ctx context.Context, customerId string) ([]CartItemsDomain, error) {
	return usecase.adminRepository.CartItemsGetAllByCustomerID(ctx, customerId)
}

func (usecase *adminUsecase) CartItemsCreate(ctx context.Context, cartItemsDomain *CartItemsDomain) (CartItemsDomain, error) {
	return usecase.adminRepository.CartItemsCreate(ctx, cartItemsDomain)
}

func (usecase *adminUsecase) CartItemsGetAll(ctx context.Context) ([]CartItemsDomain, error) {
	return usecase.adminRepository.CartItemsGetAll(ctx)
}

func (usecase *adminUsecase) CartItemsDelete(ctx context.Context, id string) error {
	return usecase.adminRepository.CartItemsDelete(ctx, id)
}

// func (usecase *adminUsecase) ItemTransactionsCreate(ctx context.Context, itemTransactionsDomain *ItemTransactionsDomain, id string) (ItemTransactionsDomain, error) {
// 	// func (usecase *adminUsecase) ItemTransactionsCreate(ctx context.Context, id string) (ItemTransactionsDomain, error) {
// 	return usecase.adminRepository.ItemTransactionsCreate(ctx, itemTransactionsDomain, id)
// }

func (uc *adminUsecase) ItemTransactionsCreate(ctx context.Context, customerId string) (ItemTransactionsDomain, error) {
	return uc.adminRepository.ItemTransactionsCreate(ctx, customerId)
}

func (usecase *adminUsecase) ItemTransactionsGetAll(ctx context.Context) ([]ItemTransactionsDomain, error) {
	return usecase.adminRepository.ItemTransactionsGetAll(ctx)
}

// func (usecase *adminUsecase) CartsCreate(ctx context.Context, cartItemsDomain *CartsDomain) (CartsDomain, error) {
// 	return usecase.adminRepository.CartsCreate(ctx, cartItemsDomain)
// }

// func (usecase *adminUsecase) CartsGetByID(ctx context.Context, id string) (CartsDomain, error) {
// 	return usecase.adminRepository.CartsGetByID(ctx, id)
// }

// func (usecase *adminUsecase) CartsGetAll(ctx context.Context) ([]CartsDomain, error) {
// 	return usecase.adminRepository.CartsGetAll(ctx)
// }

// func (usecase *adminUsecase) CartsDelete(ctx context.Context, id string) error {
// 	return usecase.adminRepository.CartsDelete(ctx, id)
// }
