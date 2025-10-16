package route

import (
	"E-Meeting/handler"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handlers struct {
	SnackHandler *handler.SnackHandler
	//tambahin buat handerl lain
	UserHandler *handler.UserHandler
	RoomHandler *handler.RoomHandler
}

func SetupRoutes(e *echo.Echo, h *Handlers) {
	// Route untuk health check
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "API E-Meeting is running!"})
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/register", h.UserHandler.Register)

	e.GET("/snacks", h.SnackHandler.GetAllSnacks)
	e.GET("/rooms", h.RoomHandler.GetAllRooms)
}
