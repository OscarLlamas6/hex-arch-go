package pg

import "gorm.io/gorm"

type (
	client struct {
		db *gorm.DB
	}
)

// Return new client instance
func NewClient() *client {
	return &client{db: connectDB()}
}

// Create new register
func (c *client) Create(value interface{}) error {
	return c.db.Create(value).Error
}
