package main

import (
	"E-Meeting/handler"
	"E-Meeting/internal/repository"
	"E-Meeting/internal/seeder"
	"E-Meeting/internal/service"
	"E-Meeting/route"
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	_ "E-Meeting/docs"
)

// @title E-Meeting API
// @version 1.0
// @description Ini adalah API server untuk aplikasi E-Meeting.
// @host localhost:8080
// @BasePath
// @BasePath
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: Tidak dapat memuat file .env")
	}

	var runSeeder bool
	flag.BoolVar(&runSeeder, "seed", false, "Jalankan database seeder")
	flag.Parse()

	dsn := repository.GetDSN()

	log.Println("Mencoba menjalankan migrasi database...")
	m, err := migrate.New("file://migration", dsn)
	if err != nil {
		log.Fatalf("Gagal membuat instance migrasi: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Gagal menjalankan migrasi 'up': %v", err)
	}
	log.Println("Migrasi database berhasil.")

	// Koneksi ke database
	db, err := repository.ConnectDB()
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}
	defer db.Close()

	// perintah --seed
	if runSeeder {
		seeder.Run(db)
		log.Println("Seeder selesai.")
		return
	}

	// login and aut
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	loginHandler := handler.NewLoginHandler(authService)

	// room reservation schedule
	roomReservationRepo := repository.NewRoomReservationScheduleRepository(db)
	roomReservationService := service.NewRoomReservationScheduleService(roomReservationRepo)
	roomReservationHandler := handler.NewRoomReservationScheduleHandler(roomReservationService)

	// ambil data snack
	snackRepo := repository.NewSnackRepository(db)
	snackService := service.NewSnackService(snackRepo)
	snackHandler := handler.NewSnackHandler(snackService)

	profileRepo := repository.NewProfileRepository(db)
	profileService := service.NewProfileService(profileRepo)
	profileHandler := handler.NewProfileHandler(profileService)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// ambil data room
	roomRepo := repository.NewRoomRepository(db)
	roomService := service.NewRoomService(roomRepo)
	roomHandler := handler.NewRoomHandler(roomService)

	reservationRepo := repository.NewReservationRepository(db)
	reservationService := service.NewReservationService(reservationRepo)
	reservationHandler := handler.NewReservationHandler(reservationService)

	reservationCalculationRepo := repository.NewReservationCalculationRepository(db)
	reservationCalculationService := service.NewReservationCalculationService(reservationCalculationRepo)
	reservationCalculationHandler := handler.NewReservationCalculationHandler(reservationCalculationService)

	reservationHistoryRepo := repository.NewReservationHistoryRepository(db)
	reservationHistoryService := service.NewReservationHistoryService(reservationHistoryRepo)
	reservationHistoryHandler := handler.NewReservationHistoryHandler(reservationHistoryService)

	const tempDir = "public/temp"
	const uploadDir = "public/uploads"
	const baseURL = "http://localhost:8080/assets"

	uploadRepo := repository.NewLocalDiskRepository(uploadDir, baseURL)
	uploadService := service.NewUploadService(uploadRepo)
	uploadHandler := handler.NewUploadHandler(uploadService)

	dashboardRepo := repository.NewDashboardRepository(db)
	dashboardService := service.NewDashboardService(dashboardRepo)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)

	//testing
	allHandlers := &route.Handlers{
		SnackHandler:                   snackHandler,
		ProfileHandler:                 profileHandler,
		UserHandler:                    userHandler,
		RoomHandler:                    roomHandler,
		ReservationHandler:             reservationHandler,
		LoginHandler:                   loginHandler,
		RoomReservationScheduleHandler: roomReservationHandler,
		UploadHandler:                  uploadHandler,
		DashboardHandler:               dashboardHandler,
		ReservationCalculationHandler:  reservationCalculationHandler,
		ReservationHistoryHandler:      reservationHistoryHandler,
		// handler lain di sini
	}

	// jalanin server
	e := echo.New()

	e.Static("/assets", uploadDir)
	e.Static("/temp_assets", tempDir)

	route.SetupRoutes(e, allHandlers)

	log.Println("🚀 Server berjalan di http://localhost:8080")
	log.Println("📚 Dokumentasi Swagger tersedia di http://localhost:8080/swagger/index.html")

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
