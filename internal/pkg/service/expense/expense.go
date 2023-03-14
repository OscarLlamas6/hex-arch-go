package expense

import (
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"
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

func (s *Service) AddExpense(expense *entity.Expense) error {
	return s.repo.AddExpense(expense)
}
