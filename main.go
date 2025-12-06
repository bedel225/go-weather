package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Weather struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Impossible de charger le fichier .env")
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	city := "Paris"

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

	fmt.Printf("La température à %s est %.1f°C\n", w.Name, w.Main.Temp)
}
