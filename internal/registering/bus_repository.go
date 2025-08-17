package registering

type BusRepsitory interface {
	Create(bus Bus) error
}
