package main

import (
	"context"
	"fmt"

	"github.com/PauliusLozys/meteolt"
)

func main() {

	// forecast, err := meteolt.GetWeatherForecast(context.Background(), "gargzdai", meteolt.ForecastTypeLongTerm)
	// if err != nil {
	// 	panic(err)
	// }

	// places, err := meteolt.GetAllPlaces(context.Background())
	// if err != nil {
	// 	panic(err)
	// }

	station, err := meteolt.GetAllStations(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(station)
}
