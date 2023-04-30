package controllers

import (
	"tabungan-api/app/models"
	"tabungan-api/app/services"
	"tabungan-api/utils"

	"github.com/gofiber/fiber/v2"
)

type TransaksiController struct {
	RekeningService  services.RekeningService
	TransaksiService services.TransaksiService
}

func (this *TransaksiController) Tabung(c *fiber.Ctx) error {
	var request models.TransaksiRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	request.Type = "C"

	saldo, err := this.RekeningService.Tabung(request)
	if err != nil {
		return utils.ErrorResp(c, err.Error())
	}

	if err := this.TransaksiService.Transaksi(request); err != nil {
		return utils.ErrorResp(c, err.Error())
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
		return utils.ErrorResp(c, err.Error())
	}

	if err := this.TransaksiService.Transaksi(request); err != nil {
		return utils.ErrorResp(c, err.Error())
	}

	var resp models.TransaksiResponseOk
	resp.Saldo = saldo
	c.Status(200)
	return c.JSON(resp)
}

func (this *TransaksiController) Mutasi(c *fiber.Ctx) error {
	noRekening := c.Params("no_rekening")
	if !this.RekeningService.IsExists(noRekening) {
		return utils.ErrorResp(c, "nomor rekening doesn't exists")
	}

	mutasiRekening, err := this.TransaksiService.Mutasi(noRekening)
	if err != nil {
		return utils.ErrorResp(c, err.Error())
	}

	var result models.MutasiResp
	result.Results = mutasiRekening
	c.Status(200)
	return c.JSON(result)
}
