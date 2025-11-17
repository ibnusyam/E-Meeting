// handler/reservation_handler.go
package handler

import (
	"net/http"
	"strconv"
	"strings"

	"E-Meeting/internal/service"
	"E-Meeting/model"

	"github.com/labstack/echo/v4"
)

type ReservationDetailHandler struct {
	service service.ReservationDetailService
}

func NewReservationDetailHandler(service service.ReservationDetailService) *ReservationDetailHandler {
	return &ReservationDetailHandler{service: service}
}

// GetReservationByID godoc
// @Summary      Get reservation by ID
// @Description  Get detailed reservation information by reservation ID. Admin can view all reservations, customers can only view their own reservations.
// @Tags         Reservations
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Reservation ID"
// @Security     BearerAuth
// @Success      200  {object}  model.ReservationResponse
// @Failure      400  {object}  model.ErrorResponse  "Invalid reservation ID"
// @Failure      401  {object}  model.ErrorResponse  "Unauthorized - Invalid token or insufficient permissions"
// @Failure      404  {object}  model.ErrorResponse  "Reservation not found"
// @Failure      500  {object}  model.ErrorResponse  "Internal server error"
// @Router       /reservation/{id} [get]
func (h *ReservationDetailHandler) GetReservationByID(c echo.Context) error {
	// Get ID from path parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "invalid reservation id",
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

	// Call service
	reservationData, err := h.service.GetReservationByID(id, int(userID), role)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, model.ErrorResponse{
				Message: "reservation not found",
			})
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "unauthorized",
			})
		}
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "internal server error" + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.ReservationResponse{
		Message: "success get reservation",
		Data:    *reservationData,
	})
}
