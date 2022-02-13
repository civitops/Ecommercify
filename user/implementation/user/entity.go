package user

import "gorm.io/gorm"

type Entity struct {
	gorm.Model
	ID              uint    `json:"id"  gorm:"primaryKey;autoIncrement"`
	Name            string  `json:"name" gorm:"not null"`
	PhoneNo         string  `json:"phoneNo"  gorm:"unique"`
	Email           string  `json:"email" gorm:"unique"`
	HomeAddress     Address `json:"homeAddress"  gorm:"embedded;embeddedPrefix:homeaddress_"`
	DeliveryAddress Address `json:"deliveryAddress"  gorm:"embedded;embeddedPrefix:deliveryaddress_"`
	IsAdmin         bool    `json:"isAdmin" goorm:"default:false"`
}

type Address struct {
	PhoneNo    string `json:"phoneNo"`
	AdressLine string `json:"addressLine1" `
	City       string `json:"city"`
	PinCode    string `json:"pinCode"`
	Landmark   string `json:"landmark"`
}
