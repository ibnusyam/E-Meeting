package handler

import (
	"E-Meeting/internal/service"
	"E-Meeting/model"
	"database/sql"
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

// UpdateRoom godoc
// @Summary      Update Room
// @Description  Update data room berdasarkan ID (Admin only)
// @Tags         Rooms
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer <access_token>"
// @Param        id path int true "Room ID"
// @Param        request body model.CreateRoomRequest true "Room update body"
// @Success      200 {object} map[string]string "update room success"
// @Failure      400 {object} map[string]string "room type is not valid / capacity must be larger more than 0"
// @Failure      401 {object} map[string]string "unauthorized"
// @Failure      404 {object} map[string]string "url not found"
// @Failure      500 {object} map[string]string "internal server error"
// @Router       /rooms/{id} [put]
func (h *RoomHandler) UpdateRoom(c echo.Context) error {

	// role check
	role := c.Get("role")
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorized",
		})
	}

	// parse room id
	idParam := c.Param("id")
	roomID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "url not found",
		})
	}

	// bind payload
	var req model.CreateRoomRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	// service update
	err = h.Service.UpdateRoom(roomID, req)
	if err != nil {

		switch err.Error() {
		case "room type is not valid", "capacity must be larger more than 0":
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})

		case "url not found":
			return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
		}

		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "url not found"})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "update room success",
	})
}

// DeleteRoom godoc
// @Summary      Menghapus room berdasarkan ID
// @Description  Menghapus room dari sistem menggunakan ID room
// @Tags         Rooms
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true   "Bearer <access_token>"
// @Param        id   path      int  true  "Room ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /rooms/{id} [delete]
// DeleteRoom handler
func (h *RoomHandler) DeleteRoom(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "url not found",
		})
	}

	err = h.Service.DeleteRoom(id)
	if err != nil {
		switch err {
		case service.ErrRoomUsed:
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "cannot delete rooms. room has reservation",
			})
		case sql.ErrNoRows:
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "url not found",
			})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "internal server error",
			})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "delete room success",
	})
}
