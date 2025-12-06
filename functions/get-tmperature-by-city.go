package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Weather struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func Temp(city, apiKey string) float64 {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Erreur :", err)
		os.Exit(1)
	}

	if resp.StatusCode == 401 {
		fmt.Println("la clé API est invalide.")
		os.Exit(1)
	}

	defer resp.Body.Close()

	var w Weather
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		fmt.Println("Erreur de décodage :", err)
		os.Exit(1)
	}
	return w.Main.Temp
}
