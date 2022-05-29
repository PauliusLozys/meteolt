package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Weather struct {
	Place struct {
		Code    string `json:"code"`
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"place"`
	ForecastTimestamps ForecastList `json:"forecastTimestamps"`
}

func (w *Weather) GetForecastByDay(day, startHour, endHour int) ForecastList {
	// For other days than today, show whole forecast, unless -r was used
	if day != time.Now().Day() && !UsedRangeArgument {
		startHour = 0
	}

	var dayForecast ForecastList
	for _, forecast := range w.ForecastTimestamps {
		t, _ := time.Parse("2006-01-02 15:04:05", forecast.ForecastTimeUtc)

		if t.Day() != day || t.Hour() < startHour || t.Hour() > endHour {
			if len(dayForecast) != 0 {
				break
			}
			continue
		}

		forecast.FormattedTime = t
		dayForecast = append(dayForecast, forecast)
	}

	return dayForecast
}

func ReadWeatherData(url string) (*Weather, error) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	var data Weather
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	// Got no data
	if data.Place.Name == "" {
		return nil, fmt.Errorf("miestas \"%v\" neturi duomenų", DefaultCity)
	}

	return &data, nil
}
