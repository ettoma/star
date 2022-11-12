package models

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Id       int    `json:"id"`
}
