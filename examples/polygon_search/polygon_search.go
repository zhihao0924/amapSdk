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

	resp, err := client.Place().SearchByPolygon(ctx, &amap.SearchByPolygonOptions{
		Polygon:    "116.3908,39.9074|116.4048,39.9074|116.4048,39.9165|116.3908,39.9165",
		Keywords:   "景点",
		Extensions: "all",
		Offset:     "10",
		Page:       "1",
	})
	if err != nil {
		log.Fatalf("SearchByPolygon 失败: %v", err)
	}

	fmt.Printf("多边形搜索结果:\n")
	fmt.Printf("总数量: %d\n", resp.Count)
	if len(resp.Pois) == 0 {
		fmt.Println("未返回 POI")
		return
	}

	limit := min(len(resp.Pois), 5)
	for i := 0; i < limit; i++ {
		poi := resp.Pois[i]
		fmt.Printf("%d. %s | %s | %s\n", i+1, poi.Name, poi.Address, poi.Location)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
