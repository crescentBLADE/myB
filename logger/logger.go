// sdk/logger.go
package logger

import "log"

type Logger interface {
	Info(msg string)
	Error(msg string)
}

var logger Logger

// 设置 SDK 中使用的日志
func SetLogger(l Logger) {
	logger = l
}

// 使用日志记录 SDK 中的操作
func LogInfo(message string) {
	if logger != nil {
		logger.Info(message)
	} else {
		log.Println(message) // 默认使用标准库日志
	}
}
func LogError(message string) {
	if logger != nil {
		logger.Error(message)
	} else {
		log.Println(message) // 默认使用标准库日志
	}
}
