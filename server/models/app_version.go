package models

import (
	"api/pkg/setting"
	"github.com/jinzhu/gorm"
)

type AppVersion struct {
	gorm.Model
	Version           string `gorm:"type:varchar(10);not null"` //不为空
	IOSMinVersion     string `gorm:"type:varchar(10)"`
	AndroidMinVersion string `gorm:"type:varchar(10)"`
	Desc              string `gorm:"type:varchar(255)"`
	AppUrl            string `gorm:"type:varchar(255)"`
	AppSize           int
}

func (appVersion *AppVersion) CreateAppVersion() error {
	err := db.Create(appVersion).Error
	return err
}

func GetVersion() *AppVersion {
	var appVersion AppVersion
	db.Last(&appVersion)
	return &appVersion
}

// AfterFind 数据库查询钩子，在数据库查询之后执行的方法
func (appVersion *AppVersion) AfterFind() {
	appVersion.AppUrl = setting.AppSetting.ImagePrefixUrl + "/" + appVersion.AppUrl //拼接完整的apk的url地址
}
