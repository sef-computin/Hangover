package dbhandler

import (
	"database/sql"
	"errors"
	"fmt"
)

type EventDB interface {
	GetAllEvents() ([]*models.Event, error)
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
		if err := rows.Scan(&f.EventID, &f.EventName, &f.StartDt, &f.EndDt, &f.IsPublic, &f.Description, &f.Geolng, &f.Geoltd); err != nil {
			return nil, fmt.Errorf("failed to execute the query: %w", err)
		}
		events = append(events, f)
	}

	defer rows.Close()

	return events, nil
}

