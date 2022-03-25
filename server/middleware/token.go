package middleware

import (
	"api/pkg/e"
	"api/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func TokenVer() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")

		if token == "" {
			util.ResponseWithJson(e.ErrorAuthToken, "", c)
			c.Abort()
			return
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				util.ResponseWithJson(e.ErrorAuthCheckTokenFail, "", c)
			} else if time.Now().Unix() > claims.ExpiresAt {
				util.ResponseWithJson(e.ErrorAuthCheckTokenTimeout, "", c)
				c.Abort()
				return
			} else {
				c.Set("ID", claims.ID)
				//c.Set("Mobile",claims.Mobile)
				c.Next()
			}
		}
	}
}
