package v1

import (
	"api/models"
	"api/pkg/e"
	"api/pkg/setting"
	"api/pkg/upload"
	"api/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// GetAppVersion 获取最新的app版本号
func GetAppVersion(c *gin.Context) {
	//客户端上传的参数
	appID := c.PostForm("app_id")     // 客户端上传的app_id参数，=1是iOS客户端，=2是Android客户端
	version := c.PostForm("version")  // 客户端上传的安装的app版本号
	appVersion := models.GetVersion() // 获取数据库中最新的版本信息

	//要返回的数据
	var responseData = gin.H{
		"needUpdate": 0,
		"apkUrl":     appVersion.AppUrl,
		"desc":       appVersion.Desc,
		"version":    appVersion.Version,
		"appSize":    appVersion.AppSize,
	}

	// 如果是iOS
	if appID == "1" {
		responseData["apkUrl"] = setting.AppSetting.AppStoreUrl // 返回应用商店地址

		a, b, c := VersionOrdinal(version), VersionOrdinal(appVersion.Version), VersionOrdinal(appVersion.IOSMinVersion)
		// 先比较是否是最新版本
		if a < b {
			responseData["needUpdate"] = 1 //不是最新版本提示可选升级
		}
		// 再比较是否是最低支持的版本号
		if a < c {
			responseData["needUpdate"] = 2 //需要强制升级
		}
	}

	// 如果是安卓
	if appID == "2" {
		a, b, c := VersionOrdinal(version), VersionOrdinal(appVersion.Version), VersionOrdinal(appVersion.AndroidMinVersion)
		// 先比较是否是最新版本
		if a < b {
			responseData["needUpdate"] = 1 // 不是最新版本提示可选升级
		}
		//再比较是否是最低支持的版本号
		if a < c {
			responseData["needUpdate"] = 2 // 需要强制升级
		}
	}

	util.ResponseWithJson(e.SUCCESS, responseData, c)
}

// GetAppVersionIndex 打开版本升级的html页面（暂时这样写，后续的话完成后台管理页面）
func GetAppVersionIndex(c *gin.Context) {
	c.HTML(e.SUCCESS, "appversion.html", gin.H{
		"title": "App版本升级",
	})
}

// CreateAppVersion 创建新的版本信息
func CreateAppVersion(c *gin.Context) {
	var appVersion models.AppVersion

	appVersion.Version = c.PostForm("version")                       // 新版本app版本号
	appVersion.IOSMinVersion = c.PostForm("ios_min_version")         // iOS最低可兼容的版本
	appVersion.AndroidMinVersion = c.PostForm("android_min_version") // 安卓最低可兼容的版本
	appVersion.Desc = c.PostForm("desc")                             // app升级文案
	appVersion.AppUrl = c.PostForm("app_url")                        // app下载地址

	// 获取上传的apk文件
	apkFile, _ := c.FormFile("apk_file")

	// 检查文件后缀
	if apkFile != nil {
		if !upload.CheckApkExt(apkFile.Filename) {
			util.ResponseWithJson(e.ERROR, "apk文件格式不正确", c)
			return
		}

		// 把上传的文件移动到指定的路径
		savePath := upload.GetApkFilePath()            // 保存的目录 upload/apks/
		dataPath := upload.GetApkDateName()            // 日期的目录 20220331/
		fullPath := upload.GetApkFullPath() + dataPath // 在项目中的目录 runtime/upload/apks/20220331/
		src := fullPath + apkFile.Filename             // 在项目中的位置 runtime/upload/apks/****.apk

		// 检查文件路径，这里面做了包括创建文件夹，检查权限等操作
		if err := upload.CheckApk(fullPath); err != nil {
			util.ResponseWithJson(e.ERROR, "apk文件有问题", c)
			return
		}

		// 使用c.SaveUploadedFile()把上传的文件移动到指定到位置
		if err := c.SaveUploadedFile(apkFile, src); err != nil {
			util.ResponseWithJson(e.ERROR, "上传apk失败", c)
			return
		}

		//设置结构体的值
		appVersion.AppSize = int(apkFile.Size)                     // 获取并设置apk文件大小
		appVersion.AppUrl = savePath + dataPath + apkFile.Filename // 数据库中保存apk文件的路径
	}

	//对参数做校验
	valid := validation.Validation{}
	valid.Required(appVersion.Version, "version").Message("版本号必须填写")
	valid.MinSize(appVersion.Desc, 1, "minVersion").Message("升级文案最少1个字符")
	valid.Required(appVersion.AppUrl, "appUrl").Message("app升级地址必须填写")
	if isOk := checkValidation(&valid, c); isOk == false { //校验不通过
		return
	}

	//数据库创建数据
	if err := appVersion.CreateAppVersion(); err != nil {
		util.ResponseWithJson(e.ERROR, "保存版本信息失败", c)
		return
	}

	//返回正确的数据
	util.ResponseWithJson(e.SUCCESS, appVersion, c)
}

// VersionOrdinal 用于比较两个字符串版本号的大小
func VersionOrdinal(version string) string {
	// ISO/IEC 14651:2011
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			panic("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)
}
