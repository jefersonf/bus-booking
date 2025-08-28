package repository

import (
	"context"

	"github.com/jefersonf/bus-booking/layered-arch/internal/domain"
)

// TripRepository defines the interface for trip-related data operations.
type TripRepository interface {
	FindAvailableTrips(ctx context.Context, origin, destination string) ([]domain.Trip, error)
	GetAllTrips(ctx context.Context) ([]domain.Trip, error)
	GetTripByID(ctx context.Context, id string) (*domain.Trip, error)
}

// TicketRepository defines the interface for ticket-related data operations.
type TicketRepository interface {
	CreateTicket(ctx context.Context, ticket *domain.Ticket) error
	FindByPassengerID(ctx context.Context, passengerID string) ([]domain.Ticket, error)
}
