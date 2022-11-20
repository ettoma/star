package database

import (
	"errors"
	"strings"
	"time"

	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
)

func AddUser(name, username string) (*models.User, error) {

	if len(name) <= 3 || len(username) <= 3 {
		return nil, errors.New("name or username too short (min. 4 char)")
	}

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
			return nil, errors.New("username already exists")
		}
	}

	ts := time.Now().Unix()

	_, err = db.Exec(`INSERT INTO users VALUES ($1,$2,$3,$4)`, name, username, id, ts)

	utils.HandleWarning(err)

	return &models.User{Name: name, Username: username, Id: id, CreatedAt: time.Unix(ts, 0)}, nil

}

func GetAllUsers() []*models.User {

	rows, err := db.Query("SELECT * FROM users")

	utils.HandleWarning(err)
	defer rows.Close()

	var users []*models.User

	for rows.Next() {

		var name string
		var username string
		var id int
		var createdAt int64
		utils.HandleWarning(rows.Scan(&name, &username, &id, &createdAt))

		users = append(users, &models.User{
			Name:      name,
			Username:  username,
			Id:        id,
			CreatedAt: time.Unix(createdAt, 0),
		})
	}
	return users

}

func GetUserById(u int) (*models.User, error) {
	rows := db.QueryRow("SELECT * FROM users WHERE id = $1", u)
	utils.HandleWarning(err)

	var name string
	var username string
	var id int
	var createdAt int64

	err := rows.Scan(&name, &username, &id, &createdAt)

	if err != nil {
		return nil, err
	} else {
		return &models.User{Name: name, Username: username, Id: id, CreatedAt: time.Unix(createdAt, 0)}, nil
	}

}

func GetUserByUsername(u string) (*models.User, error) {
	rows := db.QueryRow("SELECT * FROM users WHERE username = $1", u)
	utils.HandleWarning(err)

	var name string
	var username string
	var id int
	var createdAt int64

	err := rows.Scan(&name, &username, &id, &createdAt)

	if err != nil {
		return nil, err
	} else {
		return &models.User{Name: name, Username: username, Id: id, CreatedAt: time.Unix(createdAt, 0)}, nil
	}

}

func DeleteUserById(id int) (bool, error) {

	res, _ := db.Exec(`DELETE FROM users WHERE id = $1`, id)
	deletedRows, _ := res.RowsAffected()

	if deletedRows == 0 {

		return false, errors.New("user not found")
	} else {
		return true, nil
	}

}

func DeleteUserByUsername(username string) (bool, error) {

	users := GetAllUsers()

	for _, user := range users {
		if user.Username == username {
			_, err = db.Exec(`DELETE FROM users WHERE username = $1`, username)
			utils.HandleWarning(err)
			return true, nil
		}
	}
	return false, errors.New("user not found")

}
