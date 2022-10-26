package utils

import (
	"epic/conf"
	"fmt"
	"io"
	"log"
	"net/http"
)

func RemoteData() {
	resp, err := http.Get(conf.AppConf.RemoteURL)
	if err != nil {
		log.Fatal("远程数据请求失败")
	}
	
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("远程数据请求失败")
	}
	fmt.Println(string(body))
	//fmt.Println(resp.StatusCode)
	//if resp.StatusCode == 200 {
	//	fmt.Println("ok")
	//}
}
