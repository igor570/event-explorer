package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/igor570/eventexplorer/middlewares"
)

func RegisterRoutes(app *gin.Engine) {

	//Events
	app.GET("/events", getEvents)
	app.GET("/events/:id", getEventById)

	//Protected Events Routes
	authenticatedRoute := app.Group("/")
	authenticatedRoute.Use(middlewares.Authenticate) //use our middleware

	authenticatedRoute.POST("/events", createEvent)
	authenticatedRoute.PUT("/events/:id", updateEvent)
	authenticatedRoute.DELETE("/events/:id", deleteEvent)

	//Authentication
	app.POST("/signup", signUp)
	app.POST("/login", login)
}
