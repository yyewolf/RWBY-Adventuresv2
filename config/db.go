package config

import (
	"fmt"

	"gorm.io/gorm"

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
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("[DATABASE] Connected.")
}
