package entity

import (
	"time"
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
