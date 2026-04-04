package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zhihao0924/amapSdk"
)

func main() {
	// 创建客户端
	client, err := amap.NewClient(&amap.Config{
		Key:     "YOUR_API_KEY", // 替换为你的高德地图API Key
		Debug:   true,
		Timeout: 10,
	})
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}
	defer client.Close()

	// 创建地理编码选项
	geoOpts := &amap.GeocodeOptions{
		Address: "北京市朝阳区",
		City:    "北京",
	}

	// 调用地理编码服务
	resp, err := client.Geocode().Geo(context.Background(), geoOpts)
	if err != nil {
		log.Fatalf("地理编码失败: %v", err)
	}

	// 检查响应状态
	if resp.Status != "1" {
		log.Fatalf("API错误: %s (%s)", resp.Info, resp.Infocode)
	}

	// 输出结果
	fmt.Printf("地理编码结果:\n")
	fmt.Printf("状态: %s\n", resp.Status)
	fmt.Printf("信息: %s\n", resp.Info)

	if len(resp.Geocodes) > 0 {
		geo := resp.Geocodes[0]
		fmt.Printf("地址: %s\n", geo.FormattedAddress)
		fmt.Printf("经度: %s\n", geo.Location)
		fmt.Printf("级别: %s\n", geo.Level)
		fmt.Printf("城市: %s\n", geo.City)
		fmt.Printf("省份: %s\n", geo.Province)
	}
}
