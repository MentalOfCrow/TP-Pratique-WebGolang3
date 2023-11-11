package main

import (
	"html/template"
	"net/http"
)

type UserData struct {
	Nom           string
	Prenom        string
	DateNaissance string
	Sexe          string
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/user/init", userInitHandler)
	http.HandleFunc("/user/treatment", userTreatmentHandler)
	http.HandleFunc("/user/display", userDisplayHandler)

	http.ListenAndServe(":8080", nil)
}

func userInitHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "userInputForm", nil)
}

func userTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	nom := r.FormValue("nom")
	prenom := r.FormValue("prenom")
	dateNaissance := r.FormValue("dateNaissance")
	sexe := r.FormValue("sexe")

	// Logique de traitement des données (à personnaliser selon vos besoins)

	// Redirection vers la page d'affichage des données
	http.Redirect(w, r, "/user/display?nom="+nom+"&prenom="+prenom+"&dateNaissance="+dateNaissance+"&sexe="+sexe, http.StatusSeeOther)
}

func userDisplayHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérez les données de l'URL
	nom := r.FormValue("nom")
	prenom := r.FormValue("prenom")
	dateNaissance := r.FormValue("dateNaissance")
	sexe := r.FormValue("sexe")

	// Créez une structure de données pour les informations de l'utilisateur
	data := UserData{
		Nom:           nom,
		Prenom:        prenom,
		DateNaissance: dateNaissance,
		Sexe:          sexe,
	}

	tmpl, err := template.ParseFiles("template/userDataDisplay.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "userDataDisplay", data)
}
