package router

import (
	"main/api"

	"github.com/gin-gonic/gin"
)

func Router_Emote(c *gin.RouterGroup) {
	c.GET("/list", api.GetRemoteList)
	c.GET("/info", api.GetEmoteInfo)
}
