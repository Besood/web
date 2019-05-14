package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func UserSignup(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	println("username:" + username)
	println("email:" + email)
	println("password:" + password)

	c.Redirect(http.StatusSeeOther, "//localhost:8080/")
}

func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	println("username:" + username)
	println("password:" + password)
	c.HTML(http.StatusOK,"member.html",gin.H{"username": username})
}
