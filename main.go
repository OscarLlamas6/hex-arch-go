package main

import (
	"github.com/OscarLlamas6/hex-arch-go/internal/infra/api"
	firestore "github.com/OscarLlamas6/hex-arch-go/internal/infra/repositories/firestore"
	"github.com/OscarLlamas6/hex-arch-go/settings"
)

func main() {
	settings.SetConfig()
	firestore.NewFirestoreClient()
	api.RunServer()
}
