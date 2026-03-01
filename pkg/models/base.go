package models

// BaseResponse API基础响应
type BaseResponse struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
}

// GetStatus 实现StatusChecker接口
func (b *BaseResponse) GetStatus() string   { return b.Status }
func (b *BaseResponse) GetInfo() string     { return b.Info }
func (b *BaseResponse) GetInfocode() string { return b.Infocode }

// AddressComponent 地址组件
type AddressComponent struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Township string `json:"township"`
	Adcode   string `json:"adcode"`
	Towncode string `json:"towncode"`
	Citycode string `json:"citycode"`
}

// PoiInfo POI信息
type PoiInfo struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Tel      string  `json:"tel"`
	Address  string  `json:"address"`
	Location string  `json:"location"`
	Distance float64 `json:"distance"`
}

// Suggestion 建议信息
type Suggestion struct {
	Keywords []string `json:"keywords"`
	Cities   []string `json:"cities"`
}
