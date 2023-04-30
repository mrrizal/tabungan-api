package controllers

import (
	"tabungan-api/app/models"
	"tabungan-api/app/services"

	"github.com/gofiber/fiber/v2"
)

type NasabahController struct {
	NasabahService  services.NasabahService
	RekeningService services.RekeningService
}

func (this *NasabahController) Daftar(c *fiber.Ctx) error {
	var nasabah models.Nasabah
	if err := c.BodyParser(&nasabah); err != nil {
		return err
	}

	nasabahID, err := this.NasabahService.Daftar(nasabah)
	if err != nil {
		var result models.ErrorResponse
		result.Remark = err.Error()
		c.Status(400)
		return c.JSON(result)
	}

	noRekening, err := this.RekeningService.Daftar(nasabahID)
	if err != nil {
		var result models.ErrorResponse
		result.Remark = err.Error()
		c.Status(400)
		return c.JSON(result)
	}

	var result models.DaftarResponseOk
	result.NoRekening = noRekening
	c.Status(200)
	return c.JSON(result)
}
