package deposit

import (
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/ports"
)

type (
	Service struct {
		repo ports.DepositRepository
	}
)

func NewService(repo ports.DepositRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddDeposit(deposit *entity.Deposit) error {
	return s.repo.AddDeposit(deposit)
}
