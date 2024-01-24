package storage

import "github.com/jefersonf/bus-booking/pkg/ticket/listing"

type MemoryStorage struct {
	tickets map[string]listing.Ticket
}
