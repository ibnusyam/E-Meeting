package route

import (
	"E-Meeting/handler"
	"E-Meeting/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handlers struct {
	SnackHandler          *handler.SnackHandler
	UserHandler           *handler.UserHandler
	RoomHandler           *handler.RoomHandler
	ProfileHandler        *handler.ProfileHandler
	ReservationHandler    *handler.ReservationHandler
	ReservationRepository *repository.ReservationRepository
	//tambahin buat handerl lain
}

func SetupRoutes(e *echo.Echo, h *Handlers) {
	// Route untuk health check
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "API E-Meeting is running now!"})
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/snacks", h.SnackHandler.GetAllSnacks)

	e.GET("/users/:id", h.ProfileHandler.GetUserProfileByID)

	e.POST("/register", h.UserHandler.Register)
	e.POST("/reservations", h.ReservationHandler.CreateReservation)
	e.GET("/rooms", h.RoomHandler.GetAllRooms)

	e.GET("/profile/:id", h.ProfileHandler.GetUserProfileByID)
	e.GET("/rooms", h.RoomHandler.GetAllRooms)

	e.POST("/login", h.LoginHandler.Login)

	authGroup := e.Group("/rooms")
	authGroup.Use(middleware.JWTMiddleware)
	authGroup.GET("/:id_room/reservation", h.RoomReservationScheduleHandler.GetRoomReservationSchedules)
}
