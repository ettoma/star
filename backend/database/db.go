package database

import (
	"database/sql"
	"os"

	"github.com/ettoma/star/utils"
)

var db *sql.DB
var err error

func OpenDb() {
	connStr := os.Getenv("DBURL")
	db, err = sql.Open("postgres", connStr)
	utils.HandleFatal(err)
}
