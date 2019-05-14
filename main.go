package main

import (
	"git/web/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	user := router.Group("/user")
	{
		user.POST("/signup", routers.UserSignup)
		user.POST("/login", routers.UserLogin)
	}

	router.GET("/", routers.Home)
	router.GET("/login", routers.Login)
	router.GET("/signup", routers.Signup)
	router.NoRoute(routers.NoRoute)

	router.Run()
}
