package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	uses := router.Group("/user")
	{
		uses.POST("/signup", router.UserSignup)
		user.POST("/login", router.UserLogin)
	}

	router.GET("/", router.Home)
	router.GET("/login", router.Login)
	router.GET("/signup", router.Signup)
	router.NoRoute(router.NoRoute)

	router.Run()
}
