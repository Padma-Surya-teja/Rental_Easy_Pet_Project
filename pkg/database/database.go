package database

import (
	"github.com/jinzhu/gorm"
	"rental_easy.in/m/pkg/models"
)

type DataBase interface {
	CreateUser(models.User) int
	AddItem(models.Item) int
	GetItems() []models.Item
	GetItemById(int) models.Item
	Update_Item(item models.Item) (id int)
	GetItemsofOwner(int) []models.Item
	AddBooking(models.Booking) (id int)
	GetBookings(int) []models.Booking
	UserAlreadyBooked(int, int) bool
	AddReview(models.Review) int
	Update_Review(models.Review) int
	UserAlreadyAddedReview(int, int) bool
	DeleteReview(int) int
	GetReviews(int) []models.Review
}

type DBClient struct {
	Db *gorm.DB
}

func (db DBClient) CreateUser(user models.User) (id int) {
	db.Db.Create(&user)

	return int(user.ID)
}

func (db DBClient) AddItem(item models.Item) (id int) {
	db.Db.Create(&item)

	return int(item.ID)
}

func (db DBClient) GetItems() []models.Item {
	Items := []models.Item{}
	db.Db.Find(&Items)
	return Items
}

func (db DBClient) GetItemById(id int) models.Item {
	var item models.Item
	db.Db.First(&item, "id = ?", id)
	return item
}

func (db DBClient) Get_User_Details(id int) models.User {
	var user models.User
	db.Db.First(&user, "id = ?", id)
	return user
}

func (db DBClient) Update_Item(item models.Item) (id int) {
	db.Db.Save(&item)

	return int(item.ID)
}

func (db DBClient) GetItemsofOwner(id int) []models.Item {
	Items := []models.Item{}
	db.Db.Find(&Items, "user_id = ?", id)
	return Items
}

func (db DBClient) AddBooking(booking models.Booking) (id int) {
	db.Db.Create(&booking)

	return int(booking.ID)
}

func (db DBClient) GetBookings(id int) []models.Booking {
	BookedItems := []models.Booking{}
	db.Db.Find(&BookedItems, "user_id = ?", id)
	return BookedItems
}

// Checking User Already Booked the Item Or not
func (db DBClient) UserAlreadyBooked(userid int, itemid int) bool {
	var booking = models.Booking{UserID: userid, ItemID: itemid}
	db.Db.First(&booking)

	return booking.ID != 0
}

func (db DBClient) AddReview(review models.Review) int {
	db.Db.Create(&review)

	return int(review.ID)
}

// Updating the Review
func (db DBClient) Update_Review(review models.Review) int {
	db.Db.Save(&review)

	return int(review.ID)
}

func (db DBClient) UserAlreadyAddedReview(userid int, itemid int) bool {
	var review = models.Review{UserID: userid, ItemID: itemid}
	db.Db.First(&review)

	return review.ID != 0
}

func (db DBClient) DeleteReview(reviewId int) int {
	db.Db.Delete(&models.Review{}, reviewId)

	return reviewId
}

func (db DBClient) GetReviews(item_id int) []models.Review {
	Reviews := []models.Review{}
	db.Db.Find(&Reviews, "item_id = ?", item_id)
	return Reviews
}
