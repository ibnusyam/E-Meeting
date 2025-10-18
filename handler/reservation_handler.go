package handler

import (
	"E-Meeting/internal/service"
	"E-Meeting/model"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	Service *service.ReservationService
}

func NewReservationHandler(s *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{Service: s}
}

// CreateReservation godoc
// @Summary Create a new reservation
// @Description Create a new meeting room reservation with room details and snacks
// @Tags Reservations
// @Accept json
// @Produce json
// @Param request body model.ReservationRequest true "Reservation Request"
// @Success 201 {object} map[string]interface{} "reservation created successfully"
// @Failure 400 {object} map[string]interface{} "invalid request / room not found / room has been booked / invalid userID"
// @Failure 401 {object} map[string]interface{} "unauthorized"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /reservations [post]
func (h *ReservationHandler) CreateReservation(c echo.Context) error {
	var req model.ReservationRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("Bind error: %v", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}

	// Validate userID is a valid number
	if _, err := strconv.Atoi(req.UserID); err != nil {
		log.Printf("Invalid userID: %v", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid userID, must be a number"})
	}

	// Pass request directly to service
	err := h.Service.CreateReservation(req)
	if err != nil {
		log.Printf("Service error: %v", err)
		switch err.Error() {
		case "room not found", "room has been booked":
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		case "unauthorized":
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "internal server error",
				"error":   err.Error(),
			})
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "reservation created successfully"})
}
