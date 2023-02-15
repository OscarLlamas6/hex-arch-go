package entity

type (
	// Struct that represents email/password combination
	DefaultCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
