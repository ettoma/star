package models

type Kudos struct {
	Sender    User   `json:"sender"`
	Receiver  User   `json:"receiver"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}
