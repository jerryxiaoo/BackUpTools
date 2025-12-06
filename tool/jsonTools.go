package tool

import (
	"backupTools/common"
	"backupTools/config"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// 解析taskjson,返回任务队列切片
func LoadTaskConfig(jsonFilePath string) ([]common.Task, error) {
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

// 写回taskjson，刷新任务状态
func FlushJsonFile(taskList []common.Task) error {
	marshalIndent, err := json.MarshalIndent(taskList, "", "\t")
	if err != nil {
		return err
	}

	file, err := os.OpenFile(config.MyConfig.TaskConfig, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(marshalIndent))
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
