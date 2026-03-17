package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/zhihao0924/amapSdk/pkg/common"
)

// HTTPClient HTTP客户端
type HTTPClient struct {
	client      httpClient
	baseURL     string
	apiKey      string
	logger      common.Logger
	retryConfig *RetryConfig
	bufferPool  *sync.Pool
}

// httpClient HTTP客户端接口
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// realHttpClient 真实的HTTP客户端
type realHttpClient struct {
	client      *http.Client
	interceptor *InterceptorChain
	logger      common.Logger
	retryConfig *RetryConfig
}

// NewHTTPClient 创建新的HTTP客户端
func NewHTTPClient(
	client httpClient,
	baseURL string,
	apiKey string,
	logger common.Logger,
	retryConfig *RetryConfig,
) *HTTPClient {
	return &HTTPClient{
		client:      client,
		baseURL:     baseURL,
		apiKey:      apiKey,
		logger:      logger,
		retryConfig: retryConfig,
		bufferPool: &sync.Pool{
			New: func() any {
				return new(bytes.Buffer)
			},
		},
	}
}

// Get 执行GET请求
func (h *HTTPClient) Get(ctx context.Context, path string, params map[string]string, result any) error {
	return h.request(ctx, "GET", path, params, nil, result)
}

// Post 执行POST请求
func (h *HTTPClient) Post(ctx context.Context, path string, params map[string]string, body any, result any) error {
	return h.request(ctx, "POST", path, params, body, result)
}

// request 执行请求
func (h *HTTPClient) request(
	ctx context.Context,
	method string,
	path string,
	params map[string]string,
	body any,
	result any,
) error {
	// 构建完整URL
	fullURL, err := h.buildURL(path, params)
	if err != nil {
		return fmt.Errorf("构建URL失败: %w", err)
	}
	h.logger.Debug("请求URL: %s", fullURL)

	// 准备请求体
	var reqBody io.Reader
	if body != nil && (method == "POST" || method == "PUT") {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("序列化请求体失败: %w", err)
		}
		reqBody = bytes.NewReader(jsonData)
	}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, method, fullURL, reqBody)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	if body != nil && (method == "POST" || method == "PUT") {
		req.Header.Set("Content-Type", "application/json")
	}

	// 应用拦截器链
	if err := h.applyInterceptors(req); err != nil {
		return fmt.Errorf("请求拦截器执行失败: %w", err)
	}

	// 执行请求（带重试）
	resp, err := h.doWithRetry(req)
	if err != nil {
		return fmt.Errorf("请求执行失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应体失败: %w", err)
	}

	buf := h.bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer h.bufferPool.Put(buf)

	if _, err := buf.Write(bodyBytes); err != nil {
		return fmt.Errorf("写入缓冲区失败: %w", err)
	}

	// 记录响应
	bodyStr := buf.String()
	h.logger.Debug("响应状态: %d, 响应体长度: %d 字节", resp.StatusCode, len(bodyStr))

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API请求失败，状态码: %d, 响应: %s", resp.StatusCode, bodyStr)
	}

	// 检查响应体是否为空
	if len(bodyStr) == 0 {
		return fmt.Errorf("响应体为空，状态码: %d", resp.StatusCode)
	}

	// 解析响应
	if result != nil {
		if err := json.Unmarshal(buf.Bytes(), result); err != nil {
			// 输出响应体前500个字符用于调试
			preview := bodyStr
			if len(preview) > 500 {
				preview = preview[:500] + "..."
			}
			return fmt.Errorf("解析响应失败: %w, 响应体预览: %s", err, preview)
		}
	}

	return nil
}

// buildURL 构建完整URL
func (h *HTTPClient) buildURL(path string, params map[string]string) (string, error) {
	u, err := url.Parse(h.baseURL + path)
	if err != nil {
		return "", err
	}

	// 添加API Key
	values := u.Query()
	values.Set("key", h.apiKey)

	// 添加参数
	for k, v := range params {
		if v != "" {
			values.Set(k, v)
		}
	}

	u.RawQuery = values.Encode()
	return u.String(), nil
}

// doWithRetry 带重试的请求执行
func (h *HTTPClient) doWithRetry(req *http.Request) (*http.Response, error) {
	var lastErr error
	var resp *http.Response

	for attempt := 0; attempt < h.retryConfig.MaxRetries; attempt++ {
		// 执行请求
		resp, lastErr = h.client.Do(req)
		if lastErr == nil {
			// 检查响应状态码是否需要重试
			if resp.StatusCode < 500 || resp.StatusCode == http.StatusTooManyRequests {
				return resp, nil
			}
			// 5xx 错误可能需要重试，关闭响应体后继续
			resp.Body.Close()
			resp = nil
		}

		// 检查是否应该重试
		if !isRetryable(lastErr) {
			break
		}

		// 计算退避时间（使用指数退避）
		backoffTime := time.Duration(1<<uint(attempt)) * h.retryConfig.RetryDelay
		h.logger.Warn("请求失败，将在 %v 后重试 (尝试 %d/%d): %v", backoffTime, attempt+1, h.retryConfig.MaxRetries, lastErr)

		// 等待退避时间或上下文取消
		select {
		case <-time.After(backoffTime):
		case <-req.Context().Done():
			return nil, req.Context().Err()
		}
	}

	if lastErr != nil {
		return nil, lastErr
	}

	return resp, nil
}

// applyInterceptors 应用请求拦截器链
func (h *HTTPClient) applyInterceptors(req *http.Request) error {
	// 从真实客户端获取拦截器链
	if realClient, ok := h.client.(*realHttpClient); ok {
		return realClient.interceptor.applyRequest(req)
	}
	return nil
}

// isRetryable 判断错误是否可重试
func isRetryable(err error) bool {
	if err == nil {
		return false
	}

	// 检查是否为网络错误或超时错误
	return isNetworkError(err) || isTimeoutError(err)
}

// isNetworkError 判断是否为网络错误
func isNetworkError(err error) bool {
	return common.IsNetworkError(err)
}

// isTimeoutError 判断是否为超时错误
func isTimeoutError(err error) bool {
	return common.IsTimeoutError(err)
}

// newRealHttpClient 创建真实的HTTP客户端
func newRealHttpClient(client *http.Client, interceptor *InterceptorChain, logger common.Logger) httpClient {
	return &realHttpClient{
		client:      client,
		interceptor: interceptor,
		logger:      logger,
		retryConfig: DefaultRetryConfig,
	}
}

// Do 实现httpClient接口
func (r *realHttpClient) Do(req *http.Request) (*http.Response, error) {
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	// 应用响应拦截器
	if err := r.interceptor.applyResponse(resp); err != nil {
		return resp, err
	}

	return resp, nil
}
