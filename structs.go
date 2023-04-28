package meteolt

// ForeCastType describes what type of forecast is returned by the API.
type ForecastType string

const (
	// ForecastTypeLongTerm is a long term forecast response type from the API.
	ForecastTypeLongTerm ForecastType = "long-term"
)

// ConditionCode describes weather conditions.
type ConditionCode string

const (
	ConditionCodeClear                      ConditionCode = "clear"
	ConditionCodePartlyCloudy               ConditionCode = "partly-cloudy"
	ConditionCodeCloudyWithSunnyIntervals   ConditionCode = "cloudy-with-sunny-intervals"
	ConditionCodeCloudy                     ConditionCode = "cloudy"
	ConditionCodeThunder                    ConditionCode = "thunder"
	ConditionCodeIsolatedThunderStorms      ConditionCode = "isolated-thunderstorms"
	ConditionCodeThunderstorms              ConditionCode = "thunderstorms"
	ConditionCodeHeavyRainWithThunderstorms ConditionCode = "heavy-rain-with-thunderstorms"
	ConditionCodeLightRain                  ConditionCode = "light-rain"
	ConditionCodeRain                       ConditionCode = "rain"
	ConditionCodeHeavyRain                  ConditionCode = "heavy-rain"
	ConditionCodeLightSleet                 ConditionCode = "light-sleet"
	ConditionCodeSleet                      ConditionCode = "sleet"
	ConditionCodeFreezingRain               ConditionCode = "freezing-rain"
	ConditionCodeHail                       ConditionCode = "hail"
	ConditionCodeLightShow                  ConditionCode = "light-snow"
	ConditionCodeSnow                       ConditionCode = "snow"
	ConditionCodeHeavySnow                  ConditionCode = "heavy-snow"
	ConditionCodeFog                        ConditionCode = "fog"
	ConditionCodeNull                       ConditionCode = "null"
)

type Forecast struct {
	Place                   PlaceMetadata       `json:"place"`                   // Weather forecast place metadata.
	ForecastType            ForecastType        `json:"forecastType"`            // Type of weather forecast.
	ForecastCreationTimeUtc string              `json:"forecastCreationTimeUtc"` // Time of the weather forecast creation (UTC time zone).
	ForecastTimestamps      []ForecastTimestamp `json:"forecastTimestamps"`      // List of weather forecast timestamps.
}

type ForecastTimestamp struct {
	ForecastTimeUtc      string        `json:"forecastTimeUtc"`      // Weather forecast time (UTC time zone).
	AirTemperature       float32       `json:"airTemperature"`       // Air temperature, °C.
	FeelsLikeTemperature float32       `json:"feelsLikeTemperature"` // Sensible temperature, °C.
	WindSpeed            int           `json:"windSpeed"`            // Wind speed, m/s.
	WindGust             int           `json:"windGust"`             // Wind gust, m/s.
	WindDirection        int           `json:"windDirection"`        // Wind direction in degrees. Values: 0 - from the north, 180 - from the south, etc.
	CloudCover           int           `json:"cloudCover"`           // Cloudiness, %. Values: 0 - clear, 100 - cloudy.
	SeaLevelPressure     int           `json:"seaLevelPressure"`     // Pressure at sea level, hPa.
	RelativeHumidity     int           `json:"relativeHumidity"`     // Relative air humidity, %.
	TotalPrecipitation   float32       `json:"totalPrecipitation"`   // Amount of precipitation, mm.
	ConditionCode        ConditionCode `json:"conditionCode"`        // Weather condition code.
}

type PlaceMetadata struct {
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
	Code        string           `json:"code"`
	Name        string           `json:"name"`
	Coordinates PlaceCoordinates `json:"coordinates"`
}
