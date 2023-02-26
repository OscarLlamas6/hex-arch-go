package couchdb

import (
	"context"
	"fmt"
	"github.com/go-kivik/kivik/v3"
	"github.com/lithammer/shortuuid"
)

type (
	Client struct {
		CLI *kivik.Client
		DB  *kivik.DB
	}
)

func NewCouchDBClient() (*Client, error) {

	err := connectDB()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &Client{CLI: couchDBClient, DB: couchDB}, nil
}

func (c *Client) AddDoc(ctx context.Context, id string, doc interface{}) error {
	_, err := c.DB.Put(ctx, id, doc)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (c *Client) AddExpense(ctx context.Context, doc interface{}) error {

	uuid := shortuuid.New()
	expensesID := fmt.Sprintf("expenses-%s", uuid)

	err := c.AddDoc(ctx, expensesID, doc)
	return err
}

func (c *Client) AddDeposit(ctx context.Context, doc interface{}) error {

	uuid := shortuuid.New()
	expensesID := fmt.Sprintf("deposits-%s", uuid)

	err := c.AddDoc(ctx, expensesID, doc)
	return err
}
