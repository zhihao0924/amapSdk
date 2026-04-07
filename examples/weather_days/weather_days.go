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

	tomorrow, err := client.Weather().GetTomorrowWeather(ctx, "110101")
	if err != nil {
		log.Fatalf("GetTomorrowWeather 失败: %v", err)
	}

	fmt.Printf("明天天气:\n")
	fmt.Printf("日期: %s\n", tomorrow.Date)
	fmt.Printf("白天: %s %s°C\n", tomorrow.DayWeather, tomorrow.DayTemp)
	fmt.Printf("夜间: %s %s°C\n", tomorrow.NightWeather, tomorrow.NightTemp)

	nextDays, err := client.Weather().GetNextDaysWeather(ctx, "110101", 3)
	if err != nil {
		log.Fatalf("GetNextDaysWeather 失败: %v", err)
	}

	fmt.Printf("\n未来3天天气:\n")
	for i, cast := range nextDays {
		fmt.Printf("%d. %s | 白天: %s %s°C | 夜间: %s %s°C\n",
			i+1, cast.Date, cast.DayWeather, cast.DayTemp, cast.NightWeather, cast.NightTemp)
	}
}
