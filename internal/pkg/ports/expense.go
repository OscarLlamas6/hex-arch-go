package ports

import "github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"

type (
	ExpenseRepository interface {
		Add(expense *entity.Expense) error
	}
)
