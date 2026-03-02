package models

// IpLocationResponse IP定位响应
type IpLocationResponse struct {
	BaseResponse
	Province  string `json:"province"`  // 省份
	City      string `json:"city"`      // 城市
	Adcode    string `json:"adcode"`    // 区域编码
	Rectangle string `json:"rectangle"` // 矩形范围
	IP        string `json:"ip"`        // 查询的IP地址
}
