package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := session()
	router.Use(gin.Logger())

	if err := session().Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

// func sessionEngine() *gin.Engine {
// 	router := gin.New()
// 	router.Use(sessions.Sessions("access_session", sessions.NewCokieStore([]byte("secret"))))

// 	router.POST("/login", login)
// 	router.GET("/logout", logout)

// 	private := router.Group("/private")
// 	private.Use(AuthRequired)
// 	{
// 		private.GET("/me", me)
// 		private.GET("/status", status)
// 	}
// 	return router
// }

// func AuthRequired() {

// }

// func login() {

// }

// func logout() {

// }

// func me() {

// }

// func status() {

// }

// https://github.com/Depado/gin-auth-example
