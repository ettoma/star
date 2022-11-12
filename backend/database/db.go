package database

import (
	"database/sql"
	"errors"
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

func AddUser(name, username string) (models.User, error) {

	//TODO : normalise characters to check for lower/upper case conflicts
	var users = GetAllUsers()

	//TODO : get id from last inserted record, not from table length
	id := len(users) + 1

	for _, user := range users {
		if username == user.Username {
			return models.User{}, errors.New("username already exists")
		}
	}

	_, err = db.Exec(`INSERT INTO users VALUES ($1,$2,$3)`, name, username, id)

	utils.HandleWarning(err)

	return models.User{Name: name, Username: username, Id: id}, nil

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

		users = append(users, models.User{
			Name:     name,
			Username: username,
			Id:       id,
		})
	}
	return users

}

func DeleteAll() {

	_, err = db.Exec(`DELETE FROM users`)
	utils.HandleWarning(err)
}

func DeleteUserById(id int) {

	//TODO : implement ID check

	_, err = db.Exec(`DELETE FROM users WHERE id = $1`, id)
	utils.HandleWarning(err)

}

func DeleteUserByUsername(username string) {

	//TODO : implement username check

	_, err = db.Exec(`DELETE FROM users WHERE username = $1`, username)
	utils.HandleWarning(err)

}
