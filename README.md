# 高德地图 Go SDK

> 一个简单易用、结构清晰的 Golang 高德地图 API SDK

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-blue.svg)](https://golang.org/doc/go1.21)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

## ✨ 特性

- 🎯 **类型安全** - 强类型参数和枚举，减少运行时错误
- 🔄 **Context支持** - 完整的 context.Context 支持，支持请求取消和超时
- 🛡️ **错误处理** - 完善的错误类型和判断函数
- 🔄 **重试机制** - 可配置的指数退避重试策略
- 🔌 **拦截器** - 灵活的请求/响应拦截器链
- ⚡ **性能优化** - sync.Pool 复用 buffer，减少内存分配
- 🧵 **并发安全** - RWMutex 保护，线程安全
- 📝 **日志系统** - 可配置的日志接口
- 🏗️ **模块化设计** - 清晰的分层架构
- 🧩 **单例模式** - 服务实例缓存，减少资源消耗

## 📦 功能模块

| 模块 | 功能 |
|------|------|
| **Geocode** | 地理编码与逆地理编码 |
| **Direction** | 驾车/步行路径规划 |
| **Place** | POI关键字搜索、周边搜索、多边形搜索 |
| **Weather** | 实况天气与天气预报 |
| **IP** | IP地址定位 |

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
    "log"
    "github.com/zhihao0924/amapSdk"
)

func main() {
    // 基础配置
    client, err := amap.NewClient(&amap.Config{
        Key: "YOUR_AMAP_KEY", // 你的高德地图API Key
    })
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
}
```

### 3. 完整配置

```go
import (
    "time"
    "github.com/zhihao0924/amapSdk"
)

client, err := amap.NewClient(&amap.Config{
    Key:    "YOUR_AMAP_KEY",
    BaseURL: "https://restapi.amap.com/v3", // API基础URL（可选）
    Timeout: 10,                               // 请求超时时间（秒）
    Debug:   true,                              // 开启调试日志

    // 自定义请求头
    Headers: map[string]string{
        "User-Agent": "MyApp/1.0",
        "X-Request-ID": generateRequestID(),
    },

    // 重试配置
    RetryConfig: &amap.RetryConfig{
        MaxRetries: 3,
        RetryDelay: time.Second,
    },

    // 拦截器链
    InterceptorChain: amap.NewInterceptorChain(),
})
```

## 📖 使用示例

### 地理编码

将地址转换为经纬度坐标：

```go
// 地理编码
resp, err := client.Geocode().Geo(&amap.GeocodeOptions{
    Address: "北京市朝阳区阜通东大街6号",
    City:    "北京",
})
if err != nil {
    log.Fatal(err)
}

// 输出结果
for _, geo := range resp.Geocodes {
    fmt.Printf("地址: %s\n", geo.FormattedAddress)
    fmt.Printf("坐标: %s\n", geo.Location)
}
```

逆地理编码（坐标转地址）：

```go
// 逆地理编码
resp, err := client.Geocode().ReGeo(&amap.ReGeoOptions{
    Location:   "116.480881,39.989410",
    Radius:     "1000",
    Extensions: "all",
})
if err != nil {
    log.Fatal(err)
}

// 输出结果
if resp.Regeocode != nil {
    fmt.Printf("地址: %s\n", resp.Regeocode.FormattedAddress)
    fmt.Printf("省份: %s\n", resp.Regeocode.AddressComponent.Province)
    fmt.Printf("城市: %s\n", resp.Regeocode.AddressComponent.City)
}
```

### 路径规划

驾车路径规划：

```go
// 使用选项
resp, err := client.Direction().Driving(&amap.DrivingOptions{
    Origin:      "116.481028,39.989643",
    Destination: "116.465302,40.004717",
    Strategy:    "1", // 策略：1=速度优先
})

// 使用枚举类型
resp, err := client.Direction().Driving(&amap.DrivingOptions{
    Origin:      "116.481028,39.989643",
    Destination: "116.465302,40.004717",
    Strategy:    amap.DrivingStrategyFastest.String(),
})

if err != nil {
    log.Fatal(err)
}

