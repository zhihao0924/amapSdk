package core

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/zhihao0924/amapSdk/pkg/common"
)

const (
	// maxBodySize 最大请求体大小（10MB）
	maxBodySize = 10 * 1024 * 1024
)

// InterceptorFunc 请求拦截器函数类型
type InterceptorFunc func(req *http.Request) error

// ResponseInterceptorFunc 响应拦截器函数类型
type ResponseInterceptorFunc func(resp *http.Response) error

// InterceptorChain 拦截器链
type InterceptorChain struct {
	requestInterceptors  []InterceptorFunc
	responseInterceptors []ResponseInterceptorFunc
}

// NewInterceptorChain 创建拦截器链
func NewInterceptorChain() *InterceptorChain {
	return &InterceptorChain{
		requestInterceptors:  make([]InterceptorFunc, 0),
		responseInterceptors: make([]ResponseInterceptorFunc, 0),
	}
}

// AddRequest 添加请求拦截器
func (c *InterceptorChain) AddRequest(interceptor InterceptorFunc) {
	c.requestInterceptors = append(c.requestInterceptors, interceptor)
}

// AddResponse 添加响应拦截器
func (c *InterceptorChain) AddResponse(interceptor ResponseInterceptorFunc) {
	c.responseInterceptors = append(c.responseInterceptors, interceptor)
}

// ApplyRequest 应用所有请求拦截器
func (c *InterceptorChain) ApplyRequest(req *http.Request) error {
	for _, interceptor := range c.requestInterceptors {
		if interceptor == nil {
			continue
		}
		if err := interceptor(req); err != nil {
			return err
		}
	}
	return nil
}

// applyRequest 应用所有请求拦截器（内部方法）
func (c *InterceptorChain) applyRequest(req *http.Request) error {
	return c.ApplyRequest(req)
}

// ApplyResponse 应用所有响应拦截器
func (c *InterceptorChain) ApplyResponse(resp *http.Response) error {
	for _, interceptor := range c.responseInterceptors {
		if interceptor == nil {
			continue
		}
		if err := interceptor(resp); err != nil {
			return err
		}
	}
	return nil
}

// applyResponse 应用所有响应拦截器（内部方法）
func (c *InterceptorChain) applyResponse(resp *http.Response) error {
	return c.ApplyResponse(resp)
}

// LoggingRequestInterceptor 日志请求拦截器
func LoggingRequestInterceptor(logger common.Logger) InterceptorFunc {
	return func(req *http.Request) error {
		if req == nil {
			return nil
		}

		var body string
		if req.Body != nil {
			buf := make([]byte, 0, 1024)
			n, err := req.Body.Read(buf[:cap(buf)])
			if err != nil && err != io.EOF {
				return err
			}
			buf = buf[:n]
			body = string(buf)
			req.Body = io.NopCloser(bytes.NewReader(buf))
		}

		logger.Info("请求: %s %s, Header: %v, Body: %s", req.Method, req.URL, req.Header, body)
		return nil
	}
}

// LoggingResponseInterceptor 日志响应拦截器
func LoggingResponseInterceptor(logger common.Logger) ResponseInterceptorFunc {
	return func(resp *http.Response) error {
		if resp == nil {
			return nil
		}

		var body string
		if resp.Body != nil {
			// 使用 io.ReadAll 确保读取完整的响应体
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			body = string(bodyBytes)
			// 重建响应体供后续使用
			resp.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}

		logger.Info("响应: %d, Header: %v, Body: %s", resp.StatusCode, resp.Header, body)
		return nil
	}
}

// HeaderInterceptor 请求头拦截器
func HeaderInterceptor(headers map[string]string) InterceptorFunc {
	return func(req *http.Request) error {
		if req == nil {
			return nil
		}

		if headers == nil {
			return nil
		}

		for k, v := range headers {
			if v != "" {
				req.Header.Set(k, v)
			}
		}

		return nil
	}
}

// LoggingInterceptor 简化的日志拦截器（兼容旧版本）
func LoggingInterceptor() InterceptorFunc {
	logger := common.NewLogger(true)
	return LoggingRequestInterceptor(logger)
}

// DebugInterceptor 调试拦截器（标准库log）
func DebugInterceptor() InterceptorFunc {
	return func(req *http.Request) error {
		if req == nil {
			return nil
		}
		log.Printf("[DEBUG] Request: %s %s\n", req.Method, req.URL)
		return nil
	}
}
