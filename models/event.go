package models

import (
	"time"

	"github.com/saurabhpawar8/golang-rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {

	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err

}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			panic(err)
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		panic(err)
	}
	return &event, nil
}

func (e Event) Update() error {
	query := `UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ? `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// var id int64
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	if err != nil {
		panic(err)
	}
	return err

}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err

}
func (e *Event) Register(userId int64) error {
	query := `INSERT INTO registrations(event_id,user_id) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return err
	}
	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id=? AND user_id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return err
	}
	return err
}
