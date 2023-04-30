package repositories

import (
	"context"
	"tabungan-api/app/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransaksiRepository interface {
	Transaksi(models.TransaksiRequest) error
	Mutasi(string) ([]models.MutasiTransaksi, error)
}

type transaksiRepository struct {
	db *pgxpool.Pool
}

func NewTransaksiRepository(db *pgxpool.Pool) TransaksiRepository {
	return &transaksiRepository{db: db}
}

func (this *transaksiRepository) Transaksi(transaksiRequest models.TransaksiRequest) error {
	sqlStmt := `
		INSERT INTO transaksi (no_rekening, nominal, kode_transaksi) 
		VALUES ($1, $2, CAST($3 AS kodetransaksi))
	`
	_, err := this.db.Exec(
		context.Background(),
		sqlStmt,
		transaksiRequest.NoRekening,
		transaksiRequest.Nominal,
		transaksiRequest.Type)

	return err
}

func (this transaksiRepository) Mutasi(noRekening string) ([]models.MutasiTransaksi, error) {
	sqlStmt := "SELECT created_at, kode_transaksi, nominal FROM transaksi WHERE no_rekening = $1"
	rows, err := this.db.Query(context.Background(), sqlStmt, noRekening)
	if err != nil {
		return []models.MutasiTransaksi{}, err
	}

	var results []models.MutasiTransaksi
	for rows.Next() {
		var mutasi models.MutasiTransaksi
		err := rows.Scan(&mutasi.Waktu, &mutasi.KodeTransaksi, &mutasi.Nominal)
		if err != nil {
			return []models.MutasiTransaksi{}, err
		}
		results = append(results, mutasi)
	}

	return results, nil
}
