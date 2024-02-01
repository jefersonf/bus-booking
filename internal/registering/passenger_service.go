package registering

// PassengerService provides passenger registering operations.
type PassengerService interface {
	AddPassenger(passenger Passenger) error
	AddPassengers(p ...Passenger) []error
}

type paxService struct {
	repo PassengerRepository
}
