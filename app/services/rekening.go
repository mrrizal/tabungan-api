package services

import (
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
