package handler

import (
	"E-Meeting/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	Service *service.RoomService
}

func NewRoomHandler(service *service.RoomService) *RoomHandler {
	return &RoomHandler{Service: service}
}

// GetAllRooms godoc
// @Summary      Mendapatkan semua data room
// @Description  Mengambil seluruh daftar room yang tersedia dengan filter dan pagination
// @Tags         Rooms
// @Accept       json
// @Produce      json
// @Param        name       query     string  false  "Nama ruangan"
// @Param        type       query     string  false  "Tipe ruangan"
// @Param        capacity   query     int     false  "Kapasitas minimal"
// @Param        page       query     int     false  "Nomor halaman"  default(1)
// @Param        pageSize   query     int     false  "Jumlah data per halaman"  default(20)
// @Success      200  {object}  map[string]interface{}
// @Router       /rooms [get]
func (h *RoomHandler) GetAllRooms(c echo.Context) error {
	name := c.QueryParam("name")
	roomType := c.QueryParam("type")
	capacityStr := c.QueryParam("capacity")
	pageStr := c.QueryParam("page")
	pageSizeStr := c.QueryParam("pageSize")

	// Konversi numeric
	capacity := 0
	if capacityStr != "" {
		capacity, _ = strconv.Atoi(capacityStr)
	}

	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
		if page < 1 {
			page = 1
		}
	}

	pageSize := 20
	if pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
		if pageSize <= 0 {
			pageSize = 20
		}
	}

	rooms, err := h.Service.GetAllRooms(name, roomType, capacity, page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "gagal mengambil data rooms"})
	}

	if len(rooms) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "room type is not valid"})
	}

	return c.JSON(http.StatusOK, rooms)
}
