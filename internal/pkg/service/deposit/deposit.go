package deposit

import (
	"context"

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

func (s *Service) AddDeposit(ctx context.Context, doc interface{}) error {
	return s.repo.AddDeposit(ctx, doc)
}
