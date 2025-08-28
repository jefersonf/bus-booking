package adding

// Bus defines the properties of a bus to be added to the fleet.
type Bus struct {
	CompanyName string `json:"companyName"`
	Code        string `json:"code"`
	Seats       uint   `json:"seats"`
}
