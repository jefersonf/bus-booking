package listing

import "time"

// Ticket defines the properties of a ticket to be listed.
type Ticket struct {
	ID                 string    `json:"id"`
	OriginID           string    `json:"origin_id"`
	DestinationID      string    `json:"destination_id"`
	PurchasingLocation string    `json:"purchasing_location"`
	PassengerID        int       `json:"passenger_id"`
	SeatNumber         int       `json:"seat_number"`
	PricePaid          float32   `json:"price_paid"`
	DepartureDatetime  time.Time `json:"departure_datetime"`
}
