package routers

import (

	"net/http"

	"github.com/gin-gonic/gin"
)

var msg struct {
	Name    string `json:"user"`
	Message string
	Number  int
}

func Api(c *gin.Context) {
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123
	c.JSON(http.StatusOK, msg)
}