package repositories

import (
	"context"
	"fmt"
	"tabungan-api/app/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NasabahRepository interface {
	Daftar(nasabah models.Nasabah) (int, error)
	IsExists(field, value string) (bool, error)
}

type nasabahRepository struct {
	db *pgxpool.Pool
}

func NewNasabahRepository(db *pgxpool.Pool) NasabahRepository {
	return &nasabahRepository{db: db}
}

func (this *nasabahRepository) Daftar(nasabah models.Nasabah) (int, error) {
	sqlStmt := `
		INSERT INTO nasabah (nama, nik, no_hp) 
		VALUES ($1, $2, $3) RETURNING id
	`
	var nasabahID int
	err := this.db.QueryRow(
		context.Background(),
		sqlStmt,
		nasabah.Nama,
		nasabah.Nik,
		nasabah.NoHp).Scan(&nasabahID)

	if err != nil {
		return 0, err
	}

	return nasabahID, nil
}

func (this *nasabahRepository) IsExists(field, value string) (bool, error) {
	var nasabahID int

	sqlStmt := fmt.Sprintf("SELECT id FROM nasabah WHERE %s = '%s'", field, value)
	err := this.db.QueryRow(context.Background(), sqlStmt).Scan(&nasabahID)
	if err != nil {
		return false, nil
	}
	return nasabahID != 0, nil
}
