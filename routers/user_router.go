package routers

import (
	"net/http"
	"git/web/db"
	"github.com/gin-gonic/gin"
)

func UserSignup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag,msg:=db.InsertDB(username,password)
	if flag{
		c.Redirect(http.StatusMovedPermanently,"//localhost:8080/")
	}else{
		c.Redirect(http.StatusMovedPermanently,"//localhost:8080/signup")
	}
	print(msg)
}

func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag,msg:=db.SelectDB(username,password)
	if flag {
		c.Redirect(http.StatusMovedPermanently,"//localhost:8080/user/"+username)
	}else{
		c.Redirect(http.StatusMovedPermanently,"//localhost:8080/login")
	}
	println(username,msg)
}

func UserPage(c *gin.Context){
	name:=c.Param("username")
	c.HTML(http.StatusOK,"user.html",gin.H{"username":name})
}
