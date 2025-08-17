package registering

import (
	"github.com/jefersonf/bus-booking/internal/listing"
)

// Passenger provides access to passenger storage.
type PassengerRepository interface {
	CreatePassenger
	ReadPassenger
	UpdatePassenger
	DeletePassenger
}

type RegisteringPassenger interface {
	CreatePassenger
	ReadPassenger
	UpdatePassenger
}

type CreatePassenger interface {
	// Create persists the given passenger to storage.
	Create(passenger Passenger) error
}

type ReadPassenger interface {
	// ByID returns passenger with given ID.
	GetByID(id string) (listing.Passenger, error)
	// ByPhone returns the passenger with given phone number.
	GetByPhone(phone string) (listing.Passenger, error)
	// ByEmail returns the passenger with given email.
	GetByEmail(email string) (listing.Passenger, error)
}

type UpdatePassenger interface {
	// Update replaces passenger info with given passenger.
	Update(passenger listing.Passenger) error
}

type DeletePassenger interface {
	// Delete remove passenger from storage with given ID.
	Delete(id string) error
}
