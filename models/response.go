package models

type Response struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

type ResponseUsersList struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
	Result     []User `json:"result"`
}
