package common

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	// minLng 最小经度
	minLng = -180.0
	// maxLng 最大经度
	maxLng = 180.0
	// minLat 最小纬度
	minLat = -90.0
	// maxLat 最大纬度
	maxLat = 90.0
	// defaultPrecision 默认精度
	defaultPrecision = 6
)

// Location 位置结构
type Location struct {
	Lng float64
	Lat float64
}

// NewLocation 创建位置
func NewLocation(lng, lat float64) *Location {
	return &Location{
		Lng: roundToPrecision(lng, defaultPrecision),
		Lat: roundToPrecision(lat, defaultPrecision),
	}
}

// NewLocationFromString 从字符串创建位置
func NewLocationFromString(location string) (*Location, error) {
	return ParseLocation(location)
}

// String 返回位置字符串
func (l *Location) String() string {
	if l == nil {
		return ""
	}
	return fmt.Sprintf("%.6f,%.6f", l.Lng, l.Lat)
}

// Validate 验证位置
func (l *Location) Validate() error {
	if l == nil {
		return ErrInvalidParamsError
	}

	if l.Lng < minLng || l.Lng > maxLng {
		return fmt.Errorf("longitude must be between %f and %f", minLng, maxLng)
	}
	if l.Lat < minLat || l.Lat > maxLat {
		return fmt.Errorf("latitude must be between %f and %f", minLat, maxLat)
	}

	return nil
}

// DistanceTo 计算到另一个位置的距离（单位：米）
func (l *Location) DistanceTo(other *Location) float64 {
	if l == nil || other == nil {
		return 0
	}

	// 使用Haversine公式计算大圆距离
	const earthRadius = 6371000 // 地球半径（米）

	lat1Rad := toRadians(l.Lat)
	lat2Rad := toRadians(other.Lat)
	deltaLat := toRadians(other.Lat - l.Lat)
	deltaLng := toRadians(other.Lng - l.Lng)

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLng/2)*math.Sin(deltaLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

// ParseLocation 解析位置字符串
func ParseLocation(location string) (*Location, error) {
	if location == "" {
		return nil, ErrInvalidParamsError
	}

	parts := strings.Split(strings.TrimSpace(location), ",")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid location format, expected 'lng,lat', got: %s", location)
	}

	lng, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("invalid longitude: %w", err)
	}

	lat, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return nil, fmt.Errorf("invalid latitude: %w", err)
	}

	loc := NewLocation(lng, lat)
	if err := loc.Validate(); err != nil {
		return nil, err
	}

	return loc, nil
}

// Box 边界框
type Box struct {
	MinLng float64
	MinLat float64
	MaxLng float64
	MaxLat float64
}

// NewBox 创建边界框
func NewBox(minLng, minLat, maxLng, maxLat float64) *Box {
	return &Box{
		MinLng: minLng,
		MinLat: minLat,
		MaxLng: maxLng,
		MaxLat: maxLat,
	}
}

// NewBoxFromLocations 从多个位置创建边界框
func NewBoxFromLocations(locations []*Location) *Box {
	if len(locations) == 0 {
		return nil
	}

	minLng, maxLng := locations[0].Lng, locations[0].Lng
	minLat, maxLat := locations[0].Lat, locations[0].Lat

	for _, loc := range locations {
		if loc.Lng < minLng {
			minLng = loc.Lng
		}
		if loc.Lng > maxLng {
			maxLng = loc.Lng
		}
		if loc.Lat < minLat {
			minLat = loc.Lat
		}
		if loc.Lat > maxLat {
			maxLat = loc.Lat
		}
	}

	return NewBox(minLng, minLat, maxLng, maxLat)
}

// String 返回边界框字符串
func (b *Box) String() string {
	if b == nil {
		return ""
	}
	return fmt.Sprintf("%.6f,%.6f,%.6f,%.6f", b.MinLng, b.MinLat, b.MaxLng, b.MaxLat)
}

// Contains 检查位置是否在边界框内
func (b *Box) Contains(loc *Location) bool {
	if b == nil || loc == nil {
		return false
	}

	return loc.Lng >= b.MinLng && loc.Lng <= b.MaxLng &&
		loc.Lat >= b.MinLat && loc.Lat <= b.MaxLat
}

// Center 获取边界框中心点
func (b *Box) Center() *Location {
	if b == nil {
		return nil
	}

	return NewLocation(
		(b.MinLng+b.MaxLng)/2,
		(b.MinLat+b.MaxLat)/2,
	)
}

// Width 获取边界框宽度（经度差）
func (b *Box) Width() float64 {
	if b == nil {
		return 0
	}
	return math.Abs(b.MaxLng - b.MinLng)
}

// Height 获取边界框高度（纬度差）
func (b *Box) Height() float64 {
	if b == nil {
		return 0
	}
	return math.Abs(b.MaxLat - b.MinLat)
}

// ParseBox 解析边界框字符串
func ParseBox(box string) (*Box, error) {
	if box == "" {
		return nil, ErrInvalidParamsError
	}

	parts := strings.Split(strings.TrimSpace(box), ",")
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid box format, expected 'minLng,minLat,maxLng,maxLat', got: %s", box)
	}

	minLng, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("invalid min longitude: %w", err)
	}

	minLat, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return nil, fmt.Errorf("invalid min latitude: %w", err)
	}

	maxLng, err := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
	if err != nil {
		return nil, fmt.Errorf("invalid max longitude: %w", err)
	}

	maxLat, err := strconv.ParseFloat(strings.TrimSpace(parts[3]), 64)
	if err != nil {
		return nil, fmt.Errorf("invalid max latitude: %w", err)
	}

	// 验证边界框
	if minLng >= maxLng {
		return nil, fmt.Errorf("min longitude must be less than max longitude")
	}
	if minLat >= maxLat {
		return nil, fmt.Errorf("min latitude must be less than max latitude")
	}

	return NewBox(minLng, minLat, maxLng, maxLat), nil
}

// toRadians 将度数转换为弧度
func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// roundToPrecision 四舍五入到指定精度
func roundToPrecision(value float64, precision int) float64 {
	factor := math.Pow10(precision)
	return math.Round(value*factor) / factor
}
