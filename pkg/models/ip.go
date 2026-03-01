package models

// IpLocationResponse IP定位响应
type IpLocationResponse struct {
	BaseResponse
	Province  string `json:"province"`
	City      string `json:"city"`
	Adcode    string `json:"adcode"`
	Rectangle string `json:"rectangle"`
	IP        string `json:"ip"`
}
