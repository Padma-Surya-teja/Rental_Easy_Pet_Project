package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name         *string `gorm:"unique"`
	Email        string
	Phone_Number string
	Address      string
	District     string
	Postal_Code  string
	Country      string
	Item         []Item
	Reviews      []Review
	Booking      []Booking
}

type Item struct {
	gorm.Model
	Name           string
	Description    string
	Category       string
	Available_From time.Time
	Available_To   time.Time
	Amount_per_day int
	UserID         int
	Reviews        []Review
	Booking        []Booking
}

type Review struct {
	gorm.Model
	Comments string
	Rating   int
	UserID   int
	ItemID   int
}

type Booking struct {
	gorm.Model
	Start_Date time.Time
	End_Date   time.Time
	UserID     int
	ItemID     int
}
