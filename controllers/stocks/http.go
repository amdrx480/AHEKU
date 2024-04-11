package stocks

import (
	"backend-golang/businesses/stocks"

	"backend-golang/controllers"
	"backend-golang/controllers/stocks/request"
	"backend-golang/controllers/stocks/response"
	"net/http"

	// "github.com/skip2/go-qrcode"

	"github.com/labstack/echo/v4"
)

type StockController struct {
	stockUseCase stocks.Usecase
}

func NewStockController(authUC stocks.Usecase) *StockController {
	return &StockController{
		stockUseCase: authUC,
	}
}

func (cc *StockController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	stockID := c.Param("id")

	stock, err := cc.stockUseCase.GetByID(ctx, stockID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "category found", response.FromDomain(stock))
}

func (sc *StockController) Create(c echo.Context) error {
	input := request.Stock{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	stock, err := sc.stockUseCase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a stock", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "stock registered", response.FromDomain(stock))
}

func (sc *StockController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	categoriesData, err := sc.stockUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	categories := []response.Stock{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all categories", categories)
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
