package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igor570/eventexplorer/models"
)

func main() {
	app := gin.Default()

	app.GET("/events", getEvents)

	app.POST("/events", createEvent)

	app.Run(":3100") // localhost:3100
}

// Get all events
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

// Create an event
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event data", "error": err.Error()})
		return
	}

	// Simulate setting the ID and UserID
	event.ID = len(models.GetAllEvents()) + 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"Message": "Event has been created!"})
}
