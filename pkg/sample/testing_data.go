package server

import (
	"time"

	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
)

//User Data

var email1 = "19bd1a057d@gmail.com"

// var email2 = "surya@gmail.com"
// var email3 = "surya@gmail.com"

// Data for mock call

// 1
var mockuser1 = models.User{
	Name:        "Padma surya teja",
	Email:       &email1,
	PhoneNumber: "9999999991",
	Address:     "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad",
	District:    "Medchal",
	PostalCode:  "500029",
	Country:     "India",
}

// 2
var mockuser2 = models.User{
	Name:        "Padma surya teja",
	Email:       &email1,
	PhoneNumber: "9999999991",
	Address:     "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad",
	District:    "Medchal",
	PostalCode:  "500029",
	Country:     "India",
}

// 3
var mockuser3 = models.User{
	Name:        "Padma surya teja",
	Email:       &email1,
	PhoneNumber: "9999999991",
	Address:     "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad",
	District:    "Medchal",
	PostalCode:  "500029",
	Country:     "India",
}

// Data for Actual Call
var User1 = rental.User{Name: "Padma surya teja", Email: "19bd1a057d@gmail.com", PhoneNumber: "9999999991", Address: "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad", District: "Medchal", PostalCode: "500029", Country: "India"}

//Item Data

var mockItem1 = models.Item{
	Name:          "Iqoo Neo 6",
	Description:   "Phone is Superfast and Display is just amazing",
	Category:      "Mobile Phones",
	AvailableFrom: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
	AvailableTo:   time.Date(2023, 01, 03, 0, 0, 0, 0, time.UTC),
	AmountPerDay:  500,
	UserId:        1,
}

var mockItem2 = models.Item{
	Name:          "Asus Zenbook 14 OLED",
	Description:   "Nice Display View Experience is amazing",
	Category:      "Laptops",
	AvailableFrom: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
	AvailableTo:   time.Date(2023, 01, 31, 0, 0, 0, 0, time.UTC),
	AmountPerDay:  1000,
	UserId:        3,
}

var mockItem4 = models.Item{
	Name:          "Asus Zenbook 14",
	Description:   "Amazing Speed",
	Category:      "Laptops",
	AvailableFrom: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
	AvailableTo:   time.Date(2023, 01, 31, 0, 0, 0, 0, time.UTC),
	AmountPerDay:  1000,
	UserId:        2,
}

var mockItem3 = models.Item{
	Name:          "Iqoo Neo 6",
	Description:   "Display is just amazing",
	Category:      "Mobile Phones",
	AvailableFrom: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
	AvailableTo:   time.Date(2023, 01, 03, 0, 0, 0, 0, time.UTC),
	AmountPerDay:  300,
	UserId:        1,
}

var mockUpdatedItem = models.Item{
	Name:          "Iqoo Neo 6",
	Description:   "Phone is Superfast and Display is just amazing",
	Category:      "Mobile Phones",
	AvailableFrom: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
	AvailableTo:   time.Date(2023, 01, 03, 0, 0, 0, 0, time.UTC),
	AmountPerDay:  500,
	UserId:        1,
}

var inputItem = rental.Item{
	Name:          "Iqoo Neo 6",
	Description:   "Phone is Superfast and Display is just amazing",
	Category:      "Mobile Phones",
	AvailableFrom: "01-01-2023",
	AvailableTo:   "03-01-2023",
	AmountPerDay:  500,
	UserId:        1,
}
var Item1 = rental.Item{
	Name:          "Iqoo Neo 6",
	Description:   "Phone is Superfast and Display is just amazing",
	Category:      "Mobile Phones",
	AvailableFrom: "2023-01-01",
	AvailableTo:   "2023-01-03",
	AmountPerDay:  500,
	UserId:        1,
}

var Item2 = rental.Item{
	Name:          "Asus Zenbook 14 OLED",
	Description:   "Nice Display View Experience is amazing",
	Category:      "Laptops",
	AvailableFrom: "2023-01-01",
	AvailableTo:   "2023-01-31",
	AmountPerDay:  1000,
	UserId:        3,
}

