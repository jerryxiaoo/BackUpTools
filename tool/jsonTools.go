package tool

import (
	"backupTools/common"
	"encoding/json"
	"fmt"
	"os"
)

func ReadJson(jsonFilePath string) ([]common.Task, error) {
	//制造任务队列切片
	taskList := make([]common.Task, 10)
	//读取json文件
	taskBytes, err := os.ReadFile(jsonFilePath)

	if err != nil {
		fmt.Println("读取taskConfig.json文件失败")
		return nil, err
	}

	//解析json
	err = json.Unmarshal(taskBytes, &taskList)
	if err != nil {
		fmt.Printf("解析taskConfig.json文件出现问题")
		return nil, err
	}

	return taskList, nil

}
