## meteolt ##
Meteo.lt API library for go.

## Install ##

```sh
go get github.com/PauliusLozys/meteolt
```

## Example ##
``` go
import (
	"context"
	"log"

	"github.com/PauliusLozys/meteolt"
)

func main() {
	forecast, err := meteolt.GetWeatherForecast(context.Background(), "kaunas", meteolt.ForecastTypeLongTerm)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(forecast.Place.Name)
}
```

## Meteo LT API documentation ##
https://api.meteo.lt/