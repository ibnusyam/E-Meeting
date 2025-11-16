package handler

import (
	"E-Meeting/internal/service"
	"E-Meeting/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateRoomHandler struct {
	Service service.RoomCreateService
}

func NewRoomCreateHandler(service service.RoomCreateService) *CreateRoomHandler {
	return &CreateRoomHandler{Service: service}
}

// CreateRoom godoc
// @Summary Create Room
// @Description Create a new meeting room (Admin only)
// @Tags Rooms
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>"
// @Param request body model.CreateRoomRequest true "Room request body"
// @Success 201 {object} map[string]string "Created"
// @Failure 400 {object} map[string]string "Bad Request - room type is not valid / capacity must be larger more than 0"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Not Found - url not found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Security BearerAuth
// @Router /rooms [post]
func (h *CreateRoomHandler) CreateRoom(c echo.Context) error {

	// JWT role checking
	role := c.Get("role")
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	var req model.CreateRoomRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	err := h.Service.CreateRoom(req)
	if err != nil {
		switch err.Error() {
		case "room type is not valid", "capacity must be larger more than 0":
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		case "url not found":
			return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error" + err.Error()})
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "success"})
}
