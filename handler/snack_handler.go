package handler

import (
	"E-Meeting/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SnackHandler struct {
	Service service.SnackService
}

func NewSnackHandler(service *service.SnackService) *SnackHandler {
	return &SnackHandler{Service: *service}
}

// handler/snack_handler.go

// GetAllSnacks godoc
// @Summary      Mendapatkan semua data snack
// @Description  Mengambil seluruh daftar snack yang tersedia
// @Tags         Snacks
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.Snack
// @Router       /snacks [get]
func (h *SnackHandler) GetAllSnacks(c echo.Context) error {

	snacks, err := h.Service.GetAllSnacks()

	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "gagal mengambil data snacks"})
	}

	return c.JSON(http.StatusOK, snacks)
}
