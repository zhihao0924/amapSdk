package models

// IpLocationResponse IP定位响应
type IpLocationResponse struct {
	BaseResponse
	Province  FlexString `json:"province"`  // 省份
	City      FlexString `json:"city"`      // 城市
	Adcode    FlexString `json:"adcode"`    // 区域编码
	Rectangle FlexString `json:"rectangle"` // 矩形范围
	IP        string     `json:"ip"`        // 查询的IP地址
}
