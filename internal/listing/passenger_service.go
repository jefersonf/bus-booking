package listing

// PassengerService provides passenger listing operations.
type PassengerService interface {
	GetPassenger(id string) (Passenger, error)
	GetAllPassengers() []Passenger
}

type passengerService struct {
	repo PassengerRepository
}

// GetPassenger returns a passenger and possibly a non-nil error.
func (s *passengerService) GetPassenger(id string) (Passenger, error) {
	return s.repo.GetOne(id)
}

// GetAllPassengers returns all passengers.
func (s *passengerService) GetAllPassengers() []Passenger {
	return s.repo.GetAll()
}

// NewPassengerService creates a passenger listing service with the necessary dependencies.
func NewPassengersService(r PassengerRepository) PassengerService {
	return &passengerService{r}
}
