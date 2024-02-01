package listing

// TicketService provides ticket listing operations.
type TicketService interface {
	GetTicket(id string) (Ticket, error)
	GetTickets() []Ticket
}

type ticketService struct {
	repo TicketRepository
}

// GetTicket returns a ticket and possibly a non-nil error.
func (s *ticketService) GetTicket(id string) (Ticket, error) {
	return s.repo.GetOne(id)
}

// GetTickets returns all tickets.
func (s *ticketService) GetTickets() []Ticket {
	return s.repo.GetAll()
}

// NewTicketService creates a ticket listing service with the necessary dependencies.
func NewTicketService(r TicketRepository) TicketService {
	return &ticketService{r}
}
