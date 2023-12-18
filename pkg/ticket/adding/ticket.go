package adding

// Ticket decribes the properties of a ticket to be added.
type Ticket struct {
	OriginID           string  `json:"origin_id"`
	DestinationID      string  `json:"destination_id"`
	PurchasingLocation string  `json:"purchasing_location"`
	PassengerID        int     `json:"passenger_id"`
	PassengerName      string  `json:"passenger_name"`
	SeatNumber         int     `json:"seat_number"`
	Price              float32 `json:"price"`
}
