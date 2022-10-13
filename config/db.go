package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// Not needed for import.
	"gorm.io/driver/postgres"
)

var Database *gorm.DB

func init() {
	var err error
	Database, err = gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbhost, dbport, dbuser, dbbase, dbpswd),
		),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             time.Second,
					LogLevel:                  logger.Silent,
					Colorful:                  true,
					IgnoreRecordNotFoundError: true,
				},
			),
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("[DATABASE] Connected.")
}
