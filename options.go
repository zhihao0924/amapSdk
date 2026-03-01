package amap

import (
	"fmt"
	"net/http"

	"github.com/zhihao0924/amapSdk/pkg/common"
	"github.com/zhihao0924/amapSdk/pkg/core"
)

// ==================== 地理编码构建器 ====================

// GeocodeOptionsBuilder 地理编码选项构建器
type GeocodeOptionsBuilder struct {
	options *GeocodeOptions
}

// NewGeocodeOptionsBuilder 创建地理编码选项构建器
func NewGeocodeOptionsBuilder() *GeocodeOptionsBuilder {
	return &GeocodeOptionsBuilder{
		options: &GeocodeOptions{},
	}
}

// SetAddress 设置地址
func (b *GeocodeOptionsBuilder) SetAddress(address string) *GeocodeOptionsBuilder {
	b.options.Address = address
	return b
}

// SetCity 设置城市
func (b *GeocodeOptionsBuilder) SetCity(city string) *GeocodeOptionsBuilder {
	b.options.City = city
	return b
}

// Build 构建选项
func (b *GeocodeOptionsBuilder) Build() *GeocodeOptions {
	return b.options
}

// ReGeoOptionsBuilder 逆地理编码选项构建器
type ReGeoOptionsBuilder struct {
	options *ReGeoOptions
}

// NewReGeoOptionsBuilder 创建逆地理编码选项构建器
func NewReGeoOptionsBuilder() *ReGeoOptionsBuilder {
	return &ReGeoOptionsBuilder{
		options: &ReGeoOptions{},
	}
}

// SetLocation 设置位置
func (b *ReGeoOptionsBuilder) SetLocation(lng, lat float64) *ReGeoOptionsBuilder {
	b.options.Location = fmt.Sprintf("%.6f,%.6f", lng, lat)
	return b
}

// SetRadius 设置搜索半径
func (b *ReGeoOptionsBuilder) SetRadius(radius string) *ReGeoOptionsBuilder {
	b.options.Radius = radius
	return b
}

// SetExtensions 设置扩展信息
func (b *ReGeoOptionsBuilder) SetExtensions(extensions string) *ReGeoOptionsBuilder {
	b.options.Extensions = extensions
	return b
}

// Build 构建选项
func (b *ReGeoOptionsBuilder) Build() *ReGeoOptions {
	return b.options
}

// ==================== 路径规划构建器 ====================

// DrivingOptionsBuilder 驾车路径规划选项构建器
type DrivingOptionsBuilder struct {
	options *DrivingOptions
}

// NewDrivingOptionsBuilder 创建驾车路径规划选项构建器
func NewDrivingOptionsBuilder() *DrivingOptionsBuilder {
	return &DrivingOptionsBuilder{
		options: &DrivingOptions{},
	}
}

// SetOrigin 设置起点
func (b *DrivingOptionsBuilder) SetOrigin(lng, lat float64) *DrivingOptionsBuilder {
	b.options.Origin = fmt.Sprintf("%.6f,%.6f", lng, lat)
	return b
}

// SetDestination 设置终点
func (b *DrivingOptionsBuilder) SetDestination(lng, lat float64) *DrivingOptionsBuilder {
	b.options.Destination = fmt.Sprintf("%.6f,%.6f", lng, lat)
	return b
}

// SetStrategy 设置策略
func (b *DrivingOptionsBuilder) SetStrategy(strategy int) *DrivingOptionsBuilder {
	b.options.Strategy = fmt.Sprintf("%d", strategy)
	return b
}

// AddWaypoint 添加途经点
func (b *DrivingOptionsBuilder) AddWaypoint(lng, lat float64) *DrivingOptionsBuilder {
	if b.options.Waypoints == "" {
		b.options.Waypoints = fmt.Sprintf("%.6f,%.6f", lng, lat)
	} else {
		b.options.Waypoints += "|" + fmt.Sprintf("%.6f,%.6f", lng, lat)
	}
	return b
}

// Build 构建选项
func (b *DrivingOptionsBuilder) Build() *DrivingOptions {
	return b.options
}

// ==================== 工具函数 ====================

// NewLocation 创建位置
func NewLocation(lng, lat float64) *Location {
	return common.NewLocation(lng, lat)
}

// NewBox 创建边界框
func NewBox(minLng, minLat, maxLng, maxLat float64) *Box {
	return common.NewBox(minLng, minLat, maxLng, maxLat)
}

// NewError 创建错误
func NewError(code, message string) *Error {
	return common.NewError(code, message)
}

// NewLogger 创建日志实例
func NewLogger(debug bool) Logger {
	return common.NewLogger(debug)
}

// NewInterceptorChain 创建拦截器链
func NewInterceptorChain() *InterceptorChain {
	return core.NewInterceptorChain()
}

// LoggingRequestInterceptor 日志请求拦截器
func LoggingRequestInterceptor(logger Logger) func(req *http.Request) error {
	return core.LoggingRequestInterceptor(logger)
}

// LoggingResponseInterceptor 日志响应拦截器
func LoggingResponseInterceptor(logger Logger) func(resp *http.Response) error {
	return core.LoggingResponseInterceptor(logger)
}

// HeaderInterceptor 请求头拦截器
func HeaderInterceptor(headers map[string]string) func(req *http.Request) error {
	return core.HeaderInterceptor(headers)
}
