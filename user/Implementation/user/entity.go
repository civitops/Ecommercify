package user

type Entity struct {
	// ID string
	Name        string  `json:"name"`
	PhoneNo     string  `json:"phoneNo"`
	Email       string  `json:"email"`
	HomeAddress Address `json:"homeAddress"`
	// WorkAddress     Address `json:"workAddress"`
	// DeliveryAddress Address `json:"deliveryAddress"`
	IsAdmin bool `json:"isAdmin"`
}

type Address struct {
	PhoneNo    string `json:"phoneNo"`
	AdressLine string `json:"addressLine1"`
	City       string `json:"city"`
	PinCode    string `json:"pinCode"`
	Landmark   string `json:"landmark"`
}
