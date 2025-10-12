package main

import "E-Meeting/internal/router"

func main() {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":8080"))
}
