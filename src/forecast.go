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
	WindGust           float32 `json:"windGust"`
	RelativeHumidity   int     `json:"relativeHumidity"`
	FormattedTime      time.Time
}

type ForecastList []Forecast

func DisplayDayInfoList(
	showDetails bool,
	forecasts ForecastList,
) (float32, float32, float32) {

	var (
		topTemperature     float32 = 0.0
		averageTemperature float32 = 0.0
		averageWindSpeed   float32 = 0.0
	)

	for _, forecast := range forecasts {
		if topTemperature < forecast.AirTemperature {
			topTemperature = forecast.AirTemperature
		}
		color := reset
		if forecast.FormattedTime.Hour() == time.Now().Hour() {
			color = purple
		}

		fmt.Printf("%sLaikas: %+2vh%s %-7v %s%s\n",
			color,
			forecast.FormattedTime.Hour(),
			TemperatureColor(forecast.AirTemperature),
			fmt.Sprintf("%v°C", forecast.AirTemperature),
			reset,
			GetRainDescription(forecast.TotalParticipation),
		)

		if showDetails {
			fmt.Printf("\tVėjo gretis: %vm/s\n\tVėjo gūsiai: %vm/s\n\tSantykinė oro dregmė: %v%%\n",
				forecast.WindSpeed,
				forecast.WindGust,
				forecast.RelativeHumidity,
			)
		}
		averageTemperature += forecast.AirTemperature
		averageWindSpeed += forecast.WindSpeed
	}
	return topTemperature, averageTemperature, averageWindSpeed
}

func DisplayDayInfoColumn(
	showDetails bool,
	forecasts ForecastList,
) (float32, float32, float32) {

	var (
		topTemperature     float32 = 0.0
		averageTemperature float32 = 0.0
		averageWindSpeed   float32 = 0.0
	)

	for _, forecast := range forecasts {
		if topTemperature < forecast.AirTemperature {
			topTemperature = forecast.AirTemperature
		}
		color := reset
		if forecast.FormattedTime.Hour() == time.Now().Hour() {
			color = purple
		}
		fmt.Printf(color+" %-7v"+reset+"|", fmt.Sprintf("%vh", forecast.FormattedTime.Hour()))
		averageTemperature += forecast.AirTemperature
		averageWindSpeed += forecast.WindSpeed
	}
	fmt.Println()
	for _, forecast := range forecasts {
		fmt.Printf(TemperatureColor(forecast.AirTemperature)+" %-7v"+reset+"|",
			fmt.Sprintf("%v°C", forecast.AirTemperature))
	}
	fmt.Println()
	return topTemperature, averageTemperature, averageWindSpeed
}

func (forecasts ForecastList) DisplayDayInfo(
	showDetails bool,
	displayFn func(bool, ForecastList) (float32, float32, float32),
) {

	separator := strings.Repeat("=", 9*len(forecasts))
	if len(forecasts) == 0 {
		fmt.Println("Nėra duomenų")
		return
	}

	fmt.Println("Data:",
		MapMonthsToLithuanian(forecasts[0].FormattedTime.Month()),
		forecasts[0].FormattedTime.Day(),
	)
	fmt.Println(separator)

	topTemp, averageTemp, averageWindSpeed := displayFn(
		showDetails,
		forecasts,
	)

	fmt.Println(separator)
	fmt.Printf("Aukščiausia temperatūra: %v°C\n", topTemp)
	fmt.Printf("Vidutinė temperatūra: %.1f°C\n", averageTemp/float32(len(forecasts)))
	fmt.Printf("Vidutinis vėjo greitis: %.1fm/s\n", averageWindSpeed/float32(len(forecasts)))
	fmt.Println(separator)
}
