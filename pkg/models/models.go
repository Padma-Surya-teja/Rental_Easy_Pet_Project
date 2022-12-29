package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string
	Email       *string `gorm:"unique"`
	PhoneNumber string
	Address     string
	District    string
	PostalCode  string
	Country     string
	Item        []Item
	Reviews     []Review
	Booking     []Booking
}

type Item struct {
	gorm.Model
	Name          string
	Description   string
	Category      string
	AvailableFrom time.Time
	AvailableTo   time.Time
	AmountPerDay  int
	UserId        int
	Reviews       []Review
	Booking       []Booking
}

type Review struct {
	gorm.Model
	Comments string
	Rating   int
	UserId   int
	ItemId   int
}

type Booking struct {
	gorm.Model
	StartDate   time.Time
	EndDate     time.Time
	TotalAmount int
	UserId      int
	ItemId      int
}
