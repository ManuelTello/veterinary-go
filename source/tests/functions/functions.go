package tests_functions

import (
	"encoding/json"
	"io"
	"strings"
	"time"

	session_dto "github.com/ManuelTello/veterinary/internal/dto/session"
)

func CreateSignInBody() io.Reader {
	body := session_dto.IncomingSignIn{
		Username: "testuser",
		Password: "michurri",
	}
	bodyBytes, _ := json.Marshal(body)
	bodyParsed := strings.NewReader(string(bodyBytes))
	return bodyParsed
}

func CreateSignUpBody() io.Reader {
	dcParse, _ := time.Parse("2006-01-02 15:04:05", "2024-01-02 15:04:05")
	body := session_dto.IncomingSignUp{
		Username:    "testuser",
		Password:    "michurri",
		FirstName:   "John",
		LastName:    "Doe",
		DateCreated: dcParse,
		Email:       "testemail@mail.com",
		PhoneNumber: nil,
	}
	bodyBytes, _ := json.Marshal(body)
	bodyParsed := strings.NewReader(string(bodyBytes))
	return bodyParsed
}
