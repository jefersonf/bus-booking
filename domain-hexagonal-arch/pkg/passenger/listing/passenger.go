package listing

import "github.com/jefersonf/bus-booking/pkg/identity"

// Passenger describes the properties of passenger to be listed.
type Passenger struct {
	ID           string             `json:"id"`
	FullName     string             `json:"fullname"`
	SocialName   string             `json:"socialname"`
	IdentityInfo identity.Passenger `json:"identity_info"`
}
