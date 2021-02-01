package session

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
	session := session.Default()
	email := context.PostForm("email")
	password := context.PostForm("password")

	if strings.Trim(email, " ") == "" || strings.Trim(password, " ") == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error", "Parameters can't be empty"})
		return
	}

	if email != "test@test.com" && password != "test" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong email or password"})
		return
	}

	session.Set(userKey, email)
	if err := session.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User successfully authenticated"})
}
