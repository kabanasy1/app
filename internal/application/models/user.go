package models

type User struct {
	Id           int64  `json:"id"`
	Firstname    string `json:"name"`
	Lastname     string `json:"surname"`
	Officenumber int64  `json:"office"`
	Phonenumber  string `json:"phone_number"`
	Position     string `json:"position"`
}
