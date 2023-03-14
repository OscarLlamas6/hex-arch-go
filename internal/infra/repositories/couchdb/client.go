package couchdb

import (
	"context"
	"fmt"

	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"
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

func (c *Client) AddExpense(expense *entity.Expense) error {

	uuid := shortuuid.New()
	ctx := context.Background()
	expensesID := fmt.Sprintf("expenses-%s", uuid)

	err := c.AddDoc(ctx, expensesID, expense)
	return err
}

func (c *Client) AddDeposit(deposit *entity.Deposit) error {

	uuid := shortuuid.New()
	ctx := context.Background()
	expensesID := fmt.Sprintf("deposits-%s", uuid)

	err := c.AddDoc(ctx, expensesID, deposit)
	return err
}
