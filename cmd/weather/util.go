package main

import (
	"context"
	"fmt"
	"time"

	"github.com/PauliusLozys/meteolt"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

var monthToName = map[time.Month]string{
	1:  "Sausio",
	2:  "Vasario",
	3:  "Kovo",
	4:  "Balandžio",
	5:  "Gegužės",
	6:  "Birželio",
	7:  "Liepos",
	8:  "Rugpjūčio",
	9:  "Rugsėjo",
	10: "Spalio",
	11: "Lapkričio",
	12: "Gruodžio",
}

func TemperatureColor(temperature float32) string {
	switch {
	case temperature <= 12:
		return blue
	case temperature < 20:
		return cyan
	case temperature < 25:
		return green
	default: // 25+
		return red
	}
}

func GetRainDescription(totalParticipation float32) string {
	switch {
	case totalParticipation == 0:
		return "nelis"
	case totalParticipation <= 1:
		return "mažas lietus"
	case totalParticipation <= 2:
		return "vidutinis lietus"
	default: // 2+
		return "smarkus lietus"
	}
}

func MapMonthsToLithuanian(month time.Month) string {
	if monthName, ok := monthToName[month]; ok {
		return monthName
	}
	return "Nežinomas mėnuo"
}

func GetForecastByDay(w *meteolt.Forecast, day, startHour, endHour int) []meteolt.ForecastTimestamp {
	// For other days than today, show whole forecast, unless -r was used
	if day != time.Now().Day() && !usedRangeArgument {
		startHour = 0
	}

	var dayForecast []meteolt.ForecastTimestamp
	for _, forecast := range w.ForecastTimestamps {
		t := forecast.ForecastTimeUtc
		if t.Day() != day || t.Hour() < startHour || t.Hour() > endHour {
			if len(dayForecast) != 0 {
				break
			}
			continue
		}

		dayForecast = append(dayForecast, forecast)
	}

	return dayForecast
}

func ReadWeatherData() (*meteolt.Forecast, error) {
	forecast, err := meteolt.GetWeatherForecast(context.TODO(), defaultCity, meteolt.ForecastTypeLongTerm)
	if err != nil {
		return nil, err
	}

	// Got no data
	if forecast.Place.Name == "" {
		return nil, fmt.Errorf("miestas \"%v\" neturi duomenų", defaultCity)
	}

	return forecast, nil
}
