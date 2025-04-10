package models

import (
)

type Event struct {
	EventID     string	  `json:"event_id"`
	EventName   string    `json:"event_name"`
	IsPublic 	  bool	  	`json:"is_public"`
	StartDt		  string    `json:"start_dt"`
	FinishDt	  string    `json:"finish_dt"`
	Description string 	  `json:"description"`
	City        string    `json:"city"`
	Geolat			float32		`json:"geolat"`
	Geolng			float32   `json:"geolng"`
}

