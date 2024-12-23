package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	event, err := models.GetEventById(int64(eventId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch the event"})
		return
	}

	err = event.Register(int64(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "could not register for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"Registered!"})
	
}

func cancelRegistration(context *gin.Context)  {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return 
	}

	event, err := models.GetEventById(int64(eventId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch the event"})
		return 
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		if err.Error() == fmt.Sprintf("no registration found for event_id=%d and user_id=%d", event.ID, userId) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Registration not found"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cancel registration"})
		return}

	context.JSON(http.StatusOK, gin.H{"message" : "Registration canceled successfuly"})
}