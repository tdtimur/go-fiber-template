package models

type User struct {
	FirstName string `json:"firstName",form:"firstName",bson:"first_name"`
	LastName  string `json:"lastName",form:"lastName",bson:"last_name"`
	Email     string `json:"email",form:"email",bson:"email"`
	Password  string `json:"password",form:"password",bson:"password"`
}
