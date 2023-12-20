package registering

import (
	"errors"
	"strings"

	"github.com/jefersonf/bus-booking/pkg/passenger/listing"
)

var ErrPassengerAlreadyRegisted = errors.New("passenger already registed")

// Service provides passenger registering operations.
type Servive interface {
	AddPassenger(Passenger) error
}

// Repository provides access to passenger repository.
type Repository interface {
	// AddPassenger saves a given passenger to storage.
	AddPassenger(Passenger) error
	// GetPassengers return all the passengers registered.
	GetPassengers() []listing.Passenger
}

type passengerRegisteringService struct {
	repo Repository
}

func (s *passengerRegisteringService) AddPassenger(p Passenger) error {
	registedPassengers := s.repo.GetPassengers()
	for _, rp := range registedPassengers {
		for _, rInfo := range rp.IdentityInfo {
			for _, info := range p.IdentityInfo {
				if strings.Compare(string(rInfo.Type), string(info.Type)) != 0 {
					continue
				}

				if rInfo.Value == info.Value {
					return ErrPassengerAlreadyRegisted
				}

			}
		}
	}

	err := s.repo.AddPassenger(p)
	if err != nil {
		return err
	}

	return nil
}

func NewService(r Repository) Servive {
	return &passengerRegisteringService{r}
}
