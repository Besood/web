package routers

import (
	"net/http"
	"git/web/db"
	"github.com/gin-gonic/gin"
)

func UserSignup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	current:=db.InsertDB(username,password)
	if current != true{
		c.Redirect(http.StatusMovedPermanently,"//localhost:8080/signup")
		return
	}
	c.Redirect(http.StatusMovedPermanently,"//localhost:8080/")
}

func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	current:=db.SelectDB(username,password)
	if current!=true {
		c.Redirect(http.StatusMovedPermanently,"//localhost:8080/login")
		return
	}
	c.HTML(http.StatusOK,"index.html",gin.H{"username":username})
}

func UserPage(c *gin.Context){
	name:=c.Param("username")
	c.HTML(http.StatusOK,"user.html",gin.H{"username":name})
}