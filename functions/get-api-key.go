package functions

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ApiKey() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Impossible de charger le fichier .env")
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	return apiKey
}
