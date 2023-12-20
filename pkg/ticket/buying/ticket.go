package buying

import "github.com/jefersonf/bus-booking/pkg/trip/overview"

type Ticket struct {
	TripSummary overview.Trip `json:"trip_summary"`
	Price       float32       `json:"price"`
	Fee         float32       `json:"fee"`
}