var Item3 = rental.Item{
	Name:          "Iqoo Neo 6",
	Description:   "Phone is Superfast and Display is just amazing",
	Category:      "Mobile Phones",
	AvailableFrom: "2023-01-01",
	AvailableTo:   "2023-01-03",
	AmountPerDay:  300,
	UserId:        1,
}

var Item4 = rental.Item{
	Name:          "Asus Zenbook 14",
	Description:   "Display is just amazing",
	Category:      "Laptops",
	AvailableFrom: "2023-01-01",
	AvailableTo:   "2023-01-31",
	AmountPerDay:  1000,
	UserId:        3,
}

//Data for mock call

//Data for Actual Call

//Booking Data

// Data for mock call
var mockBooking1 = models.Booking{
	StartDate:   time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
	EndDate:     time.Date(2023, 01, 03, 0, 0, 0, 0, time.UTC),
	TotalAmount: 1500,
	UserId:      1,
	ItemId:      1,
}
var mockBooking2 = models.Booking{
	StartDate:   time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
	EndDate:     time.Date(2023, 01, 03, 0, 0, 0, 0, time.UTC),
	TotalAmount: 1500,
	UserId:      1,
	ItemId:      1,
}

// Data for Actual Call
var booking1 = rental.Booking{
	StartDate: "01-01-2023",
	EndDate:   "03-01-2023",
	UserId:    1,
	ItemId:    1,
}

//Reviews Data

// Data for mock call
var mockReview1 = models.Review{
	Comments: "The Phone is too good",
	Rating:   4,
	UserId:   2,
	ItemId:   1,
}

var mockReview2 = models.Review{
	Comments: "The Phone is too good",
	Rating:   4,
	UserId:   2,
	ItemId:   1000,
}

// Data for Actual Call
var Review1 = rental.Review{
	Comment: "The Phone is too good",
	Rating:  4,
	UserId:  2,
	ItemId:  1,
}

var Review2 = rental.Review{
	Comment: "The Phone is too good",
	Rating:  4,
	UserId:  2,
	ItemId:  1000,
}

//Searching

var Item5 = rental.Item{
	Id:            5,
	Name:          "Iqoo Neo 6",
	Description:   "Phone is Superfast and Display is just amazing",
	Category:      "Mobile Phones",
	AvailableFrom: "2023-01-01",
	AvailableTo:   "2023-01-03",
	AmountPerDay:  500,
	UserId:        1,
}

var Item6 = rental.Item{
	Id:            6,
	Name:          "Asus Zenbook 14 OLED",
	Description:   "Nice Display View Experience is amazing",
	Category:      "Laptops",
	AvailableFrom: "2023-01-01",
	AvailableTo:   "2023-01-31",
	AmountPerDay:  1000,
	UserId:        3,
}

var Item7 = rental.Item{
	Id:            7,
	Name:          "Iqoo Neo 6",
	Description:   "Phone is Superfast and Display is just amazing",
	Category:      "Mobile Phones",
	AvailableFrom: "2023-01-01",
	AvailableTo:   "2023-01-03",
	AmountPerDay:  300,
	UserId:        1,
}

var Item8 = rental.Item{
	Id:            8,
	Name:          "Asus Zenbook 14",
	Description:   "Display is just amazing",
	Category:      "Laptops",
	AvailableFrom: "2023-01-01",
	AvailableTo:   "2023-01-31",
	AmountPerDay:  1000,
	UserId:        3,
}

// Data for mock call
var req1 = rental.ItemRequest{
	Request: "Iqoo Neo 6",
}

var req2 = rental.ItemRequest{
	Category: 2,
}

var search1 = []models.Item{
	mockItem1,
	mockItem3,
}

var search2 = []models.Item{
	mockItem2,
	mockItem4,
}

// Data for actual call
var result1 = []rental.Item{
	Item5,
	Item7,
}

var result2 = []rental.Item{
	Item6,
	Item8,
}
