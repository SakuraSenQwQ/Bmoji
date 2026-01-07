package router

import (
	"main/config"

	"github.com/gin-gonic/gin"
)

func StartUp() {
	r := gin.Default()
	v1 := r.Group("v1")
	v1.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", CheckCors(c.GetHeader("origin")))
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.JSON(200, gin.H{"message": "OK"})
			return
		}
		c.Next()
	})
	e := v1.Group("emote")
	Router_Emote(e)
	r.Run(":1270")
}
func CheckCors(url string) string {
	list := config.Cfg.App.Cors
	for _, v := range list {
		if v == url {
			return url
		}
	}
	return ""
}
