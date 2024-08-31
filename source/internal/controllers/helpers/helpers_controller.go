package helpers_controller

import (
	"net/http"

	helpers_dto "github.com/ManuelTello/veterinary/internal/dto/helpers"
	helpers_service "github.com/ManuelTello/veterinary/internal/services/helpers"
	gin "github.com/gin-gonic/gin"
)

func SearchIfUsernameIsRepeated(service helpers_service.HelpersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := new(helpers_dto.RepeatedUsername)
		c.ShouldBindBodyWithJSON(body)
		userExists, searchErr := service.SearchIfUsernameExists(body.Username)
		if searchErr != nil {
			c.JSON(http.StatusInternalServerError, "500 internal server error")
		} else {
			if userExists {
				c.JSON(http.StatusOK, true)
			} else {
				c.JSON(http.StatusOK, false)
			}
		}
	}
}

func SearchIfEmailIsRepeated(service helpers_service.HelpersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := new(helpers_dto.RepeatedEmail)
		c.ShouldBindBodyWithJSON(body)
		emailExists, searchErr := service.SearchIfEmailExists(body.Email)
		if searchErr != nil {
			c.JSON(http.StatusInternalServerError, "500 internal server error")
		} else {
			if emailExists {
				c.JSON(http.StatusOK, true)
			} else {
				c.JSON(http.StatusOK, false)
			}
		}
	}
}
