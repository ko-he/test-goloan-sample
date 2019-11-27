package handler

import (
	"github.com/gin-gonic/gin"
)

// ClientHandler interface
type ClientHandler interface {
	GetClients(c *gin.Context)
}

type clientHandler struct{}

func (ch *clientHandler) GetClients(c *gin.Context) {
	c.JSON(200, gin.H{
		"hoge": "fuga",
	})
}
