package util

import (
	"api/pkg/e"
	"github.com/gin-gonic/gin"
)

func ResponseWithJson(code int, data interface{}, c *gin.Context) {
	c.JSON(code, gin.H{
		"Code": code,
		"Msg":  e.GetMsg(code),
		"Data": data,
	})
}
