package repositories

import (
	"context"
	"fmt"
	"tabungan-api/app/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RekeningRepository interface {
	Daftar(nasabahID int, nomorRekening string) error
	IsExists(nomorRekening string) bool
	GetSaldo(nomorRekening string) (float64, error)
	Transaksi(models.TransaksiRequest) (float64, error)
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

func (this *rekeningRepository) IsExists(nomorRekening string) bool {
	var nasabahID int

	sqlStmt := fmt.Sprintf("SELECT nasabah_id FROM rekening WHERE no_rekening = '%s'", nomorRekening)
	err := this.db.QueryRow(context.Background(), sqlStmt).Scan(&nasabahID)
	if err != nil {
		return false
	}
	return nasabahID != 0
}

func (this *rekeningRepository) GetSaldo(nomorRekening string) (float64, error) {
	var saldo float64
	sqlStmt := fmt.Sprintf("SELECT saldo FROM rekening WHERE no_rekening = '%s'", nomorRekening)
	err := this.db.QueryRow(context.Background(), sqlStmt).Scan(&saldo)
	if err != nil {
		return 0, err
	}
	return saldo, nil
}

func (this *rekeningRepository) Transaksi(transaksiRequest models.TransaksiRequest) (float64, error) {
	var saldo float64
	if transaksiRequest.Type == "C" {
		saldo = transaksiRequest.CurrentSaldo + transaksiRequest.Nominal
	} else if transaksiRequest.Type == "D" {
		saldo = transaksiRequest.CurrentSaldo - transaksiRequest.Nominal
	}

	sqlStmt := "UPDATE rekening SET saldo = $1 WHERE no_rekening = $2"
	_, err := this.db.Exec(context.Background(), sqlStmt, saldo, transaksiRequest.NoRekening)
	if err != nil {
		return 0, err
	}
	return saldo, nil
}
