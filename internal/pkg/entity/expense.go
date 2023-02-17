package entity

import (
	"fmt"
	"time"

	"github.com/OscarLlamas6/hex-arch-go/settings"
)

type (
	// Struct that represents a expense entity
	Expense struct {
		UserID      int
		Category    string
		Name        string
		Description string
		Amount      float32
		Date        time.Time
	}
)

func (Expense) TableName() string {
	return fmt.Sprintf("%s.expenses", settings.AppConfig.DBSchema)
}
