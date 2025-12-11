package frontweather

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/bedel225/go-weather/functions"
)

type PageData struct {
	Countries []string
	Cities    map[string][]string
}

var SelectedCity string
var SelectedCountry string

// JSON convertit une structure Go en JSON utilisable dans le template
func JSON(v interface{}) template.JS {
	a, _ := json.Marshal(v)
	return template.JS(a)
}

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	countries := []string{"France", "Allemagne", "Canada", "Japon", "Brésil"}

	cities := map[string][]string{
		"France":    {"Paris", "Lyon", "Marseille"},
		"Allemagne": {"Berlin", "Munich", "Hambourg"},
		"Canada":    {"Montréal", "Toronto", "Vancouver"},
		"Japon":     {"Tokyo", "Osaka", "Kyoto"},
		"Brésil":    {"Rio de Janeiro", "São Paulo", "Brasília"},
	}

	data := PageData{
		Countries: countries,
		Cities:    cities,
	}

	tmpl := template.New("countries.html").Funcs(template.FuncMap{
		"json": JSON,
	})

	tmpl, err := tmpl.ParseFiles("templates/countries.html")
	if err != nil {
		http.Error(w, "Erreur template", http.StatusInternalServerError)
		log.Println("Erreur template:", err)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Erreur exécution:", err)
	}
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer la ville et le pays depuis le formulaire
	SelectedCity = r.FormValue("city")
	SelectedCountry = r.FormValue("country")
	apiKey := functions.ApiKey()

	temperature := functions.Temp(SelectedCity, apiKey)

	fmt.Fprintf(w, "Ville: %s\nPays: %s\nTemperature: %.1f°C", SelectedCity, SelectedCountry, temperature)
}
