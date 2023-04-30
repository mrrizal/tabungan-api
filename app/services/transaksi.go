package services

import (
	"tabungan-api/app/models"
	"tabungan-api/app/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransaksiService struct {
	TransaksiRepository repositories.TransaksiRepository
}

func NewTransaksiService(db *pgxpool.Pool) TransaksiService {
	return TransaksiService{
		TransaksiRepository: repositories.NewTransaksiRepository(db),
	}
}

func (this *TransaksiService) Transaksi(transaksiRequest models.TransaksiRequest) error {
	return this.TransaksiRepository.Transaksi(transaksiRequest)
}

func (this *TransaksiService) Mutasi(noRekening string) ([]models.MutasiTransaksi, error) {
	return this.TransaksiRepository.Mutasi(noRekening)
}
