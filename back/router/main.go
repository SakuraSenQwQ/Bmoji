package router

import (
	"fmt"
	"main/config"

	"github.com/gin-gonic/gin"
)

func StartUp() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		fmt.Println(c.GetHeader("origin"))
		c.Writer.Header().Set("Access-Control-Allow-Origin", CheckCors(c.GetHeader("origin")))
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.JSON(200, gin.H{"message": "OK"})
			return
		}
		c.Next()
	})
	v1 := r.Group("v1")

	e := v1.Group("emote")
	Router_Emote(e)
	fmt.Println("已启动")
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
