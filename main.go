package main

import (
	"context"
	"fmt"
	"time"

	"github.com/OscarLlamas6/hex-arch-go/internal/infra/api"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"
	"github.com/lithammer/shortuuid"

	couchDB "github.com/OscarLlamas6/hex-arch-go/internal/infra/repositories/couchdb"
	"github.com/OscarLlamas6/hex-arch-go/settings"
)

func main() {
	settings.SetConfig()
	client, err := couchDB.NewCouchDBClient()
	if err != nil {
		fmt.Println(err)
	}

	goodConn, err := client.CLI.Ping(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	if goodConn {
		fmt.Println("Connected to CouchDB :D")
	}

	newExpense := &entity.Expense{
		UserID:      1,
		Category:    "Renta",
		Name:        "Pago de renta",
		Description: "Mes de Marzo",
		Amount:      340.00,
		Date:        time.Now(),
	}

	uuid := shortuuid.New()

	expensesID := fmt.Sprintf("expenses-%s", uuid)

	err = client.AddDoc(context.Background(), expensesID, newExpense)
	if err != nil {
		fmt.Println(err)
	}

	api.RunServer()
}
