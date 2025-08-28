package listing

// Service provides ticket listing operations.
type Service interface {
	GetTicket(string) (Ticket, error)
	GetTickets() []Ticket
}

// Repository provides access to the ticket storage.
type Repository interface {
	// GetTicket returns the ticket with given ID.
	GetTicket(string) (Ticket, error)
	// GetTickets returns all tickets saved in storage.
	GetTickets() []Ticket
}

type ticketListingService struct {
	repo Repository
}

// GetTicket returns a ticket.
func (s *ticketListingService) GetTicket(id string) (Ticket, error) {
	return s.repo.GetTicket(id)
}

// GetTickets returns all tickets.
func (s *ticketListingService) GetTickets() []Ticket {
	return s.repo.GetTickets()
}

// NewService creates a ticket listing service with the necessary dependencies.
func NewService(r Repository) Service {
	return &ticketListingService{r}
}
