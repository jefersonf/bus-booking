package passenger

import (
	"github.com/google/uuid"
)

// Passenger describes the properties of passenger to be listed.
type Passenger struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"fullname"`
	BirthYear int       `json:"byr"`
}
