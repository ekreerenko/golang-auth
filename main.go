package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := engine()
	router.Use(gin.Logger())

	if err := engine().Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func engine() *gin.Engine {
	router := gin.New()
	router.Use(sessions.Sessions("access_session", sessions.NewCokieStore([]byte("secret"))))

	router.POST("/login", login)
	router.GET("/logout", logout)

	private := router.Group("/private")
	private.Use(AuthRequired) {
		private.GET("/me", me)
		private.GET("/status", status)
	}
	return router
}

// https://github.com/Depado/gin-auth-example
