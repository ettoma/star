package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ettoma/star/utils"
)

var db *sql.DB
var err error

func OpenDb() {
	connStr := os.Getenv("DBURL")
	db, err = sql.Open("postgres", connStr)
	utils.HandleFatal(err)
}

func ping() error {
	return db.Ping()
}

func KeepAlive() {
	for {
		err := ping()
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Println("keep alive: success")
			fmt.Printf("open connections: %v \n", db.Stats().OpenConnections)
			time.Sleep(10 * time.Second)
		}
	}
}
