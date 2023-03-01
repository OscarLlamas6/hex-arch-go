package main

import (
	"context"
	"fmt"
	"time"

	cdb "github.com/OscarLlamas6/hex-arch-go/internal/infra/repositories/couchdb"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/service/deposit"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/service/expense"
	"github.com/OscarLlamas6/hex-arch-go/settings"
)

func main() {
	settings.SetConfig()

	// userRepository, err := pg.NewClient()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// userService := user.NewService(userRepository)

	// jwtToken, err := userService.Login(&entity.DefaultCredentials{
	// 	Email:    "oscar@example.com",
	// 	Password: "admin123",
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("Token: %s\n", jwtToken)
	// }

	couchRepository, err := cdb.NewCouchDBClient()
	if err != nil {
		fmt.Println(err)
	}

	depositService := deposit.NewService(couchRepository)

	err = depositService.AddDeposit(context.Background(), &entity.Deposit{
		UserID:      6,
		Category:    "Salarios",
		Name:        "Deposito en GyT",
		Description: "Pago salario Abril",
		Amount:      1000000,
		Date:        time.Now(),
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deposito salario en GyT exitoso :D")
	}

	expenseService := expense.NewService(couchRepository)

	err = expenseService.AddExpense(context.Background(), &entity.Expense{
		UserID:      6,
		Category:    "Vacaciones",
		Name:        "Suiza",
		Description: "Luna de miel",
		Amount:      250000,
		Date:        time.Now(),
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Gasto realizado exitosamente :D")
	}

	//api.RunServer()
}
