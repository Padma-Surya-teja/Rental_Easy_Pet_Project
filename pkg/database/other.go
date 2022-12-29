package database

import (
	"rental_easy.in/m/pkg/models"
)

func (db DBClient) UserAlreadyAddedReview(userid int, itemid int) (bool, error) {
	var review = models.Review{}
	result := db.Db.First(&review, "user_id = ? and item_id = ?", userid, itemid)

	return review.ID != 0, result.Error
}

// Checking User Already Booked the Item Or not
func (db DBClient) UserAlreadyBooked(userid int, itemid int) (bool, error) {
	var booking = models.Booking{}
	result := db.Db.First(&booking, "user_id = ? and item_id = ?", userid, itemid)

	return booking.ID != 0, result.Error
}

func (db DBClient) GetUserEmail(user_id int) (string, error) {
	email := models.User{}
	result := db.Db.Select("email").First(&email, user_id)

	return *email.Email, result.Error
}
