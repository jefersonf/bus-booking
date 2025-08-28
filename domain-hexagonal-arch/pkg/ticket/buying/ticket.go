package buying

import (
	"time"

	"github.com/jefersonf/bus-booking/pkg/trip/listing"
)

// Ticket describes the properties of a bought ticket.
type Ticket struct {
	TripSummary        listing.Ticket `json:"trip_summary"`
	Price              float32        `json:"price"`
	Fee                float32        `json:"fee"`
	PurchasingLocation string         `json:"purchasing_location"`
	PassengerID        int            `json:"passenger_id"`
	PassengerName      string         `json:"passenger_name"`
	SeatNumber         int            `json:"seat_number"`
	ReturnDate         time.Time      `json:"return_date"`
	BoughtAt           time.Time      `json:"bought_at"`
	DepartureDatetime  time.Time      `json:"departure_datetime"`
	CreatedAt          time.Time      `json:"created_at"`
}
