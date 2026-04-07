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

	resp, err := client.Place().AroundSearch(ctx, &amap.AroundSearchOptions{
		Location:   "116.397428,39.90923",
		Keywords:   "餐厅",
		Radius:     "1000",
		Sortrule:   "distance",
		Extensions: "all",
	})
	if err != nil {
		log.Fatalf("AroundSearch 失败: %v", err)
	}

	fmt.Printf("周边搜索结果:\n")
	fmt.Printf("总数量: %d\n", resp.Count)
	printPois(resp.Pois)

	helperResp, err := client.Place().AroundSearchByLocation(
		ctx,
		amap.NewLocation(116.397428, 39.90923),
		"咖啡",
		800,
	)
	if err != nil {
		log.Fatalf("AroundSearchByLocation 失败: %v", err)
	}

	fmt.Printf("\nAroundSearchByLocation 总数量: %d\n", helperResp.Count)
	printPois(helperResp.Pois)
}

func printPois(pois []amap.PoiInfo) {
	if len(pois) == 0 {
		fmt.Println("未返回 POI")
		return
	}

	limit := min(len(pois), 3)
	for i := 0; i < limit; i++ {
		poi := pois[i]
		fmt.Printf("%d. %s | %s | %s\n", i+1, poi.Name, poi.Address, poi.Location)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
