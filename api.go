package meteolt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseAPIURL = "https://api.meteo.lt/v1"

// GetWeatherObservation returns weather observation data for a specific place.
//
// Internally calls '/places/{place}/forecasts/{forecast-type}' endpoint.
func GetWeatherForecast(ctx context.Context, place string, forecastType ForecastType) (*Forecast, error) {
	url := fmt.Sprintf("%s/places/%s/forecasts/%s", baseAPIURL, place, forecastType)
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

// GetAllPlaces returns a list of places for which weather forecast is provided.
//
// Internally calls '/places' endpoint.
func GetAllPlaces(ctx context.Context) ([]Place, error) {
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

	var places []Place
	if err := json.NewDecoder(resp.Body).Decode(&places); err != nil {
		return nil, err
	}

	return places, nil
}

// GetAllStation returns a list of meteorological stations for which observation data are provided is returned.
//
// Internally calls '/stations' endpoint.
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

// GetStationInfo returns information about the meteorological station.
//
// Internally calls '/stations/{station-code}' endpoint.
func GetStationInfo(ctx context.Context, stationCode string) (*StationExtended, error) {
	url := fmt.Sprintf("%s/stations/%s", baseAPIURL, stationCode)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var station StationExtended
	if err := json.NewDecoder(resp.Body).Decode(&station); err != nil {
		return nil, err
	}

	return &station, nil
}

// GetStationObservationInfo returns information about the meteorological station's observations.
//
// Internally calls '/stations/{station-code}/observations' endpoint.
func GetStationObservationInfo(ctx context.Context, stationCode string) (*StationObservation, error) {
	url := fmt.Sprintf("%s/stations/%s/observations", baseAPIURL, stationCode)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var observation StationObservation
	if err := json.NewDecoder(resp.Body).Decode(&observation); err != nil {
		return nil, err
	}

	return &observation, nil
}

// GetStationObservationsData returns the observations measured by the station at specified date.
//
// Internally calls '/stations/{station-code}/observations/{date}' endpoint.
func GetStationObservationsDataByDate(ctx context.Context, stationCode string, date time.Time) (*Station, []Observation, error) {
	url := fmt.Sprintf("%s/stations/%s/observations/%s", baseAPIURL, stationCode, date.Format("2006-01-02"))
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	tmp := struct {
		Station      Station       `json:"station"`
		Observations []Observation `json:"observations"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&tmp); err != nil {
		return nil, nil, err
	}

	return &tmp.Station, tmp.Observations, nil
}

// GetStationObservationsDataLatest returns the observations measured by the station from the last 24 hours.
//
// Internally calls '/stations/{station-code}/observations/latest' endpoint.
func GetStationObservationsDataLatest(ctx context.Context, stationCode string) (*Station, []Observation, error) {
	url := fmt.Sprintf("%s/stations/%s/observations/latest", baseAPIURL, stationCode)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	tmp := struct {
		Station      Station       `json:"station"`
		Observations []Observation `json:"observations"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&tmp); err != nil {
		return nil, nil, err
	}

	return &tmp.Station, tmp.Observations, nil
}
