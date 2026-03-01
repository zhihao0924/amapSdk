package core

import (
	"errors"
	"net/http"
	"sync"

	"github.com/zhihao0924/amapSdk/pkg/common"
)

// Client 核心客户端
type Client struct {
	config      *Config
	http        *HTTPClient
	logger      common.Logger
	interceptor *InterceptorChain

	// 服务缓存（单例模式）
	geocodeService   interface{}
	directionService interface{}
	placeService     interface{}
	weatherService   interface{}
	ipService        interface{}

	mu sync.RWMutex
}

// NewClient 创建新的核心客户端
func NewClient(config *Config) (*Client, error) {
	if config == nil || config.Key == "" {
		return nil, errors.New("config or config.Key is required")
	}

	// 标准化配置
	config.Normalize()

	// 创建日志实例
	logger := config.GetLogger()

	// 创建拦截器链
	interceptor := config.GetInterceptorChain()

	// 添加默认拦截器
	if config.Debug {
		interceptor.AddRequest(LoggingRequestInterceptor(logger))
		interceptor.AddResponse(LoggingResponseInterceptor(logger))
	}

	// 添加自定义请求头
	if len(config.Headers) > 0 {
		interceptor.AddRequest(HeaderInterceptor(config.Headers))
	}

	// 创建HTTP客户端
	httpClient := &http.Client{
		Timeout: config.GetTimeout(),
	}

	// 创建HTTP服务
	httpService := NewHTTPClient(
		newRealHttpClient(httpClient, interceptor, logger),
		config.BaseURL,
		config.Key,
		logger,
		config.RetryConfig,
	)

	client := &Client{
		config:      config,
		http:        httpService,
		logger:      logger,
		interceptor: interceptor,
	}

	logger.Info("AMap SDK client initialized with key: %s", maskKey(config.Key))

	return client, nil
}

// maskKey 掩码API Key
func maskKey(key string) string {
	if len(key) <= 8 {
		return "****"
	}
	return key[:4] + "****" + key[len(key)-4:]
}

// GetConfig 获取配置
func (c *Client) GetConfig() *Config {
	return c.config
}

// GetLogger 获取日志实例
func (c *Client) GetLogger() common.Logger {
	return c.logger
}

// GetHTTP 获取HTTP客户端
func (c *Client) GetHTTP() *HTTPClient {
	return c.http
}

// Close 关闭客户端，释放资源
func (c *Client) Close() error {
	c.logger.Info("AMap SDK client closed")
	return nil
}
