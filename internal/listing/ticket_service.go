package listing

// TicketService provides ticket listing operations.
type TicketService interface {
	GetTicket(string) (Ticket, error)
	GetTickets() []Ticket
}

// TicketRepository provides access to the ticket storage.
type TicketRepository interface {
	// GetTicket returns the ticket with given ID.
	GetTicket(string) (Ticket, error)
	// GetTickets returns all tickets saved in storage.
	GetTickets() []Ticket
}

type ticketService struct {
	repo TicketRepository
}

// GetTicket returns a ticket.
func (s *ticketService) GetTicket(id string) (Ticket, error) {
	return s.repo.GetTicket(id)
}

// GetTickets returns all tickets.
func (s *ticketService) GetTickets() []Ticket {
	return s.repo.GetTickets()
}

// NewTicketService creates a ticket listing service with the necessary dependencies.
func NewTicketService(r TicketRepository) TicketService {
	return &ticketService{r}
}
