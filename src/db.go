package soutien

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "MyGameList.db")
	if err != nil {
		panic(err)
	}
}

func CreateDB() {
	InitDB()
	createTableusers := `
	CREATE TABLE IF NOT EXISTS Users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT,
	email  TEXT,
	mdp VARCHAR(40)
	);
	`
	_, err := db.Exec(createTableusers)

	if err != nil {
		panic(err)
	}

	createTablecartes := `
	CREATE TABLE IF NOT EXISTS Cartes(
	carte_id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	type TEXT,
	FOREIGN KEY (user_id) REFERENCES Users(id)
	);
     `
	_, err = db.Exec(createTablecartes)

	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func InsertValue(nom, email, mdp string) int {
	InitDB()

	insertQuery := `INSERT INTO Users(username, email, mdp) VALUES(?,?,?)`
	res, err := db.Exec(insertQuery, nom, email, mdp)

	if err != nil {
		panic(err)
	}
	id, _ := res.LastInsertId()
	fmt.Println(id)

	defer db.Close()
	return int(id)

}
