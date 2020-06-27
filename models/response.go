package models

type Response struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}
