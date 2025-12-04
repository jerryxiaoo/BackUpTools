package log

import (
	"fmt"
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
