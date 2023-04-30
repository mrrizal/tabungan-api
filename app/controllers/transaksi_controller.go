package controllers

import (
	"tabungan-api/app/models"
	"tabungan-api/app/services"

	"github.com/gofiber/fiber/v2"
)

type TransaksiController struct {
	RekeningService services.RekeningService
}

func (this *TransaksiController) Tabung(c *fiber.Ctx) error {
	var request models.TransaksiRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	request.Type = "C"

	saldo, err := this.RekeningService.Tabung(request)
	if err != nil {
		var result models.ErrorResponse
		result.Remark = err.Error()
		c.Status(400)
		return c.JSON(result)
	}

	var resp models.TransaksiResponseOk
	resp.Saldo = saldo
	c.Status(200)
	return c.JSON(resp)
}

func (this *TransaksiController) Tarik(c *fiber.Ctx) error {
	var request models.TransaksiRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	request.Type = "D"

	saldo, err := this.RekeningService.Tarik(request)
	if err != nil {
		var result models.ErrorResponse
		result.Remark = err.Error()
		c.Status(400)
		return c.JSON(result)
	}

	var resp models.TransaksiResponseOk
	resp.Saldo = saldo
	c.Status(200)
	return c.JSON(resp)
}
