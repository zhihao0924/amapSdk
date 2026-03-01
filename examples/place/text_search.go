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

	// 创建关键字搜索选项
	searchOpts := &amap.TextSearchOptions{
		Keywords:   "肯德基",
		City:       "北京",
		CityLimit:  "1",
		Children:   "1",
		Offset:     "20",
		Page:       "1",
		Extensions: "all",
	}

	// 调用POI搜索服务
	resp, err := client.Place().TextSearch(searchOpts)
	if err != nil {
		log.Fatalf("POI搜索失败: %v", err)
	}

	// 检查响应状态
	if resp.Status != "1" {
		log.Fatalf("API错误: %s (%s)", resp.Info, resp.Infocode)
	}

	// 输出结果
	fmt.Printf("POI关键字搜索结果:\n")
	fmt.Printf("状态: %s\n", resp.Status)
	fmt.Printf("信息: %s\n", resp.Info)
	fmt.Printf("总数量: %d\n", resp.Count)

	// 输出POI信息
	if len(resp.Pois) > 0 {
		fmt.Printf("\n找到 %d 个POI:\n", len(resp.Pois))
		for i, poi := range resp.Pois {
			if i >= 5 { // 只显示前5个
				break
			}
			fmt.Printf("\n%d. %s\n", i+1, poi.Name)
			fmt.Printf("   地址: %s\n", poi.Address)
			fmt.Printf("   电话: %s\n", poi.Tel)
			fmt.Printf("   类型: %s\n", poi.Type)
			fmt.Printf("   距离: %.0f 米\n", poi.Distance)

			// 输出位置信息
			if poi.Location != "" {
				fmt.Printf("   坐标: %s\n", poi.Location)
			}
		}
	}
}
