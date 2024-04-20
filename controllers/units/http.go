package units

import (
	"backend-golang/businesses/units"

	"backend-golang/controllers"
	"backend-golang/controllers/units/request"
	"backend-golang/controllers/units/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UnitsController struct {
	unitsUseCase units.Usecase
}

func NewUnitsController(authUC units.Usecase) *UnitsController {
	return &UnitsController{
		unitsUseCase: authUC,
	}
}

func (cc *UnitsController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	unitsID := c.Param("id")

	unit, err := cc.unitsUseCase.GetByID(ctx, unitsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "unit not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "unit found", response.FromDomain(unit))
}

func (pc *UnitsController) Create(c echo.Context) error {
	input := request.Units{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	unit, err := pc.unitsUseCase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a unit", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "unit registered", response.FromDomain(unit))
}

func (pc *UnitsController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	unitsData, err := pc.unitsUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	units := []response.Units{}

	for _, unit := range unitsData {
		units = append(units, response.FromDomain(unit))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all units", units)
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
