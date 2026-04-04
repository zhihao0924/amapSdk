package weather

import (
	"context"
	"errors"
	"fmt"

	"github.com/zhihao0924/amapSdk/pkg/common"
	"github.com/zhihao0924/amapSdk/pkg/core"
	"github.com/zhihao0924/amapSdk/pkg/models"
)

// Service 天气服务
type Service struct {
	http   *core.HTTPClient
	logger common.Logger
}

// New 创建天气服务
func New(http *core.HTTPClient, logger common.Logger) *Service {
	return &Service{
		http:   http,
		logger: logger,
	}
}

// Options 天气查询选项
type Options struct {
	City       string `json:"city"`       // 城市编码或名称
	Extensions string `json:"extensions"` // 返回结果控制: base-实况, all-预报
}

// Validate 验证选项
func (o *Options) Validate() error {
	if o.City == "" {
		return errors.New("city is required")
	}
	return nil
}

// Query 查询天气
func (s *Service) Query(ctx context.Context, opts *Options) (*models.WeatherResponse, error) {
	if ctx == nil {
		return nil, common.ErrInvalidParamsError
	}
	if opts == nil {
		return nil, common.ErrInvalidParamsError
	}
	if err := opts.Validate(); err != nil {
		s.logger.Error("Weather options validation failed: %v", err)
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	extensions := opts.Extensions
	if extensions == "" {
		extensions = "base"
	}

	params := map[string]string{
		"city":       opts.City,
		"extensions": extensions,
	}

	var resp models.WeatherResponse
	err := s.http.Get(ctx, "/weather/weatherInfo", params, &resp)
	if err != nil {
		s.logger.Error("Weather request failed: %v", err)
		return nil, err
	}

	s.logger.Info("Weather query success: city=%s, type=%s", opts.City, extensions)
	return &resp, nil
}

// Base 查询实况天气
func (s *Service) Base(ctx context.Context, city string) (*models.WeatherLive, error) {
	opts := &Options{
		City:       city,
		Extensions: "base",
	}

	resp, err := s.Query(ctx, opts)
	if err != nil {
		return nil, err
	}

	if len(resp.Lives) == 0 {
		return nil, errors.New("no weather data found")
	}

	s.logger.Info("Weather base success: city=%s, weather=%s", city, resp.Lives[0].Weather)
	return &resp.Lives[0], nil
}

// Forecast 查询预报天气
func (s *Service) Forecast(ctx context.Context, city string) (*models.WeatherForecast, error) {
	opts := &Options{
		City:       city,
		Extensions: "all",
	}

	resp, err := s.Query(ctx, opts)
	if err != nil {
		return nil, err
	}

	if len(resp.Forecasts) == 0 {
		return nil, errors.New("no forecast data found")
	}

	s.logger.Info("Weather forecast success: city=%s, days=%d", city, len(resp.Forecasts[0].Casts))
	return &resp.Forecasts[0], nil
}

// GetTomorrowWeather 获取明天天气
func (s *Service) GetTomorrowWeather(ctx context.Context, city string) (*models.WeatherCast, error) {
	forecast, err := s.Forecast(ctx, city)
	if err != nil {
		return nil, err
	}

	if len(forecast.Casts) < 2 {
		return nil, errors.New("insufficient forecast data")
	}

	return &forecast.Casts[1], nil
}

// GetNextDaysWeather 获取未来N天天气预报
func (s *Service) GetNextDaysWeather(ctx context.Context, city string, days int) ([]models.WeatherCast, error) {
	if days <= 0 {
		return nil, errors.New("days must be greater than 0")
	}

	forecast, err := s.Forecast(ctx, city)
	if err != nil {
		return nil, err
	}

	if len(forecast.Casts) < days {
		return nil, fmt.Errorf("insufficient forecast data, requested %d days but got %d", days, len(forecast.Casts))
	}

	return forecast.Casts[:days], nil
}
