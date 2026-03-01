package models

// WeatherResponse 天气查询响应
type WeatherResponse struct {
	BaseResponse
	Lives     []WeatherLive     `json:"lives"`
	Forecasts []WeatherForecast `json:"forecasts"`
}

// WeatherLive 实况天气
type WeatherLive struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	Adcode        string `json:"adcode"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	WindDirection string `json:"winddirection"`
	WindPower     string `json:"windpower"`
	Humidity      string `json:"humidity"`
	ReportTime    string `json:"reporttime"`
}

// WeatherForecast 预报天气
type WeatherForecast struct {
	City       string        `json:"city"`
	Adcode     string        `json:"adcode"`
	Province   string        `json:"province"`
	ReportTime string        `json:"reporttime"`
	Casts      []WeatherCast `json:"casts"`
}

// WeatherCast 天气预报
type WeatherCast struct {
	Date         string `json:"date"`
	Week         string `json:"week"`
	DayWeather   string `json:"dayweather"`
	NightWeather string `json:"nightweather"`
	DayTemp      string `json:"daytemp"`
	NightTemp    string `json:"nighttemp"`
	DayWind      string `json:"daywind"`
	NightWind    string `json:"nightwind"`
	DayPower     string `json:"daypower"`
	NightPower   string `json:"nightpower"`
}
