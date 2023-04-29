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
	ConditionCodePartlyVariableCloudiness   ConditionCode = "variable-cloudiness"
	ConditionCodeCloudyWithSunnyIntervals   ConditionCode = "cloudy-with-sunny-intervals"
	ConditionCodeCloudy                     ConditionCode = "cloudy"
	ConditionCodeThunder                    ConditionCode = "thunder"
	ConditionCodeIsolatedThunderStorms      ConditionCode = "isolated-thunderstorms"
	ConditionCodeThunderstorms              ConditionCode = "thunderstorms"
	ConditionCodeHeavyRainWithThunderstorms ConditionCode = "heavy-rain-with-thunderstorms"
	ConditionCodeLightRain                  ConditionCode = "light-rain"
	ConditionCodeRain                       ConditionCode = "rain"
	ConditionCodeRainShowers                ConditionCode = "rain-showers"
	ConditionCodeLightRainAtTimes           ConditionCode = "light-rain-at-times"
	ConditionCodeRainAtTimes                ConditionCode = "rain-at-times"
	ConditionCodeHeavyRain                  ConditionCode = "heavy-rain"
	ConditionCodeLightSleet                 ConditionCode = "light-sleet"
	ConditionCodeSleet                      ConditionCode = "sleet"
	ConditionCodeSleetAtTimes               ConditionCode = "sleet-at-times"
	ConditionCodeSleetShowers               ConditionCode = "sleet-showers"
	ConditionCodeFreezingRain               ConditionCode = "freezing-rain"
	ConditionCodeHail                       ConditionCode = "hail"
	ConditionCodeLightShow                  ConditionCode = "light-snow"
	ConditionCodeSnow                       ConditionCode = "snow"
	ConditionCodeSnowShowers                ConditionCode = "snow-showers"
	ConditionCodeSnowAtTimes                ConditionCode = "snow-at-times"
	ConditionCodeLightSnowAtTimes           ConditionCode = "light-snow-at-times"
	ConditionCodeHeavySnow                  ConditionCode = "heavy-snow"
	ConditionCodeShowStorm                  ConditionCode = "snowstorm"
	ConditionCodeMist                       ConditionCode = "mist"
	ConditionCodeSquall                     ConditionCode = "squall"
	ConditionCodeFog                        ConditionCode = "fog"
	ConditionCodeNull                       ConditionCode = "null"
)
