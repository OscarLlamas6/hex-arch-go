package ports

import "context"

type (
	DepositRepository interface {
		AddDeposit(ctx context.Context, doc interface{}) error
	}
)
