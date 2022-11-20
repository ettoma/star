package database

import (
	"errors"
	"strings"
	"time"

	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
)

func AddKudos(sender, receiver string, content string) (*models.Kudos, error) {

	var kudos = GetAllKudos()
	var id int
	senderLower := strings.ToLower(sender)
	receiverLower := strings.ToLower(receiver)

	if len(kudos) > 0 {
		id = kudos[len(kudos)-1].Id + 1
	} else {
		id = 1
	}

	ts := time.Now().Unix()

	_, err = db.Exec(`INSERT INTO kudos VALUES ($1,$2,$3,$4,$5)`, id, senderLower, receiverLower, ts, content)

	utils.HandleWarning(err)

	return &models.Kudos{Sender: senderLower, Receiver: receiverLower, Content: content, Timestamp: ts, Id: id}, nil
}

func GetAllKudos() []*models.Kudos {
	rows, err := db.Query("SELECT * FROM kudos")

	utils.HandleWarning(err)
	defer rows.Close()

	var kudos []*models.Kudos

	for rows.Next() {

		var sender string
		var receiver string
		var id int
		var createdAt int64
		var content string
		utils.HandleWarning(rows.Scan(&id, &sender, &receiver, &createdAt, &content))

		kudos = append(kudos, &models.Kudos{
			Sender:    sender,
			Receiver:  receiver,
			Id:        id,
			Content:   content,
			Timestamp: createdAt,
		})
	}
	return kudos
}

func GetKudosPerReceiver(user string) ([]*models.Kudos, error) {
	rows, err := db.Query("SELECT * FROM kudos WHERE receiver = $1", user)
	defer rows.Close()
	utils.HandleWarning(err)

	var kudos []*models.Kudos

	for rows.Next() {

		var sender string
		var receiver string
		var id int
		var createdAt int64
		var content string
		utils.HandleWarning(rows.Scan(&id, &sender, &receiver, &createdAt, &content))

		kudos = append(kudos, &models.Kudos{
			Sender:    sender,
			Receiver:  receiver,
			Id:        id,
			Content:   content,
			Timestamp: createdAt,
		})
	}

	if len(kudos) == 0 {
		return nil, errors.New("No kudos found")
	}
	return kudos, nil
}

func GetKudosPerSender(user string) ([]*models.Kudos, error) {
	rows, err := db.Query("SELECT * FROM kudos WHERE sender = $1", user)
	defer rows.Close()
	utils.HandleWarning(err)

	var kudos []*models.Kudos

	for rows.Next() {

		var sender string
		var receiver string
		var id int
		var createdAt int64
		var content string
		utils.HandleWarning(rows.Scan(&id, &sender, &receiver, &createdAt, &content))

		kudos = append(kudos, &models.Kudos{
			Sender:    sender,
			Receiver:  receiver,
			Id:        id,
			Content:   content,
			Timestamp: createdAt,
		})
	}

	if len(kudos) == 0 {
		return nil, errors.New("No kudos found")
	}
	return kudos, nil
}

func DeleteKudos(id int) (bool, error) {
	res, _ := db.Exec(`DELETE FROM kudos WHERE id = $1`, id)
	deletedRows, _ := res.RowsAffected()

	if deletedRows == 0 {

		return false, errors.New("kudos not found")
	} else {
		return true, nil
	}

}
