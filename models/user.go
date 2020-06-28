package models

type User struct {
	FirstName string `json:"firstname",form:"firstname",bson:"firstname"`
	LastName  string `json:"lastname",form:"lastname",bson:"lastname"`
	Email     string `json:"email",form:"email",bson:"email"`
	Password  string `json:"password",form:"password",bson:"password"`
}
