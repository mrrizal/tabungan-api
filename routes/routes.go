package routes

import (
	// "project/app/controllers"
	// "project/app/services"

	"tabungan-api/app/controllers"
	"tabungan-api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(app *fiber.App, db *pgxpool.Pool) {
	// Initialize services
	nasabahService := services.NewNasabahService(db)
	rekeningService := services.NewRekeningService(db)

	// Initialize controllers
	nasabahController := &controllers.NasabahController{
		NasabahService:  nasabahService,
		RekeningService: rekeningService,
	}

	rekeningController := &controllers.TransaksiController{
		RekeningService: rekeningService,
	}

	// Set up routes
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/daftar/", nasabahController.Daftar)
	v1.Post("/tabung/", rekeningController.Tabung)
	v1.Post("/tarik/", rekeningController.Tarik)
}
