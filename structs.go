package meteolt

import (
	"encoding/json"
	"strconv"
	"time"
)

type jsonTime time.Time

type Forecast struct {
	Place                   Place               `json:"place"`                   // Weather forecast place metadata.
	ForecastType            ForecastType        `json:"forecastType"`            // Type of weather forecast.
	ForecastCreationTimeUtc time.Time           `json:"forecastCreationTimeUtc"` // Time of the weather forecast creation (UTC time zone).
	ForecastTimestamps      []ForecastTimestamp `json:"forecastTimestamps"`      // List of weather forecast timestamps.
}

type ForecastTimestamp struct {
	ForecastTimeUtc      time.Time     `json:"forecastTimeUtc"`      // Weather forecast time (UTC time zone).
	AirTemperature       float32       `json:"airTemperature"`       // Air temperature, °C.
	FeelsLikeTemperature float32       `json:"feelsLikeTemperature"` // Sensible temperature, °C.
	WindSpeed            float32       `json:"windSpeed"`            // Wind speed, m/s.
	WindGust             float32       `json:"windGust"`             // Wind gust, m/s.
	WindDirection        int           `json:"windDirection"`        // Wind direction in degrees. Values: 0 - from the north, 180 - from the south, etc.
	CloudCover           int           `json:"cloudCover"`           // Cloudiness, %. Values: 0 - clear, 100 - cloudy.
	SeaLevelPressure     float32       `json:"seaLevelPressure"`     // Pressure at sea level, hPa.
	RelativeHumidity     int           `json:"relativeHumidity"`     // Relative air humidity, %.
	TotalPrecipitation   float32       `json:"totalPrecipitation"`   // Amount of precipitation, mm.
	ConditionCode        ConditionCode `json:"conditionCode"`        // Weather condition code.
}

type Place struct {
	Code                   string           `json:"code"`                   // Place code.
	Name                   string           `json:"name"`                   // Place name.
	AdministrativeDivision string           `json:"administrativeDivision"` // The administrative unit to which the area belongs.
	CountryCode            string           `json:"countryCode"`            // Country code in ISO-3166-1 alpha-2 format.
	Country                string           `json:"country"`                // Country name.
	Coordinates            PlaceCoordinates `json:"coordinates"`            // Location coordinates (WGS 84 in decimal degrees).
}

type PlaceCoordinates struct {
	Latitude  float32 `json:"latitude"`  // Latitude, decimal degrees.
	Longitude float32 `json:"longitude"` // Longitude, decimal degrees.
}

type Station struct {
	Code        string           `json:"code"`        // Station code.
	Name        string           `json:"name"`        // Station name.
	Coordinates PlaceCoordinates `json:"coordinates"` // Station coordinates (WGS 84 in decimal degrees).
}

type StationExtended struct {
	Station

	Type string `json:"type"` // Station type.
}

// StationObservation stores information about the station's observation data.
type StationObservation struct {
	Code                  string                `json:"code"`                  // Station code.
	ObservationsDataRange ObservationsDataRange `json:"observationsDataRange"` // Time interval of the stations observation data.
}

type ObservationsDataRange struct {
	StartTimeUtc time.Time `json:"startTimeUtc"` // The start of the time interval (UTC time zone) of the stored station observation data.
	EndTimeUtc   time.Time `json:"endTimeUtc"`   // The end of the time interval (UTC time zone) of the stored station observation data.
}

type Observation struct {
	ObservationTimeUtc   time.Time     `json:"observationTimeUtc"`   // Weather forecast time (UTC time zone).
	AirTemperature       float32       `json:"airTemperature"`       // Air temperature, °C.
	FeelsLikeTemperature float32       `json:"feelsLikeTemperature"` // Sensible temperature, °C.
	WindSpeed            float32       `json:"windSpeed"`            // Wind speed, m/s.
	WindGust             float32       `json:"windGust"`             // Wind gust, m/s.
	WindDirection        int           `json:"windDirection"`        // Wind direction in degrees. Values: 0 - from the north, 180 - from the south, etc.
	CloudCover           int           `json:"cloudCover"`           // Cloudiness, %. Values: 0 - clear, 100 - cloudy.
	SeaLevelPressure     float32       `json:"seaLevelPressure"`     // Pressure at sea level, hPa.
	RelativeHumidity     int           `json:"relativeHumidity"`     // Relative air humidity, %.
	Precipitation        float32       `json:"precipitation"`        // Amount of precipitation, mm.
	ConditionCode        ConditionCode `json:"conditionCode"`        // Weather condition code.
}

