package models

// GeocodeResponse 地理编码响应
type GeocodeResponse struct {
	BaseResponse
	Geocodes []GeocodeResult `json:"geocodes"` // 地理编码结果列表
}

// GeocodeResult 地理编码结果
type GeocodeResult struct {
	Province         string     `json:"province"`          // 省份
	City             string     `json:"city"`              // 城市
	District         string     `json:"district"`          // 区县
	Township         FlexString `json:"township"`          // 乡镇
	FormattedAddress string     `json:"formatted_address"` // 格式化地址
	Location         string     `json:"location"`          // 坐标（经度,纬度）
	Adcode           string     `json:"adcode"`            // 区域编码
	Level            string     `json:"level"`             // 地址等级
	Citycode         string     `json:"citycode"`          // 城市编码
	Bounds           string     `json:"bounds"`           // 坐标范围
	Accurate         bool       `json:"accurate"`          // 是否精确
}

// ReGeocodeResponse 逆地理编码响应
type ReGeocodeResponse struct {
	BaseResponse
	Regeocode *ReGeocodeResult `json:"regeocode"` // 逆地理编码结果
}

// ReGeocodeResult 逆地理编码结果
type ReGeocodeResult struct {
	FormattedAddress string           `json:"formatted_address"` // 格式化地址
	AddressComponent AddressComponent `json:"addressComponent"` // 地址组件
	Pois             []PoiInfo        `json:"pois"`              // 附近POI列表
}
