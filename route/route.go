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
	ReservationHistoryHandler      *handler.ReservationHistoryHandler
	ReservationDetailHandler       *handler.ReservationDetailHandler
	PasswordResetHandler           *handler.PasswordResetHandler
	PasswordResetbyIdHandler       *handler.PasswordResetHandler
	DeleteRoomHandler              *handler.RoomHandler
	CreateRoomHandler              *handler.CreateRoomHandler
	UpdateRoomHandler              *handler.RoomHandler
	//tambahin buat handerl lain
}

func SetupRoutes(e *echo.Echo, h *Handlers) {
	// Route untuk health check
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "API E-Meeting is running now!"})
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/login", h.LoginHandler.Login)
	e.POST("/register", h.UserHandler.Register)
	e.POST("/password/reset", h.PasswordResetHandler.ResetRequest)
	e.PUT("/password/reset/:id", h.PasswordResetHandler.ResetPassword)

	authGroup := e.Group("")
	authGroup.Use(middleware.JWTMiddleware)

	authGroup.GET("/rooms/:id_room/reservation", h.RoomReservationScheduleHandler.GetRoomReservationSchedules)

	authGroup.GET("/snacks", h.SnackHandler.GetAllSnacks)
	authGroup.GET("/rooms", h.RoomHandler.GetAllRooms)
	authGroup.DELETE("/rooms/:id", h.DeleteRoomHandler.DeleteRoom)
	authGroup.POST("/rooms", h.CreateRoomHandler.CreateRoom)
	authGroup.PUT("/rooms/:id", h.UpdateRoomHandler.UpdateRoom)

	authGroup.GET("/users/:id", h.ProfileHandler.GetUserProfileByID)
	authGroup.PATCH("/users/:id", h.ProfileHandler.UpdateUserHandler)
	authGroup.GET("/profile/:id", h.ProfileHandler.GetUserProfileByID)

	authGroup.POST("/reservations", h.ReservationHandler.CreateReservation)
	authGroup.PATCH("/reservation/status/:id", h.ReservationHandler.UpdateReservationStatusHandler)
	authGroup.GET("/reservation/calculation", h.ReservationCalculationHandler.GetReservationCalculation)
	authGroup.GET("/reservation/history", h.ReservationHistoryHandler.GetHistory)
	authGroup.GET("/reservation/:id", h.ReservationDetailHandler.GetReservationByID)

	authGroup.POST("/uploads", h.UploadHandler.UploadFile)
	authGroup.GET("/dashboard", h.DashboardHandler.GetDashboard)
}
