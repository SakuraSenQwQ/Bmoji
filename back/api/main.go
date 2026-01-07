package api

import "github.com/gin-gonic/gin"

func Send(code int, msg string) gin.H {
	return gin.H{"code": code, "msg": msg}
}
