// package sessionAuth

// import (
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// )

// func ApiGroup() *gin.Engine {
// 	router := gin.Default()

// 	authRouter := router.Group("/session")
// 	{
// 		authRouter.POST("/login", loginHandler)
// 	}

// 	return router
// 	// router := gin.Default()

// 	// router := gin.New()
// 	// router.Use(sessions.Sessions("access_session", sessions.NewCokieStore([]byte("secret"))))

// 	// router.POST("/login", login)
// 	// router.GET("/logout", logout)

// 	// private := router.Group("/auth-session")
// 	// private.Use(AuthRequired)
// 	// {
// 	// 	private.GET("/me", me)
// 	// 	private.GET("/status", status)
// 	// }
// 	// return router
// }

// func loginHandler(context *gin.Context) {
// 	fmt.Println("loginHandler")
// }

// // func AuthRequired(context *gin.Context) {
// // 	session = session.Default(context)
// // 	user := sessiong.Get(userKey)

// // 	if user == nil {
// // 		context.AbortWithStatusJson(http.StatusUnauthorized, gin.H{"error": "unanthorized"})
// // 		return
// // 	}
// // 	context.Next()
// // }
