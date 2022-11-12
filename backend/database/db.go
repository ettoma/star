package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

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

	var users = GetAllUsers()
	var id int
	usernameLower := strings.ToLower(username)

	if len(users) > 0 {
		id = users[len(users)-1].Id + 1
	} else {
		id = 1
	}

	for _, user := range users {
		user.Username = strings.ToLower(user.Username)
		if usernameLower == user.Username {
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

func deleteAll() {

	_, err = db.Exec(`DELETE FROM users`)
	utils.HandleWarning(err)
}

func DeleteUserById(id int) (bool, error) {

	users := GetAllUsers()

	for _, user := range users {
		if user.Id == id {
			_, err = db.Exec(`DELETE FROM users WHERE id = $1`, id)
			utils.HandleWarning(err)
			fmt.Print("user deleted")
			return true, nil
		}
	}
	return false, errors.New("user not found")

}

func DeleteUserByUsername(username string) (bool, error) {

	users := GetAllUsers()

	for _, user := range users {
		if user.Username == username {
			_, err = db.Exec(`DELETE FROM users WHERE username = $1`, username)
			utils.HandleWarning(err)
			fmt.Print("user deleted")
			return true, nil
		}
	}
	return false, errors.New("user not found")

}
