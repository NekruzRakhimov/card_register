package models

type Info struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
