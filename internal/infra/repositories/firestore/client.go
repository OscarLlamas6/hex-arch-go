package firestore

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	firestore "cloud.google.com/go/firestore"
	"github.com/OscarLlamas6/hex-arch-go/settings"
	logger "github.com/sirupsen/logrus"
)

type (
	client struct {
		cli *firestore.Client
	}
)

var (
	onceFirestoreClient sync.Once
	firestoreClient     *client
)

const (
	ExpensesCollection = "expenses"
	DepositsCollection = "deposits"
)

func NewFirestoreClient() *client {

	var err error
	var cli *firestore.Client

	onceFirestoreClient.Do(func() {
		ctx := context.Background()
		var i int
		for {
			if i >= 30 {
				log.Fatal("Error connecting to Firestore")
				panic("Error connecting to Firestore")
			}
			time.Sleep(3 * time.Second)

			// CREATING FIRESTORE CLIENT FOR FIREBASE EMULATOR!!!!!!!!!!!!!!!!!

			cli, err = firestore.NewClient(ctx, settings.AppConfig.FirestoreProject)
			if err != nil {
				fmt.Println(err.Error())
				logger.Info("Retrying in 3 seconds...")
				i++
				continue
			} else {
				logger.Info("Firestore client configured correctly :D")
				fmt.Println("Firestore client configured correctly :D")
			}
			break
		}

	})

	firestoreClient = &client{
		cli: cli,
	}

	return firestoreClient
}

// Add document method for Firestore client
func (c *client) AddDoc(ctx context.Context, collection string, data interface{}) error {
	_, _, err := c.cli.Collection(collection).Add(ctx, data)

	return err
}
