package models

type Review struct {
	Sender    User
	Receiver  User
	Content   string
	Timestamp int64
}
