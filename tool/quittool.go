package tool

import (
	"bufio"
	"fmt"
	"os"
)

func QuitTool() {
	fmt.Println("程序执行完成！按任意键退出...")
	// 关键：关闭输入缓冲，实现“按任意键立即响应”
	reader := bufio.NewReader(os.Stdin)
	reader.ReadByte() // 读取单个字节（任意按键都会触发）
}
