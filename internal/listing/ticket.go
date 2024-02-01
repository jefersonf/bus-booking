package listing

import (
	"time"

	"github.com/google/uuid"
)

// Ticket describes the properties of a ticket to be listed.
type Ticket struct {
	ID                 uuid.UUID `json:"id"`
	OriginID           string    `json:"orig_id"`
	DestinationID      string    `json:"dest_id"`
	PurchasingLocation string    `json:"purchasing_location"`
	PassengerID        int       `json:"pax_id"`
	SeatNumber         int       `json:"seat_number"`
	PricePaid          float32   `json:"price_paid"`
	DepartureDatetime  time.Time `json:"departure_datetime"`
	CreatedAt          time.Time `json:"created_at"`
}
