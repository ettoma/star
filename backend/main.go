package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
)

var users = []models.User{
	{
		Name:     "Borat",
		Username: "mrborat",
	},
	{
		Name:     "Joaquin Phoenix",
		Username: "totallyJoker",
	},
}

func getUserData(username string) (models.User, error) {
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	errMsg := fmt.Sprintf("user: %s not found", username)
	return models.User{}, errors.New(errMsg)

}

func submitReview(sender, receiver, content string) {
	review := models.Review{
		Content:   content,
		Timestamp: time.Now().Unix(),
	}

	r, err := getUserData(receiver)
	utils.HandleFatal(err)

	s, err := getUserData(sender)
	utils.HandleFatal(err)

	review.Receiver = r
	review.Sender = s

	fmt.Println(review)

}

func addUser(name, username string) error {
	newUser := models.User{
		Name:     name,
		Username: username,
	}

	for _, u := range users {
		if u.Username == newUser.Username {
			errMsg := fmt.Sprintf("Username: %s already exists", newUser.Username)
			return errors.New(errMsg)
		}
	}

	users = append(users, newUser)
	return nil
}

func main() {
	// err := addUser("Borat", "mrborat")
	// utils.HandleFatal(err)
	submitReview("mrborat", "totallyJoker", "Jammin'")
}
