package route

import (
	"E-Meeting/handler"
	"E-Meeting/internal/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handlers struct {
	SnackHandler                   *handler.SnackHandler
	LoginHandler                   *handler.LoginHandler
	RoomReservationScheduleHandler *handler.RoomReservationScheduleHandler
	//tambahin buat handerl lain
}

func SetupRoutes(e *echo.Echo, h *Handlers) {
	// Route untuk health check
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "API E-Meeting is running!"})
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/snacks", h.SnackHandler.GetAllSnacks)

	e.POST("/login", h.LoginHandler.Login)

	authGroup := e.Group("/rooms")
	authGroup.Use(middleware.JWTMiddleware)
	authGroup.GET("/:id_room/reservation", h.RoomReservationScheduleHandler.GetRoomReservationSchedules)

}
