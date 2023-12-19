package listing

// Service provides ticket listing operations.
type Service interface {
	GetTicket(string) (Ticket, error)
	GetTickets() []Ticket
}
