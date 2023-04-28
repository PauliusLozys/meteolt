package meteolt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseAPIURL = "https://api.meteo.lt/v1"

func GetWeatherForecast(ctx context.Context, city string, forecastType ForecastType) (*Forecast, error) {
	url := fmt.Sprintf("%s/places/%s/forecasts/%s", baseAPIURL, city, forecastType)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var forecast Forecast
	if err := json.NewDecoder(resp.Body).Decode(&forecast); err != nil {
		return nil, err
	}

	return &forecast, nil
}

func GetAllPlaces(ctx context.Context) ([]PlaceMetadata, error) {
	url := fmt.Sprintf("%s/places", baseAPIURL)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var places []PlaceMetadata
	if err := json.NewDecoder(resp.Body).Decode(&places); err != nil {
		return nil, err
	}

	return places, nil
}

func GetAllStations(ctx context.Context) ([]Station, error) {
	url := fmt.Sprintf("%s/stations", baseAPIURL)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var stations []Station
	if err := json.NewDecoder(resp.Body).Decode(&stations); err != nil {
		return nil, err
	}

	return stations, nil
}
