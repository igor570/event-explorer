package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	MAX_CONNECTIONS      = 10
	MAX_IDLE_CONNECTIONS = 5
)

var Database *sql.DB

func InitDB() {
	var err error
	Database, err = sql.Open("sqlite3", "api.db") // Open a connection to SQLite

	if err != nil {
		panic("No database connection")
	}

	Database.SetMaxOpenConns(MAX_CONNECTIONS)
	Database.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL
	)
	`

	_, err := Database.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Could not create events table: %v", err)
	}
}
