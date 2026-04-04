package amap

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/zhihao0924/amapSdk/pkg/common"
	"github.com/zhihao0924/amapSdk/pkg/core"
	"github.com/zhihao0924/amapSdk/pkg/models"
	directionservice "github.com/zhihao0924/amapSdk/pkg/services/direction"
	geocodeservice "github.com/zhihao0924/amapSdk/pkg/services/geocode"
	ipservice "github.com/zhihao0924/amapSdk/pkg/services/ip"
	placeservice "github.com/zhihao0924/amapSdk/pkg/services/place"
	weatherservice "github.com/zhihao0924/amapSdk/pkg/services/weather"
)

// Client 高德地图API客户端
type Client struct {
	coreClient *core.Client

	// 服务缓存（单例模式）
	geocodeService   *geocodeservice.Service
	directionService *directionservice.Service
	placeService     *placeservice.Service
	weatherService   *weatherservice.Service
	ipService        *ipservice.Service

	mu sync.RWMutex
}

// Config SDK配置（公开）
type (
	Config           = core.Config
	RetryConfig      = core.RetryConfig
	InterceptorChain = core.InterceptorChain

	// 枚举类型
	DrivingStrategy  = core.DrivingStrategy
	GeocodeExtension = core.GeocodeExtension
	WeatherType      = core.WeatherType
	SearchType       = core.SearchType
	SortType         = core.SortType
)

// 重新导出类型
type (
	Logger   = common.Logger
	Error    = common.Error
	Location = common.Location
	Box      = common.Box

	// 基础模型
	BaseResponse     = models.BaseResponse
	AddressComponent = models.AddressComponent
	PoiInfo          = models.PoiInfo
	Suggestion       = models.Suggestion

	// 地理编码模型
	GeocodeResponse   = models.GeocodeResponse
	GeocodeResult     = models.GeocodeResult
	ReGeocodeResponse = models.ReGeocodeResponse
	ReGeocodeResult   = models.ReGeocodeResult

	// 路径规划模型
	DrivingResponse = models.DrivingResponse
	DrivingRoute    = models.DrivingRoute
	DrivingPath     = models.DrivingPath
	DrivingStep     = models.DrivingStep
	WalkingResponse = models.WalkingResponse
	WalkingRoute    = models.WalkingRoute
	WalkingPath     = models.WalkingPath
	WalkingStep     = models.WalkingStep

	// POI模型
	TextSearchResponse      = models.TextSearchResponse
	AroundSearchResponse    = models.TextSearchResponse
	SearchByPolygonResponse = models.SearchByPolygonResponse

	// 天气模型
	WeatherResponse = models.WeatherResponse
	WeatherLive     = models.WeatherLive
	WeatherForecast = models.WeatherForecast
	WeatherCast     = models.WeatherCast

	// IP模型
	IpLocationResponse = models.IpLocationResponse
)

// 服务类型
type (
	GeocodeService   = geocodeservice.Service
	DirectionService = directionservice.Service
	PlaceService     = placeservice.Service
	WeatherService   = weatherservice.Service
	IPService        = ipservice.Service

	GeocodeOptions = geocodeservice.Options
	ReGeoOptions   = geocodeservice.ReGeoOptions

	DrivingOptions = directionservice.DrivingOptions
	WalkingOptions = directionservice.WalkingOptions

	TextSearchOptions      = placeservice.TextSearchOptions
	AroundSearchOptions    = placeservice.AroundSearchOptions
	SearchByPolygonOptions = placeservice.SearchByPolygonOptions

	WeatherOptions = weatherservice.Options

	LocationOptions = ipservice.LocationOptions
)

// NewClient 创建新的客户端
func NewClient(config *Config) (*Client, error) {
	if config == nil || config.Key == "" {
		return nil, common.NewErrorf(common.ErrInvalidConfig, "config or config.Key is required")
	}

	// 验证配置
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("配置验证失败: %w", err)
	}

	coreClient, err := core.NewClient(config)
	if err != nil {
		return nil, common.WrapError(err, "创建核心客户端失败")
	}

	return &Client{
		coreClient: coreClient,
	}, nil
}

// GetConfig 获取配置
func (c *Client) GetConfig() *Config {
	return c.coreClient.GetConfig()
}

// GetLogger 获取日志实例
func (c *Client) GetLogger() Logger {
	return c.coreClient.GetLogger()
}

// Geocode 地理编码服务（单例）
func (c *Client) Geocode() *GeocodeService {
	c.mu.RLock()
	if c.geocodeService != nil {
		c.mu.RUnlock()
		return c.geocodeService
	}
	c.mu.RUnlock()

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.geocodeService == nil {
		c.geocodeService = geocodeservice.New(c.coreClient.GetHTTP(), c.coreClient.GetLogger())
	}
	return c.geocodeService
}

// Direction 路径规划服务（单例）
func (c *Client) Direction() *DirectionService {
	c.mu.RLock()
	if c.directionService != nil {
		c.mu.RUnlock()
		return c.directionService
	}
	c.mu.RUnlock()

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.directionService == nil {
		c.directionService = directionservice.New(c.coreClient.GetHTTP(), c.coreClient.GetLogger())
	}
	return c.directionService
}

// Place POI搜索服务（单例）
func (c *Client) Place() *PlaceService {
	c.mu.RLock()
	if c.placeService != nil {
		c.mu.RUnlock()
		return c.placeService
	}
	c.mu.RUnlock()

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.placeService == nil {
		c.placeService = placeservice.New(c.coreClient.GetHTTP(), c.coreClient.GetLogger())
	}
	return c.placeService
}

// Weather 天气服务（单例）
func (c *Client) Weather() *WeatherService {
	c.mu.RLock()
	if c.weatherService != nil {
		c.mu.RUnlock()
		return c.weatherService
	}
	c.mu.RUnlock()

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.weatherService == nil {
		c.weatherService = weatherservice.New(c.coreClient.GetHTTP(), c.coreClient.GetLogger())
	}
	return c.weatherService
}

// IP IP定位服务（单例）
func (c *Client) IP() *IPService {
	c.mu.RLock()
	if c.ipService != nil {
		c.mu.RUnlock()
		return c.ipService
	}
	c.mu.RUnlock()

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.ipService == nil {
		c.ipService = ipservice.New(c.coreClient.GetHTTP(), c.coreClient.GetLogger())
	}
	return c.ipService
}

// Close 关闭客户端，释放资源
func (c *Client) Close() error {
	if err := c.coreClient.Close(); err != nil {
		return fmt.Errorf("关闭客户端失败: %w", err)
	}
	return nil
}

// CloseWithTimeout 带超时的关闭方法
func (c *Client) CloseWithTimeout(timeout time.Duration) error {
	if c.coreClient == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- c.Close()
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return fmt.Errorf("关闭客户端超时: %w", ctx.Err())
	}
}

// IsClosed 检查客户端是否已关闭
func (c *Client) IsClosed() bool {
	return c.coreClient == nil
}

// String 返回客户端字符串表示
func (c *Client) String() string {
	config := c.GetConfig()
	if config == nil {
		return "AMap Client (未初始化)"
	}
	return fmt.Sprintf("AMap Client (Key: %s, BaseURL: %s)", maskKey(config.Key), config.BaseURL)
}

// maskKey 掩码API Key
func maskKey(key string) string {
	if len(key) <= 8 {
		return "****"
	}
	return key[:4] + "****" + key[len(key)-4:]
}
