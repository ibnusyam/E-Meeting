package router

import (
	"E-Meeting/internal/handler/http"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/", http.HomeHandler)

	return e
}
