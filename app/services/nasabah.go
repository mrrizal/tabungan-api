package services

import (
	"database/sql"
	"tabungan-api/app/models"
	"tabungan-api/app/repositories"
)

type NasabahService struct {
	NasabahRepository repositories.NasabahRepository
}

func NewNasabahService(db *sql.DB) NasabahService {
	return NasabahService{
		NasabahRepository: repositories.NewNasabahRepository(db),
	}
}

func (this *NasabahService) Daftar(nasabah *models.Nasabah) error {
	// todo: do bussiness logic here like check nik and no_hp, also return the no_rekening
	return nil
}
