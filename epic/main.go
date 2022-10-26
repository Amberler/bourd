package main

import (
	"epic/conf"
	"epic/utils"
	"log"
)

func main() {
	log.Println("Hello,epic 正在启动")
	conf.SetUp()
	utils.RemoteData()
}
