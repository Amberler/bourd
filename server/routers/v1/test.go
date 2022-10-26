package v1

import (
	"api/pkg/e"
	"api/pkg/util"
	"github.com/gin-gonic/gin"
)

func A(c *gin.Context) {
	var mobile, passwd = "", ""
	if c.Request.Method == "GET" {
		mobile = c.DefaultQuery("mobile", "13011110000") //取出参数手机号mobile
		passwd = c.DefaultQuery("passwd", "123456")      //取出密码
	} else {
		mobile = c.PostForm("mobile") //取出参数手机号mobile
		passwd = c.PostForm("passwd") //取出密码
	}

	util.ResponseWithJson(e.SUCCESS, gin.H{
		"User": map[string]interface{}{
			"mobile": mobile,
			"passwd": passwd,
			"hello":  "Hello Word",
		},
	}, c)
}
