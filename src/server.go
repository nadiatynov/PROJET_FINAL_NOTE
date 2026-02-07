package soutien

import (
	"fmt"
	"net/http"
)

func Server() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/inscription", Inscription)
	http.HandleFunc("/login", Connexion)
	http.HandleFunc("/setinfo", SetInscription)

	fmt.Println("Serveur lancé sur localhost 8080")
	http.ListenAndServe(":8080", nil)
}
