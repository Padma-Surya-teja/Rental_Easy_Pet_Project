package database

import (
	"fmt"

	"rental_easy.in/m/pkg/models"
)

func (db DBClient) UserAlreadyAddedReview(userid int, itemid int) bool {
	var review = models.Review{}
	db.Db.First(&review, "user_id = ? and item_id = ?", userid, itemid)

	fmt.Println(review)
	return review.ID != 0
}

// Checking User Already Booked the Item Or not
func (db DBClient) UserAlreadyBooked(userid int, itemid int) bool {
	var booking = models.Booking{}
	db.Db.First(&booking, "user_id = ? and item_id = ?", userid, itemid)

	return booking.ID != 0
}

func (db DBClient) GetUserEmail(user_id int) string {
	email := models.User{}
	db.Db.Select("email").First(&email, user_id)

	return email.Email
}
