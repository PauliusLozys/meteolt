package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	Separator = "================================"
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

func GetInput() int {
	fmt.Print("> ")
	var input string
	fmt.Scanf("%v", &input)
	if input == "h" {
		return 1
	}
	if input == "sd" {
		return 2
	}
	if input == "sv" {
		return 3
	}
	if input == "q" {
		return 4
	}
	return 0
}

func GetNumericInput() int {
	fmt.Print("> ")
	var input string
	fmt.Scanf("%v", &input)
	val, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return val
}

func HandleInput(input int, days [][]Forecast) {
	switch input {
	case 1:
		fmt.Println("Available commands:")
		fmt.Println("\th - display this page\n\tsd - switch day\n\tsv - switch view\n\tq - quit\nPress Enter to continue")
		GetInput()
	case 2:
		fmt.Println("Select which day to display:")
		for idx, d := range days {
			fmt.Printf("\t%v - %v-%v\n", idx, d[0].FormattedTime.Month(), d[0].FormattedTime.Day())
		}
		val := GetNumericInput()
		if val <= 6 && val >= 0 {
			DefaultDayIndex = val
		}
		time.Sleep(1*time.Second)
	case 3:
		fmt.Println("What view to use?:\n\t1 - Column view\n\t2 - List view")
		val := GetNumericInput()
		if(val == 1) {
			DefaultColumnView = true
		} else if val == 2 {
			DefaultColumnView = false
		}
	}
}


func Draw(seperated [][]Forecast) {
	if(DefaultColumnView) {
		DisplayDayInfoColumn(DefaultDayIndex, seperated)
	} else {
		DisplayDayInfoList(DefaultDayIndex, seperated)
	}
}

func ClearConsole() {
	cmd := exec.Command("clear") // Linux only because Windows just doesn't want to work
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	url := "https://api.meteo.lt/v1/places/gargzdai/forecasts/long-term"
	weather := ReadData(url)
	seperated := weather.SeparateByDay()

	for {
		ClearConsole()
		fmt.Println("Miestas:", weather.Place.Name)
		Draw(seperated)
		input := GetInput()
		if input == 4 {
			return
		}
		HandleInput(input, seperated)

	}

	//DisplayDayInfoList(0, seperated)
	//DisplayDayInfoColumn(0, seperated)

}


