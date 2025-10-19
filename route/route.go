package route

import (
	"E-Meeting/handler"
	"E-Meeting/internal/middleware"
	"E-Meeting/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handlers struct {
	SnackHandler                   *handler.SnackHandler
	UserHandler                    *handler.UserHandler
	RoomHandler                    *handler.RoomHandler
	ProfileHandler                 *handler.ProfileHandler
	ReservationHandler             *handler.ReservationHandler
	ReservationRepository          *repository.ReservationRepository
	LoginHandler                   *handler.LoginHandler
	RoomReservationScheduleHandler *handler.RoomReservationScheduleHandler
	UploadHandler                  *handler.UploadHandler
	DashboardHandler               *handler.DashboardHandler
	ReservationCalculationHandler  *handler.ReservationCalculationHandler
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
	e.PATCH("/reservation/status/:id", h.ReservationHandler.UpdateReservationStatusHandler)

	e.GET("/rooms", h.RoomHandler.GetAllRooms)

	e.GET("/profile/:id", h.ProfileHandler.GetUserProfileByID)
	e.GET("/rooms", h.RoomHandler.GetAllRooms)
	e.GET("/reservation/calculation", h.ReservationCalculationHandler.GetReservationCalculation)

	e.POST("/login", h.LoginHandler.Login)

	e.PATCH("/users/:id", h.ProfileHandler.UpdateUserHandler)

	e.POST("/uploads", h.UploadHandler.UploadFile)

	authGroup := e.Group("/rooms")
	authGroup.Use(middleware.JWTMiddleware)
	authGroup.GET("/:id_room/reservation", h.RoomReservationScheduleHandler.GetRoomReservationSchedules)
}
