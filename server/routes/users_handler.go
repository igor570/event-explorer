package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igor570/eventexplorer/models"
	"github.com/igor570/eventexplorer/utils"
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

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "Failed to authenticate the user"})
		return
	}

	token, err := utils.CreateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Failed to get token for user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Login successful!", "Authorization": token})

}
