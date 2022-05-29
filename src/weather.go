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
	ForecastTimestamps []Forecast `json:"forecastTimestamps"`
}

func (w *Weather) GetDefaultDayForecast() ForecastList {
	// For other days than today, show whole forecast, unless -r was used
	if DefaultDay != time.Now().Day() && !UsedRangeArgument {
		DefaultStartHour = 0
	}

	var dayForecast []Forecast
	for _, day := range w.ForecastTimestamps {
		t, _ := time.Parse("2006-01-02 15:04:05", day.ForecastTimeUtc)

		if t.Day() != DefaultDay || t.Hour() < DefaultStartHour || t.Hour() > DefaultEndHour {
			if len(dayForecast) != 0 {
				break
			}
			continue
		}

		day.FormattedTime = t
		dayForecast = append(dayForecast, day)
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
