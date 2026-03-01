package amap

import (
	"github.com/zhihao0924/amapSdk/pkg/common"
)

// ==================== 错误相关函数 ====================

// GetAPIErrorMessage 根据错误代码获取错误信息
func GetAPIErrorMessage(code string) string {
	return common.GetAPIErrorMessage(code)
}

// IsAPIError 判断是否为API错误
func IsAPIError(err error) bool {
	return common.IsAPIError(err)
}

// IsNetworkError 判断是否为网络错误
func IsNetworkError(err error) bool {
	return common.IsNetworkError(err)
}

// IsTimeoutError 判断是否为超时错误
func IsTimeoutError(err error) bool {
	return common.IsTimeoutError(err)
}

// IsNotFoundError 判断是否为未找到错误
func IsNotFoundError(err error) bool {
	return common.IsNotFoundError(err)
}

// IsRateLimitError 判断是否为频率限制错误
func IsRateLimitError(err error) bool {
	return common.IsRateLimitError(err)
}

// IsAuthError 判断是否为认证错误
func IsAuthError(err error) bool {
	return common.IsAuthError(err)
}
