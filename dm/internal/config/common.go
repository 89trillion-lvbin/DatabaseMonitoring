package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// 读Json文件
func ReadConfigFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// json文件反序列化
func myUnmarshal(path string, v interface{}) {
	buf, err := ReadConfigFile(path)
	if err != nil {
		log.Printf("read file error: %s", err.Error())
		os.Exit(1)
	}
	err = json.Unmarshal(buf, v)
	if err != nil {
		log.Printf("loads json error: %s", err.Error())
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func NacosConfigUnmarshal(data string, v interface{}) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		panic(err)
	}
}
