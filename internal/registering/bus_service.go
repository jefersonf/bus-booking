package registering

// BusService provides bus registering operations.
type BusService interface {
	AddBus(bus Bus) error
	AddBuses(bb ...Bus) []error
}

type busService struct {
}
