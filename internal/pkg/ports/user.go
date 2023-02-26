package ports

type (
	UserRepository interface {
		Create(value interface{}) error
		First(out interface{}, conditions ...interface{}) error
	}
)
