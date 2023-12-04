package adding

import (
	"errors"

	"github.com/jefersonf/bus-booking/pkg/fleet/listing"
)

var ErrBusAlreadyExists = errors.New("bus already exists")

// Service provides bus adding operations.
type Service interface {
	AddBus(...Bus) error
}

// Repository provides access to bus repository.
type Repository interface {
	// AddBus saves a given bus to the storage.
	AddBus(Bus) error
	// GetFeet returns all the buses saved.
	GetFleet() []listing.Bus
}

type busAddingService struct {
	repo Repository
}

// AddBus persists the given bus(es) to storage.
func (s *busAddingService) AddBus(b ...Bus) error {
	fleet := s.repo.GetFleet()
	for _, nb := range b {
		for _, bus := range fleet {
			if nb.Code == bus.Code && nb.CompanyName == bus.CompanyName {
				return ErrBusAlreadyExists
			}
		}
	}

	for _, bus := range b {
		err := s.repo.AddBus(bus)
		if err != nil {
			return err
		}
	}

	return nil
}

// NewService creates a bus adding service with the necessary dependencies.
func NewService(r Repository) Service {
	return &busAddingService{r}
}
