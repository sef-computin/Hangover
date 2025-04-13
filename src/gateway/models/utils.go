package models

import "github.com/google/uuid"

func validateUUID(id string) bool{
	_, err := uuid.Parse(id)
	if err != nil{
		return false
	}
	return true
}

func validateEvent(e Event) bool{
  // TODO proper validation
	return true
}
