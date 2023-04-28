package main

import (
	"database/sql"
	"tabungan-api/routes"

	"github.com/gofiber/fiber/v2"
)

type AppInstance struct {
	app *fiber.App
	db  *sql.DB
}

func NewAppInstance() AppInstance {
	return AppInstance{app: fiber.New(), db: &sql.DB{}}
}

func main() {
	app := NewAppInstance()
	routes.SetupRoutes(app.app, app.db)
}
