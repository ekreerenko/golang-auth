package sessionAuth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiGroup() *gin.Engine {
	router := gin.New()
	router.Use(sessions.Sessions("access_session", sessions.NewCokieStore([]byte("secret"))))

	router.POST("/login", login)
	router.GET("/logout", logout)

	private := router.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/me", me)
		private.GET("/status", status)
	}
	return router
}

func AuthRequired(context *gin.Context) {
	session = session.Default(context)
	user := sessiong.Get(userKey)

	if user == nil {
		context.AbortWithStatusJson(http.StatusUnauthorized, gin.H{"error": "unanthorized"})
		return
	}
	context.Next()
}
