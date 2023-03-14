package ports

import (
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"
)

type (
	DepositRepository interface {
		AddDeposit(deposit *entity.Deposit) error
	}

	DepositService interface {
		AddDeposit(deposit *entity.Deposit) error
	}
)
