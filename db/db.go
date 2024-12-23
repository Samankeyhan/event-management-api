package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to db")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTables()

}

func CreateTables() {

	CreateUserTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(CreateUserTable)
	
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
		`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		log.Fatalf("Could not create events table: %v", err)
	}

	createRegistirationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES  events(id)
		FOREIGN KEY(user_id) REFERENCES  users(id)
	)
	`

	_, err = DB.Exec(createRegistirationTable)
	if err != nil {
		log.Fatalf("Error creating registrations table: %v", err)
    
	}
}