// jsonTime implements the json.Unmarshaler interface.
func (t *jsonTime) UnmarshalJSON(data []byte) error {
	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	nt, err := time.Parse("2006-01-02 15:04:05", unquoted)
	if err != nil {
		return err
	}

	*t = jsonTime(nt)

	return nil
}

// Forecast implements the json.Unmarshaler interface.
func (f *Forecast) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Place                   Place               `json:"place"`
		ForecastType            ForecastType        `json:"forecastType"`
		ForecastCreationTimeUtc jsonTime            `json:"forecastCreationTimeUtc"`
		ForecastTimestamps      []ForecastTimestamp `json:"forecastTimestamps"`
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*f = Forecast{
		Place:                   tmp.Place,
		ForecastType:            tmp.ForecastType,
		ForecastCreationTimeUtc: time.Time(tmp.ForecastCreationTimeUtc),
		ForecastTimestamps:      tmp.ForecastTimestamps,
	}

	return nil
}

// ForecastTimeStamp implements the json.Unmarshaler interface.
func (f *ForecastTimestamp) UnmarshalJSON(data []byte) error {
	var tmp struct {
		ForecastTimeUtc      jsonTime      `json:"forecastTimeUtc"`
		AirTemperature       float32       `json:"airTemperature"`
		FeelsLikeTemperature float32       `json:"feelsLikeTemperature"`
		WindSpeed            float32       `json:"windSpeed"`
		WindGust             float32       `json:"windGust"`
		WindDirection        int           `json:"windDirection"`
		CloudCover           int           `json:"cloudCover"`
		SeaLevelPressure     float32       `json:"seaLevelPressure"`
		RelativeHumidity     int           `json:"relativeHumidity"`
		TotalPrecipitation   float32       `json:"totalPrecipitation"`
		ConditionCode        ConditionCode `json:"conditionCode"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*f = ForecastTimestamp{
		ForecastTimeUtc:      time.Time(tmp.ForecastTimeUtc),
		AirTemperature:       tmp.AirTemperature,
		FeelsLikeTemperature: tmp.FeelsLikeTemperature,
		WindSpeed:            tmp.WindSpeed,
		WindGust:             tmp.WindGust,
		WindDirection:        tmp.WindDirection,
		CloudCover:           tmp.CloudCover,
		SeaLevelPressure:     tmp.SeaLevelPressure,
		RelativeHumidity:     tmp.RelativeHumidity,
		TotalPrecipitation:   tmp.TotalPrecipitation,
		ConditionCode:        tmp.ConditionCode,
	}
	return nil
}

// ObservationsDataRange implements the json.Unmarshaler interface.
func (o *ObservationsDataRange) UnmarshalJSON(data []byte) error {
	var tmp struct {
		StartTimeUtc jsonTime `json:"startTimeUtc"`
		EndTimeUtc   jsonTime `json:"endTimeUtc"`
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*o = ObservationsDataRange{
		StartTimeUtc: time.Time(tmp.StartTimeUtc),
		EndTimeUtc:   time.Time(tmp.EndTimeUtc),
	}

	return nil
}

// Observation implements the json.Unmarshaler interface.
func (o *Observation) UnmarshalJSON(data []byte) error {
	var tmp struct {
		ObservationTimeUtc   jsonTime      `json:"observationTimeUtc"`
		AirTemperature       float32       `json:"airTemperature"`
		FeelsLikeTemperature float32       `json:"feelsLikeTemperature"`
		WindSpeed            float32       `json:"windSpeed"`
		WindGust             float32       `json:"windGust"`
		WindDirection        int           `json:"windDirection"`
		CloudCover           int           `json:"cloudCover"`
		SeaLevelPressure     float32       `json:"seaLevelPressure"`
		RelativeHumidity     int           `json:"relativeHumidity"`
		Precipitation        float32       `json:"precipitation"`
		ConditionCode        ConditionCode `json:"conditionCode"`
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*o = Observation{
		ObservationTimeUtc:   time.Time(tmp.ObservationTimeUtc),
		AirTemperature:       tmp.AirTemperature,
		FeelsLikeTemperature: tmp.FeelsLikeTemperature,
		WindSpeed:            tmp.WindSpeed,
		WindGust:             tmp.WindGust,
		WindDirection:        tmp.WindDirection,
		CloudCover:           tmp.CloudCover,
		SeaLevelPressure:     tmp.SeaLevelPressure,
		RelativeHumidity:     tmp.RelativeHumidity,
		Precipitation:        tmp.Precipitation,
		ConditionCode:        tmp.ConditionCode,
	}

	return nil
}
