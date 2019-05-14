package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func NoRoute(c *gin.Context) {
	c.HTML(http.StatusNotFound,"NotFound.html", gin.H{})
}
