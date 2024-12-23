package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "new user created succesfully."})
}

func login(context *gin.Context)  {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data."})
	return 
	}

	err = user.ValidateUser()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message" : "could not authenticate user"})
		return 
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message" :"could not authenticate user"})
		return 
	}


	context.JSON(http.StatusOK, gin.H{"message" : "login successful", "token": token} )
}