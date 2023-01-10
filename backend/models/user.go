package models

import "time"

type User struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type NewUser struct {
	Id       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
