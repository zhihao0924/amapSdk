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

	resp, err := client.Geocode().ReGeo(ctx, &amap.ReGeoOptions{
		Location:   "116.397428,39.90923",
		Extensions: "all",
	})
	if err != nil {
		log.Fatalf("逆地理编码失败: %v", err)
	}

	if resp.Regeocode == nil {
		fmt.Println("未返回逆地理编码结果")
		return
	}

	fmt.Printf("逆地理编码结果:\n")
	fmt.Printf("格式化地址: %s\n", resp.Regeocode.FormattedAddress)
	fmt.Printf("省份: %s\n", resp.Regeocode.AddressComponent.Province)
	fmt.Printf("城市: %s\n", resp.Regeocode.AddressComponent.City)
	fmt.Printf("区县: %s\n", resp.Regeocode.AddressComponent.District)
	fmt.Printf("区域编码: %s\n", resp.Regeocode.AddressComponent.Adcode)

	if len(resp.Regeocode.Pois) > 0 {
		poi := resp.Regeocode.Pois[0]
		fmt.Printf("附近POI: %s\n", poi.Name)
		fmt.Printf("POI地址: %s\n", poi.Address)
	}
}
