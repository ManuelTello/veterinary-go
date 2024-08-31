package session_controller

import (
	"net/http"

	session_dto "github.com/ManuelTello/veterinary/internal/dto/session"
	session_service "github.com/ManuelTello/veterinary/internal/services/session"
	session_viewmodel "github.com/ManuelTello/veterinary/internal/viewmodels/session"
	gin "github.com/gin-gonic/gin"
)

func SignInTemplate() gin.HandlerFunc {
	return func(c *gin.Context) {
		viewmodel := session_viewmodel.SginInViewModel{Title: "Sign in"}
		c.HTML(http.StatusOK, "signin.html", viewmodel)
	}
}

func SignUpTemplate() gin.HandlerFunc {
	viewmodel := session_viewmodel.SginInViewModel{Title: "Sign up"}
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", viewmodel)
	}
}

func ProcessSignUp(service session_service.SessionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		signupBody := new(session_dto.IncomingSignUp)
		c.ShouldBindBodyWithJSON(signupBody)
		err := service.ProcessSignUp(*signupBody)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "500 internal server error")
		} else {
			c.JSON(http.StatusCreated, "201 user created")
		}
	}
}

func ProcessSignIn(service session_service.SessionService) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
