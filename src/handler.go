package soutien

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func LoadGames() ([]Game, error) {
	resp, err := http.Get("https://www.freetogame.com/api/games")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var data []Game
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil //bien mettre templ
}

func PlayerHandler(w http.ResponseWriter, r *http.Request) {

	games, err := LoadGames()
	if err != nil {
		http.Error(w, "errreur api", http.StatusInternalServerError) //diff erreur pr savoir d'ou vient
		return
	}

	tmpl, err := template.ParseFiles("pages/api.html", "pages/template/data.html")
	if err != nil {
		http.Error(w, "erreur du template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, games)
	if err != nil {
		http.Error(w, "errer executon du template", http.StatusInternalServerError)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func Inscription(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/inscription.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func Connexion(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/connexion.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func SetInscription(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	email := r.FormValue("emai")
	mdp := r.FormValue("password")

	id := InsertValue(username, email, mdp)
	cookie := &http.Cookie{
		Name:  "user",
		Value: strconv.Itoa(id),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusFound)
}
