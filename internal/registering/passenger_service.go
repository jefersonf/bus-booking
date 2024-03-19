package registering

import (
	"errors"
	"fmt"
)

var (
	ErrPassengerEmailAlreadyRegistered = errors.New("passenger email already registered")
	ErrPassengerPhoneAlreadyRegistered = errors.New("passenger phone already registered")
)

// PassengerService provides passenger registering operations.
type PassengerService interface {
	AddPassenger(Passenger) error
	AddPassengers(...Passenger) []error
}

type passengerService struct {
	repo PassengerRepository
}

// AddPassenger adds a new passenger and possibly returns a non-nil error.
func (s *passengerService) AddPassenger(p Passenger) error {
	_, err := s.repo.GetByEmail(p.Email)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrPassengerEmailAlreadyRegistered, err)
	}
	_, err = s.repo.GetByPhone(p.Phone)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrPassengerPhoneAlreadyRegistered, err)
	}

	return s.repo.Create(p)
}

// AddPassengers add multiple passengers and possibly retruns a list of non-nil errors.
func (s *passengerService) AddPassengers(pp ...Passenger) []error {
	errors := make([]error, 0, len(pp))
	for _, p := range pp {
		err := s.AddPassenger(p)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

// NewPassengerService creates a passenger service with necessary dependencies.
func NewPassengerService(r PassengerRepository) PassengerService {
	return &passengerService{r}
}
