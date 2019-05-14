package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin{})
}

func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup", gin.H{})
}

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
