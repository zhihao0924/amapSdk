# 高德地图 Go SDK

> 简单易用、类型安全的 Golang 高德地图 API SDK

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-blue.svg)](https://golang.org/doc/go1.21)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

## ✨ 特性

- 🎯 **类型安全** - 强类型参数和枚举，减少运行时错误
- 🔄 **Context 支持** - 完整的 `context.Context` 支持，支持请求取消和超时
- 🛡️ **错误处理** - 完善的错误类型和判断函数
- 🔄 **重试机制** - 可配置的指数退避重试策略
- 🔌 **拦截器** - 灵活的请求/响应拦截器链
- ⚡ **性能优化** - `sync.Pool` 复用 buffer，减少内存分配
- 🧵 **并发安全** - `RWMutex` 保护，线程安全
- 📝 **日志系统** - 可配置的日志接口
- 🏗️ **模块化设计** - 清晰的分层架构

## 📦 功能模块

| 模块 | 功能 | API |
|------|------|-----|
| **Geocode** | 地理编码/逆地理编码 | `/v3/geocode/geo`, `/v3/geocode/regeo` |
| **Direction** | 驾车/步行路径规划 | `/v3/direction/driving`, `/v3/direction/walking` |
| **Place** | POI 关键字/周边/多边形搜索 | `/v3/place/text`, `/v3/place/around`, `/v3/place/polygon` |
| **Weather** | 实况天气与天气预报 | `/v3/weather/weatherInfo` |
| **IP** | IP 地址定位 | `/v3/ip` |

## 📥 安装

```bash
go get github.com/zhihao0924/amapSdk
```

## 🚀 快速开始

### 1. 获取 API Key

1. 访问 [高德开放平台](https://lbs.amap.com/)
2. 注册并登录账号
3. 进入控制台，创建应用
4. 获取 Web 服务 API Key

### 2. 初始化客户端

```go
package main

import (
    "context"
    "log"
    "github.com/zhihao0924/amapSdk"
)

func main() {
    // 基础配置
    client, err := amap.NewClient(&amap.Config{
        Key: "YOUR_AMAP_KEY",
    })
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // 使用示例...
}
```

### 3. 完整配置

```go
client, err := amap.NewClient(&amap.Config{
    Key:      "YOUR_AMAP_KEY",
    BaseURL:  "https://restapi.amap.com/v3",
    Timeout:  10,
    Debug:    true,
    Headers:  map[string]string{"User-Agent": "MyApp/1.0"},
    RetryConfig: &amap.RetryConfig{
        MaxRetries: 3,
        RetryDelay: time.Second,
    },
})
```

## 📖 使用示例

### 地理编码

```go
ctx := context.Background()

// 地址转坐标
resp, err := client.Geocode().Geo(ctx, &amap.GeocodeOptions{
    Address: "北京市朝阳区阜通东大街6号",
    City:    "北京",
})

// 坐标转地址
resp, err := client.Geocode().ReGeo(ctx, &amap.ReGeoOptions{
    Location:   "116.480881,39.989410",
    Extensions: "all",
})
```

### 路径规划

```go
ctx := context.Background()

// 驾车路径
resp, err := client.Direction().Driving(ctx, &amap.DrivingOptions{
    Origin:      "116.481028,39.989643",
    Destination: "116.465302,40.004717",
    Strategy:    amap.DrivingStrategyFastest.String(),
})

// 步行路径
resp, err := client.Direction().Walking(ctx, &amap.WalkingOptions{
    Origin:      "116.481028,39.989643",
    Destination: "116.465302,40.004717",
})
```

### POI 搜索

```go
ctx := context.Background()

// 关键字搜索
resp, err := client.Place().TextSearch(ctx, &amap.TextSearchOptions{
    Keywords:   "肯德基",
    City:       "北京",
    CityLimit:  "1",
    Extensions: "all",
})

// 周边搜索
resp, err := client.Place().AroundSearch(ctx, &amap.AroundSearchOptions{
    Location: "116.481028,39.989643",
    Keywords: "餐厅",
    Radius:   "1000",
})
```

### 天气查询

```go
// 实时天气
live, err := client.Weather().Base(context.Background(), "110101")

// 天气预报
forecast, err := client.Weather().GetNextDaysWeather(context.Background(), "110101", 4)
```

### IP 定位

```go
ctx := context.Background()

// 指定 IP
resp, err := client.IP().Location(ctx, &amap.LocationOptions{IP: "8.8.8.8"})

// 当前 IP
resp, err := client.IP().Current(ctx)
```

## 🔧 配置说明

### Config 配置项

| 字段 | 类型 | 必需 | 默认值 | 说明 |
|------|------|------|--------|------|
| `Key` | string | ✅ | - | 高德地图 API Key |
| `BaseURL` | string | ❌ | `https://restapi.amap.com/v3` | API 基础 URL |
| `Timeout` | int | ❌ | 10 | 请求超时时间（秒） |
| `Debug` | bool | ❌ | false | 是否开启调试日志 |
| `Headers` | map[string]string | ❌ | nil | 自定义请求头 |
| `RetryConfig` | *RetryConfig | ❌ | DefaultRetryConfig | 重试配置 |
| `InterceptorChain` | *InterceptorChain | ❌ | nil | 拦截器链 |

### 枚举类型

```go
// 驾车策略
amap.DrivingStrategyFastest        // 速度优先
amap.DrivingStrategyCost            // 费用优先
amap.DrivingStrategyDistance        // 距离优先
amap.DrivingStrategyAvoidCongestion // 避免拥堵

// 天气类型
amap.WeatherTypeLive     // 实时天气
amap.WeatherTypeForecast // 天气预报

// POI 排序
amap.SortTypeDistance // 按距离排序
amap.SortTypeWeight  // 按权重排序
```

## 🛠️ 错误处理

```go
resp, err := client.Geocode().Geo(context.Background(), &amap.GeocodeOptions{
    Address: "不存在的地址",
})
if err != nil {
    // 错误类型判断
    switch {
    case amap.IsNetworkError(err):
        log.Println("网络连接失败")
    case amap.IsTimeoutError(err):
        log.Println("请求超时")
    case amap.IsAuthError(err):
        log.Println("API Key 无效")
    default:
        log.Println(err)
    }
}
```

## 🎨 高级特性

### 自定义拦截器

```go
interceptor := amap.NewInterceptorChain()

// 请求拦截器
interceptor.AddRequest(func(req *http.Request) error {
    req.Header.Set("X-Request-ID", generateRequestID())
    return nil
})

// 响应拦截器
interceptor.AddResponse(func(resp *http.Response) error {
    fmt.Printf("状态码: %d\n", resp.StatusCode)
    return nil
})

client, err := amap.NewClient(&amap.Config{
    Key:            "YOUR_AMAP_KEY",
    InterceptorChain: interceptor,
})
```

### 自定义重试策略

```go
retryConfig := amap.NewRetryConfig(5, 2*time.Second).
    WithMaxRetries(6).
    WithRetryDelay(time.Second * 2).
    WithRetryable(func(err error) bool {
        return amap.IsNetworkError(err)
    })

client, err := amap.NewClient(&amap.Config{
    Key:         "YOUR_AMAP_KEY",
    RetryConfig: retryConfig,
})
```

### Context 超时控制

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// SDK 内部支持 context，可以随时取消请求
resp, err := client.Geocode().Geo(ctx, &amap.GeocodeOptions{
    Address: "北京市朝阳区",
})
```

## 📁 项目结构

```
amap-sdk/
├── client.go              # 客户端主入口
├── options.go             # 选项构建器
├── errors_export.go       # 错误处理导出
│
├── pkg/
│   ├── common/            # 通用工具
│   │   ├── errors.go
│   │   ├── location.go
│   │   └── logger.go
│   │
│   ├── core/              # 核心功能
│   │   ├── client.go      # 核心客户端
│   │   ├── config.go      # 配置管理
│   │   ├── enums.go       # 枚举类型
│   │   ├── http.go        # HTTP 客户端
│   │   ├── middleware.go  # 拦截器链
│   │   └── retry.go       # 重试机制
│   │
│   ├── models/            # 数据模型
│   │   ├── base.go        # 基础模型
│   │   ├── types.go       # 自定义类型
│   │   ├── direction.go   # 路径规划
│   │   ├── geocode.go     # 地理编码
│   │   ├── ip.go          # IP 定位
│   │   ├── place.go       # POI 搜索
│   │   └── weather.go     # 天气查询
│   │
│   └── services/          # 服务实现
│       ├── direction/
│       ├── geocode/
│       ├── ip/
│       ├── place/
│       └── weather/
│
└── examples/              # 使用示例
    ├── basic/
    ├── direction/
    ├── geocode/
    ├── ip/
    ├── place/
    └── weather/
```

## 📚 更多示例

```bash
cd examples/basic && go run basic.go
cd examples/geocode && go run geocode.go
cd examples/direction && go run driving.go
cd examples/place && go run text_search.go
cd examples/weather && go run weather.go
cd examples/ip && go run ip.go
```

## 🔗 相关链接

- [高德开放平台](https://lbs.amap.com/)
- [Web 服务 API 文档](https://lbs.amap.com/api/webservice/summary)
- [获取 API Key](https://console.amap.com/dev/key/app)

## 📄 许可证

[MIT License](LICENSE)

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

## 📞 联系方式

- Issues: [GitHub Issues](https://github.com/zhihao0924/amapSdk/issues)

---

**Made with ❤️ by zhihao**
