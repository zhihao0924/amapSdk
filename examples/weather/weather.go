package main

import (
	"fmt"
	"log"

	"github.com/zhihao0924/amapSdk/examples/internal/exampleutil"
)

func main() {
	// 创建客户端
	client, err := exampleutil.NewClient()
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}
	defer client.Close()
	ctx, cancel := exampleutil.NewRequestContext()
	defer cancel()

	// 调用天气查询服务（实时天气）
	liveResp, err := client.Weather().Base(ctx, "110101")
	if err != nil {
		log.Fatalf("天气查询失败: %v", err)
	}

	// 输出实时天气结果
	fmt.Printf("实时天气查询结果:\n")
	fmt.Printf("  城市: %s\n", liveResp.City)
	fmt.Printf("  天气: %s\n", liveResp.Weather)
	fmt.Printf("  温度: %s°C\n", liveResp.Temperature)
	fmt.Printf("  风向: %s\n", liveResp.WindDirection)
	fmt.Printf("  风力: %s级\n", liveResp.WindPower)
	fmt.Printf("  湿度: %s%%\n", liveResp.Humidity)
	fmt.Printf("  更新时间: %s\n", liveResp.ReportTime)

	fmt.Println("========================================")

	// 调用天气查询服务（预报天气）
	forecastResp, err := client.Weather().Forecast(ctx, "110101")
	if err != nil {
		log.Fatalf("天气查询失败: %v", err)
	}

	// 输出天气预报结果
	fmt.Printf("天气预报查询结果:\n")
	fmt.Printf("  城市: %s\n", forecastResp.City)
	fmt.Printf("  省份: %s\n", forecastResp.Province)
	fmt.Printf("  报告时间: %s\n", forecastResp.ReportTime)
	fmt.Printf("\n未来%d天预报:\n", len(forecastResp.Casts))

	for i, cast := range forecastResp.Casts {
		fmt.Printf("  %d. %s (%s)\n", i+1, cast.Week, cast.Date)
		fmt.Printf("     白天: %s, %s°C\n", cast.DayWeather, cast.DayTemp)
		fmt.Printf("     夜间: %s, %s°C\n", cast.NightWeather, cast.NightTemp)
		fmt.Printf("     白天风: %s, %s级\n", cast.DayWind, cast.DayPower)
		fmt.Printf("     夜间风: %s, %s级\n", cast.NightWind, cast.NightPower)
	}
}
