package database

import (
	"database/sql"
	"fmt"
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

func InsertUser(name, username string, id int) {

	_, err = db.Exec(`INSERT INTO users VALUES ($1,$2,$3)`, name, username, id)

	utils.HandleWarning(err)

}

func GetAllUsers() {

	result := db.QueryRow("SELECT * FROM users")

	utils.HandleWarning(result.Err())
	var name string
	var username string
	var id int
	utils.HandleWarning(result.Scan(&name, &username, &id))

	fmt.Printf("Name: %s \nUsername: %s \nID: %d", name, username, id)
}

func DeleteUserByUsername(username string) {

	_, err = db.Exec(`DELETE FROM users WHERE username = $1`, username)
	utils.HandleWarning(err)

}

func DeleteUserById(id int) {

	_, err = db.Exec(`DELETE FROM users WHERE id = $1`, id)
	utils.HandleWarning(err)

}
