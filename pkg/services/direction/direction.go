package direction

import (
	"context"
	"errors"
	"fmt"

	"github.com/zhihao0924/amapSdk/pkg/common"
	"github.com/zhihao0924/amapSdk/pkg/core"
	"github.com/zhihao0924/amapSdk/pkg/models"
)

// Service 路径规划服务
type Service struct {
	http   *core.HTTPClient
	logger common.Logger
}

// New 创建路径规划服务
func New(http *core.HTTPClient, logger common.Logger) *Service {
	return &Service{
		http:   http,
		logger: logger,
	}
}

// DrivingOptions 驾车路径规划选项
type DrivingOptions struct {
	Origin      string `json:"origin"`      // 起点经纬度
	Destination string `json:"destination"` // 终点经纬度
	Strategy    string `json:"strategy"`    // 驾车策略
	Waypoints   string `json:"waypoints"`   // 途经点
	Avoidpoly   string `json:"avoidpoly"`   // 避让区域
	PlateNumber string `json:"platenumber"` // 车牌号
}

// Validate 验证选项
func (o *DrivingOptions) Validate() error {
	if o.Origin == "" {
		return errors.New("origin is required")
	}
	if o.Destination == "" {
		return errors.New("destination is required")
	}
	return nil
}

// Driving 驾车路径规划
func (s *Service) Driving(ctx context.Context, opts *DrivingOptions) (*models.DrivingResponse, error) {
	if ctx == nil {
		return nil, common.ErrInvalidParamsError
	}
	if opts == nil {
		return nil, common.ErrInvalidParamsError
	}
	if err := opts.Validate(); err != nil {
		s.logger.Error("Driving options validation failed: %v", err)
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	params := map[string]string{
		"origin":      opts.Origin,
		"destination": opts.Destination,
		"strategy":    opts.Strategy,
		"waypoints":   opts.Waypoints,
		"avoidpoly":   opts.Avoidpoly,
		"platenumber": opts.PlateNumber,
	}

	var resp models.DrivingResponse
	err := s.http.Get(ctx, "/direction/driving", params, &resp)
	if err != nil {
		s.logger.Error("Driving request failed: %v", err)
		return nil, err
	}

	if len(resp.Route.Paths) > 0 {
		s.logger.Info("Driving success: origin=%s, destination=%s, distance=%s, duration=%s",
			opts.Origin, opts.Destination, resp.Route.Paths[0].Distance, resp.Route.Paths[0].Duration)
	}
	return &resp, nil
}

// DrivingByLocations 使用Location对象进行驾车路径规划
func (s *Service) DrivingByLocations(ctx context.Context, origin, destination *common.Location, strategy int) (*models.DrivingResponse, error) {
	if origin == nil || destination == nil {
		return nil, common.ErrInvalidParamsError
	}
	return s.Driving(ctx, &DrivingOptions{
		Origin:      origin.String(),
		Destination: destination.String(),
		Strategy:    fmt.Sprintf("%d", strategy),
	})
}

// WalkingOptions 步行路径规划选项
type WalkingOptions struct {
	Origin      string `json:"origin"`      // 起点经纬度
	Destination string `json:"destination"` // 终点经纬度
}

// Validate 验证选项
func (o *WalkingOptions) Validate() error {
	if o.Origin == "" {
		return errors.New("origin is required")
	}
	if o.Destination == "" {
		return errors.New("destination is required")
	}
	return nil
}

// Walking 步行路径规划
func (s *Service) Walking(ctx context.Context, opts *WalkingOptions) (*models.WalkingResponse, error) {
	if ctx == nil {
		return nil, common.ErrInvalidParamsError
	}
	if opts == nil {
		return nil, common.ErrInvalidParamsError
	}
	if err := opts.Validate(); err != nil {
		s.logger.Error("Walking options validation failed: %v", err)
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	params := map[string]string{
		"origin":      opts.Origin,
		"destination": opts.Destination,
	}

	var resp models.WalkingResponse
	err := s.http.Get(ctx, "/direction/walking", params, &resp)
	if err != nil {
		s.logger.Error("Walking request failed: %v", err)
		return nil, err
	}

	if len(resp.Route.Paths) > 0 {
		s.logger.Info("Walking success: origin=%s, destination=%s, distance=%s, duration=%s",
			opts.Origin, opts.Destination, resp.Route.Paths[0].Distance, resp.Route.Paths[0].Duration)
	}
	return &resp, nil
}

// WalkingByLocations 使用Location对象进行步行路径规划
func (s *Service) WalkingByLocations(ctx context.Context, origin, destination *common.Location) (*models.WalkingResponse, error) {
	if origin == nil || destination == nil {
		return nil, common.ErrInvalidParamsError
	}
	return s.Walking(ctx, &WalkingOptions{
		Origin:      origin.String(),
		Destination: destination.String(),
	})
}
