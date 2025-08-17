package registering

import "github.com/jefersonf/bus-booking/pkg/identity"

// Passenger describes the properties of passenger to be registered.
type Passenger struct {
	FullName     string             `json:"fullname"`
	SocialName   string             `json:"socialname"`
	IdentityInfo identity.Passenger `json:"identity_info"`
}

/*

pkg/passenger/listing
pkg/ticket/listing

listing.Passanger{}
listing.Ticket{}

internal/listing
internal/listing
internal/registering

listing.Ticket{}
listing.Passenger{}
registering.Passenger{}

*/
