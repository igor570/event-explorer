package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igor570/eventexplorer/models"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not save error"})

	}

	context.JSON(http.StatusCreated, gin.H{"Message": "User was created!"})

}
