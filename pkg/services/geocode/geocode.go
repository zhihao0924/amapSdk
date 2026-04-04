package geocode

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/zhihao0924/amapSdk/pkg/common"
	"github.com/zhihao0924/amapSdk/pkg/core"
	"github.com/zhihao0924/amapSdk/pkg/models"
)

const (
	// defaultRadius 默认搜索半径（米）
	defaultRadius = "1000"
	// baseExtensions 基础扩展信息
	baseExtensions = "base"
	// allExtensions 详细扩展信息
	allExtensions = "all"
)

// Service 地理编码服务
type Service struct {
	http   *core.HTTPClient
	logger common.Logger
}

// New 创建地理编码服务
func New(http *core.HTTPClient, logger common.Logger) *Service {
	if http == nil || logger == nil {
		return nil
	}
	return &Service{
		http:   http,
		logger: logger,
	}
}

// Options 地理编码选项
type Options struct {
	Address string `json:"address"` // 待解析的结构化地址信息
	City    string `json:"city"`    // 指定查询的城市
}

// Validate 验证选项
func (o *Options) Validate() error {
	o.Address = strings.TrimSpace(o.Address)
	if o.Address == "" {
		return errors.New("address is required")
	}
	if len(o.Address) > 1000 {
		return errors.New("address too long, maximum 1000 characters")
	}
	return nil
}

// buildParams 构建请求参数
func (o *Options) buildParams() map[string]string {
	params := make(map[string]string)
	if o.Address != "" {
		params["address"] = o.Address
	}
	if o.City != "" {
		params["city"] = o.City
	}
	return params
}

// Geo 地理编码 - 将地址转换为经纬度
func (s *Service) Geo(ctx context.Context, opts *Options) (*models.GeocodeResponse, error) {
	if ctx == nil {
		return nil, common.ErrInvalidParamsError
	}
	if opts == nil {
		return nil, common.ErrInvalidParamsError
	}

	if err := opts.Validate(); err != nil {
		s.logger.Error("Geocode options validation failed: %v", err)
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	var resp models.GeocodeResponse
	err := s.http.Get(ctx, "/geocode/geo", opts.buildParams(), &resp)
	if err != nil {
		s.logger.Error("Geocode request failed: address=%s, error=%v", opts.Address, err)
		return nil, err
	}

	s.logger.Info("Geocode success: address=%s, count=%d", opts.Address, len(resp.Geocodes))
	return &resp, nil
}

// ReGeoOptions 逆地理编码选项
type ReGeoOptions struct {
	Location   string `json:"location"`   // 经纬度坐标，格式为：经度,纬度
	Radius     string `json:"radius"`     // 搜索半径
	Extensions string `json:"extensions"` // 返回结果控制: base-基本信息, all-详细信息
	Batch      string `json:"batch"`      // 是否支持批量查询
	RoadLevel  string `json:"roadlevel"`  // 道路等级
}

// Validate 验证选项
func (o *ReGeoOptions) Validate() error {
	o.Location = strings.TrimSpace(o.Location)
	if o.Location == "" {
		return errors.New("location is required")
	}

	// 验证坐标格式
	if _, err := common.ParseLocation(o.Location); err != nil {
		return fmt.Errorf("invalid location format: %w", err)
	}

	return nil
}

// buildParams 构建请求参数
func (o *ReGeoOptions) buildParams() map[string]string {
	params := make(map[string]string)
	params["location"] = o.Location

	if o.Radius != "" {
		params["radius"] = o.Radius
	}
	if o.Extensions != "" {
		params["extensions"] = o.Extensions
	}
	if o.Batch != "" {
		params["batch"] = o.Batch
	}
	if o.RoadLevel != "" {
		params["roadlevel"] = o.RoadLevel
	}

	return params
}

// ReGeo 逆地理编码 - 将经纬度转换为地址
func (s *Service) ReGeo(ctx context.Context, opts *ReGeoOptions) (*models.ReGeocodeResponse, error) {
	if ctx == nil {
		return nil, common.ErrInvalidParamsError
	}
	if opts == nil {
		return nil, common.ErrInvalidParamsError
	}

	if err := opts.Validate(); err != nil {
		s.logger.Error("ReGeocode options validation failed: %v", err)
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	var resp models.ReGeocodeResponse
	err := s.http.Get(ctx, "/geocode/regeo", opts.buildParams(), &resp)
	if err != nil {
		s.logger.Error("ReGeocode request failed: location=%s, error=%v", opts.Location, err)
		return nil, err
	}

	if resp.Regeocode != nil {
		s.logger.Info("ReGeocode success: location=%s, address=%s",
			opts.Location, resp.Regeocode.FormattedAddress)
	}
	return &resp, nil
}

// ReGeoByLocation 使用Location对象进行逆地理编码
func (s *Service) ReGeoByLocation(ctx context.Context, loc *common.Location, extensions string) (*models.ReGeocodeResponse, error) {
	if loc == nil {
		return nil, common.ErrInvalidParamsError
	}

	return s.ReGeo(ctx, &ReGeoOptions{
		Location:   loc.String(),
		Extensions: extensions,
	})
}
