package controllers

import (
	"tabungan-api/app/models"
	"tabungan-api/app/services"

	"github.com/gofiber/fiber/v2"
)

type NasabahController struct {
	NasabahService services.NasabahService
}

func (this *NasabahController) Daftar(c *fiber.Ctx) error {
	// todo: parse request, send into to nasabah service
	if err := this.NasabahService.Daftar(&models.Nasabah{}); err != nil {
		return err
	}
	// todo return json
	return c.JSON(make(map[string]bool))
}
