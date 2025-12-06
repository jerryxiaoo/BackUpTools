package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var MyConfig Config = Config{}

type Config struct {
	LogFilePath string `json:"logFilePath"`
	TaskConfig  string `json:"taskConfig"`
}

func NewConfig() Config {
	return Config{}

}

// 加载config,返回MyConfig对象
func LoadConfig(configPath string) (err error) {

	//读取json文件
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println(err, "读取json文件错误")
		return err
	}

	err = json.Unmarshal(configBytes, &MyConfig)
	if err != nil {
		fmt.Println(err, "json文件解析错误")
		return err
	}

	return nil
}
