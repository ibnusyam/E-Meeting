package handler

import (
	"E-Meeting/internal/service"
	"E-Meeting/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ReservationCalculationHandler struct {
	service service.ReservationCalculationService
}

func NewReservationCalculationHandler(service service.ReservationCalculationService) *ReservationCalculationHandler {
	return &ReservationCalculationHandler{service: service}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// GetReservationCalculation godoc
// @Summary      Reservation Calculation
// @Description  Hitung total biaya reservasi berdasarkan room, snack, waktu, dan peserta
// @Tags         Reservation
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true   "Bearer <access_token>"
// @Param        room_id        query   int     true   "ID ruangan"  example(1)
// @Param        snack_id       query   int     false  "ID snack (optional)"  example(2)
// @Param        startTime      query   string  true   "Waktu mulai (format RFC3339)"  example(2023-01-01T10:00:00Z)
// @Param        endTime        query   string  true   "Waktu selesai (format RFC3339)"  example(2023-01-01T12:00:00Z)
// @Param        participant    query   int     true   "Jumlah peserta"  example(10)
// @Param        user_id        query   int     false  "ID user (optional)"  example(123)
// @Param        name           query   string  true  "Nama pemesan"  example(John Doe)
// @Param        phoneNumber    query   string  true  "Nomor telepon pemesan"  example(+628123456789)
// @Param        company        query   string  true  "Nama perusahaan"  example(Example Corp)
// @Success      200  {object}  model.ReservationCalculationResponse  "Success"
// @Failure      400  {object}  ErrorResponse  "Bad Request - over capacity / booking bentrok"
// @Failure      401  {object}  ErrorResponse  "Unauthorized"
// @Failure      404  {object}  ErrorResponse  "Not Found"
// @Failure      500  {object}  ErrorResponse  "Internal Server Error"
// @Router       /reservation/calculation [get]
func (h *ReservationCalculationHandler) GetReservationCalculation(c echo.Context) error {

	log.Printf("Received calculation request: room_id=%s, startTime=%s, endTime=%s",
		c.QueryParam("room_id"),
		c.QueryParam("startTime"),
		c.QueryParam("endTime"))
	// Parse room_id
	roomIDStr := c.QueryParam("room_id")
	if roomIDStr == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "room_id is required"})
	}
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid room_id format"})
	}

	// Parse snack_id (optional)
	snackID := 0
	snackIDStr := c.QueryParam("snack_id")
	if snackIDStr != "" {
		snackID, err = strconv.Atoi(snackIDStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid snack_id format"})
		}
	}

	// Parse startTime
	startTimeStr := c.QueryParam("startTime")
	if startTimeStr == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "startTime is required"})
	}
	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid startTime format, use RFC3339"})
	}

	// Parse endTime
	endTimeStr := c.QueryParam("endTime")
	if endTimeStr == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "endTime is required"})
	}
	endTime, err := time.Parse(time.RFC3339, endTimeStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid endTime format, use RFC3339"})
	}

	// Parse participant
	participantStr := c.QueryParam("participant")
	if participantStr == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "participant is required"})
	}
	participant, err := strconv.Atoi(participantStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid participant format"})
	}

	// Parse user_id (optional)
	userID := 0
	userIDStr := c.QueryParam("user_id")
	if userIDStr != "" {
		userID, err = strconv.Atoi(userIDStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid user_id format"})
		}
	}

	// Parse personal data
	name := c.QueryParam("name")
	phoneNumber := c.QueryParam("phoneNumber")
	company := c.QueryParam("company")

	// Create request object
	req := model.ReservationCalculationRequest{
		RoomID:      roomID,
		SnackID:     snackID,
		StartTime:   startTime,
		EndTime:     endTime,
		Participant: participant,
		UserID:      userID,
		Name:        name,
		PhoneNumber: phoneNumber,
		Company:     company,
	}

	// Log before calling service
	log.Printf("Calling service with req: %+v", req)

	// Call service
	result, err := h.service.CalculateReservation(c.Request().Context(), req)
	if err != nil {
		log.Printf("Service error: %v", err)
		// Handle specific error cases
		switch err.Error() {
		case "over capacity", "booking bentrok":
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		case "room not found":
			return c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		case "unauthorized":
			return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "unauthorized"})
		default:
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "internal server error"})
		}
	}

	log.Printf("Calculation successful, total: %.2f", result.Total)

	// Success response
	return c.JSON(http.StatusOK, result)
}
