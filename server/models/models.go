package models

import (
	"api/pkg/logging"
	"api/pkg/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetUp() {
	var (
		err          error
		databaseType = setting.DatabaseSetting.Type
		user         = setting.DatabaseSetting.User
		pass         = setting.DatabaseSetting.Password
		host         = setting.DatabaseSetting.Host
		name         = setting.DatabaseSetting.Name
	)
	// 使用gorm链接数据库
	db, err = gorm.Open(databaseType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", user, pass, host, name))
	if err != nil {
		logging.Fatal("数据库连接失败", err)
	}

	// 设置表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	// 设置禁用表名的复数形式
	db.SingularTable(true)

	// 打印日志，本地调试的时候可以看执行语句
	db.LogMode(true)

	////自动检查 Tag 结构是否变化，变化则进行迁移，需要的参数为数据库模型结构体
	db.AutoMigrate(&User{})

	// 设置空闲时候的最大链接数
	db.DB().SetMaxIdleConns(10)

	// 设置数据库的最大打开链接数
	db.DB().SetMaxOpenConns(100)
}
