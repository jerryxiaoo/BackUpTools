package main

import (
	"backupTools/task"
	"backupTools/tool"
	"bufio"
	"fmt"
	"os"
)

/*

1. 从taskConfig.json 读取每个任务的源地址目录、目标地址目录
   从taskQueue.json  读取当前队列的状态，以及后续覆写任务咋黄台
2. 拷贝文件
3. 日志记录：记录日志到run.log、Error.log



*/

func main() {

	// 处理taskConfig配置文件
	taskList, err := tool.LoadTaskConfig("config/taskConfigTest.json")
	if err != nil {
		fmt.Println(err)
	}

	// 处理copy逻辑
	err = task.CopyToTargetPath(&taskList)
	if err != nil {
		fmt.Println(err)
	}
	quitTool()

}

func quitTool() {
	fmt.Println("程序执行完成！按任意键退出...")
	// 关键：关闭输入缓冲，实现“按任意键立即响应”
	reader := bufio.NewReader(os.Stdin)
	reader.ReadByte() // 读取单个字节（任意按键都会触发）
}
