package models

type DefaultResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Success bool   `json:"success"`
}

type TokenResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Token   string `json:"token"`
}
