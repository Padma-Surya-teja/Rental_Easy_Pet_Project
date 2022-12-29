package database

import (
	"rental_easy.in/m/pkg/models"
)

func (db DBClient) AddBooking(booking models.Booking) (models.Booking, error) {

	var bookings []models.Booking
	result := db.Db.Find(models.Item{}, booking.ItemId)
	if result.Error != nil {
		return models.Booking{}, result.Error
	}
	result = db.Db.Where("(start_date between ? and ?) and (end_date between ? and ?) and item_id = ?", booking.StartDate, booking.EndDate, booking.StartDate, booking.EndDate, booking.ItemId).Find(&bookings)
	if result.Error != nil || len(bookings) != 0 {
		return models.Booking{}, result.Error
	}

	result = db.Db.Where("(? between available_from and available_to) and (? between available_from and available_to)", booking.StartDate, booking.EndDate).Find(&models.Item{}, booking.ItemId)
	if result.Error != nil {
		return models.Booking{}, result.Error
	}

	result = db.Db.Create(&booking)
	return booking, result.Error
}

func (db DBClient) GetBookings(id int) ([]models.Booking, error) {
	bookings := []models.Booking{}
	result := db.Db.Find(&bookings, "user_id = ?", id)
	return bookings, result.Error
}
