package service

import (
	"context"

	"github.com/jefersonf/bus-booking/layered-arch/internal/domain"
	"github.com/jefersonf/bus-booking/layered-arch/internal/repository"
)

type TicketService struct {
	tripRepo   repository.TripRepository
	ticketRepo repository.TicketRepository
}

// NewTicketService creates a new instance of TicketService.
func NewTicketService(tripRepo repository.TripRepository, ticketRepo repository.TicketRepository) *TicketService {
	return &TicketService{
		tripRepo:   tripRepo,
		ticketRepo: ticketRepo,
	}
}

// ListAvailableTrips lists available trips based on origin and destination.
func (s *TicketService) ListAvailableTrips(ctx context.Context, origin, destination string) ([]domain.Trip, error) {
	return s.tripRepo.FindAvailableTrips(ctx, origin, destination)
}

// PurchaseTicket processes the ticket purchase for a given trip and passenger.
func (s *TicketService) PurchaseTicket(ctx context.Context, tripID, passengerID, seatNumber string) (*domain.Ticket, error) {
	// TODO 1. Validate trip existence and seat availability.
	// TODO 2. Process payment
	// TODO 3. Create and store the ticket.
	// TODO 4. Deduct the seat from available seats.
	// TODO 5. Return the ticket details.
	return nil, nil
}
