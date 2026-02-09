package soutien

import (
	"encoding/json"
	"fmt"
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

	games, err := LoadGames() //charm de api changer pr mettre dashboard
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	cookie := &http.Cookie{
		Name:  "user",
		Value: strconv.Itoa(userId),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusFound) //!!!!!!!!!!!!!! a changer quand inscription marche pr redirect vers /dashboard

}

func Verifconnect(username, mdp string) int { //int car return id pas string att
	InitDB()

	row := db.QueryRow("SELECT id FROM Users WHERE username = ? AND mdp = ?", username, mdp) //requete pr chercher user a username et mpd

	var id int
	err := row.Scan(&id) //on recup id avec * car garder en memoire (revoir note soutien pointeur)
	if err != nil {
		return 0
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

	http.Redirect(w, r, "/", http.StatusFound) //!!!!!!!!!!!!!! a changer quand inscription marche pr redirect vers /dashboard
}

type Card struct {
	Nom   string
	Type  string
	Image string
}

func Dashboard(w http.ResponseWriter, r *http.Request) { //btn marche pas car pas encore de connexion
	cookie, err := r.Cookie("user")       //tester quand inscription et connexion amrche pr voir si gestion err fonctionne
	if err != nil || cookie.Value == "" { //enlevé le 2eme if et mis avec le ou logique revu du cours immersions + rapide
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	userId, err := strconv.Atoi(cookie.Value) //bloc fait IA car besoin de id + tard pr connaitre nbr de carte, collection etc par user
	if err != nil || userId <= 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	total := CartePossede(userId) //car pr chaque user grace a ID donc met var (userId) recup nbr carte vu que c compter pas oublier count ds requete

	collection := CollectionPerso(userId) //recup collection

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
	db.Close()
	return cards
}

func Pack(w http.ResponseWriter, r *http.Request) {

}
