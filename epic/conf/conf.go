package conf

import (
	"github.com/go-ini/ini"
	"log"
)

type App struct {
	RemoteURL string
}

var AppConf = &App{}

func SetUp() {

	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("获取.ini配置失败")
	}

	err = Cfg.Section("app").MapTo(AppConf)
	if err != nil {
		log.Fatalf("Cfg配置文件映射 AppConf 错误: %v", err)
	}

	log.Println("本地配置读取成功")
}
