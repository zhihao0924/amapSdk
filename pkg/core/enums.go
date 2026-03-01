package core

// DrivingStrategy 驾车路径规划策略
type DrivingStrategy int

const (
	// DrivingStrategyFastest 速度优先
	DrivingStrategyFastest DrivingStrategy = iota + 1
	// DrivingStrategyCost 费用优先
	DrivingStrategyCost
	// DrivingStrategyDistance 距离优先
	DrivingStrategyDistance
	// DrivingStrategyGeneral 不走高速
	DrivingStrategyGeneral
	// DrivingStrategyAvoidCongestion 避免拥堵
	DrivingStrategyAvoidCongestion
	// DrivingStrategyAvoidHighSpeed 不走高速
	DrivingStrategyAvoidHighSpeed
	// DrivingStrategyAvoidHighSpeedCongestion 不走高速且避免拥堵
	DrivingStrategyAvoidHighSpeedCongestion
)

// String 返回策略字符串
func (s DrivingStrategy) String() string {
	return string(rune('0' + s))
}

// GeocodeExtension 地理编码扩展类型
type GeocodeExtension int

const (
	// GeocodeExtensionNone 不返回扩展信息
	GeocodeExtensionNone GeocodeExtension = iota
	// GeocodeExtensionBase 返回基本地址信息
	GeocodeExtensionBase
	// GeocodeExtensionAll 返回所有扩展信息
	GeocodeExtensionAll
)

// String 返回扩展类型字符串
func (e GeocodeExtension) String() string {
	return []string{"", "base", "all"}[e]
}

// WeatherType 天气查询类型
type WeatherType int

const (
	// WeatherTypeLive 实时天气
	WeatherTypeLive WeatherType = iota
	// WeatherTypeForecast 预报天气
	WeatherTypeForecast
)

// String 返回天气类型字符串
func (t WeatherType) String() string {
	return []string{"base", "all"}[t]
}

// SearchType POI搜索类型
type SearchType int

const (
	// SearchTypeText 关键字搜索
	SearchTypeText SearchType = iota
	// SearchTypeAround 周边搜索
	// SearchTypePolygon 多边形搜索
	SearchTypePolygon
)

// SortType 排序类型
type SortType int

const (
	// SortTypeDefault 默认排序
	SortTypeDefault SortType = iota
	// SortTypeDistance 距离排序
	SortTypeDistance
	// SortTypeWeight 权重排序
	SortTypeWeight
)

// String 返回排序类型字符串
func (s SortType) String() string {
	return []string{"", "distance", "weight"}[s]
}
