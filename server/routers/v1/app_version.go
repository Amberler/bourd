package v1

import (
	"api/pkg/e"
	"github.com/gin-gonic/gin"
)

func GetAppVersion(c *gin.Context) {
	c.JSON(e.SUCCESS, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": "1.0.0",
	})
}
