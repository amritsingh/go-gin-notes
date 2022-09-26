package main

import (
	"gin_notes/controllers"
	controller_helpers "gin_notes/controllers/helpers"
	"gin_notes/middlewares"
	"gin_notes/models"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/*")

	models.ConnectDatabase()
	models.DBMigrate()

	// Sessions Init
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))

	r.Use(middlewares.AuthenticateUser())

	// Route Group - Notes
	notes := r.Group("/notes")
	{
		notes.GET("/", controllers.NotesIndex)
		notes.GET("/new", controllers.NotesNew)
		notes.POST("/", controllers.NotesCreate)
		notes.GET("/:id", controllers.NotesShow)
		notes.GET("/edit/:id", controllers.NotesEditPage)
		notes.POST("/:id", controllers.NotesUpdate)
		notes.DELETE("/:id", controllers.NotesDelete)
	}

	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", controller_helpers.SetPayload(c, gin.H{
			"title":     "Notes application",
			"logged_in": controller_helpers.IsUserLoggedIn(c),
		}))
	})

	log.Println("Server started!")
	r.Run() // Default Port 8080
}
