package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RekeningRepository interface {
	Daftar(nasabahID int, nomorRekening string) error
}

type rekeningRepository struct {
	db *pgxpool.Pool
}

func NewRekeningRepository(db *pgxpool.Pool) RekeningRepository {
	return &rekeningRepository{db: db}
}

func (this *rekeningRepository) Daftar(nasabahID int, nomorRekening string) error {
	sqlStmt := `
		INSERT INTO rekening (no_rekening, nasabah_id, saldo) 
		VALUES ($1, $2, $3)
	`
	_, err := this.db.Exec(context.Background(), sqlStmt, nomorRekening, nasabahID, 0)
	if err != nil {
		return err
	}

	return nil
}
