package frontweather

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Countries []string
	Cities    map[string][]string
}

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
