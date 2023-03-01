package ports

import "context"

type (
	ExpenseRepository interface {
		AddExpense(ctx context.Context, doc interface{}) error
	}
)
