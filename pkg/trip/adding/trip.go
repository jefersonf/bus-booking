package adding

import "time"

// Trip describes a trip.
type Trip struct {
	OriginID      string
	DestinationID string
	CreatedAt     time.Time
}
