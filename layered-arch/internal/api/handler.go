package api

import (
	"encoding/json"
	"net/http"

	"github.com/jefersonf/bus-booking/layered-arch/internal/service"
)

type TicketHandler struct {
	service *service.TicketService
}

// NewTicketHandler creates a new instance of TicketHandler.
func NewTicketHandler(s *service.TicketService) *TicketHandler {
	return &TicketHandler{service: s}
}

// RegisterRoutes registers the HTTP routes for ticket handling.
func (h *TicketHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/trips", h.handleListAvailableTrips)
}

// handleListAvailableTrips handles the HTTP request to list available trips.
func (h *TicketHandler) handleListAvailableTrips(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	origin := r.URL.Query().Get("origin")
	destination := r.URL.Query().Get("destination")
	if origin == "" || destination == "" {
		http.Error(w, "Missing origin or destination", http.StatusBadRequest)
		return
	}

	trips, err := h.service.ListAvailableTrips(r.Context(), origin, destination)
	if err != nil {
		http.Error(w, "Failed to fetch trips", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trips)
}
