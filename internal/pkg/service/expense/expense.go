package expense

import (
	"context"

	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/ports"
)

type (
	Service struct {
		repo ports.ExpenseRepository
	}
)

func NewService(repo ports.ExpenseRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddExpense(ctx context.Context, doc interface{}) error {
	return s.repo.AddExpense(ctx, doc)
}
