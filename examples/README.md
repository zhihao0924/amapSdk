# 高德地图 Go SDK 示例

本目录包含当前仓库中可直接运行的示例程序。

## 目录结构

```text
examples/
├── README.md
├── around_search/
│   └── around_search.go
├── basic/
│   └── basic.go
├── direction/
│   └── driving.go
├── geocode/
│   └── geocode.go
├── internal/
│   └── exampleutil/
│       └── client.go
├── ip/
│   └── ip.go
├── place/
│   └── text_search.go
├── polygon_search/
│   └── polygon_search.go
├── regeo/
│   └── regeo.go
├── regeo_location/
│   └── regeo_location.go
├── walking/
│   └── walking.go
└── weather/
    └── weather.go
├── weather_days/
│   └── weather_days.go
└── ip_tools/
    └── ip_tools.go
```

## 运行前准备

示例通过环境变量读取 API Key，不再在源码里硬编码密钥。

```bash
export AMAP_KEY=你的高德_Web服务_Key
```

## 运行方式

```bash
go run ./examples/basic
go run ./examples/geocode
go run ./examples/regeo
go run ./examples/regeo_location
go run ./examples/direction
go run ./examples/walking
go run ./examples/place
go run ./examples/around_search
go run ./examples/polygon_search
go run ./examples/weather
go run ./examples/weather_days
go run ./examples/ip
go run ./examples/ip_tools
```

## 示例说明

### basic

`basic/basic.go` 演示客户端初始化、配置读取、日志接口和工具函数。

### geocode

`geocode/geocode.go` 演示地址转经纬度：

```go
ctx, cancel := exampleutil.NewRequestContext()
defer cancel()

resp, err := client.Geocode().Geo(ctx, &amap.GeocodeOptions{
    Address: "北京市朝阳区",
    City:    "北京",
})
```

### regeo

`regeo/regeo.go` 演示经纬度转地址：

```go
ctx, cancel := exampleutil.NewRequestContext()
defer cancel()

resp, err := client.Geocode().ReGeo(ctx, &amap.ReGeoOptions{
    Location:   "116.397428,39.90923",
    Extensions: "all",
})
```

### regeo_location

`regeo_location/regeo_location.go` 演示 `ReGeoByLocation` 辅助方法。

### direction

`direction/driving.go` 演示驾车路径规划：

```go
ctx, cancel := exampleutil.NewRequestContext()
defer cancel()

resp, err := client.Direction().Driving(ctx, &amap.DrivingOptions{
    Origin:      "116.481181,39.989792",
    Destination: "116.397428,39.90923",
    Strategy:    "1",
})
```

### walking

`walking/walking.go` 演示 `Walking` 和 `WalkingByLocations`。

### place

`place/text_search.go` 演示关键字搜索 POI：

```go
ctx, cancel := exampleutil.NewRequestContext()
defer cancel()

resp, err := client.Place().TextSearch(ctx, &amap.TextSearchOptions{
    Keywords: "肯德基",
    City:     "北京",
})
```

### around_search

`around_search/around_search.go` 演示 `AroundSearch` 和 `AroundSearchByLocation`。

### polygon_search

`polygon_search/polygon_search.go` 演示 `SearchByPolygon`。

### weather

`weather/weather.go` 分别演示实时天气和天气预报：

```go
live, err := client.Weather().Base(ctx, "110101")
forecast, err := client.Weather().Forecast(ctx, "110101")
```

### weather_days

`weather_days/weather_days.go` 演示 `GetTomorrowWeather` 和 `GetNextDaysWeather`。

### ip

`ip/ip.go` 演示指定 IP 定位：

```go
resp, err := client.IP().Location(ctx, &amap.LocationOptions{
    IP: "114.247.50.2",
})
```

### ip_tools

`ip_tools/ip_tools.go` 演示 `Current`、`GetIPInfo` 和 `BatchLocation`。

## 调试说明

- 所有请求示例都使用了带超时的 `context.Context`
- API 业务错误现在会直接通过 `err` 返回，示例中不再额外判断 `resp.Status`
- 如果未设置 `AMAP_KEY`，示例会直接给出明确报错
