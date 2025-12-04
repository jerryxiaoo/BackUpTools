package tool

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type log struct {
	logLevel logLevel
}

func NewLog(level logLevel) *log {
	return &log{
		logLevel: level,
	}
}

type logLevel uint8

const (
	UNKNOW logLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
)

// 接受信息和可能的对象
func Debug(msg string, param ...interface{}) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("【%s】%s\n", nowTime, msg)

}

func Info(msg string, param ...interface{}) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("【%s】%s\n", nowTime, msg)
}

func Warning(msg string, param ...interface{}) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("【%s】%s\n", nowTime, msg)
}
func Error(msg string, param ...interface{}) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("【%s】%s\n", nowTime, msg)
}

// 写进日志文件
func WriteLogFile(msg string, configPath string) {
	logFile, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer logFile.Close()
	writer := bufio.NewWriter(logFile)
	formatTime := time.Now().Format("2006-01-02 15:04:05")
	msg = fmt.Sprintf("[%s]%s\n", formatTime, msg)
	_, err = writer.WriteString(msg)
	if err != nil {
		fmt.Println(err)
	}

	writer.Flush()
}
