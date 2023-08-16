package main

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

func main() {
	apiKey := "e4257af80a583839c67ce91bebf68fe0" // à remplacer
	city := "Paris"

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur :", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var w Weather
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		fmt.Println("Erreur de décodage :", err)
		os.Exit(1)
	}

	fmt.Printf("La température à %s est %.1f°C\n", w.Name, w.Main.Temp)
}
