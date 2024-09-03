package session_dto

import "time"

type IncomingSignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type IncomingSignUp struct {
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	FirstName         string    `json:"firs_tname"`
	LastName          string    `json:"last_name"`
	DateCreated       time.Time `json:"date_created"`
	PhoneNumber       *int      `json:"phone_number"`
	AlternativeNumber *int      `json:"alternative_number"`
}

type SessionToken struct {
	Token string `json:"token"`
}
