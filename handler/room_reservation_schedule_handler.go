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

// GetRoomReservationSchedules godoc
// @Summary      Mendapatkan jadwal reservasi room
// @Description  Mengambil jadwal reservasi untuk room tertentu berdasarkan tanggal mulai
// @Tags         Room Reservation Schedules
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true   "Bearer <access_token>"
// @Param        id_room     path     int     true  "ID Room"
// @Param        startDate   query     string  true  "Tanggal mulai dalam format YYYY-MM-DD"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /rooms/:id/reservation [get]
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
