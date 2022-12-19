package database

import "rental_easy.in/m/pkg/models"

func (db DBClient) AddBooking(booking models.Booking) (id int) {
	db.Db.Create(&booking)

	return int(booking.ID)
}

func (db DBClient) GetBookings(id int) []models.Booking {
	BookedItems := []models.Booking{}
	db.Db.Find(&BookedItems, "user_id = ?", id)
	return BookedItems
}
