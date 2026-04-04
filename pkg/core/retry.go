package core

import (
	"errors"
	"net"
	"syscall"
	"time"
)

// RetryConfig 重试配置
type RetryConfig struct {
	// MaxRetries 最大重试次数
	MaxRetries int
	// RetryDelay 重试延迟
	RetryDelay time.Duration
	// Retryable 判断是否可重试的函数
	Retryable func(error) bool
}

// DefaultRetryConfig 默认重试配置
var DefaultRetryConfig = &RetryConfig{
	MaxRetries: 3,
	RetryDelay: 1 * time.Second,
	Retryable:  DefaultRetryable,
}

// Clone 拷贝重试配置，避免多个客户端共享同一个可变实例
func (r *RetryConfig) Clone() *RetryConfig {
	if r == nil {
		return nil
	}

	cloned := *r
	if cloned.Retryable == nil {
		cloned.Retryable = DefaultRetryable
	}
	return &cloned
}

// DefaultRetryable 默认的可重试判断函数
func DefaultRetryable(err error) bool {
	if err == nil {
		return false
	}

	// 检查是否为网络错误
	var netErr net.Error
	if errors.As(err, &netErr) {
		return true
	}

	// 检查是否为超时错误
	if netErr, ok := err.(net.Error); ok {
		if netErr.Timeout() {
			return true
		}
	}

	// 检查是否为临时错误
	if errors.Is(err, syscall.ECONNRESET) {
		return true
	}

	if errors.Is(err, syscall.ECONNABORTED) {
		return true
	}

	return false
}

// Validate 验证重试配置
func (r *RetryConfig) Validate() error {
	if r.MaxRetries < 0 {
		return errors.New("MaxRetries 不能为负数")
	}

	if r.MaxRetries > 10 {
		return errors.New("MaxRetries 不能超过10")
	}

	if r.RetryDelay < 0 {
		return errors.New("RetryDelay 不能为负数")
	}

	return nil
}

// NewRetryConfig 创建新的重试配置
func NewRetryConfig(maxRetries int, delay time.Duration) *RetryConfig {
	return &RetryConfig{
		MaxRetries: maxRetries,
		RetryDelay: delay,
		Retryable:  DefaultRetryable,
	}
}

// WithMaxRetries 设置最大重试次数
func (r *RetryConfig) WithMaxRetries(maxRetries int) *RetryConfig {
	r.MaxRetries = maxRetries
	return r
}

// WithRetryDelay 设置重试延迟
func (r *RetryConfig) WithRetryDelay(delay time.Duration) *RetryConfig {
	r.RetryDelay = delay
	return r
}

// WithRetryable 设置自定义重试判断函数
func (r *RetryConfig) WithRetryable(fn func(error) bool) *RetryConfig {
	r.Retryable = fn
	return r
}
