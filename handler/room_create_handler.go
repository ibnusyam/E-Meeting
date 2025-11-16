package handler

import (
	"E-Meeting/internal/service"
	"E-Meeting/model"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

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
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer <access_token>"
// @Param name formData string true "Room name"
// @Param pricePerHour formData number true "Price per hour"
// @Param capacity formData int true "Capacity"
// @Param type formData string true "Room type (small, medium, large)"
// @Param image formData file true "Image (PNG/JPG/JPEG) max 1MB"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /rooms [post]
func (h *CreateRoomHandler) CreateRoom(c echo.Context) error {

	// role check
	role := c.Get("role")
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	// parse form fields
	req := new(model.CreateRoomRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid form data"})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "image is required"})
	}

	// open file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "failed to open file"})
	}
	defer src.Close()

	// create folder
	savePath := "public/uploads/rooms/"
	os.MkdirAll(savePath, 0755)

	fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	fullPath := savePath + fileName

	// save
	dst, err := os.Create(fullPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "failed to create file"})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "failed to save file"})
	}

	// generate URL
	imageURL := "/static/rooms/" + fileName

	// panggil service
	err = h.Service.CreateRoom(req, imageURL)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "success"})
}
