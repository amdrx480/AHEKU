package customers

import (
	"backend-golang/businesses/customers"

	"backend-golang/controllers"
	"backend-golang/controllers/customers/request"
	"backend-golang/controllers/customers/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomersController struct {
	customersUseCase customers.Usecase
}

func NewCustomersController(authUC customers.Usecase) *CustomersController {
	return &CustomersController{
		customersUseCase: authUC,
	}
}

func (cc *CustomersController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	customersID := c.Param("id")

	customers, err := cc.customersUseCase.GetByID(ctx, customersID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, true, "customer not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, false, "customer found", response.FromDomain(customers))
}

func (cc *CustomersController) Create(c echo.Context) error {
	input := request.Customers{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "invalid request", "")
	}

	customers, err := cc.customersUseCase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to add a customer", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "customer registered", response.FromDomain(customers))
}

func (cc *CustomersController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	categoriesData, err := cc.customersUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	categories := []response.Customers{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all customers", categories)
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
