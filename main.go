package main

import (
	"fmt"

	frontweather "github.com/bedel225/go-weather/front-weather"
	"github.com/bedel225/go-weather/functions"
)

type Weather struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	apiKey := functions.ApiKey()
	//city := "Paris"
	city := "abidjan"

	temperature := functions.Temp(city, apiKey)

	message := fmt.Sprintf("La température à %s est %.1f°C\n", city, temperature)

	frontweather.Index(message)
}
