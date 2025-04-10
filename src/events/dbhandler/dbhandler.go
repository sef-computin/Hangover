package dbhandler

import (
	"database/sql"
	"fmt"

	"github.com/sef-comp/Hangover/events/models"
)

type EventDB interface {
	GetAllEvents() ([]*models.Event, error)
	CreateEvent(*models.Event) error
	DeleteEventByID(string) error
}

type DBHandler struct {
	db *sql.DB
}

func InitDBHandler(db *sql.DB) *DBHandler {
	return &DBHandler{
		db: db,
	}
}

func (dbhand *DBHandler) GetAllEvents() ([]*models.Event, error) {

	var events []*models.Event
	rows, err := dbhand.db.Query(`SELECT * FROM business.events;`)
	if err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	for rows.Next() {
		f := new(models.Event)
		if err := rows.Scan(&f.EventID, &f.EventName, &f.IsPublic, &f.StartDt, &f.FinishDt, &f.Description, &f.City, &f.Geolng, &f.Geolat); err != nil {
			return nil, fmt.Errorf("failed to execute the query: %w", err)
		}
		events = append(events, f)
	}

	defer rows.Close()

	return events, nil
}

func (dbhand *DBHandler) CreateEvent(event *models.Event) error {

	_, err := dbhand.db.Query(
		`INSERT INTO business.events (event_id, event_name, is_public, start_dt, finish_dt, description, city, geolat, geolng) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`,
		event.EventID,
		event.EventName,
		event.IsPublic,
		event.StartDt,
		event.FinishDt,
		event.Description,
		event.City,
		event.Geolat,
		event.Geolng,
	)

	return err
}

func (dbhand *DBHandler) DeleteEventByID(event_id string) error {
	_, err := dbhand.db.Query(
		`DELETE FROM business.events WHERE event_id=$1`,
		event_id,
	)
	return err
}
