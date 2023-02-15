package entity

import "time"

type (

	// Struct that represents a user entity
	User struct {
		ID        int
		Name      string
		Lastname  string
		Email     string
		Password  string
		CreatedAt *time.Time // Cuando se creo el registro
		UpdatedAt *time.Time // Cuando se actualizo el registro
		DeletedAt *time.Time // Cuando se borro el registro
	}
)
