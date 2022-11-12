package models

import "time"

type User struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
