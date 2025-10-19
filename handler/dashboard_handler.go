package handler

import (
	"E-Meeting/internal/service"
	"E-Meeting/internal/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
	service service.DashboardService
}

func NewDashboardHandler(service service.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: service}
}

func (h *DashboardHandler) GetDashboard(c echo.Context) error {
	startDate := c.QueryParam("startDate")
	endDate := c.QueryParam("endDate")

	data, err := h.service.GetDashboard(startDate, endDate)
	if err != nil {
		log.Printf("Service error: %v ", err)
		if err.Error() == "start date must be smaller than end date" {
			return c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("internal server error"))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessResponse("get dashboard data success", data))
}
