package upload

import (
	"api/pkg/file"
	"api/pkg/logging"
	"api/pkg/setting"
	"fmt"
	"os"
	"strings"
	"time"
)

// GetApkFilePath 获取apk保存路径
func GetApkFilePath() string {
	return setting.AppSetting.ApkSavePath
}

// GetApkDateName 日期文件夹如：20190730/
func GetApkDateName() string {
	t := time.Now()
	return fmt.Sprintf("%d%02d%02d/", t.Year(), t.Month(), t.Day())
}

// GetApkFullUrl 获取apk文件完整访问URL 如：http://127.0.0.1:9999/upload/apks/20190730/******.apk
func GetApkFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetApkFilePath() + GetApkDateName() + name
}

// GetApkFullPath 获取apk文件在项目中的目录  如：runtime/upload/apks/
func GetApkFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetApkFilePath()
}

// CheckApkExt 检查文件后缀，是否属于配置中允许的后缀名
func CheckApkExt(fileName string) bool {
	ext := file.GetExt(fileName)
	if strings.ToLower(ext) == strings.ToLower(setting.AppSetting.ApkAllowExt) {
		return true
	}
	return false
}

// CheckApk 检查apk文件
func CheckApk(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		logging.Warn("pkg/upload/apk.go文件CheckApk方法os.Getwd出错", err)
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src) //如果不存在则新建文件夹
	if err != nil {
		logging.Warn("pkg/upload/apk.go文件CheckApk方法file.IsNotExistMkDir出错", err)
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src) //检查文件权限
	if perm == true {
		logging.Warn("pkg/upload/apk.go文件CheckApk方法file.CheckPermission出错", err)
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
