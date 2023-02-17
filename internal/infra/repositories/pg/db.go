package pg

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"

	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var onceDBLoad sync.Once

var tables = []interface{}{
	&entity.User{},
	&entity.Deposit{},
	&entity.Expense{},
}

func connectDB() *gorm.DB {

	var err error

	onceDBLoad.Do(func() {
		var i int
		for {
			if i >= 30 {
				log.Fatal("Error connecting to database")
				err = errors.New("Error connecting to database")
				break
			}
			time.Sleep(3 * time.Second)
			db, err = gorm.Open(postgres.Open(GetDSN()), &gorm.Config{})
			if err != nil {
				logger.Info("Retrying in 3 seconds...")
				i++
				continue
			} else {
				runMigrations()
				logger.Info("Connected to database :D")
			}
			break
		}
	})

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return db
}

func runMigrations() {

	for _, table := range tables {
		err := db.AutoMigrate(table)
		if err != nil {
			logger.Info("Error while running migration")
			return
		}
	}

	logger.Info("Migrations ran successfully :P")
}
