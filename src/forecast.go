package main

import (
	"fmt"
	"strings"
	"time"
)

type Forecast struct {
	ForecastTimeUtc    string  `json:"forecastTimeUtc"`
	AirTemperature     float32 `json:"airTemperature"`
	TotalParticipation float32 `json:"totalPrecipitation"`
	WindSpeed          float32 `json:"windSpeed"`
	FormattedTime      time.Time
}

type ForecastList []Forecast

func (forecast ForecastList) DisplayDayInfoList() {
	var (
		topTemperature     float32 = 0.0
		averageTemperature float32 = 0.0
		averageWindSpeed   float32 = 0.0
		separator          string  = strings.Repeat("=", 40)
	)

	fmt.Println("Data:",
		MapMonthsToLithuanian(forecast[0].FormattedTime.Month()),
		forecast[0].FormattedTime.Day(),
	)
	fmt.Println(separator)

	for _, hour := range forecast {
		if topTemperature < hour.AirTemperature {
			topTemperature = hour.AirTemperature
		}
		weatherDescription := GetRainDescription(hour.TotalParticipation)
		if hour.FormattedTime.Hour() == time.Now().Hour() {
			fmt.Printf(Purple+"Laikas: %+2vh "+TemperatureColor(hour.AirTemperature)+" %-7v"+Reset+"%v\n",
				hour.FormattedTime.Hour(), fmt.Sprintf("%v°C", hour.AirTemperature), weatherDescription)
		} else {
			fmt.Printf(Reset+"Laikas: %+2vh "+TemperatureColor(hour.AirTemperature)+" %-7v"+Reset+"%v\n",
				hour.FormattedTime.Hour(), fmt.Sprintf("%v°C", hour.AirTemperature), weatherDescription)
		}
		averageTemperature += hour.AirTemperature
		averageWindSpeed += hour.WindSpeed
	}
	fmt.Println(separator)
	fmt.Printf("Aukščiausia temperatūra: %v°C\n", topTemperature)
	fmt.Printf("Vidutinė temperatūra: %.1f°C\n", averageTemperature/float32(len(forecast)))
	fmt.Printf("Vidutinis vėjo greitis: %.1fm/s\n", averageWindSpeed/float32(len(forecast)))
	fmt.Println(separator)
}

func (forecast ForecastList) DisplayDayInfoColumn() {
	var (
		topTemperature     float32 = 0.0
		averageTemperature float32 = 0.0
		averageWindSpeed   float32 = 0.0
		separator          string  = strings.Repeat("=", 9*len(forecast))
	)

	fmt.Println("Data:",
		MapMonthsToLithuanian(forecast[0].FormattedTime.Month()),
		forecast[0].FormattedTime.Day())

	fmt.Println(separator)

	for _, hour := range forecast {
		if topTemperature < hour.AirTemperature {
			topTemperature = hour.AirTemperature
		}
		if hour.FormattedTime.Hour() == time.Now().Hour() {
			fmt.Printf(Purple+" %-7v"+Reset+"|", fmt.Sprintf("%vh", hour.FormattedTime.Hour()))
		} else {
			fmt.Printf(" %-7v|", fmt.Sprintf("%vh", hour.FormattedTime.Hour()))
		}
		averageTemperature += hour.AirTemperature
		averageWindSpeed += hour.WindSpeed
	}
	fmt.Println()
	for _, hour := range forecast {
		if hour.FormattedTime.Hour() == time.Now().Hour() {
			fmt.Printf(TemperatureColor(hour.AirTemperature)+" %-7v"+Reset+"|",
				fmt.Sprintf("%v°C", hour.AirTemperature))
		} else {
			fmt.Printf(TemperatureColor(hour.AirTemperature)+" %-7v"+Reset+"|",
				fmt.Sprintf("%v°C", hour.AirTemperature))
		}
	}

	fmt.Println("\n" + separator)
	fmt.Printf("Aukščiausia temperatūra: %v°C\n", topTemperature)
	fmt.Printf("Vidutinė temperatūra: %.1f°C\n", averageTemperature/float32(len(forecast)))
	fmt.Printf("Vidutinis vėjo greitis: %.1fm/s\n", averageWindSpeed/float32(len(forecast)))
	fmt.Println(separator)
}
