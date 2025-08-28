package listing

// TicketRepository provides access to the ticket storage.
type TicketRepository interface {
	// GetOne returns the ticket with given ID.
	GetOne(id string) (Ticket, error)
	// GetAll returns all tickets saved in storage.
	GetAll() []Ticket
}
