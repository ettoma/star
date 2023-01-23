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

func DbInit() {
	openDb()
	go keepAlive()
}

func openDb() {
	connStr := os.Getenv("DBURL")
	db, err = sql.Open("postgres", connStr)
	utils.HandleFatal(err)
}

func keepAlive() {
	for {
		err := db.Ping()
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Println("keep alive: success")
			fmt.Printf("open connections: %v \n", db.Stats().OpenConnections)
			time.Sleep(60 * time.Second)
		}
	}
}
