package main

import (
	"backupTools/Tool"
	task "backupTools/task"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

/*

1. 从taskConfig.json 读取每个任务的源地址目录、目标地址目录
   从taskQueue.json  读取当前队列的状态，以及后续覆写任务咋黄台
2. 拷贝文件
3. 日志记录：记录日志到run.log、Error.log


*/

func main() {

	//获取TaskConfig.json数据

	//获取taskQueue对象
	//taskQueue := task.NewTaskQueue()
	//获取task对象
	//taskObj := task.NewTask()
	//制造任务队列切片
	taskList := make([]task.Task, 10)
	//读取json文件
	taskBytes, err := os.ReadFile("config/taskConfig.json")

	if err != nil {
		fmt.Println("读取taskConfig.json文件失败")
		return
	}

	//解析json
	err = json.Unmarshal(taskBytes, &taskList)
	if err != nil {
		fmt.Printf("解析taskConfig.json文件出现问题")
		fmt.Printf(err.Error())
		return
	}

	//开启循环，拿到对象进行操作
	for _, task := range taskList {
		//拿到多个地址
		for _, singleTargetPath := range task.TargetPath {
			//filepath.WalkDir ：功能是遍历指定目录，并对每个文件和子目录执行给定的函数。(func)
			err := filepath.WalkDir(task.SourcePath, Tool.NewWalkDirFunc(task.SourcePath, singleTargetPath))
			if err != nil {
				fmt.Printf(err.Error())
				return
			}
		}

	}

}
