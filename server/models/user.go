package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model        //会自动添加id,created_at,updated_at,deleted_at四个字段，可以进入看下类型
	Mobile     string `gorm:"type:char(11);index;unique;not null;" json:"mobile,omitempty"` //手机号,加索引，唯一，不为空
	Name       string `gorm:"type:varchar(12);"`                                            //用户昵称，3-12个字符
	Desc       string `gorm:"type:varchar(100);"`
	Sex        int    `gorm:"type:tinyint(1);default:0;"`
	Age        string `gorm:"type:char(13);"` //用户年龄，存储的是时间戳字符串
	Avatar     string `gorm:"type:varchar(255);"`
	FollowNum  int    `gorm:"default:0;"`
	FansNum    int    `gorm:"default:0;"`
	State      int    `gorm:"type:tinyint(1);default:0;" json:"-"` //用户状态，比如=1账号冻结，=2不允许聊天之类的,默认=0,为json时不返回
}

// FindUserByMobile 根据手机号查找用户
func FindUserByMobile(mobile string) (*User, error) {
	var user User
	err := db.Where("mobile = ?", mobile).First(&user).Error
	return &user, err
}

// CreateUser 根据手机号创建用户
func CreateUser(mobile string) (*User, error) {
	user := User{
		Mobile: mobile,
	}
	err := db.Create(&user).Error
	return &user, err
}
