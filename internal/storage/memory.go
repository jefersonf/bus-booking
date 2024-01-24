package storage

import (
	busListing "github.com/jefersonf/bus-booking/pkg/bus/listing"
	ticketListing "github.com/jefersonf/bus-booking/pkg/ticket/listing"
)

type MemoryStorage struct {
	tickets map[string]ticketListing.Ticket
	buses   map[string]busListing.Bus
}
