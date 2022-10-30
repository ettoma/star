package models

type Review struct {
	Sender    User
	Receiver  string
	Content   string
	Timestamp int64
}
