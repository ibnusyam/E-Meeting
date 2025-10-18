package handler

import (
	"E-Meeting/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomReservationScheduleHandler struct {
	Service *service.RoomReservationScheduleService
}

func NewRoomReservationScheduleHandler(service *service.RoomReservationScheduleService) *RoomReservationScheduleHandler {
	return &RoomReservationScheduleHandler{Service: service}
}

func (h *RoomReservationScheduleHandler) GetRoomReservationSchedules(c echo.Context) error {
	roomID, err := strconv.Atoi(c.Param("id_room"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid room id"})
	}

	startDate := c.QueryParam("startDate")
	if startDate == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "startDate query parameter is required"})
	}

	data, err := h.Service.GetRoomReservationSchedules(roomID, startDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error"})
	}

	if data.RoomName == "" {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "url not found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get room reservation schedules data success",
		"data":    data,
	})
}
