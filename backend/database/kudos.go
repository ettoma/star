package database

import (
	"time"

	"github.com/ettoma/star/models"
)

func AddKudos(sender, receiver models.User, content string) (models.Kudos, error) {

	// users := GetAllUsers()

	return models.Kudos{Sender: sender, Receiver: receiver, Content: content, Timestamp: time.Now().Unix()}, nil
}
