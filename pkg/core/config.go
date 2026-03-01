package core

import (
	"time"

	"github.com/zhihao0924/amapSdk/pkg/common"
)

const (
	// DefaultBaseURL 默认API基础URL
	DefaultBaseURL = "https://restapi.amap.com/v3"
	// DefaultTimeout 默认请求超时时间（秒）
	DefaultTimeout = 10
	// MaxTimeout 最大超时时间（秒）
	MaxTimeout = 300 // 5分钟
)

// Config SDK配置
type Config struct {
	// Key 高德地图API Key (必需)
	Key string
	// BaseURL API基础URL，默认为官方地址
	BaseURL string
	// Timeout 请求超时时间（秒），默认10秒
	Timeout int
	// Debug 是否开启调试日志
	Debug bool
	// Headers 自定义请求头
	Headers map[string]string
	// RetryConfig 重试配置
	RetryConfig *RetryConfig
	// InterceptorChain 拦截器链
	InterceptorChain *InterceptorChain
}

// Normalize 标准化配置，填充默认值并验证
func (c *Config) Normalize() {
	// 设置默认值
	if c.BaseURL == "" {
		c.BaseURL = DefaultBaseURL
	}
	if c.Timeout == 0 {
		c.Timeout = DefaultTimeout
	}
	if c.RetryConfig == nil {
		c.RetryConfig = DefaultRetryConfig
	}

	// 验证配置
	if c.Timeout < 1 {
		c.Timeout = DefaultTimeout
	}
	if c.Timeout > MaxTimeout {
		c.Timeout = MaxTimeout
	}
}

// GetTimeout 获取超时时间
func (c *Config) GetTimeout() time.Duration {
	return time.Duration(c.Timeout) * time.Second
}

// GetLogger 获取日志实例
func (c *Config) GetLogger() common.Logger {
	return common.NewLogger(c.Debug)
}

// GetInterceptorChain 获取拦截器链
func (c *Config) GetInterceptorChain() *InterceptorChain {
	if c.InterceptorChain == nil {
		c.InterceptorChain = NewInterceptorChain()
	}
	return c.InterceptorChain
}

// Validate 验证配置
func (c *Config) Validate() error {
	if c.Key == "" {
		return common.ErrInvalidConfigError
	}
	return nil
}
