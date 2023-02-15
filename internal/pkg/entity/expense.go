package entity

import "time"

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
