package models

// BaseResponse API基础响应
type BaseResponse struct {
	Status   string `json:"status"`    // 状态码，1=成功，0=失败
	Info     string `json:"info"`      // 返回的状态说明
	Infocode string `json:"infocode"` // 返回状态说明
}

// GetStatus 实现StatusChecker接口
func (b *BaseResponse) GetStatus() string   { return b.Status }
func (b *BaseResponse) GetInfo() string     { return b.Info }
func (b *BaseResponse) GetInfocode() string { return b.Infocode }

// AddressComponent 地址组件
type AddressComponent struct {
	Country  string `json:"country"`  // 国家
	Province string `json:"province"` // 省份
	City     string `json:"city"`     // 城市
	District string `json:"district"` // 区县
	Township string `json:"township"` // 乡镇
	Adcode   string `json:"adcode"`   // 区域编码
	Towncode string `json:"towncode"` // 乡镇编码
	Citycode string `json:"citycode"` // 城市编码
}

// PoiInfo POI信息
type PoiInfo struct {
	ID       string          `json:"id"`       // POI ID
	Name     string          `json:"name"`     // POI名称
	Type     string          `json:"type"`     // POI类型
	Tel      string          `json:"tel"`      // 电话
	Address  string          `json:"address"`  // 地址
	Location string          `json:"location"` // 坐标（经度,纬度）
	Distance Float64OrString `json:"distance"` // 距离（米）
}

// Suggestion 搜索建议信息
type Suggestion struct {
	Keywords []string `json:"keywords"` // 关键词建议
	Cities   []string `json:"cities"`   // 城市建议
}
