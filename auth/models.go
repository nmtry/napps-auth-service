package auth

import (
	"time"
)

type RegisterRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type RegisterResponse struct {
	Id string `json:"id"`
}

type User struct {
	Id          string
	Firstname   string
	Lastname    string
	Email       string
	Username    string
	Password    string
	CreatedTime time.Time
}
