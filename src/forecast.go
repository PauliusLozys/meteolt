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

func (forecasts ForecastList) DisplayDayInfoList() {
	var (
		topTemperature     float32 = 0.0
		averageTemperature float32 = 0.0
		averageWindSpeed   float32 = 0.0
		separator          string  = strings.Repeat("=", 40)
	)

	fmt.Println("Data:",
		MapMonthsToLithuanian(forecasts[0].FormattedTime.Month()),
		forecasts[0].FormattedTime.Day(),
	)
	fmt.Println(separator)

	for _, forecast := range forecasts {
		if topTemperature < forecast.AirTemperature {
			topTemperature = forecast.AirTemperature
		}
		weatherDescription := GetRainDescription(forecast.TotalParticipation)
		color := Reset
		if forecast.FormattedTime.Hour() == time.Now().Hour() {
			color = Purple
		}
		fmt.Printf(color+"Laikas: %+2vh "+TemperatureColor(forecast.AirTemperature)+" %-7v"+Reset+"%v\n",
			forecast.FormattedTime.Hour(), fmt.Sprintf("%v°C", forecast.AirTemperature), weatherDescription)
		averageTemperature += forecast.AirTemperature
		averageWindSpeed += forecast.WindSpeed
	}

	fmt.Println(separator)
	fmt.Printf("Aukščiausia temperatūra: %v°C\n", topTemperature)
	fmt.Printf("Vidutinė temperatūra: %.1f°C\n", averageTemperature/float32(len(forecasts)))
	fmt.Printf("Vidutinis vėjo greitis: %.1fm/s\n", averageWindSpeed/float32(len(forecasts)))
	fmt.Println(separator)
}

func (forecasts ForecastList) DisplayDayInfoColumn() {
	var (
		topTemperature     float32 = 0.0
		averageTemperature float32 = 0.0
		averageWindSpeed   float32 = 0.0
		separator          string  = strings.Repeat("=", 9*len(forecasts))
	)

	fmt.Println("Data:",
		MapMonthsToLithuanian(forecasts[0].FormattedTime.Month()),
		forecasts[0].FormattedTime.Day())

	fmt.Println(separator)

	for _, forecast := range forecasts {
		if topTemperature < forecast.AirTemperature {
			topTemperature = forecast.AirTemperature
		}
		color := Reset
		if forecast.FormattedTime.Hour() == time.Now().Hour() {
			color = Purple
		}
		fmt.Printf(color+" %-7v"+Reset+"|", fmt.Sprintf("%vh", forecast.FormattedTime.Hour()))
		averageTemperature += forecast.AirTemperature
		averageWindSpeed += forecast.WindSpeed
	}
	fmt.Println()
	for _, forecast := range forecasts {
		fmt.Printf(TemperatureColor(forecast.AirTemperature)+" %-7v"+Reset+"|",
			fmt.Sprintf("%v°C", forecast.AirTemperature))
	}

	fmt.Println("\n" + separator)
	fmt.Printf("Aukščiausia temperatūra: %v°C\n", topTemperature)
	fmt.Printf("Vidutinė temperatūra: %.1f°C\n", averageTemperature/float32(len(forecasts)))
	fmt.Printf("Vidutinis vėjo greitis: %.1fm/s\n", averageWindSpeed/float32(len(forecasts)))
	fmt.Println(separator)
}
