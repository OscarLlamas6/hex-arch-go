package ports

import "github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"

type (
	UserRepository interface {
		Create(value interface{}) error
		First(out interface{}, conditions ...interface{}) error
	}

	UserService interface {
		Create(user *entity.User) error
		Login(credentials *entity.DefaultCredentials) (string, error)
	}
)
