package listing

// Bus defines the properties of a bus to be listed.
type Bus struct {
	ID          string `json:"id"`
	CompanyName string `json:"companyName"`
	Code        string `json:"code"`
	Seats       uint   `json:"seats"`
}
