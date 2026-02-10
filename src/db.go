package soutien

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "MyAppList.db")
	if err != nil {
		panic(err)
	}
	fmt.Println("InitDB called")
}

func CreateDB() {
	InitDB()
	createTableusers := `
	CREATE TABLE IF NOT EXISTS Users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	email  TEXT NOT NULL,
	mdp TEXT NOT NULL
	);
	`
	_, err := db.Exec(createTableusers)

	if err != nil {
		panic(err)
	}

	createTablecartes := `
	CREATE TABLE IF NOT EXISTS Cartes(
	carte_id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	type TEXT NOT NULL,
	image TEXT 
	);
     ` //ajout de image par marjane
	_, err = db.Exec(createTablecartes)

	if err != nil {
		panic(err)
	}

	createTableUserCarte := `
	CREATE TABLE IF NOT EXISTS UserCarte(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER,
	carte_id INTEGER,
	FOREIGN KEY (user_id) REFERENCES Users(id),
	FOREIGN KEY (carte_id) REFERENCES Cartes(carte_id)
	);
	` //ajout de la table par marjane

	_, err = db.Exec(createTableUserCarte)

	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func InsertValue(nom, email, mdp string) int {
	InitDB()

	hash, err := HashPassword(mdp)

	if err != nil {
		panic(err)
	}

	insertQuery := `INSERT INTO Users(username, email, mdp) VALUES(?,?,?)`
	res, err := db.Exec(insertQuery, nom, email, hash)

	if err != nil {
		panic(err)
	}
	id, _ := res.LastInsertId()
	fmt.Println(id)

	defer db.Close()
	return int(id)

}
