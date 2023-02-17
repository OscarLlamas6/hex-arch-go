package entity

import (
	"fmt"
	"time"

	"github.com/OscarLlamas6/hex-arch-go/settings"
)

type (
	// Struct that represents a deposit entity
	Deposit struct {
		UserID      int
		Category    string
		Name        string
		Description string
		Amount      float32
		Date        time.Time
	}
)

func (Deposit) TableName() string {
	return fmt.Sprintf("%s.deposits", settings.AppConfig.DBSchema)
}
