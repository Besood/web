package routers

import (
	"net/http"
	"git/web/db"
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
	flag := db.InsertDB(username,password)
	if !flag{
		c.HTML(http.StatusBadRequest,"login.html",gin.H{})
	}else{
		c.HTML(http.StatusOK,"index.html",gin.H{})
	}
}
