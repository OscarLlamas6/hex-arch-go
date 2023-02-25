package couchdb

import (
	"context"
	"fmt"
	"github.com/go-kivik/kivik/v3"
)

type (
	client struct {
		CLI *kivik.Client
		DB  *kivik.DB
	}
)

func NewCouchDBClient() (*client, error) {

	err := connectDB()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &client{CLI: couchDBClient, DB: couchDB}, nil
}

func (c *client) AddDoc(ctx context.Context, id string, doc interface{}) error {
	_, err := c.DB.Put(ctx, id, doc)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
