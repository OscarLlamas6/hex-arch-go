package main

import (
	"github.com/OscarLlamas6/hex-arch-go/internal/infra/api"
	"github.com/OscarLlamas6/hex-arch-go/internal/infra/repositories/pg"
	"github.com/OscarLlamas6/hex-arch-go/settings"
)

func main() {
	settings.SetConfig()
	pg.ConnectDB()
	api.RunServer()
}
