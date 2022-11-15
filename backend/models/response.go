package models

type DefaultResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Success bool   `json:"success"`
}
