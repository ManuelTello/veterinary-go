package session_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	application "github.com/ManuelTello/veterinary/internal/application"
	tests_functions "github.com/ManuelTello/veterinary/tests/functions"
)

var router http.Handler = application.New().GetServerHandler()

func TestSignIn(t *testing.T) {
	request := httptest.NewRequest("POST", "/api/v1/session/signin", tests_functions.CreateSignInBody())
	request.Header.Set("Authorization", "ApiKey "+os.Getenv("API_KEY"))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	statusCode := recorder.Result().StatusCode
	if statusCode != http.StatusOK {
		t.Fatal("Wrong status code, given: ", statusCode, ", expected 200")
	}
}

func TestSignUp(t *testing.T) {
	request := httptest.NewRequest("POST", "/api/v1/session/signup", tests_functions.CreateSignUpBody())
	request.Header.Set("Authorization", "ApiKey "+os.Getenv("API_KEY"))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	statusCode := recorder.Result().StatusCode
	if statusCode != http.StatusCreated {
		t.Fatal("Wrong status code, given: ", statusCode, ", expected 201")
	}
}
