package config

import (
	"fmt"

	"gorm.io/gorm"
	// Not needed for import.
	"gorm.io/driver/postgres"
)

// Database is the object uses by the models for accessing
// database tables and executing queries.
var Database *gorm.DB

// func init() {
// 	var err error
// 	Database, err = gorm.Open("postgres", fmt.Sprintf("host=admin.rwbyadventures.com port=5432 user=admin dbname=rwbyorm password=%s sslmode=disable", dbpswd))

// 	if err != nil {
// 		panic(err)
// 	}

// 	// set this to 'true' to see sql logs
// 	Database.LogMode(false)

// 	fmt.Println("[DATABASE] Connected.")
// }

func init() {
	var err error
	Database, err = gorm.Open(postgres.Open("host=admin.rwbyadventures.com port=5432 user=admin dbname=rwbyorm password="+dbpswd+" sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("[DATABASE] Connected.")
}
