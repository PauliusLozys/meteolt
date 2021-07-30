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
	Separator = "================================"
	DefaultCity = "gargzdai"
	DefaultColumnView = true
	DefaultDayIndex = 0
	Gray      = "\033[37m"
	Reset     = "\033[0m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Purple    = "\033[35m"
	Cyan      = "\033[36m"
	White     = "\033[97m"
)

type Forecast struct {
	ForecastTimeUtc string  `json:"forecastTimeUtc"`
	AirTemperature  float32 `json:"airTemperature"`
	CloudCover      int     `json:"cloudCover"`
	ConditionCode   string  `json:"conditionCode"`
	FormattedTime   time.Time
}

type Weather struct {
	Place struct {
		Code    string `json:"code"`
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"place"`
	ForecastTimestamps []Forecast `json:"forecastTimestamps"`
}

func (w *Weather) SeparateByDay() [][]Forecast {
	var result [][]Forecast
	var dayForecast []Forecast
	current, _ := time.Parse("2006-01-02 15:04:05", w.ForecastTimestamps[0].ForecastTimeUtc)

	for _, day := range w.ForecastTimestamps {
		t, _ := time.Parse("2006-01-02 15:04:05", day.ForecastTimeUtc)
		if current.Day() != t.Day() {
			current = t
			result = append(result, dayForecast)
			dayForecast = []Forecast{}
		}
		day.FormattedTime = t
		dayForecast = append(dayForecast, day)
	}

	return result
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

func TemperatureColor(tempreture float32) string {
	switch {
	case tempreture <= 12:
		return Blue
	case tempreture > 12 && tempreture < 20:
		return Cyan
	case tempreture >= 20 && tempreture < 25:
		return Green
	case tempreture >= 25:
		return Red
	default:
		return Reset
	}
}

func DisplayDayInfoList(day int, forecasts [][]Forecast) {
	fmt.Println("Diena:", forecasts[day][0].FormattedTime.Day())
	var topTemperature float32 = 0
	var averageTempreture float32 = 0
	for _, hour := range forecasts[day] {
		if topTemperature < hour.AirTemperature {
			topTemperature = hour.AirTemperature
		}
		if hour.FormattedTime.Hour() == time.Now().Hour() {
			fmt.Printf(Purple+"Laikas: %vh:"+TemperatureColor(hour.AirTemperature)+" %v\n"+Reset,
				hour.FormattedTime.Hour(), hour.AirTemperature)
		} else {
			fmt.Printf(Reset+"Laikas: %vh:"+TemperatureColor(hour.AirTemperature)+" %v\n"+Reset,
				hour.FormattedTime.Hour(), hour.AirTemperature)
		}
		averageTempreture += hour.AirTemperature
	}
	fmt.Println(Separator)
	fmt.Printf("Aukščiausia temperatūra: %v°C\n", topTemperature)
	fmt.Printf("Vidutinė temperatūra: %.1f°C\n", averageTempreture/float32(len(forecasts[day])))
	fmt.Println(Separator)
}

func DisplayDayInfoColumn(day int, forecasts [][]Forecast) {
	Separator := strings.Repeat("=", 9*len(forecasts[day]))
	fmt.Println("Diena:", forecasts[day][0].FormattedTime.Day())
	fmt.Println(Separator)

	var topTemperature float32 = 0
	var averageTempreture float32 = 0

	for _, hour := range forecasts[day] {
		if topTemperature < hour.AirTemperature {
			topTemperature = hour.AirTemperature
		}
		if hour.FormattedTime.Hour() == time.Now().Hour() {
			fmt.Printf(Purple+" %-7v"+Reset+"|", fmt.Sprintf("%vh", hour.FormattedTime.Hour()))
		} else {
			fmt.Printf(" %-7v|", fmt.Sprintf("%vh", hour.FormattedTime.Hour()))
		}
		averageTempreture += hour.AirTemperature
	}
	fmt.Println()
	for _, hour := range forecasts[day] {
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
	fmt.Printf("Vidutinė temperatūra: %.1f°C\n", averageTempreture/float32(len(forecasts[day])))
	fmt.Println(Separator)
}

func HandleArguments() {
	args := os.Args[1:]
	//example:  weather -c gargzdai -lv -d 0
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-c":
			i++
			DefaultCity = args[i]
		case "-lv":
			DefaultColumnView = false
		case "-d":
			i++
			day, _ := strconv.Atoi(args[i])
			DefaultDayIndex = day
		case "-h":
			fmt.Println("Usage: weather [arguments]")
			fmt.Println("Arguments\n\t-c CITYNAME - change default city\n\t-lv - change to a list view\n\t" +
				"-d NUMBER - display day (0 - today, 1 - tomorrow, ...). Range 0..6 <Default = 0>")
			os.Exit(0)
		}
	}
}

func main() {

	HandleArguments()
	url := "https://api.meteo.lt/v1/places/"+ DefaultCity +"/forecasts/long-term"
	weather := ReadData(url)
	seperated := weather.SeparateByDay()

	fmt.Println("Miestas:", weather.Place.Name)

	if DefaultColumnView {
		DisplayDayInfoColumn(DefaultDayIndex, seperated)
	} else {
		DisplayDayInfoList(DefaultDayIndex, seperated)
	}
}