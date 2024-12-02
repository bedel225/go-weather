package frontweather

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Countries []string
}

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Données
	countries := []string{"France", "Allemagne", "Canada", "Japon", "Brésil"}
	data := PageData{Countries: countries}

	// 2. Charger template
	tmpl, err := template.ParseFiles("templates/countries.html")
	if err != nil {
		http.Error(w, "Erreur template", http.StatusInternalServerError)
		log.Println("Erreur template:", err)
		return
	}

	// 3. Exécuter template
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Erreur exécution:", err)
	}
}
