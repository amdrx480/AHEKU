package history

import (
	"backend-golang/businesses/history"
	"backend-golang/controllers"
	"backend-golang/controllers/history/request"
	"backend-golang/controllers/history/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HistoryController struct {
	historyUsecase history.Usecase
}

func NewHistoryController(historyUC history.Usecase) *HistoryController {
	return &HistoryController{
		historyUsecase: historyUC,
	}
}

func (hc *HistoryController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	historiesData, err := hc.historyUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	histories := []response.History{}

	for _, course := range historiesData {
		histories = append(histories, response.FromDomain(course))
	}

	return controllers.NewResponse(c, http.StatusOK, false, "all histories", histories)
}

func (hc *HistoryController) Create(c echo.Context) error {
	input := request.History{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, true, "validation failed", "")
	}

	course, err := hc.historyUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, true, "failed to create a course", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, false, "course created", response.FromDomain(course))
}
