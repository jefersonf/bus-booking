package overview

import "time"

// Trip describes the properties of a trip overview.
type Trip struct {
	Origin              string        `json:"departing_from"`
	Destination         string        `json:"going_to"`
	EstimatedTravelTime time.Duration `json:"ett"`
}

func (t *Trip) String() string {
	return fmt.Sprintf("%s to %s in %v", t.Origin, t.Destination, t.EstimatedTravelTime)
}
