package controllers

import (
	"tabungan-api/app/models"
	"tabungan-api/app/services"

	"github.com/gofiber/fiber/v2"
)

type RekeningController struct {
	RekeningService services.RekeningService
}

func (this *RekeningController) Saldo(c *fiber.Ctx) error {
	noRekening := c.Params("no_rekening")
	saldo, err := this.RekeningService.GetSaldo(noRekening)
	if err != nil {
		var result models.ErrorResponse
		result.Remark = err.Error()
		c.Status(400)
		return c.JSON(result)
	}

	var result models.TransaksiResponseOk
	result.Saldo = saldo
	c.Status(200)
	return c.JSON(result)
}
