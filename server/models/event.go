package models

import (
	"time"

	"github.com/igor570/eventexplorer/db"
)

type Event struct {
	ID          int64     `json:"id"`
	UserID      int       `json:"user_id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time"`
}

func (e *Event) Save() error {
	// Send data to database
	query := `
	INSERT INTO events(user_id, name, description, location, dateTime)
	VALUES(?, ?, ?, ?, ?)
	`
	val, err := db.Database.Prepare(query)

	if err != nil {
		return err
	}

	defer val.Close() //defer execution until end

	result, err := val.Exec(e.UserID, e.Name, e.Description, e.Location, e.DateTime) //pass in arguments to query, exec used for writing to db

	if err != nil {
		return err
	}

	id, err := result.LastInsertId() //get the last inserted id in DB

	if err != nil {
		return err
	}

	e.ID = id //set the structs id to last id in DB

	return err
}

func GetAllEvents() ([]Event, error) {
	// Get all events from DB
	query := `SELECT id, user_id, name, description, location, dateTime FROM events`
	rows, err := db.Database.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event //instantiate empty events

	//loops as long as there are rows to read
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.UserID, &event.Name, &event.Description, &event.Location, &event.DateTime) //read contents of rows
		if err != nil {
			return nil, err
		}

		events = append(events, event) //add read rows to empty events
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.Database.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.UserID, &event.Name, &event.Description, &event.Location, &event.DateTime)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
