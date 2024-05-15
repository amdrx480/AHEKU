package sales

import (
	"backend-golang/businesses/sales"

	"backend-golang/controllers"
	"backend-golang/controllers/sales/request"
	"backend-golang/controllers/sales/response"

	_reqHistory "backend-golang/controllers/history/request"
	// _resHistory "backend-golang/controllers/history/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type SalesController struct {
	salesUseCase sales.Usecase
}

func NewSalesController(authUC sales.Usecase) *SalesController {
	return &SalesController{
		salesUseCase: authUC,
	}
}

func (sc *SalesController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	salesID := c.Param("id")

	sales, err := sc.salesUseCase.GetByID(ctx, salesID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "sales not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "sales found", response.FromDomain(sales))
}

func (sc *SalesController) Create(c echo.Context) error {
	input := request.Sales{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	sales, err := sc.salesUseCase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a sales", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "sales registered", response.FromDomain(sales))
}

func (sc *SalesController) ToHistory(c echo.Context) error {
	historyID := c.Param("id")
	input := _reqHistory.History{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	// history, err := sc.salesUseCase.ToHistory(ctx, input.ToDomain(), historyID)
	_, err = sc.salesUseCase.ToHistory(ctx, input.ToDomain(), historyID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "Missing cart data", "")
	}

	// return controllers.NewResponse(c, http.StatusCreated, false, "Success ToHistory Data", _resHistory.FromDomain(history))
	// return controllers.NewResponse(c, http.StatusCreated, false, "Success ToHistory Data", " ")
	return controllers.NewResponseWithoutData(c, http.StatusCreated, false, "Success ToHistory Data")

}

func (sc *SalesController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	salesData, err := sc.salesUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	salesResponse := []response.Sales{}

	for _, sales := range salesData {
		salesResponse = append(salesResponse, response.FromDomain(sales))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all sales", salesResponse)
}

func (sc *SalesController) Delete(c echo.Context) error {
	categoryID := c.Param("id")
	ctx := c.Request().Context()

	err := sc.salesUseCase.Delete(ctx, categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, false, "failed to delete a sales", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "sales deleted", "")
}

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
