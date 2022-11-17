package models

type Kudos struct {
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Content   string `json:"content"`
	Id        int    `json:"id"`
	Timestamp int64  `json:"timestamp"`
}
