package router

import (
	"github.com/gin-gonic/gin"
)

// NewRouter func
func NewRouter(r *gin.Engine, ch *handler.clientsRouter) {
	clientsRouter := r.Group("/clients")
	clientsRouter.GET("", ch.GetClients)
}
