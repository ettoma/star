package database

import (
	"database/sql"
	"os"

	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
)

var db *sql.DB
var err error

func OpenDb() {
	connStr := os.Getenv("DBURL")
	db, err = sql.Open("postgres", connStr)
	utils.HandleFatal(err)
}

func AddUser(name, username string) {

	id := len(GetAllUsers()) + 1

	_, err = db.Exec(`INSERT INTO users VALUES ($1,$2,$3)`, name, username, id)

	utils.HandleWarning(err)

}

func GetAllUsers() []models.User {

	rows, err := db.Query("SELECT * FROM users")

	utils.HandleWarning(err)
	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var name string
		var username string
		var id int
		utils.HandleWarning(rows.Scan(&name, &username, &id))
		// fmt.Printf("Name: %s \nUsername: %s \nID: %d\n", name, username, id)

		users = append(users, models.User{
			Name:     name,
			Username: username,
			Id:       id,
		})
	}
	return users

}

func DeleteUserByUsername(username string) {

	_, err = db.Exec(`DELETE FROM users WHERE username = $1`, username)
	utils.HandleWarning(err)

}

func DeleteUserById(id int) {

	_, err = db.Exec(`DELETE FROM users WHERE id = $1`, id)
	utils.HandleWarning(err)

}
