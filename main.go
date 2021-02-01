package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	routes := router.Group("/api")
	{
		routes.POST("/login", login)
	}

	router.Run(":8080")
}

// https://github.com/Depado/gin-auth-example
// https://github.com/ektagarg/gin-gorm-todo-app
