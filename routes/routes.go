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

	transaksiController := &controllers.TransaksiController{
		RekeningService: rekeningService,
	}

	rekeningController := &controllers.RekeningController{
		RekeningService: rekeningService,
	}

	// Set up routes
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/daftar/", nasabahController.Daftar)
	v1.Post("/tabung/", transaksiController.Tabung)
	v1.Post("/tarik/", transaksiController.Tarik)
	v1.Get("/saldo/:no_rekening", rekeningController.Saldo)
}
