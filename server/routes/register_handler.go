package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/igor570/eventexplorer/models"
)

func registerForEvent(context *gin.Context) {
	//get id for event
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event id from params"})
		return
	}

	//get user id from context
	userId := context.GetInt64("userId")

	//find the event in DB
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not find event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not register the user to event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Message": "Successfully registered user"})

}

func cancelRegistration(context *gin.Context) {
	//get id for event
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event id from params"})
		return
	}

	//get user id from context
	userId := context.GetInt64("userId")

	var event models.Event

	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Successfully canceled registration"})

}
