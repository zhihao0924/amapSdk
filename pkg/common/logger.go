package common

import (
	"fmt"
	"log"
)

// Logger 日志接口
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
}

// StatusChecker 状态检查接口
type StatusChecker interface {
	GetStatus() string
	GetInfo() string
	GetInfocode() string
}

// defaultLogger 默认日志实现
type defaultLogger struct{}

func (l *defaultLogger) Debug(format string, args ...interface{}) {
	log.Printf("[DEBUG] "+format, args...)
}

func (l *defaultLogger) Info(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func (l *defaultLogger) Warn(format string, args ...interface{}) {
	log.Printf("[WARN] "+format, args...)
}

func (l *defaultLogger) Error(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

// noOpLogger 无操作日志实现（用于生产环境）
type noOpLogger struct{}

func (l *noOpLogger) Debug(format string, args ...interface{}) {}
func (l *noOpLogger) Info(format string, args ...interface{})  {}
func (l *noOpLogger) Warn(format string, args ...interface{})  {}
func (l *noOpLogger) Error(format string, args ...interface{}) {}

// NewLogger 创建日志实例
func NewLogger(debug bool) Logger {
	if debug {
		return &defaultLogger{}
	}
	return &noOpLogger{}
}

// FormatLog 格式化日志
func FormatLog(level string, format string, args ...interface{}) string {
	return fmt.Sprintf("[%s] %s", level, fmt.Sprintf(format, args...))
}
