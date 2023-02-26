package entity

import (
	"fmt"
	"time"

	"github.com/OscarLlamas6/hex-arch-go/settings"
)

type (

	// Struct that represents a user entity
	User struct {
		ID        int
		Name      string
		Lastname  string
		Email     string `gorm:"unique"`
		Password  string
		CreatedAt *time.Time // Cuando se creo el registro
		UpdatedAt *time.Time // Cuando se actualizo el registro
		DeletedAt *time.Time // Cuando se borro el registro
	}
)

func (User) TableName() string {
	return fmt.Sprintf("%s.users", settings.AppConfig.DBSchema)
}
