package listing

// PassengerRepository provides access to passenger storage.
type PassengerRepository interface {
	// GetOne returns the passenger with given ID.
	GetOne(id string) (Passenger, error)
	// GetAll returns all passengers saved in storage.
	GetAll() []Passenger
}
