package main

import (
	"fmt"
	"github.com/OscarLlamas6/hex-arch-go/internal/infra/repositories/pg"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/service/user"
	"github.com/OscarLlamas6/hex-arch-go/settings"
)

func main() {
	settings.SetConfig()

	userRepository, err := pg.NewClient()
	if err != nil {
		fmt.Println(err)
	}

	userService := user.NewService(userRepository)

	jwtToken, err := userService.Login(&entity.DefaultCredentials{
		Email:    "oscar@example.com",
		Password: "admin123",
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Token: %s\n", jwtToken)
	}

	//api.RunServer()
}
