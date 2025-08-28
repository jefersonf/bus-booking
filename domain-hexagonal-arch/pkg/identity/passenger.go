package identity

type IdentityType string

// IdentityInfo describes the properties of a single passenger identity type.
type IdentityInfo struct {
	Description  string       `json:"description"`
	Value        string       `json:"value"`
	Type         IdentityType `json:"type"`

}

// Passenger describes the identity types registered by the passenger.
type Passenger map[IdentityType]IdentityInfo
