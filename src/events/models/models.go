package models

import (
	"time"
)

type Event struct {
	EventID     string	  `json:"event_id"`
	EventName   string    `json:"event_name"`
	IsPublic 	  bool		  `json:"is_public"`
	StartDt		  time.Time `json:"start_dt"`
	FinishDt	  time.Time `json:"finish_dt"`
	Description string 	  `json:"description"`
	Geolat			float32		`json:"geolat"`
	Geolng			float32   `json:"geolng"`
}

