package user

import "gorm.io/gorm"

type Entity struct {
	gorm.Model
	ID              uint    `json:"id" mapstructure:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	Name            string  `json:"name" mapstructure:"name,omitempty"  gorm:"not null"`
	PhoneNo         string  `json:"phoneNo" mapstructure:"phone_no,omitempty" gorm:"unique"`
	Email           string  `json:"email" mapstructure:"email,omitempty" gorm:"unique"`
	HomeAddress     Address `json:"homeAddress" mapstructure:"homeaddress,omitempty,squash" gorm:"embedded;embeddedPrefix:homeaddress_"`
	DeliveryAddress Address `json:"deliveryAddress" mapstructure:"deliveryaddress,omitempty,squash" gorm:"embedded;embeddedPrefix:deliveryaddress_"`
	IsAdmin         bool    `json:"isAdmin" mapstructure:"is_admin,omitempty"`
}

type Address struct {
	PhoneNo    string `json:"phoneNo" mapstructure:"homeaddress_phoneno,omitempty"`
	AdressLine string `json:"addressLine1" mapstructure:"homeaddress_address_line,omitempty"`
	City       string `json:"city" mapstructure:"homeaddress_city,omitempty"`
	PinCode    string `json:"pinCode" mapstructure:"homeaddress_pin_code,omitempty"`
	Landmark   string `json:"landmark" mapstructure:"homeaddress_landmark,omitempty"`
}
