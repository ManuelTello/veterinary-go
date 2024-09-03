package application

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	helpers_handlers "github.com/ManuelTello/veterinary/internal/handlers/helpers"
	notfound_handlers "github.com/ManuelTello/veterinary/internal/handlers/notfound"
	session_handlers "github.com/ManuelTello/veterinary/internal/handlers/session"
	audit_model "github.com/ManuelTello/veterinary/internal/models/audit"
	session_model "github.com/ManuelTello/veterinary/internal/models/session"
	providers_store "github.com/ManuelTello/veterinary/internal/providers/store"
	helpers_service "github.com/ManuelTello/veterinary/internal/services/helpers"
	session_service "github.com/ManuelTello/veterinary/internal/services/session"
	gin "github.com/gin-gonic/gin"
)

type Application struct {
	ginServer *gin.Engine
	sqlStore  *sql.DB
}

func (application Application) GetServerHandler() http.Handler {
	return application.ginServer.Handler()
}

func (application Application) StartServer() {
	runErr := application.ginServer.Run(":" + os.Getenv("PORT"))
	defer func() {
		application.sqlStore.Close()
	}()
	if runErr != nil {
		log.Fatal("Gracefully stopping server, following:\n", runErr)
	}
}

func (application Application) SetUpRoutes() {
	application.ginServer.StaticFS("/public", http.Dir(filepath.Join(os.Getenv("GO_CWD"), "www", "static")))
	application.ginServer.StaticFile("/favicon.ico", filepath.Join(os.Getenv("GO_CWD"), "www", "static", "favicon.ico"))
	application.ginServer.NoRoute(notfound_handlers.NotFound())

	application.ginServer.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.html", nil)
	})

	// Globally accesible HTML template routes
	sessionGroup := application.ginServer.Group("/session")
	sessionGroup.Use()
	{
		sessionGroup.GET("signin", session_handlers.SignInTemplate())
		sessionGroup.GET("signup", session_handlers.SignUpTemplate())
	}

	// Specific API routes
	apiGroup := application.ginServer.Group("/api")
	apiGroup.Use( /*middlewares.AuthValidation()*/ )
	{
		/*
			API VERSION 1.0
		*/
		apiVersion_01 := apiGroup.Group("v1")
		apiVersion_01.Use()
		{
			apiSession := apiVersion_01.Group("session")
			apiSession.Use()
			{
				service := session_service.New(session_model.New(application.sqlStore), audit_model.New(application.sqlStore))
				apiSession.POST("signin", session_handlers.ProcessSignIn(service))
				apiSession.POST("signup", session_handlers.ProcessSignUp(service))
			}

			apiHelpers := apiVersion_01.Group("helper")
			apiHelpers.Use()
			{
				service := helpers_service.New(session_model.New(application.sqlStore))
				apiHelpers.POST("repeatedemail", helpers_handlers.SearchIfEmailIsRepeated(service))
			}
		}
	}
}

func New() Application {
	store, storeErr := providers_store.NewStore()
	if storeErr != nil {
		log.Fatal("error trying to create store:\n", storeErr)
	}

	application := Application{
		ginServer: gin.Default(),
		sqlStore:  store,
	}

	application.ginServer.LoadHTMLGlob(filepath.Join(os.Getenv("GO_CWD"), "www", "views") + "/**/*")
	application.SetUpRoutes()

	return application
}
