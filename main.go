package main

import (
	"log"
	"net/http"

	frontweather "github.com/bedel225/go-weather/front-weather"
)

type Weather struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {

	http.HandleFunc("/", frontweather.AccueilHandler)

	// Route pour récupérer le pays et la ville
	http.HandleFunc("/weather", frontweather.WeatherHandler)

	// Servir fichiers statiques (CSS, images…)
	http.Handle("/static/css/",
		http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css"))),
	)

	log.Println("Serveur lancé sur http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
