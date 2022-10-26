package utils

import (
	"encoding/json"
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

	body, e := io.ReadAll(resp.Body)
	if e != nil {
		log.Fatal("远程数据请求失败")
	}

	m := make(map[string]interface{})

	error := json.Unmarshal(body, &m)

	if error != nil {
		panic(error)
	}

	fmt.Printf("%+v\n", m)

}
