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
