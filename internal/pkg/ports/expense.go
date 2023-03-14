package ports

import (
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"
)

type (
	ExpenseRepository interface {
		AddExpense(expense *entity.Expense) error
	}

	ExpenseService interface {
		AddExpense(expense *entity.Expense) error
	}
)
