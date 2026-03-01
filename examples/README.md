# 高德地图 Go SDK 示例

本目录包含高德地图 Go SDK 的各种使用示例。

## 目录结构

```
examples/
├── README.md                  # 本文件
├── basic/                     # 基础示例
│   └── basic.go              # 客户端初始化和基础用法
├── geocode/                   # 地理编码示例
│   ├── geocode.go            # 地理编码（地址转坐标）
│   └── regeocode.go          # 逆地理编码（坐标转地址）
├── direction/                 # 路径规划示例
│   ├── driving.go            # 驾车路径规划
│   └── walking.go            # 步行路径规划
├── place/                     # POI搜索示例
│   ├── text_search.go        # 关键字搜索
│   └── around_search.go      # 周边搜索
├── weather/                   # 天气查询示例
│   └── weather.go            # 实时天气和天气预报
├── ip/                        # IP定位示例
│   └── ip.go                 # IP定位服务
└── advanced/                  # 高级特性示例
    ├── error_handling.go     # 错误处理
    ├── interceptor.go        # 拦截器使用
    └── retry.go              # 重试机制
```

## 快速开始

### 1. 替换 API Key

在运行示例之前，请将代码中的 `YOUR_API_KEY` 替换为你的高德地图 API Key。

获取 API Key 的步骤：
1. 访问 [高德开放平台](https://lbs.amap.com/)
2. 注册并登录账号
3. 进入控制台，创建应用
4. 获取 Web 服务 API Key

### 2. 运行示例

```bash
# 进入示例目录
cd examples/basic

# 运行基础示例
go run basic.go

# 或者先编译再运行
go build -o basic
./basic
```

## 示例说明

### 基础示例

- **basic.go**: 演示如何创建客户端、获取配置信息、使用工具函数

### 地理编码

- **geocode.go**: 将地址转换为经纬度坐标
  ```go
  opts := options.NewGeocodeOptionsBuilder().
      SetAddress("北京市朝阳区").
      SetCity("北京").
      Build()
  resp, err := client.Geocode().Get(opts)
  ```

- **regeocode.go**: 将经纬度坐标转换为地址信息
  ```go
  opts := options.NewReGeoOptionsBuilder().
      SetLocation(116.397428, 39.90923).
      SetRadius("1000").
      Build()
  resp, err := client.Geocode().ReGet(opts)
  ```

### 路径规划

- **driving.go**: 驾车路径规划
  ```go
  opts := options.NewDrivingOptionsBuilder().
      SetOrigin(116.481181, 39.989792).
      SetDestination(116.397428, 39.90923).
      SetStrategy(int(amap.DrivingStrategyFastest)).
      Build()
  resp, err := client.Direction().Driving(opts)
  ```

- **walking.go**: 步行路径规划
  ```go
  opts := options.NewWalkingOptionsBuilder().
      SetOrigin(116.481181, 39.989792).
      SetDestination(116.397428, 39.90923).
      Build()
  resp, err := client.Direction().Walking(opts)
  ```

### POI搜索

- **text_search.go**: 关键字搜索POI
  ```go
  opts := &amap.TextSearchOptions{
      Keywords:  "肯德基",
      City:      "北京",
      CityLimit: true,
  }
  resp, err := client.Place().TextSearch(opts)
  ```

- **around_search.go**: 周边搜索POI
  ```go
  opts := &amap.AroundSearchOptions{
      Keywords: "美食",
      Location: "116.397428,39.90923",
      Radius:   "1000",
      SortType: "distance",
  }
  resp, err := client.Place().AroundSearch(opts)
  ```

### 天气查询

- **weather.go**: 查询实时天气和天气预报
  ```go
  // 实时天气
  opts := &amap.WeatherOptions{
      City:       "110101",
      Extensions: "base",
  }
  resp, err := client.Weather().Query(opts)

  // 天气预报
  opts.Extensions = "all"
  resp, err = client.Weather().Query(opts)
  ```

### IP定位

- **ip.go**: 根据IP地址获取位置信息
  ```go
  opts := &amap.LocationOptions{
      IP: "114.247.50.2",
  }
  resp, err := client.IP().Location(opts)
  ```

### 高级特性

- **error_handling.go**: 错误处理最佳实践
  - 基本错误处理
  - 错误类型判断（网络错误、超时错误等）
  - 错误解包
  - 超时处理
  - API错误消息获取

- **interceptor.go**: 使用拦截器
  - 日志拦截器
  - 请求头拦截器
  - 自定义拦截器
  - 多个拦截器组合

- **retry.go**: 重试机制
  - 使用默认重试配置
  - 自定义重试配置
  - 自定义重试判断函数
  - 重试配置构建器

## 配置选项

### 基本配置

```go
config := &amap.Config{
    Key:     "YOUR_API_KEY",     // 必需：API Key
    BaseURL: "https://restapi.amap.com/v3", // 可选：API基础URL
    Timeout: 10,                  // 可选：请求超时时间（秒）
    Debug:   true,                // 可选：是否开启调试日志
}
```

### 高级配置

```go
config := &amap.Config{
    Key:     "YOUR_API_KEY",
    Debug:   true,
    Timeout: 10,

    // 自定义请求头
    Headers: map[string]string{
        "User-Agent": "MyApp/1.0",
    },

    // 重试配置
    RetryConfig: &amap.RetryConfig{
        MaxRetries: 3,
        RetryDelay: 1 * time.Second,
    },

    // 拦截器链
    InterceptorChain: amap.NewInterceptorChain(),
}
```

## 错误处理

SDK 提供了丰富的错误处理机制：

```go
// 判断错误类型
if amap.IsNetworkError(err) {
    // 处理网络错误
} else if amap.IsTimeoutError(err) {
    // 处理超时错误
} else if amap.IsAuthError(err) {
    // 处理认证错误
}

// 获取API错误消息
if amap.IsAPIError(err) {
    // 获取错误代码
    if apiErr, ok := err.(*amap.Error); ok {
        message := amap.GetAPIErrorMessage(apiErr.Code)
        fmt.Println("API错误:", message)
    }
}

// 错误包装
err = amap.WrapError(err, "操作失败")

// 错误解包
originalErr := amap.UnwrapError(err)
```

## 枚举类型

SDK 提供类型安全的枚举：

```go
// 驾车策略
strategy := amap.DrivingStrategyFastest     // 速度优先
strategy = amap.DrivingStrategyCost         // 费用优先
strategy = amap.DrivingStrategyAvoidCongestion // 避免拥堵

// 地理编码扩展
extension := amap.GeocodeExtensionBase      // 基本信息扩展
extension = amap.GeocodeExtensionAll         // 所有信息扩展

// 天气类型
weatherType := amap.WeatherTypeLive          // 实时天气
weatherType = amap.WeatherTypeForecast       // 天气预报

// 排序类型
sortType := amap.SortTypeDistance            // 按距离排序
sortType = amap.SortTypeWeight               // 按权重排序
```

## 注意事项

1. **API Key 安全**: 请勿将 API Key 提交到公共代码仓库
2. **请求频率**: 注意高德API的QPS限制
3. **错误处理**: 所有API调用都应该进行错误处理
4. **资源释放**: 使用完毕后记得调用 `client.Close()` 关闭客户端
5. **并发安全**: 客户端是并发安全的，可以在多个goroutine中共享使用

## 更多信息

- [高德开放平台](https://lbs.amap.com/)
- [高德地图Web服务API文档](https://lbs.amap.com/api/webservice/summary)
- [Go SDK 源码](https://github.com/zhihao0924/amapSdk)

## 问题反馈

如有问题或建议，请提交 [Issue](https://github.com/zhihao0924/amapSdk/issues)
