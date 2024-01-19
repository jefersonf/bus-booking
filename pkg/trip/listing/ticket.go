package listing

import "time"

// Ticket describes a trip ticket.
type Ticket struct {
	OriginID      string
	DestinationID string
	CreatedAt     time.Time 
}
