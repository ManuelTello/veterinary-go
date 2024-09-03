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
		Email:    "myaddress@mail.com",
		Password: "michurri",
	}
	bodyBytes, _ := json.Marshal(body)
	bodyParsed := strings.NewReader(string(bodyBytes))
	return bodyParsed
}

func CreateSignUpBody() io.Reader {
	dcParse, _ := time.Parse("2006-01-02 15:04:05", "2024-01-02 15:04:05")
	phoneNumber := 111111
	body := session_dto.IncomingSignUp{
		Password:          "michurri",
		FirstName:         "John",
		LastName:          "Doe",
		DateCreated:       dcParse,
		Email:             "myaddress@mail.com",
		PhoneNumber:       &phoneNumber,
		AlternativeNumber: nil,
	}
	bodyBytes, _ := json.Marshal(body)
	bodyParsed := strings.NewReader(string(bodyBytes))
	return bodyParsed
}
