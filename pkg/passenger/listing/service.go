package listing

// Service provides the passenger listing operations.
type Service interface {
	GetPassenger(string) (Passenger, error)
	GetPassengers() []Passenger
}

// Repository provides access to the passenger storage.
type Repository interface {
	// GetPassenger returns the passenger with given ID.
	GetPassenger(string) (Passenger, error)
	// GetPassengers returns all passengers saved in storage.
	GetPassengers() []Passenger
}

type passengerListingService struct {
	repo Repository
}

func (s *passengerListingService) GetPassenger(id string) (Passenger, error) {
	return s.repo.GetPassenger(id)
}

func (s *passengerListingService) GetPassengers() []Passenger {
	return s.repo.GetPassengers()
}

func NewService(r Repository) Service {
	return &passengerListingService{r}
}
