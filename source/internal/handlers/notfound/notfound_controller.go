package notfound_handlers

import (
	"net/http"

	notfound_viewmodel "github.com/ManuelTello/veterinary/internal/viewmodels/notfound"
	"github.com/gin-gonic/gin"
)

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		viewmodel := notfound_viewmodel.NotFoundViewModel{Title: "Not found"}
		c.HTML(http.StatusNotFound, "notfound.html", viewmodel)
	}
}
