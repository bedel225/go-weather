package main

import (
	"fmt"
	"log"
	"net/http"

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

	http.HandleFunc("/", frontweather.AccueilHandler)

	// Servir fichiers statiques (CSS, images…)
	http.Handle("/static/css/",
		http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css"))),
	)

	log.Println("Serveur lancé sur http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	apiKey := functions.ApiKey()
	//city := "Paris"
	city := "abidjan"

	temperature := functions.Temp(city, apiKey)

	message := fmt.Sprintf("La température à %s est %.1f°C\n", city, temperature)

	frontweather.Weather(message)
}
