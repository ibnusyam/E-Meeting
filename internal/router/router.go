package router

import (
	"E-Meeting/internal/handler/http"

	"github.com/labstack/echo/v4"
)

// Ubah fungsi agar menerima SnackHandler
func NewRouter(snackHandler *http.SnackHandler) *echo.Echo {
	e := echo.New()

	e.GET("/", http.HomeHandler)

	// Daftarkan method dari instance handler, bukan fungsi biasa
	e.GET("/snacks", snackHandler.GetAllSnacks)

	return e
}
