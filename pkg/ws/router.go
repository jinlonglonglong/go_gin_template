package ws

import (
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, wsh *WebSocketHandler) {
	router.GET("/ws", func(c *gin.Context) {
		wsh.HandleConn(c.Writer, c.Request)
	})
}
