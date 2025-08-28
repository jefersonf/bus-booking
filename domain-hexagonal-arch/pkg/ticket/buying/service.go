package buying

import (
	"errors"
	"math"

	"github.com/jefersonf/bus-booking/pkg/ticket/listing"
)

const minSecondsDurationBetweenDuplicateTicketCreation = 30.

var ErrTicketAlreadyExists = errors.New("ticket already exists")

// Service provides ticket buying operations.
type Service interface {
	AddTicket(Ticket) error
	GetTickets() []listing.Ticket
	GetTripDetails(string) listing.Ticket
}

// Repository provides access to ticket repository.
type Repository interface {
	// AddTicket saves the given ticket to the storage.
	AddTicket(Ticket) error
	// GetTickets return all the tickets saved.
	GetTickets() []listing.Ticket
	// GetTripDetails return trip details of a given ticket ID.
	GetTripDetails(string) listing.Ticket
}

type buyingTicketService struct {
	repo Repository
}

// AddTicket persists the given ticket to storage.
func (s *buyingTicketService) AddTicket(t Ticket) error {
	tickets := s.repo.GetTickets()
	for _, et := range tickets {
		if et.PassengerID != t.PassengerID {
			continue
		}
		if et.OriginID != t.TripSummary.OriginID {
			continue
		}
		if et.DestinationID != t.TripSummary.DestinationID {
			continue
		}
		if et.PurchasingLocation != t.PurchasingLocation {
			continue
		}
		if et.SeatNumber != t.SeatNumber {
			continue
		}
		if math.Abs(et.CreatedAt.Sub(t.CreatedAt).Seconds()) > minSecondsDurationBetweenDuplicateTicketCreation {
			continue
		}

		return ErrTicketAlreadyExists
	}

	err := s.repo.AddTicket(t)
	if err != nil {
		return err
	}
	return nil
}

// GetTickets returns all saved tickets.
func (s *buyingTicketService) GetTickets() []listing.Ticket {
	return s.repo.GetTickets()
}

// GetTripDetails returns the trip ticket details from storage.
func (s *buyingTicketService) GetTripDetails(id string) listing.Ticket {
	return s.repo.GetTripDetails(id)
}

// NewService creates a ticket buying service with the necessary dependencies.
func NewService(r Repository) Service {
	return &buyingTicketService{r}
}
