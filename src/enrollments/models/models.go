package models

import (

	"github.com/google/uuid"
)

type Enrollment struct{
	EnrollmentID		uuid.UUID `json:"enroll_id"`
	UserID          uuid.UUID `json:"user_id"`
	EventID         uuid.UUID `json:"event_id"`
	Status 					int8			`json:"status"`
}