// 输出结果
if len(resp.Route.Paths) > 0 {
    path := resp.Route.Paths[0]
    fmt.Printf("距离: %s 米\n", path.Distance)
    fmt.Printf("时长: %s 秒\n", path.Duration)
}
```

步行路径规划：

```go
resp, err := client.Direction().Walking(&amap.WalkingOptions{
    Origin:      "116.481028,39.989643",
    Destination: "116.465302,40.004717",
})
if err != nil {
    log.Fatal(err)
}
```

### POI 搜索

关键字搜索：

```go
resp, err := client.Place().TextSearch(&amap.TextSearchOptions{
    Keywords:   "肯德基",
    City:       "北京",
    CityLimit:  "1",
    Offset:     "20",
    Extensions: "all",
})
if err != nil {
    log.Fatal(err)
}

// 输出结果
for _, poi := range resp.Pois {
    fmt.Printf("名称: %s\n", poi.Name)
    fmt.Printf("地址: %s\n", poi.Address)
    fmt.Printf("电话: %s\n", poi.Tel)
}
```

周边搜索：

```go
resp, err := client.Place().AroundSearch(&amap.AroundSearchOptions{
    Location:   "116.481028,39.989643",
    Keywords:   "餐厅",
    Radius:     "1000",
    Extensions: "all",
})
if err != nil {
    log.Fatal(err)
}
```

### 天气查询

实时天气：

```go
live, err := client.Weather().Base("110101")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("天气: %s\n", live.Weather)
fmt.Printf("温度: %s°C\n", live.Temperature)
fmt.Printf("湿度: %s%%\n", live.Humidity)
```

天气预报：

```go
// 查询实时天气
resp, err := client.Weather().Query(&amap.WeatherOptions{
    City:       "110101",
    Extensions: "base", // 或 "all" 获取预报
})

// 获取未来N天天气
forecast, err := client.Weather().GetNextDaysWeather("110101", 4)
if err != nil {
    log.Fatal(err)
}

for _, cast := range forecast.Casts {
    fmt.Printf("%s: %s, %s°C\n",
        cast.Date, cast.DayWeather, cast.DayTemp)
}
```

### IP 定位

```go
// 查询指定IP
resp, err := client.IP().Location(&amap.LocationOptions{
    IP: "8.8.8.8",
})
if err != nil {
    log.Fatal(err)
}

// 查询当前IP
resp, err := client.IP().Current()

// 批量IP定位
ips := []string{"8.8.8.8", "114.114.114.114"}
responses, err := client.IP().BatchLocation(ips)
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

### RetryConfig 重试配置

```go
type RetryConfig struct {
    MaxRetries  int                   // 最大重试次数（默认3）
    RetryDelay time.Duration          // 重试延迟（默认1秒）
    Retryable  func(error) bool      // 自定义重试判断函数
}
```

### 枚举类型

SDK 提供类型安全的枚举：

```go
// 驾车策略
amap.DrivingStrategyFastest     // 速度优先
amap.DrivingStrategyCost         // 费用优先
amap.DrivingStrategyDistance     // 距离优先
amap.DrivingStrategyAvoidCongestion // 避免拥堵

// 地理编码扩展
amap.GeocodeExtensionBase       // 基本信息
amap.GeocodeExtensionAll        // 所有信息

// 天气类型
amap.WeatherTypeLive           // 实时天气
amap.WeatherTypeForecast       // 天气预报

// POI 排序
amap.SortTypeDistance          // 按距离排序
amap.SortTypeWeight           // 按权重排序
```

## 🛠️ 错误处理

SDK 提供完善的错误处理机制：

### 基础错误处理

```go
resp, err := client.Geocode().Geo(&amap.GeocodeOptions{
    Address: "不存在的地址",
})
if err != nil {
    log.Fatal(err)
}
```

### 错误类型判断

```go
if amap.IsNetworkError(err) {
    // 处理网络错误
    log.Println("网络连接失败，请检查网络")
} else if amap.IsTimeoutError(err) {
    // 处理超时错误
    log.Println("请求超时，请稍后重试")
} else if amap.IsRateLimitError(err) {
    // 处理频率限制错误
    log.Println("请求过于频繁，请降低调用频率")
} else if amap.IsAuthError(err) {
    // 处理认证错误
    log.Println("API Key 无效或已过期")
} else if amap.IsAPIError(err) {
    // 处理 API 错误
    if apiErr, ok := err.(*amap.Error); ok {
        msg := amap.GetAPIErrorMessage(apiErr.Code)
        log.Printf("API错误: %s (代码: %s)", msg, apiErr.Code)
    }
}
```

### 错误信息查询

```go
// 根据错误码获取中文错误信息
msg := amap.GetAPIErrorMessage("10001")
fmt.Println(msg) // 输出: key不正确或过期
```

