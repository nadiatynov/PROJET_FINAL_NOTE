package soutien

import (
	"fmt"
	"net/http"
)

type Game struct {
	Nom   string `json:"title"` //att bien metttre les champs de l'api et pas ceux du sujet
	Type  string `json:"genre"`
	Image string `json:"thumbnail"`
}

func Server() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/inscription", Inscription)
	http.HandleFunc("/login", Connexion)
	http.HandleFunc("/setconnect", SetConnexion)
	http.HandleFunc("/setinfo", SetInscription)
	http.HandleFunc("/player", PlayerHandler)
	http.HandleFunc("/dashboard", Dashboard)
	http.HandleFunc("/deconnex", Deconnexion)
	http.HandleFunc("/pack", Pack)

	fmt.Println("Serveur lancé sur localhost 8080")
	http.ListenAndServe(":8080", nil)
}
