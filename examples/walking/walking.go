package main

import (
	"fmt"
	"log"

	"github.com/zhihao0924/amapSdk"
	"github.com/zhihao0924/amapSdk/examples/internal/exampleutil"
)

func main() {
	client, err := exampleutil.NewClient()
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}
	defer client.Close()

	ctx, cancel := exampleutil.NewRequestContext()
	defer cancel()

	resp, err := client.Direction().Walking(ctx, &amap.WalkingOptions{
		Origin:      "116.397428,39.90923",
		Destination: "116.404006,39.915119",
	})
	if err != nil {
		log.Fatalf("步行路径规划失败: %v", err)
	}

	if len(resp.Route.Paths) == 0 {
		fmt.Println("Walking 未返回路径")
		return
	}

	path := resp.Route.Paths[0]
	fmt.Printf("步行路径规划结果:\n")
	fmt.Printf("距离: %s 米\n", path.Distance)
	fmt.Printf("时长: %s 秒\n", path.Duration)
	fmt.Printf("步骤数: %d\n", len(path.Steps))

	origin := amap.NewLocation(116.397428, 39.90923)
	destination := amap.NewLocation(116.404006, 39.915119)
	helperResp, err := client.Direction().WalkingByLocations(ctx, origin, destination)
	if err != nil {
		log.Fatalf("WalkingByLocations 失败: %v", err)
	}

	if len(helperResp.Route.Paths) > 0 {
		fmt.Printf("WalkingByLocations 距离: %s 米\n", helperResp.Route.Paths[0].Distance)
	}
}
