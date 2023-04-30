package controllers

import (
	"tabungan-api/app/models"
	"tabungan-api/app/services"
	"tabungan-api/utils"

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
		return utils.ErrorResp(c, err.Error())
	}

	noRekening, err := this.RekeningService.Daftar(nasabahID)
	if err != nil {
		return utils.ErrorResp(c, err.Error())
	}

	var result models.DaftarResponseOk
	result.NoRekening = noRekening
	c.Status(200)
	return c.JSON(result)
}
