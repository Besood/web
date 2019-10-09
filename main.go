package main

import (
	"git/web/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("assets","views/assets")
	router.Static("image","views/image")

	user := router.Group("/user")
	{
		user.GET("/:username",routers.UserPage)
	}
	router.GET("/", routers.Home)

	router.GET("/login", routers.Login)
	router.POST("/login", routers.UserLogin)

	router.GET("/signup", routers.Signup)
	router.POST("/signup", routers.UserSignup)
	
	router.NoRoute(routers.NoRoute)

	router.Run()
}