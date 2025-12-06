package main

import (
	"fmt"

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

	fmt.Printf("La température à %s est %.1f°C\n", city, temperature)
}
