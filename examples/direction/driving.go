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

	// 创建驾车路径规划选项
	drivingOpts := &amap.DrivingOptions{
		Origin:      "116.481181,39.989792", // 起点：北京站
		Destination: "116.397428,39.90923",  // 终点：天安门
		Strategy:    "1",                    // 策略：速度优先 (对应 DrivingStrategyFastest)
		Waypoints:   "116.407526,39.904030", // 途经点1：故宫
	}

	// 调用驾车路径规划服务
	resp, err := client.Direction().Driving(context.Background(), drivingOpts)
	if err != nil {
		log.Fatalf("驾车路径规划失败: %v", err)
	}

	// 检查响应状态
	if resp.Status != "1" {
		log.Fatalf("API错误: %s (%s)", resp.Info, resp.Infocode)
	}

	// 输出结果
	fmt.Printf("驾车路径规划结果:\n")
	fmt.Printf("状态: %s\n", resp.Status)
	fmt.Printf("信息: %s\n", resp.Info)

	if len(resp.Route.Paths) > 0 {
		path := resp.Route.Paths[0]
		fmt.Printf("\n路线信息:\n")
		fmt.Printf("  距离: %s 米\n", path.Distance)
		fmt.Printf("  时长: %s 秒\n", path.Duration)
		fmt.Printf("  步骤数: %d\n", len(path.Steps))
		fmt.Printf("  策略: %s\n", path.Strategy)

		// 输出出租车费用信息
		if resp.Route.TaxiCost != "" {
			fmt.Printf("  出租车费用: %s 元\n", resp.Route.TaxiCost)
		}

		// 输出前3个步骤
		fmt.Printf("\n前3个行驶步骤:\n")
		for i, step := range path.Steps {
			if i >= 3 {
				break
			}
			fmt.Printf("  %d. %s\n", i+1, step.Instruction)
			fmt.Printf("     距离: %s 米, 耗时: %s 秒\n", step.Distance, step.Duration)
		}
	}
}
