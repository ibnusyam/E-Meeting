package handler

import (
	"E-Meeting/internal/service"
	"E-Meeting/model"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"errors"

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
// @Param Authorization header string true "Bearer <JWT Token>"
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

type UpdateStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

// UpdateReservationStatusHandler menangani PATCH /reservation/status/:id
func (h *ReservationHandler) UpdateReservationStatusHandler(c echo.Context) error {
	reservationID := c.Param("id")

	var req UpdateStatusRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	if req.Status == "" {
		// Logika validasi input dasar
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Status field is required"})
	}

	// Panggil Service
	err := h.Service.UpdateStatus(c.Request().Context(), reservationID, req.Status)

	// Petakan Error ke HTTP Status Code
	if err != nil {
		if errors.Is(err, service.ErrReservationNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "url not found"}) // 404
		}
		if errors.Is(err, service.ErrStatusConflict) {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "reservation already canceled/paid"}) // 400
		}

		fmt.Fprintf(os.Stderr, "FATAL ERROR: %v\n", err)

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error"}) // 500
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "update status success"}) // 200
}
