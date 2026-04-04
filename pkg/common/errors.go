package common

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// 错误代码常量
const (
	ErrInvalidConfig  = "INVALID_CONFIG"
	ErrInvalidParams  = "INVALID_PARAMS"
	ErrRequestFailed  = "REQUEST_FAILED"
	ErrResponseFailed = "RESPONSE_FAILED"
	ErrAPIError       = "API_ERROR"
	ErrNetworkError   = "NETWORK_ERROR"
	ErrTimeout        = "TIMEOUT"
	ErrParseFailed    = "PARSE_FAILED"
	ErrNotFound       = "NOT_FOUND"
	ErrRateLimit      = "RATE_LIMIT"
	ErrAuthFailed     = "AUTH_FAILED"
)

// API错误映射
var errorMap = map[string]string{
	"10000": "请求正常",
	"10001": "key不正确或过期",
	"10002": "没有权限使用相应的服务",
	"10003": "访问已超出日访问量限制",
	"10004": "访问超出QPS限制",
	"10005": "IP白名单校验失败",
	"10006": "域名绑定失败",
	"10007": "签名不匹配",
	"10008": "MD5加密错误",
	"10009": "MCODE绑定异常",
	"10010": "权限不足，服务被禁用",
	"10011": "Key的绑定类型与调用接口不一致",
	"20000": "请求参数非法",
	"20001": "缺少必填参数",
	"20002": "非法协议",
	"20003": "其他未知错误",
	"20004": "访问来源IP未被授权",
	"20005": "非法域名",
}

// Error SDK错误
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error 实现error接口
func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

// NewError 创建新错误
func NewError(code, message string) *Error {
	if code == "" {
		code = ErrAPIError
	}
	if message == "" {
		message = "unknown error"
	}
	return &Error{
		Code:    code,
		Message: message,
	}
}

// NewErrorf 创建格式化错误
func NewErrorf(code, format string, args ...any) *Error {
	return NewError(code, fmt.Sprintf(format, args...))
}

// GetAPIErrorMessage 根据错误代码获取错误信息
func GetAPIErrorMessage(code string) string {
	if msg, ok := errorMap[strings.TrimSpace(code)]; ok {
		return msg
	}
	return "未知错误"
}

// NewAPIError 根据高德返回生成SDK错误
func NewAPIError(infoCode, info string) *Error {
	infoCode = strings.TrimSpace(infoCode)
	info = strings.TrimSpace(info)

	if info == "" {
		info = GetAPIErrorMessage(infoCode)
	}

	return NewError(mapAPIErrorCode(infoCode), fmt.Sprintf("amap api error [%s]: %s", infoCode, info))
}

// ValidateAPIResponse 校验高德接口业务响应状态
func ValidateAPIResponse(statusChecker StatusChecker) error {
	if statusChecker == nil {
		return nil
	}

	status := strings.TrimSpace(statusChecker.GetStatus())
	if status == "" || status == "1" {
		return nil
	}

	return NewAPIError(statusChecker.GetInfocode(), statusChecker.GetInfo())
}

// 预定义错误
var (
	ErrInvalidConfigError = errors.New("invalid configuration")
	ErrInvalidParamsError = errors.New("invalid parameters")
	ErrNetworkErrorError  = errors.New("network error")
	ErrTimeoutError       = errors.New("request timeout")
	ErrParseFailedError   = errors.New("parse failed")
	ErrNotFoundError      = errors.New("resource not found")
	ErrRateLimitError     = errors.New("rate limit exceeded")
	ErrAuthFailedError    = errors.New("authentication failed")
)

// IsAPIError 判断是否为API错误
func IsAPIError(err error) bool {
	if err == nil {
		return false
	}
	var sdkErr *Error
	return errors.As(err, &sdkErr)
}

// IsNetworkError 判断是否为网络错误
func IsNetworkError(err error) bool {
	if err == nil {
		return false
	}

	// 检查是否为预定义的网络错误
	if errors.Is(err, ErrNetworkErrorError) {
		return true
	}

	// 检查是否为网络相关的错误类型
	var netErr net.Error
	return errors.As(err, &netErr)
}

// IsTimeoutError 判断是否为超时错误
func IsTimeoutError(err error) bool {
	if err == nil {
		return false
	}

	// 检查是否为预定义的超时错误
	if errors.Is(err, ErrTimeoutError) {
		return true
	}

	// 检查是否为网络超时
	var netErr net.Error
	if errors.As(err, &netErr) {
		return netErr.Timeout()
	}

	return false
}

// IsNotFoundError 判断是否为未找到错误
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, ErrNotFoundError)
}

// IsRateLimitError 判断是否为频率限制错误
func IsRateLimitError(err error) bool {
	if err == nil {
		return false
	}
	var sdkErr *Error
	if errors.As(err, &sdkErr) {
		return sdkErr.Code == ErrRateLimit
	}
	return errors.Is(err, ErrRateLimitError)
}

// IsAuthError 判断是否为认证错误
func IsAuthError(err error) bool {
	if err == nil {
		return false
	}
	var sdkErr *Error
	if errors.As(err, &sdkErr) {
		return sdkErr.Code == ErrAuthFailed
	}
	return errors.Is(err, ErrAuthFailedError)
}

// WrapError 包装错误
func WrapError(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// UnwrapError 解包错误，获取原始错误
func UnwrapError(err error) error {
	return errors.Unwrap(err)
}

func mapAPIErrorCode(infoCode string) string {
	switch strings.TrimSpace(infoCode) {
	case "10001", "10002", "10005", "10006", "10007", "10008", "10009", "10010", "10011":
		return ErrAuthFailed
	case "10003", "10004":
		return ErrRateLimit
	case "20000", "20001", "20002":
		return ErrInvalidParams
	case "20004":
		return ErrAuthFailed
	default:
		return ErrAPIError
	}
}
