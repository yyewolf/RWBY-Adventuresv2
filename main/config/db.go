package config

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// Not needed for import.
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database is the object uses by the models for accessing
// database tables and executing queries.
var Database *gorm.DB

func init() {
	var err error
	Database, err = gorm.Open("postgres", "host=admin.rwbyadventures.com port=5432 user=admin dbname=rwbyorm password=ftT6A4MrF6hPt sslmode=disable")

	if err != nil {
		panic(err)
	}

	// set this to 'true' to see sql logs
	Database.LogMode(false)

	fmt.Println("[DATABASE] Connected.")
}
