package route

import (
	"E-Meeting/handler"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handlers struct {
	SnackHandler   *handler.SnackHandler
	ProfileHandler *handler.ProfileHandler
	DB             *sql.DB
}

func SetupRoutes(e *echo.Echo, h *Handlers) {
	// Route untuk health check
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "API E-Meeting is running!"})
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/snacks", h.SnackHandler.GetAllSnacks)

	e.GET("/profile/:id", h.ProfileHandler.GetUserProfileByID)
}
