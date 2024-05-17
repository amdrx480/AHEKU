package items

import (
	"backend-golang/businesses/items"

	"backend-golang/controllers"
	"backend-golang/controllers/items/request"
	"backend-golang/controllers/items/response"

	// _reqHistory "backend-golang/controllers/history/request"
	// _resHistory "backend-golang/controllers/history/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type ItemsController struct {
	itemsUseCase items.Usecase
	// domainCart   items.DomainCart
}

func NewItemsController(authUC items.Usecase) *ItemsController {
	return &ItemsController{
		itemsUseCase: authUC,
	}
}

func (sc *ItemsController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	itemsID := c.Param("id")

	items, err := sc.itemsUseCase.GetByID(ctx, itemsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "items not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "items found", response.FromDomain(items))
}

func (sc *ItemsController) Create(c echo.Context) error {
	input := request.Items{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	items, err := sc.itemsUseCase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a items", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "items registered", response.FromDomain(items))
}

func (sc *ItemsController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	itemsData, err := sc.itemsUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	itemsResponse := []response.Items{}

	for _, items := range itemsData {
		itemsResponse = append(itemsResponse, response.FromDomain(items))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all items", itemsResponse)
}

func (sc *ItemsController) Delete(c echo.Context) error {
	categoryID := c.Param("id")
	ctx := c.Request().Context()

	err := sc.itemsUseCase.Delete(ctx, categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, false, "failed to delete a items", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "items deleted", "")
}

func (sc *ItemsController) GetByIDCart(c echo.Context) error {
	ctx := c.Request().Context()

	itemsID := c.Param("id")

	items, err := sc.itemsUseCase.GetByIDCart(ctx, itemsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "cart not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "cart found", response.FromDomainCart(items))
}

func (sc *ItemsController) CreateCart(c echo.Context) error {
	input := request.Cart{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	items, err := sc.itemsUseCase.CreateCart(ctx, input.ToDomainCart())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a cart", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "cart registered", response.FromDomainCart(items))
}

func (sc *ItemsController) GetAllCart(c echo.Context) error {
	ctx := c.Request().Context()

	itemsData, err := sc.itemsUseCase.GetAllCart(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	itemsResponse := []response.Cart{}

	for _, items := range itemsData {
		itemsResponse = append(itemsResponse, response.FromDomainCart(items))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all cart", itemsResponse)
}

func (sc *ItemsController) DeleteCart(c echo.Context) error {
	categoryID := c.Param("id")
	ctx := c.Request().Context()

	err := sc.itemsUseCase.DeleteCart(ctx, categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, false, "failed to delete a cart", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "cart deleted", "")
}

// func (sc *ItemsController) ToHistory(c echo.Context) error {
// 	historyID := c.Param("id")
// 	input := _reqHistory.History{}
// 	ctx := c.Request().Context()

// 	if err := c.Bind(&input); err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
// 	}

// 	err := input.Validate()

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
// 	}

// 	// history, err := sc.itemsUseCase.ToHistory(ctx, input.ToDomain(), historyID)
// 	_, err = sc.itemsUseCase.ToHistory(ctx, input.ToDomain(), historyID)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, true, "Missing cart data", "")
// 	}

// 	// return controllers.NewResponse(c, http.StatusCreated, false, "Success ToHistory Data", _resHistory.FromDomain(history))
// 	// return controllers.NewResponse(c, http.StatusCreated, false, "Success ToHistory Data", " ")
// 	return controllers.NewResponseWithoutData(c, http.StatusCreated, false, "Success ToHistory Data")

// }

// func (cc *StockController) DownloadBarcodeByID(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	stockID := c.Param("id")

// 	// Check if stock is found after generating barcode
// 	stock, err := cc.stockUseCase.DownloadBarcodeByID(ctx, stockID)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusNotFound, http.StatusNotFound, true, "Stock not found", "")
// 	}

// 	// Generate QR code with stockID
// 	// tanpa total stock agar barcodenya tidak terjadi perubahan ketika update data
// 	qrContent := fmt.Sprintf("Stock ID: %s\n"+
// 		"Created At: %s\n"+
// 		"Stock Location: %s\n"+
// 		"Stock Code: %s\n"+
// 		"Stock Name: %s\n"+
// 		"Unit: %s\n"+
// 		stockID, stock.CreatedAt, stock.Stock_Location, stock.Stock_Code, stock.Stock_Name, stock.Stock_Unit)

// 	qrCode, err := qrcode.New(qrContent, qrcode.Medium)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error generating barcode", "")
// 	}

// 	// Set response headers
// 	c.Response().Header().Set("Content-Type", "image/png")
// 	c.Response().Header().Set("Content-Disposition", "attachment; filename=Kode Stok "+stock.Stock_Code+" - Nama Stok "+stock.Stock_Name+".png")
// 	c.Response().WriteHeader(http.StatusOK)

// 	// Encode QR code as PNG and write to response
// 	if err := qrCode.Write(200, c.Response()); err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error encoding barcode", "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, http.StatusOK, false, "Download", response.FromDomain(stock))
// }
