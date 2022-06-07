package session;

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
	fmt.Println("login handler")
	context.JSON(http.StatusOK, gin.H{"message": "login test post message"})
}
