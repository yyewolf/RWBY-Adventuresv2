package main

import (
	"database/sql"
	"time"

	"gopkg.in/mgutz/dat.v2/dat"
	runner "gopkg.in/mgutz/dat.v2/sqlx-runner"
)

var database *runner.DB

func init() {
	// create a normal database connection through database/sql
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", "dbname=rwby user=admin password=ftT6A4MrF6hPt host=admin.rwbyadventures.com")
	if err != nil {
		panic(err)
	}

	// ensures the database can be pinged with an exponential backoff (15 min)
	runner.MustPing(db)

	// set to reasonable values for production
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(50)

	// set this to enable interpolation
	dat.EnableInterpolation = true

	// set to check things like sessions closing.
	// Should be disabled in production/release builds.
	dat.Strict = false

	// Log any query over 10ms as warnings. (optional)
	runner.LogQueriesThreshold = 500 * time.Millisecond

	database = runner.NewDB(db, "postgres")
}
