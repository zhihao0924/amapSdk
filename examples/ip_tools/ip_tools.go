package main

import (
	"fmt"
	"log"

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

	current, err := client.IP().Current(ctx)
	if err != nil {
		log.Fatalf("Current 失败: %v", err)
	}
	fmt.Printf("当前出口 IP 定位: %s %s (%s)\n", current.Province, current.City, current.Adcode)

	info, err := client.IP().GetIPInfo(ctx, "114.247.50.2")
	if err != nil {
		log.Fatalf("GetIPInfo 失败: %v", err)
	}
	fmt.Printf("114.247.50.2 定位: %s %s (%s)\n", info.Province, info.City, info.Adcode)

	batch, err := client.IP().BatchLocation(ctx, []string{"114.247.50.2"})
	if err != nil {
		log.Fatalf("BatchLocation 失败: %v", err)
	}
	fmt.Printf("批量 IP 定位结果:\n")
	for i, item := range batch {
		fmt.Printf("%d. %s -> %s %s (%s)\n", i+1, item.IP, item.Province, item.City, item.Adcode)
	}
}
