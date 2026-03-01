package main

import (
	"fmt"
	"log"

	"github.com/zhihao0924/amapSdk"
)

func main() {
	// 创建客户端配置
	config := &amap.Config{
		Key:     "YOUR_API_KEY", // 替换为你的高德地图API Key
		Debug:   true,           // 开启调试日志
		Timeout: 10,             // 请求超时时间（秒）
	}

	// 创建客户端
	client, err := amap.NewClient(config)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}
	defer client.Close()

	fmt.Printf("客户端初始化成功: %s\n", client)

	// 获取配置信息
	configInfo := client.GetConfig()
	fmt.Printf("BaseURL: %s, Timeout: %d秒\n", configInfo.BaseURL, configInfo.Timeout)

	// 获取日志实例
	logger := client.GetLogger()
	logger.Info("这是一个示例消息")

	// 检查客户端是否已关闭
	if client.IsClosed() {
		fmt.Println("客户端已关闭")
	} else {
		fmt.Println("客户端运行中")
	}

	// 使用构建器模式创建选项
	location := amap.NewLocation(116.397428, 39.90923)
	fmt.Printf("创建位置: 经度=%.6f, 纬度=%.6f\n", location.Lng, location.Lat)

	// 创建错误
	errMsg := amap.NewError("TEST_ERROR", "这是一个测试错误")
	fmt.Printf("创建错误: %v\n", errMsg)
}
