package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	Separator         = "================================"
	DefaultCity       = "gargzdai"
	DefaultColumnView = true
	UsedRangeArgument = false
	DefaultDay        = time.Now().Day()
	DefaultStartHour  = time.Now().Hour()
	DefaultEndHour    = 24
	Gray              = "\033[37m"
	Reset             = "\033[0m"
	Red               = "\033[31m"
	Green             = "\033[32m"
	Yellow            = "\033[33m"
	Blue              = "\033[34m"
	Purple            = "\033[35m"
	Cyan              = "\033[36m"
	White             = "\033[97m"
)

type Forecast struct {
	ForecastTimeUtc    string  `json:"forecastTimeUtc"`
	AirTemperature     float32 `json:"airTemperature"`
	TotalParticipation float32 `json:"totalPrecipitation"`
	WindSpeed          float32 `json:"windSpeed"`
	FormattedTime      time.Time
}

type Weather struct {
	Place struct {
		Code    string `json:"code"`
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"place"`
	ForecastTimestamps []Forecast `json:"forecastTimestamps"`
}

func (w *Weather) GetDefaultDayForecast() []Forecast {
	var dayForecast []Forecast
	// For other days than today, show whole forecast, unless -r was used
	if DefaultDay != time.Now().Day() && !UsedRangeArgument {
		DefaultStartHour = 0
	}

	for _, day := range w.ForecastTimestamps {
		t, _ := time.Parse("2006-01-02 15:04:05", day.ForecastTimeUtc)
		if t.Day() != DefaultDay {
			if len(dayForecast) != 0 {
				break
			}
			continue
		}

		if t.Hour() < DefaultStartHour || t.Hour() > DefaultEndHour {
			continue
		}

		day.FormattedTime = t
		dayForecast = append(dayForecast, day)
	}

	return dayForecast
}

func ReadData(url string) (*Weather, error) {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	var data Weather
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	// Got no data
	if data.Place.Name == "" {
		return nil, fmt.Errorf("Miestas \"%v\" neturi duomenų", DefaultCity)
	}

	return &data, nil
}

func TemperatureColor(temperature float32) string {
	switch {
	case temperature <= 12:
		return Blue
	case temperature > 12 && temperature < 20:
		return Cyan
	case temperature >= 20 && temperature < 25:
		return Green
	case temperature >= 25:
		return Red
	default:
		return Reset
	}
}

func GetRainDescription(totalParticipation float32) string {
	switch {
	case totalParticipation == 0:
		return "nelis"
	case totalParticipation <= 1:
		return "mažas lietus"
	case totalParticipation > 1 && totalParticipation <= 2:
		return "vidutinis lietus"
	case totalParticipation > 2:
		return "smarkus lietus"
	}
	return ""
}

func MapMonthsToLithuanian(month time.Month) string {
	m := month - 1
	if m < 0 || m > 11 {
		return "Nežinomas mėnuo"
	}
	tt := []string{"Sausio", "Vasario", "Kovo", "Balandžio", "Gegužės", "Birželio", "Liepos", "Rugpjūčio", "Rugsėjo", "Spalio", "Lapkričio", "Gruodžio"}
	return tt[m]
}

func DisplayDayInfoList(forecast []Forecast) {
	fmt.Println("Data:",
		MapMonthsToLithuanian(forecast[0].FormattedTime.Month()),
		forecast[0].FormattedTime.Day())

	var topTemperature float32 = 0
	var averageTemperature float32 = 0
	var averageWindSpeed float32 = 0
	for _, hour := range forecast {
		if topTemperature < hour.AirTemperature {
			topTemperature = hour.AirTemperature
		}
		weatherDescription := GetRainDescription(hour.TotalParticipation)
		if hour.FormattedTime.Hour() == time.Now().Hour() {
			fmt.Printf(Purple+"Laikas: %+2vh "+TemperatureColor(hour.AirTemperature)+" %-7v"+Reset+"%v\n",
				hour.FormattedTime.Hour(), fmt.Sprintf("%v°C", hour.AirTemperature), weatherDescription)
		} else {
			fmt.Printf(Reset+"Laikas: %+2vh "+TemperatureColor(hour.AirTemperature)+" %-7v"+Reset+"%v\n",
				hour.FormattedTime.Hour(), fmt.Sprintf("%v°C", hour.AirTemperature), weatherDescription)
		}
		averageTemperature += hour.AirTemperature
		averageWindSpeed += hour.WindSpeed
	}
	fmt.Println(Separator)
	fmt.Printf("Aukščiausia temperatūra: %v°C\n", topTemperature)
	fmt.Printf("Vidutinė temperatūra: %.1f°C\n", averageTemperature/float32(len(forecast)))
	fmt.Printf("Vidutinis vėjo greitis: %.1fm/s\n", averageWindSpeed/float32(len(forecast)))
	fmt.Println(Separator)
}

func DisplayDayInfoColumn(forecast []Forecast) {
	Separator := strings.Repeat("=", 9*len(forecast))
	fmt.Println("Data:",
		MapMonthsToLithuanian(forecast[0].FormattedTime.Month()),
		forecast[0].FormattedTime.Day())

	fmt.Println(Separator)

	var topTemperature float32 = 0
	var averageTemperature float32 = 0
	var averageWindSpeed float32 = 0

	for _, hour := range forecast {
		if topTemperature < hour.AirTemperature {
			topTemperature = hour.AirTemperature
		}
		if hour.FormattedTime.Hour() == time.Now().Hour() {
			fmt.Printf(Purple+" %-7v"+Reset+"|", fmt.Sprintf("%vh", hour.FormattedTime.Hour()))
		} else {
			fmt.Printf(" %-7v|", fmt.Sprintf("%vh", hour.FormattedTime.Hour()))
		}
		averageTemperature += hour.AirTemperature
		averageWindSpeed += hour.WindSpeed
	}
	fmt.Println()
	for _, hour := range forecast {
		if hour.FormattedTime.Hour() == time.Now().Hour() {
			fmt.Printf(TemperatureColor(hour.AirTemperature)+" %-7v"+Reset+"|",
				fmt.Sprintf("%v°C", hour.AirTemperature))
		} else {
			fmt.Printf(TemperatureColor(hour.AirTemperature)+" %-7v"+Reset+"|",
				fmt.Sprintf("%v°C", hour.AirTemperature))
		}
	}

	fmt.Println("\n" + Separator)
	fmt.Printf("Aukščiausia temperatūra: %v°C\n", topTemperature)
	fmt.Printf("Vidutinė temperatūra: %.1f°C\n", averageTemperature/float32(len(forecast)))
	fmt.Printf("Vidutinis vėjo greitis: %.1fm/s\n", averageWindSpeed/float32(len(forecast)))
	fmt.Println(Separator)
}

func HandleArguments() {
	args := os.Args[1:]
	//example:  weather -c gargzdai -lv -d 0
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-r":
			start := DefaultStartHour
			end := DefaultEndHour
			i++
			if args[i] != "." {
				start, _ = strconv.Atoi(args[i])
			}
			i++
			if args[i] != "." {
				end , _ = strconv.Atoi(args[i])
			}

			DefaultStartHour = start
			DefaultEndHour = end
			UsedRangeArgument = true
		case "-c":
			i++
			DefaultCity = args[i]
		case "-lv":
			DefaultColumnView = false
		case "-d":
			i++
			day, _ := strconv.Atoi(args[i])
			DefaultDay = time.Now().AddDate(0, 0, day%7).Day()
		case "-h":
			fmt.Println("Usage: weather [arguments]")
			fmt.Println("Arguments:")
			fmt.Println("	-r (START|.) (END|.) - set hour display range <Default = 8 24>")
			fmt.Println("	-c CITYNAME - change default city")
			fmt.Println("	-lv - change to a list view")
			fmt.Println("	-d NUMBER - display day (0 - today, 1 - tomorrow, ...). Range 0..6 <Default = 0>")
			os.Exit(0)
		}
	}
}

func main() {

	HandleArguments()
	url := "https://api.meteo.lt/v1/places/" + DefaultCity + "/forecasts/long-term"
	weather, err := ReadData(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	forecast := weather.GetDefaultDayForecast()

	fmt.Println("Miestas:", weather.Place.Name)

	if DefaultColumnView {
		DisplayDayInfoColumn(forecast)
	} else {
		DisplayDayInfoList(forecast)
	}
}
