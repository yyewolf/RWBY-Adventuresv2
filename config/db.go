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
			fmt.Sprintf("host=192.168.1.145 port=51234 user=postgres dbname=postgres password=postgres sslmode=disable"),
		),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("[DATABASE] Connected.")
}
