package models

// GeocodeResponse 地理编码响应
type GeocodeResponse struct {
	BaseResponse
	Geocodes []GeocodeResult `json:"geocodes"`
}

// GeocodeResult 地理编码结果
type GeocodeResult struct {
	Province         string `json:"province"`
	City             string `json:"city"`
	District         string `json:"district"`
	Township         string `json:"township"`
	FormattedAddress string `json:"formatted_address"`
	Location         string `json:"location"`
	Adcode           string `json:"adcode"`
	Level            string `json:"level"`
	Citycode         string `json:"citycode"`
	Bounds           string `json:"bounds"`
	Accurate         bool   `json:"accurate"`
}

// ReGeocodeResponse 逆地理编码响应
type ReGeocodeResponse struct {
	BaseResponse
	Regeocode *ReGeocodeResult `json:"regeocode"`
}

// ReGeocodeResult 逆地理编码结果
type ReGeocodeResult struct {
	FormattedAddress string           `json:"formatted_address"`
	AddressComponent AddressComponent `json:"addressComponent"`
	Pois             []PoiInfo        `json:"pois"`
}
