package domain

import "time"

type Trip struct {
	ID             string
	Origin         string
	Destination    string
	Departure      time.Time
	Arrival        time.Time
	Price          float64
	AvailableSeats int
}
