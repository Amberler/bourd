package routers

import (
	"api/middleware"
	"api/pkg/setting"
	v1 "api/routers/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 创建gin框架路由实例
	r := gin.New()
	// 使用gin框架中的打印中间件
	r.Use(gin.Logger())
	// 使用gin框架中的恢复中间件，可以从任何恐慌中恢复，如果有，则写入500
	r.Use(gin.Recovery())

	// 设置运行模式，debug或release,如果放在gin.New或者gin.Default之后，还是会打印一些信息的。放之前则不会
	gin.SetMode(setting.ServerSetting.RunMode)

	// 路由分组，apiV1代表v1版本的路由组
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("version", v1.GetAppVersion) // app版本升级
		apiV1.POST("login", v1.Login)          // 登陆
		apiV1.POST("register", v1.Register)    // 注册
		apiV1Token := apiV1.Group("token/")    // 创建使用token中间件的路由组
		apiV1Token.Use(middleware.TokenVer())  // 使用token鉴权中间件
		{

		}
	}
	return r
}
