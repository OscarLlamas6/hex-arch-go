package pg

import (
	"log"

	"gorm.io/gorm"
)

type (
	client struct {
		db *gorm.DB
	}
)

// Return new client instance
func NewClient() (*client, error) {

	err := connectDB()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &client{db: db}, nil
}

// Create new register
func (c *client) Create(value interface{}) error {
	return c.db.Create(value).Error
}

func (c *client) First(out interface{}, conditions ...interface{}) error {
	return c.db.First(out, conditions...).Error
}
