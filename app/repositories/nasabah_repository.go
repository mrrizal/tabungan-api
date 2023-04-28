package repositories

import (
	"database/sql"
	"tabungan-api/app/models"
)

type NasabahRepository interface {
	Daftar(nasabah *models.Nasabah) error
}

type nasabahRepository struct {
	db *sql.DB
}

func NewNasabahRepository(db *sql.DB) NasabahRepository {
	return &nasabahRepository{db: db}
}

func (this *nasabahRepository) Daftar(nasabah *models.Nasabah) error {
	// todo: implement this
	return nil
}
