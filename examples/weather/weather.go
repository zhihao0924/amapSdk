package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zhihao0924/amapSdk"
	"github.com/zhihao0924/amapSdk/pkg/core"
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

	// 创建天气查询选项（实时天气）
	liveOpts := &amap.WeatherOptions{
		City:       "110101", // 城市编码，北京东城区
		Extensions: core.WeatherTypeLive.String(),
	}

	// 调用天气查询服务（实时天气）
	liveResp, err := client.Weather().Query(context.Background(), liveOpts)
	if err != nil {
		log.Fatalf("天气查询失败: %v", err)
	}

	// 检查响应状态
	if liveResp.Status != "1" {
		log.Fatalf("API错误: %s (%s)", liveResp.Info, liveResp.Infocode)
	}

	// 输出实时天气结果
	fmt.Printf("实时天气查询结果:\n")
	fmt.Printf("状态: %s\n", liveResp.Status)
	fmt.Printf("信息: %s\n", liveResp.Info)
	fmt.Printf("城市: %s\n", liveResp.Lives[0].City)

	if len(liveResp.Lives) > 0 {
		live := liveResp.Lives[0]
		fmt.Printf("\n实时天气:\n")
		fmt.Printf("  城市: %s\n", live.City)
		fmt.Printf("  天气: %s\n", live.Weather)
		fmt.Printf("  温度: %s°C\n", live.Temperature)
		fmt.Printf("  风向: %s\n", live.WindDirection)
		fmt.Printf("  风力: %s级\n", live.WindPower)
		fmt.Printf("  湿度: %s%%\n", live.Humidity)
		fmt.Printf("  更新时间: %s\n", live.ReportTime)
	}

	fmt.Println("========================================")

	// 创建天气查询选项（预报天气）
	forecastOpts := &amap.WeatherOptions{
		City:       "110101",
		Extensions: core.WeatherTypeForecast.String(),
	}

	// 调用天气查询服务（预报天气）
	forecastResp, err := client.Weather().Query(context.Background(), forecastOpts)
	if err != nil {
		log.Fatalf("天气查询失败: %v", err)
	}

	// 检查响应状态
	if forecastResp.Status != "1" {
		log.Fatalf("API错误: %s (%s)", forecastResp.Info, forecastResp.Infocode)
	}

	// 输出天气预报结果
	fmt.Printf("天气预报查询结果:\n")
	fmt.Printf("状态: %s\n", forecastResp.Status)
	fmt.Printf("信息: %s\n", forecastResp.Info)
	fmt.Printf("城市: %s\n", forecastResp.Forecasts[0].City)

	if len(forecastResp.Forecasts) > 0 {
		forecast := forecastResp.Forecasts[0]
		fmt.Printf("\n预报天气:\n")
		fmt.Printf("  城市: %s\n", forecast.City)
		fmt.Printf("  省份: %s\n", forecast.Province)
		fmt.Printf("  报告时间: %s\n", forecast.ReportTime)
		fmt.Printf("\n未来4天预报:\n")

		for i, cast := range forecast.Casts {
			fmt.Printf("  %d. %s (%s)\n", i+1, cast.Week, cast.Date)
			fmt.Printf("     白天: %s, %s°C\n", cast.DayWeather, cast.DayTemp)
			fmt.Printf("     夜间: %s, %s°C\n", cast.NightWeather, cast.NightTemp)
			fmt.Printf("     白天风: %s, %s级\n", cast.DayWind, cast.DayPower)
			fmt.Printf("     夜间风: %s, %s级\n", cast.NightWind, cast.NightPower)
		}
	}
}
