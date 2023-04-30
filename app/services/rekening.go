package services

import (
	"errors"
	"tabungan-api/app/models"
	"tabungan-api/app/repositories"
	"tabungan-api/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RekeningService struct {
	RekeningRepository repositories.RekeningRepository
}

func NewRekeningService(db *pgxpool.Pool) RekeningService {
	return RekeningService{
		RekeningRepository: repositories.NewRekeningRepository(db),
	}
}

func (this *RekeningService) Daftar(nasabahID int) (string, error) {
	// todo: handle if no rekening already exists
	noRekening := utils.GenerateAccountNumber(10)
	if err := this.RekeningRepository.Daftar(nasabahID, noRekening); err != nil {
		return "", err
	}

	return noRekening, nil
}

func (this *RekeningService) IsExists(nomorRekening string) bool {
	return this.RekeningRepository.IsExists(nomorRekening)
}

func (this *RekeningService) Tabung(tabungRequest models.TransaksiRequest) (float64, error) {
	if !this.RekeningRepository.IsExists(tabungRequest.NoRekening) {
		return 0, errors.New("no rekening doesn't exists")
	}

	saldo, err := this.RekeningRepository.GetSaldo(tabungRequest.NoRekening)
	if err != nil {
		return 0, err
	}

	tabungRequest.CurrentSaldo = saldo

	return this.RekeningRepository.Transaksi(tabungRequest)
}

func (this *RekeningService) Tarik(tabungRequest models.TransaksiRequest) (float64, error) {
	if !this.RekeningRepository.IsExists(tabungRequest.NoRekening) {
		return 0, errors.New("no rekening doesn't exists")
	}

	saldo, err := this.RekeningRepository.GetSaldo(tabungRequest.NoRekening)
	if err != nil {
		return 0, err
	}

	tabungRequest.CurrentSaldo = saldo

	if tabungRequest.CurrentSaldo < tabungRequest.Nominal {
		return 0, errors.New("saldo tidak mencukupi")
	}

	return this.RekeningRepository.Transaksi(tabungRequest)
}
