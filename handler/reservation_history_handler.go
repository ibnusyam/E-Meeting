package handler

import (
	"E-Meeting/internal/service"
	"E-Meeting/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReservationHistoryHandler struct {
	service service.ReservationHistoryService
}

func NewReservationHistoryHandler(service service.ReservationHistoryService) *ReservationHistoryHandler {
	return &ReservationHistoryHandler{service: service}
}

// GetHistory godoc
// @Summary Get reservation history
// @Description Get reservation history with filters and pagination
// @Tags Reservation
// @Accept json
// @Produce json
// @Param        Authorization  header  string  true   "Bearer <access_token>"
// @Param startDate query string false "Start date filter (YYYY-MM-DD)"
// @Param endDate query string false "End date filter (YYYY-MM-DD)"
// @Param type query string false "Room type filter (small, medium, large, hall)"
// @Param status query string false "Status filter"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} model.ReservationHistoryResponse
// @Failure 400 {object} map[string]string "Bad Request - room type is not valid"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Not Found - url not found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Security BearerAuth
// @Router /reservation/history [get]
func (h *ReservationHistoryHandler) GetHistory(c echo.Context) error {
	// Parse query parameters
	startDate := c.QueryParam("startDate")
	endDate := c.QueryParam("endDate")
	roomType := c.QueryParam("type")
	status := c.QueryParam("status")

	username, ok := c.Get("username").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorized Username",
		})
	}

	role, ok := c.Get("role").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorized Role",
		})
	}

	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorized UserID",
		})
	}

	// Cek role
	if role == "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": "Forbidden: you don't have access to this resource",
		})
	}

	// Parse pagination
	page := 1
	if p := c.QueryParam("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil {
			page = val
		}
	}

	pageSize := 10
	if ps := c.QueryParam("pageSize"); ps != "" {
		if val, err := strconv.Atoi(ps); err == nil {
			pageSize = val
		}
	}

	filter := model.ReservationHistoryFilter{
		StartDate: startDate,
		EndDate:   endDate,
		Type:      roomType,
		Status:    status,
		Page:      page,
		PageSize:  pageSize,
		Username:  username,
		UserID:    userID,
	}

	// Get history from service
	response, err := h.service.GetHistory(c.Request().Context(), filter)
	if err != nil {
		if err.Error() == "room type is not valid" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Room type is not valid: " + err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Internal Server Error: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}
