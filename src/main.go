package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	DefaultCity       = "gargzdai"
	DefaultColumnView = true
	DetailedListView  = false
	UsedRangeArgument = false
	DefaultDay        = time.Now().Day()
	DefaultStartHour  = time.Now().Hour()
	DefaultEndHour    = 24
)

func HandleArguments() {
	args := os.Args[1:]
	//example:  weather -c gargzdai -lv -d 0
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-r":
			i++
			if args[i] != "." {
				DefaultStartHour, _ = strconv.Atoi(args[i])
			}
			i++
			if args[i] != "." {
				DefaultEndHour, _ = strconv.Atoi(args[i])
			}
			UsedRangeArgument = true
		case "-c":
			i++
			DefaultCity = args[i]
		case "-lv":
			DefaultColumnView = false
		case "-lvi":
			DefaultColumnView = false
			DetailedListView = true
		case "-n":
			DefaultDay = time.Now().AddDate(0, 0, 1).Day()
		case "-d":
			i++
			days, _ := strconv.Atoi(args[i])
			DefaultDay = time.Now().AddDate(0, 0, days%7).Day()
		case "-h":
			fmt.Println("Usage: weather [arguments]")
			fmt.Println("Arguments:")
			fmt.Println("	-r (START|.) (END|.) - set hour display range <Default = 8 24>")
			fmt.Println("	-c CITYNAME - change default city")
			fmt.Println("	-lv - change to a list view")
			fmt.Println("	-lvi - change to a list view with more information")
			fmt.Println("	-d NUMBER - display day (0 - today, 1 - tomorrow, ...). Range 0..6 <Default = 0>")
			fmt.Println("	-n - display next day weather")
			os.Exit(0)
		}
	}
}

func main() {
	HandleArguments()
	url := fmt.Sprintf("https://api.meteo.lt/v1/places/%s/forecasts/long-term", DefaultCity)
	weather, err := ReadWeatherData(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	forecast := weather.GetForecastByDay(DefaultDay, DefaultStartHour, DefaultEndHour)

	fmt.Println("Miestas:", weather.Place.Name)

	if DefaultColumnView {
		forecast.DisplayDayInfoColumn()
	} else {
		forecast.DisplayDayInfoList(DetailedListView)
	}
}
