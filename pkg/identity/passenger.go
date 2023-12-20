package identity

type IdentityType string

// Identity describes the properties of a single passenger identity type.
type Identity struct {
	Type  IdentityType `json:"type"`
	Desc  string       `json:"description"`
	Value string       `json:"value"`
}

// Passenger describes the identity types registered by the passenger.
type Passenger map[IdentityType]Identity
