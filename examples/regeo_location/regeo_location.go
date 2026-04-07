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

	loc := amap.NewLocation(116.397428, 39.90923)
	resp, err := client.Geocode().ReGeoByLocation(ctx, loc, "all")
	if err != nil {
		log.Fatalf("ReGeoByLocation 失败: %v", err)
	}

	if resp.Regeocode == nil {
		fmt.Println("未返回逆地理编码结果")
		return
	}

	fmt.Printf("ReGeoByLocation 结果:\n")
	fmt.Printf("地址: %s\n", resp.Regeocode.FormattedAddress)
	fmt.Printf("区县: %s\n", resp.Regeocode.AddressComponent.District)
	if len(resp.Regeocode.Pois) > 0 {
		fmt.Printf("附近POI: %s\n", resp.Regeocode.Pois[0].Name)
	}
}
