package soutien

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func LoadGames() ([]Game, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/{id_du_pokemon}") // l'api ne renvoie pas un simple tableau
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

	games, err := LoadGames() //charm de api changer pr mettre dashboard
	if err != nil {
		http.Error(w, "erreur api", http.StatusInternalServerError) //diff erreur pr savoir d'ou vient
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
		http.Error(w, "erreur template connex", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func SetConnexion(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	mdp := r.FormValue("password")

	userId := Verifconnect(username, mdp)
	if userId == 0 { //si pas de user ou mauvais identfiant on revoie a P princiapl
		http.Redirect(w, r, "/", http.StatusSeeOther) //err pas inscription avant renvoie a p /
		return
	}

	cookie := &http.Cookie{
		Name:  "user",
		Value: strconv.Itoa(userId),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/dashboard", http.StatusFound) //!!!!!!!!!!!!!! a changer quand inscription marche pr redirect vers /dashboard

}

var id int //en glob car besoin ds pack remettre ds verifconnect si prblm

func Verifconnect(username, pwd string) int { //int car return id pas string att
	InitDB()

	row := db.QueryRow("SELECT id , mdp FROM Users WHERE username = ?", username) //requete pr chercher user a username et mpd
	var id int
	var hashedFromDB string
	err := row.Scan(&id, &hashedFromDB) //on recup id avec * car garder en memoire (revoir note soutien pointeur)
	if err != nil {
		db.Close()
		return 0
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedFromDB), []byte(pwd)) != nil {
		db.Close()
		return 0 // mauvais mot de passe
	}

	db.Close()
	return id //trouvé
}

func SetInscription(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	email := r.FormValue("email")
	mdp := r.FormValue("password")

	id := InsertValue(username, email, mdp)

	cookie := &http.Cookie{
		Name:  "user",
		Value: strconv.Itoa(id),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/?msg=signup_ok", http.StatusSeeOther) //!!!!!!!!!!!!!! a changer quand inscription marche pr redirect vers /dashboard
}

type Card struct { //ps oublier appeler pr handler pack
	Nom   string
	Type  string
	Image string
}

func Dashboard(w http.ResponseWriter, r *http.Request) { //btn marche pas car pas encore de connexion
	cookie, err := r.Cookie("user")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	userID, _ := strconv.Atoi(cookie.Value)

	total := CartePossede(userID) //car pr chaque user grace a ID donc met var (userId) recup nbr carte vu que c compter pas oublier count ds requete

	collection := CollectionPerso(userID) //recup collection

	Data := struct { //autre maniere de faire struct revoir tp soutien 1
		Total      int
		Collection []Card
	}{
		Total:      total,
		Collection: collection,
	}

	tmpl, err := template.ParseFiles("pages/dashboard.html")
	if err != nil {
		http.Error(w, "erreru template dashboard", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, Data)

}

func CartePossede(userId int) int {
	InitDB() //appelle la BDD

	row := db.QueryRow("SELECT COUNT(*) FROM UserCarte WHERE user_id = ?", userId)

	var count int
	row.Scan(&count) // IA pointeur

	db.Close()
	return count
}

func CollectionPerso(userId int) []Card {
	InitDB()

	query := `
	SELECT C.name, C.type, C.image
	FROM UserCarte UC
	JOIN Cartes C ON UC.carte_id = C.carte_id
	WHERE UC.user_id = ?
	`

	rows, err := db.Query(query, userId) //ds var query car + long et appeler jsute var = -long
	if err != nil {
		fmt.Println("erreur func collection", err)
		return nil
	}
	var cards []Card

	for rows.Next() {
		var c Card
		rows.Scan(&c.Nom, &c.Type, &c.Image) // IA pointeur
		cards = append(cards, c)
	}
	return cards
}

func Deconnexion(w http.ResponseWriter, r *http.Request) { //coller de ancien tp revoir au cas ou
	cookie := &http.Cookie{
		Name:   "user",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/login", http.StatusFound)

}
