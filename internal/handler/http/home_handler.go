package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Selamat Datang Di Aplikasi E-Meeting")
}
