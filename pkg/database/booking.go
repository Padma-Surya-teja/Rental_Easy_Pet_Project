package database

import (
	"errors"
	"log"

	"rental_easy.in/m/pkg/models"
)

func (db DBClient) AddBooking(booking models.Booking, item models.Item) (models.Booking, error) {

	var bookings []models.Booking
	result := db.Db.Where("(start_date between ? and ?) and (end_date between ? and ?) and item_id = ?", booking.StartDate, booking.EndDate, booking.StartDate, booking.EndDate, booking.ItemId).Find(&bookings)
	log.Println(len(bookings))
	if len(bookings) != 0 {
		return models.Booking{}, result.Error
	}

	if booking.StartDate.Before(item.AvailableFrom) ||
		booking.StartDate.After(item.AvailableTo) ||
		booking.EndDate.After(item.AvailableTo) ||
		booking.StartDate.After(booking.EndDate) {
		return models.Booking{}, errors.New("given dates are not valid to book the item")
	}

	result = db.Db.Create(&booking)
	return booking, result.Error
}

func (db DBClient) GetBookings(id int) ([]models.Booking, error) {
	bookings := []models.Booking{}
	result := db.Db.Find(&bookings, "user_id = ?", id)
	return bookings, result.Error
}
