package registering

import "github.com/jefersonf/bus-booking/pkg/passenger/listing"

// PassengerRepository provides access to passenger storage.
type PassengerRepository interface {
	// Create persists the given passenger to storage.
	Create(passenger Passenger) error
	// GetByPhone returns the passenger with given phone number.
	GetByPhone(phone string) (listing.Passenger, error)
	// GetByEmail returns the passenger with given email.
	GetByEmail(email string) (listing.Passenger, error)
	// Update replaces passenger info with given passenger.
	Update(passenger listing.Passenger) error
}