## 🎨 高级特性

### 自定义拦截器

```go
import "net/http"

// 创建拦截器链
interceptor := amap.NewInterceptorChain()

// 添加请求拦截器
interceptor.AddRequest(func(req *http.Request) error {
    req.Header.Set("X-Request-ID", generateRequestID())
    req.Header.Set("X-Request-Time", time.Now().Format(time.RFC3339))
    return nil
})

// 添加响应拦截器
interceptor.AddResponse(func(resp *http.Response) error {
    fmt.Printf("响应状态: %d\n", resp.StatusCode)
    return nil
})

// 使用拦截器
client, err := amap.NewClient(&amap.Config{
    Key:            "YOUR_AMAP_KEY",
    InterceptorChain: interceptor,
})
```

### 自定义重试策略

```go
// 使用默认重试配置
client, err := amap.NewClient(&amap.Config{
    Key:   "YOUR_AMAP_KEY",
    RetryConfig: amap.DefaultRetryConfig, // 重试3次，每次1秒
})

// 自定义重试配置
retryConfig := amap.NewRetryConfig(5, 2*time.Second).
    WithMaxRetries(6).
    WithRetryDelay(time.Second * 2).
    WithRetryable(func(err error) bool {
        // 只对网络错误重试
        return amap.IsNetworkError(err)
    })

client, err := amap.NewClient(&amap.Config{
    Key:         "YOUR_AMAP_KEY",
    RetryConfig: retryConfig,
})
```

### Context 超时控制

```go
import "context"

// 创建带超时的 context
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// SDK 内部支持 context，可以随时取消请求
resp, err := client.Geocode().Geo(&amap.GeocodeOptions{
    Address: "北京市朝阳区",
})
```

## 📁 项目结构

```
amap-sdk/
├── client.go              # 客户端主入口
├── options.go             # 选项构建器
├── errors_export.go       # 错误处理导出
├── go.mod                # Go 模块文件
├── go.sum                # 依赖锁定文件
│
├── pkg/                  # 内部包
│   ├── common/           # 通用工具层
│   │   ├── errors.go    # 错误定义和处理
│   │   ├── location.go  # 位置工具
│   │   └── logger.go    # 日志接口
│   │
│   ├── core/            # 核心功能层
│   │   ├── client.go    # 核心客户端
│   │   ├── config.go    # 配置管理
│   │   ├── enums.go     # 类型安全枚举
│   │   ├── http.go      # HTTP客户端（支持context）
│   │   ├── middleware.go # 拦截器链
│   │   └── retry.go     # 重试机制
│   │
│   ├── models/          # 数据模型层
│   │   ├── base.go      # 基础响应模型
│   │   ├── direction.go # 路径规划模型
│   │   ├── geocode.go  # 地理编码模型
│   │   ├── ip.go       # IP定位模型
│   │   ├── place.go     # POI搜索模型
│   │   └── weather.go   # 天气查询模型
│   │
│   └── services/        # 服务实现层
│       ├── direction/   # 路径规划服务
│       ├── geocode/    # 地理编码服务
│       ├── ip/         # IP定位服务
│       ├── place/       # POI搜索服务
│       └── weather/     # 天气查询服务
│
└── examples/             # 使用示例
    ├── README.md          # 示例说明文档
    ├── basic/            # 基础示例
    ├── geocode/          # 地理编码示例
    ├── direction/        # 路径规划示例
    ├── place/            # POI搜索示例
    ├── weather/          # 天气查询示例
    └── ip/               # IP定位示例
```

## 📚 更多示例

详细的使用示例请查看 `examples/` 目录：

```bash
# 运行基础示例
cd examples/basic
go run basic.go

# 运行地理编码示例
cd examples/geocode
go run geocode.go

# 运行路径规划示例
cd examples/direction
go run driving.go

# 运行POI搜索示例
cd examples/place
go run text_search.go
```

## 🔗 相关链接

- [高德开放平台](https://lbs.amap.com/)
- [高德地图 Web 服务 API 文档](https://lbs.amap.com/api/webservice/summary)
- [获取 API Key](https://console.amap.com/dev/key/app)

## 📄 许可证

[MIT License](LICENSE)

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

### 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

## 📞 联系方式

- Issues: [GitHub Issues](https://github.com/zhihao0924/amapSdk/issues)

---

**Made with ❤️ by zhihao**
