package main

import (
	"E-Meeting/internal/database"
	"E-Meeting/internal/handler/http"
	"E-Meeting/internal/router"
	"log"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Tidak dapat terhubung ke database: %v", err)
	}
	defer db.Close()

	snackRepo := database.NewSnackRepository(db)

	snackHandler := http.NewSnackHandler(snackRepo)

	e := router.NewRouter(snackHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
