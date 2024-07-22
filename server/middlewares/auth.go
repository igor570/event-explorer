package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igor570/eventexplorer/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization") //check if jwt token exists on request headers

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Could not find Authorization token on request"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "token failed verification"})
		return
	}

	context.Set("userId", userId)
	context.Next() //ensuring the next request handler executes correctly.
}
