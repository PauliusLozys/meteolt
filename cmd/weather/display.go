package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PauliusLozys/meteolt"
)

func DisplayDayInfoList(
	showDetails bool,
	forecasts []meteolt.ForecastTimestamp,
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
		if forecast.ForecastTimeUtc.Hour() == time.Now().Hour() {
			color = purple
		}

		fmt.Printf("%sLaikas: %+2vh%s %-7v %s%s\n",
			color,
			forecast.ForecastTimeUtc.Hour(),
			TemperatureColor(forecast.AirTemperature),
			fmt.Sprintf("%v°C", forecast.AirTemperature),
			reset,
			GetRainDescription(forecast.TotalPrecipitation),
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
	forecasts []meteolt.ForecastTimestamp,
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
		if forecast.ForecastTimeUtc.Hour() == time.Now().Hour() {
			color = purple
		}
		fmt.Printf(color+" %-7v"+reset+"|", fmt.Sprintf("%vh", forecast.ForecastTimeUtc.Hour()))
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

func DisplayDayInfo(
	forecasts []meteolt.ForecastTimestamp,
	showDetails bool,
	displayFn func(bool, []meteolt.ForecastTimestamp) (float32, float32, float32),
) {

	separator := strings.Repeat("=", 9*len(forecasts))
	if len(forecasts) == 0 {
		fmt.Println("Nėra duomenų")
		return
	}

	fmt.Println("Data:",
		MapMonthsToLithuanian(forecasts[0].ForecastTimeUtc.Month()),
		forecasts[0].ForecastTimeUtc.Day(),
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
