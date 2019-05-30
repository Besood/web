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
		user.GET("/:username",routers.UserPage)
	}
	router.GET("/", routers.Home)
	router.GET("/file",routers.FileLoad)
	router.POST("/file",routers.Filelog)

	router.GET("/login", routers.Login)
	router.POST("/login", routers.UserLogin)

	router.GET("/signup", routers.Signup)
	router.POST("/signup", routers.UserSignup)

	router.GET("/api",routers.Api)
	
	router.NoRoute(routers.NoRoute)

	router.Run()
}
