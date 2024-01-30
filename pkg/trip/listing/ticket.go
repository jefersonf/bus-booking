package listing

import (
	"time"

	"github.com/google/uuid"
)

// Ticket describes a trip ticket.
type Ticket struct {
	ID            uuid.UUID
	OriginID      string
	DestinationID string
	CreatedAt     time.Time
}
