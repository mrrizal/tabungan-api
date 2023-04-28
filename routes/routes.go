package routes

import (
	// "project/app/controllers"
	// "project/app/services"

	"database/sql"
	"tabungan-api/app/controllers"
	"tabungan-api/app/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	// Initialize services
	nasabahService := services.NewNasabahService(db)

	// Initialize controllers
	nasabahController := &controllers.NasabahController{
		NasabahService: nasabahService,
	}

	// Set up routes
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/daftar/", nasabahController.Daftar)
}
