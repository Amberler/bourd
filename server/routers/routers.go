package routers

import (
	"api/middleware"
	"api/pkg/setting"
	v1 "api/routers/v1"
	"github.com/gin-gonic/gin"
	"net/http"
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
		apiV1.POST("login", v1.Login)         // 登陆
		apiV1.POST("register", v1.Register)   // 注册
		apiV1Token := apiV1.Group("token/")   // 创建使用token中间件的路由组
		apiV1Token.Use(middleware.TokenVer()) // 使用token鉴权中间件
		{
			apiV1Token.POST("version", v1.GetAppVersion) //app版本更新
		}

		// 测试api
		apiV1.POST("a", v1.A) // 注册
		apiV1.GET("a", v1.A)  // 注册

	}
	r.LoadHTMLGlob("templates/*")
	appManage := r.Group("manager")
	{
		appManage.GET("appversion", v1.GetAppVersionIndex) //app版本升级网页文件
		appManage.POST("appversion", v1.CreateAppVersion)  //app版本升级api接口
	}

	/*
	   当访问 $HOST/upload/apks 时，将会读取到 项目/runtime/upload/apks 下的文件
	   这样就能让外部访问到图片资源了
	*/
	r.StaticFS(setting.AppSetting.ApkSavePath, http.Dir(setting.AppSetting.RuntimeRootPath+setting.AppSetting.ApkSavePath))

	return r
}
