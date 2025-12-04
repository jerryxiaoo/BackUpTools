package main

import (
	"backupTools/task"
	"backupTools/tool"
	"fmt"
)

/*

1. 从taskConfig.json 读取每个任务的源地址目录、目标地址目录
   从taskQueue.json  读取当前队列的状态，以及后续覆写任务咋黄台
2. 拷贝文件
3. 日志记录：记录日志到run.log、Error.log



*/

func main() {

	// 处理配置文件
	taskList, err := tool.ReadJson("config/taskConfig.json")
	if err != nil {
		fmt.Println(err)
	}
	// 处理copy逻辑
	err = task.CopyToTargetPath(taskList)
	if err != nil {
		fmt.Println(err)
	}

}
