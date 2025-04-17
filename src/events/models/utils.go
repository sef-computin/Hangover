package models

import "github.com/google/uuid"

func validateUUID(id string) bool{
	err := uuid.Validate(id)
	if err != nil{
		return false
	}
	return true
}

func validateEvent(e Event) bool{
  // TODO proper validation
	if e.EventID == [16]byte{} || e.EventName == ""{
		return false
	}

	return true
}
