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

// GetDashboard godoc
// @Summary      Mendapatkan data dashboard
// @Description  Mengambil data statistik untuk dashboard berdasarkan rentang tanggal
// @Tags         Dashboard
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true   "Bearer <access_token>"
// @Param        startDate   query     string  true  "Tanggal mulai dalam format YYYY-MM-DD"
// @Param        endDate     query     string  true  "Tanggal akhir dalam format YYYY-MM-DD"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /dashboard [get]
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
