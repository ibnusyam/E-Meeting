package http

import (
	"E-Meeting/internal/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SnackHandler struct {
	Repo *database.SnackRepository
}

// Buat constructor untuk handler
func NewSnackHandler(repo *database.SnackRepository) *SnackHandler {
	return &SnackHandler{Repo: repo}
}

func (h *SnackHandler) GetAllSnacks(c echo.Context) error {

	snacks, err := h.Repo.GetAllSnacks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "gagal mengambil data snacks"})
	}

	return c.JSON(http.StatusOK, snacks)
}
