package auth

type Entity struct {
	ID          uint    `json:"id" mapstructure:"id,omitempty" db:"id"`
	Name        string  `json:"name" mapstructure:"name,omitempty"  db:"name"`
	PhoneNo     string  `json:"phoneNo" mapstructure:"phone_no,omitempty"`
	Email       string  `json:"email" mapstructure:"email,omitempty"`
	HomeAddress Address `json:"homeAddress" mapstructure:"homeaddress,omitempty,squash"`
	// WorkAddress     Address `json:"workAddress"`
	// DeliveryAddress Address `json:"deliveryAddress"`
	IsAdmin bool `json:"isAdmin" mapstructure:"is_admin,omitempty"`
}

type Address struct {
	PhoneNo    string `json:"phoneNo" mapstructure:"homeaddress_phoneno,omitempty"`
	AdressLine string `json:"addressLine1" mapstructure:"homeaddress_address_line,omitempty"`
	City       string `json:"city" mapstructure:"homeaddress_city,omitempty"`
	PinCode    string `json:"pinCode" mapstructure:"homeaddress_pin_code,omitempty"`
	Landmark   string `json:"landmark" mapstructure:"homeaddress_landmark,omitempty"`
}
