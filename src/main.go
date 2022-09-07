package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	defaultCity       = "gargzdai"
	listViewEnabled   = false
	detailedListView  = false
	usedRangeArgument = false
	defaultDay        = time.Now().Day()
	defaultStartHour  = time.Now().Hour()
	defaultEndHour    = 24
)

func HandleArguments() {
	args := os.Args[1:]
	//example:  weather -c gargzdai -lv -d 0
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-r":
			i++
			if args[i] != "." {
				defaultStartHour, _ = strconv.Atoi(args[i])
			}
			i++
			if args[i] != "." {
				defaultEndHour, _ = strconv.Atoi(args[i])
			}
			usedRangeArgument = true
		case "-c":
			i++
			defaultCity = args[i]
		case "-lv":
			listViewEnabled = true
		case "-lvi":
			listViewEnabled = true
			detailedListView = true
		case "-n":
			defaultDay = time.Now().AddDate(0, 0, 1).Day()
		case "-d":
			i++
			days, _ := strconv.Atoi(args[i])
			defaultDay = time.Now().AddDate(0, 0, days%7).Day()
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
	url := fmt.Sprintf("https://api.meteo.lt/v1/places/%s/forecasts/long-term", defaultCity)
	weather, err := ReadWeatherData(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	forecast := weather.GetForecastByDay(defaultDay, defaultStartHour, defaultEndHour)

	fmt.Println("Miestas:", weather.Place.Name)

	displayFn := DisplayDayInfoColumn
	if listViewEnabled {
		displayFn = DisplayDayInfoList
	}

	forecast.DisplayDayInfo(detailedListView, displayFn)
}
