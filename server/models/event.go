package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time"`
}

// Storing the events in a slice
var events = []Event{}

func (e *Event) Save() {
	// Send data to database
	events = append(events, *e)
}

func GetAllEvents() []Event {
	// Get all events from DB
	return events
}
