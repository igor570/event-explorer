package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.Engine) {

	//Events
	app.GET("/events", getEvents)
	app.GET("/events/:id", getEventById)
	//Protected Events Routes
	app.POST("/events", createEvent)
	app.PUT("/events/:id", updateEvent)
	app.DELETE("/events/:id", deleteEvent)

	//Authentication
	app.POST("/signup", signUp)
	app.POST("/login", login)
}
