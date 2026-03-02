package models

// WeatherResponse 天气查询响应
type WeatherResponse struct {
	BaseResponse
	Lives     []WeatherLive     `json:"lives"`     // 实况天气列表
	Forecasts []WeatherForecast `json:"forecasts"` // 预报天气列表
}

// WeatherLive 实况天气
type WeatherLive struct {
	Province      string `json:"province"`      // 省份
	City          string `json:"city"`          // 城市
	Adcode        string `json:"adcode"`        // 区域编码
	Weather       string `json:"weather"`       // 天气现象
	Temperature   string `json:"temperature"`   // 实时温度
	WindDirection string `json:"winddirection"` // 风向
	WindPower     string `json:"windpower"`     // 风力
	Humidity      string `json:"humidity"`      // 湿度
	ReportTime    string `json:"reporttime"`    // 发布时间
}

// WeatherForecast 预报天气
type WeatherForecast struct {
	City       string        `json:"city"`       // 城市
	Adcode     string        `json:"adcode"`     // 区域编码
	Province   string        `json:"province"`   // 省份
	ReportTime string        `json:"reporttime"` // 发布时间
	Casts      []WeatherCast `json:"casts"`      // 天气预报列表
}

// WeatherCast 天气预报
type WeatherCast struct {
	Date         string `json:"date"`         // 日期
	Week         string `json:"week"`         // 星期
	DayWeather   string `json:"dayweather"`   // 白天天气
	NightWeather string `json:"nightweather"` // 夜间天气
	DayTemp      string `json:"daytemp"`      // 白天温度
	NightTemp    string `json:"nighttemp"`    // 夜间温度
	DayWind      string `json:"daywind"`      // 白天风向
	NightWind    string `json:"nightwind"`    // 夜间风向
	DayPower     string `json:"daypower"`     // 白天风力
	NightPower   string `json:"nightpower"`   // 夜间风力
}
