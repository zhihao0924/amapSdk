package main

import (
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

	// 创建IP定位选项
	ipOpts := &amap.LocationOptions{
		IP: "114.247.50.2", // 可以指定IP地址，为空则查询当前IP
	}

	// 调用IP定位服务
	resp, err := client.IP().Location(ipOpts)
	if err != nil {
		log.Fatalf("IP定位失败: %v", err)
	}

	// 检查响应状态
	if resp.Status != "1" {
		log.Fatalf("API错误: %s (%s)", resp.Info, resp.Infocode)
	}

	// 输出结果
	fmt.Printf("IP定位结果:\n")
	fmt.Printf("状态: %s\n", resp.Status)
	fmt.Printf("信息: %s\n", resp.Info)

	if resp.IP != "" {
		fmt.Printf("查询IP: %s\n", resp.IP)
	}

	if resp.Province != "" {
		fmt.Printf("省份: %s\n", resp.Province)
	}

	if resp.City != "" {
		fmt.Printf("城市: %s\n", resp.City)
	}

	if resp.Adcode != "" {
		fmt.Printf("区域代码: %s\n", resp.Adcode)
	}

	if resp.Rectangle != "" {
		fmt.Printf("矩形范围: %s\n", resp.Rectangle)
	}

	fmt.Printf("\n定位信息: %s省%s市\n", resp.Province, resp.City)
}
