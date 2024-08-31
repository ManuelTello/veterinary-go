package session_dto

import "time"

type IncomingSignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type IncomingSignUp struct {
	Password    string    `json:"password"`
	Username    string    `json:"username"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	DateCreated time.Time `json:"datecreated"`
	Email       string    `json:"email"`
	PhoneNumber *int      `json:"phonenumber"`
}

type SessionToken struct {
	Token string `json:"token"`
}
