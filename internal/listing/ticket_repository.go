package listing

// TicketRepository provides access to the ticket storage.
type TicketRepository interface {
	// GetTicket returns the ticket with given ID.
	GetTicket(string) (Ticket, error)
	// GetTickets returns all tickets saved in storage.
	GetTickets() []Ticket
}
