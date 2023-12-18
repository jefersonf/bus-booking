package adding

import "github.com/jefersonf/bus-booking/pkg/ticket/listing"

// Service provides ticket adding operations.
type Service interface {
	AddTicket(Ticket) error
}

// Repository provides access to ticket repository.
type Repository interface {
	// AddTicket saves the given ticket to the storage.
	AddTicket(Ticket) error
	// GetTickets return all the tickets saved.
	GetTickets() []listing.Ticket
}

type ticketAddingService struct {
	repo Repository
}

// AddTicket persists the given ticket to storage.
func (s *ticketAddingService) AddTicket(t Ticket) error {
	s.repo.AddTicket(t)
	return nil
}

// NewService creates a ticket adding service with the necessary dependencies.
func NewService(r Repository) Service {
	return &ticketAddingService{r}
}
