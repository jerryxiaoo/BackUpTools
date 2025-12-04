package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var configPath, _ = LoadConfig("./config/config.json")

type Config struct {
	LogFilePath string `json:"logFilePath"`
}

func NewConfig() Config {
	return Config{}

}

// 加载config
func LoadConfig(configPath string) (config Config, err error) {

	//读取json文件
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println(err, "读取json文件错误")
		return config, err
	}

	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		fmt.Println(err, "json文件解析错误")
	}
	return config, nil
}

// 获取configPath
func GetLogConfigPath() string {
	return configPath.LogFilePath
}
