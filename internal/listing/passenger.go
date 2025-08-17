package listing

import (
	"github.com/google/uuid"
)

// Passenger describes the properties of passenger to be listed.
type Passenger struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"fullname"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	BirthYear int       `json:"byr"`
}
