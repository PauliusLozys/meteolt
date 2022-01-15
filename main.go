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
	UsedRangeArgument = false
	DefaultDay        = time.Now().Day()
	DefaultStartHour  = time.Now().Hour()
)

type Forecast struct {
	ForecastTimeUtc    string  `json:"forecastTimeUtc"`
	AirTemperature     float32 `json:"airTemperature"`
	TotalParticipation float32 `json:"totalPrecipitation"`
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

func ReadData(url string) *Weather {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return nil
	}
	body, _ := ioutil.ReadAll(response.Body)
	var data Weather
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return &data
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
	m := month-1
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
	}
	fmt.Println(Separator)
	fmt.Printf("Aukščiausia temperatūra: %v°C\n", topTemperature)
	fmt.Printf("Vidutinė temperatūra: %.1f°C\n", averageTemperature/float32(len(forecast)))
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
	fmt.Println(Separator)
}

func HandleArguments() {
	args := os.Args[1:]
	//example:  weather -c gargzdai -lv -d 0
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-r":
			i++
			start, _ := strconv.Atoi(args[i])
			i++
			end, _ := strconv.Atoi(args[i])
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
		case "-h":
			fmt.Println("Usage: weather [arguments]")
			fmt.Println("Arguments\n\t-r START END - set hour display range <Default = 8 24>\n\t-c CITYNAME - change default city\n\t-lv - change to a list view\n\t" +
				"-d NUMBER - display day (0 - today, 1 - tomorrow, ...). Range 0..6 <Default = 0>")
			os.Exit(0)
		}
	}
}

func main() {

	HandleArguments()
	url := "https://api.meteo.lt/v1/places/" + DefaultCity + "/forecasts/long-term"
	weather := ReadData(url)
	forecast := weather.GetDefaultDayForecast()

	fmt.Println("Miestas:", weather.Place.Name)

	if DefaultColumnView {
		DisplayDayInfoColumn(forecast)
	} else {
		DisplayDayInfoList(forecast)
	}
}
