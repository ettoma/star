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
}

func getSender() models.User {
	return users[0]
}

func submitReview() {
	review := models.Review{

		Content:   "Test",
		Timestamp: time.Now().Unix(),
	}
	review.Sender = getSender()
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
	err := addUser("Borat", "mrborat")
	utils.HandleFatal(err)
	submitReview()
}
