package registering

// Bus defines the properties of a bus to be registered.
type Bus struct {
	Code    string `json:"code"`
	Company string `json:"company"`
	Seats   int    `json:"seats"`
}
