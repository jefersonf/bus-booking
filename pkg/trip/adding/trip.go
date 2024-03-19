package adding

import "time"

// Trip describes a trip to be added.
type Trip struct {
	OriginID      string
	DestinationID string
	CreatedAt     time.Time
}
