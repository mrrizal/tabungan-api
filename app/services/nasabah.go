package services

import (
	"errors"
	"tabungan-api/app/models"
	"tabungan-api/app/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NasabahService struct {
	NasabahRepository repositories.NasabahRepository
}

func NewNasabahService(db *pgxpool.Pool) NasabahService {
	return NasabahService{
		NasabahRepository: repositories.NewNasabahRepository(db),
	}
}

func (this *NasabahService) Daftar(nasabah models.Nasabah) (int, error) {
	nikIsExists, err := this.NasabahRepository.IsExists("nik", nasabah.Nik)
	if err != nil {
		return 0, err
	}

	if nikIsExists {
		return 0, errors.New("nik already exists")
	}

	noHpIsExists, err := this.NasabahRepository.IsExists("no_hp", nasabah.NoHp)
	if err != nil {
		return 0, err
	}

	if noHpIsExists {
		return 0, errors.New("no_hp already exists")
	}

	nasabahID, err := this.NasabahRepository.Daftar(nasabah)
	if err != nil {
		return 0, err
	}
	return nasabahID, nil
}
