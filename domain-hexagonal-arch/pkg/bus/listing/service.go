package listing

// Service provides bus listing operations.
type Service interface {
	GetBus(string) (Bus, error)
	GetFleet() []Bus
}

// Repository provides access to the bus storage.
type Repository interface {
	// GetBus returns the bus with given ID.
	GetBus(string) (Bus, error)
	// GetFleet returns all buses saved in storage.
	GetFleet() []Bus
}

type busListingService struct {
	repo Repository
}

// GetBus returns a bus.
func (s *busListingService) GetBus(id string) (Bus, error) {
	return s.repo.GetBus(id)
}

// GetFleet returns all buses.
func (s *busListingService) GetFleet() []Bus {
	return s.repo.GetFleet()
}

// NewService creates a bus listing service with the necessary dependencies.
func NewService(r Repository) Service {
	return &busListingService{r}
}
