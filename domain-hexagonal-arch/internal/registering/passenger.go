package registering

// Passenger describes the properties of passenger to be registered.
type Passenger struct {
	FullName  string `json:"fullname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	BirthYear int    `json:"byr"`
}
