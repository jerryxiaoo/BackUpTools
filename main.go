package main

import (
	"backupTools/config"
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

	//读总配置文件,写入到MyConfig,从而拿到TaskConfig、LogFilePath
	err := config.LoadConfig("config/config.json")
	if err != nil {
		fmt.Println(err)
	}

	// 处理taskConfig配置文件,拿到任务对象切片
	fmt.Println(config.MyConfig)
	taskList, err := tool.LoadTaskConfig(config.MyConfig.TaskConfig)
	if err != nil {
		fmt.Println(err)
	}

	// 开始处理任务对象
	err = task.CopyToTargetPath(&taskList)
	if err != nil {
		fmt.Println(err)
	}
	// 程序结束
	tool.QuitTool()

}
