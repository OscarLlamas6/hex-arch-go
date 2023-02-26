package ports

type (
	User interface {
		Create(value interface{}) error
		First(out interface{}, conditions ...interface{}) error
	}
)
