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
	DeleteItem(int) int
	GetUserEmail(int) string
	Get_Item_Name(int) string
	SearchItems(string) []models.Item
	SearchByCategory(string) []models.Item
}

type DBClient struct {
	Db *gorm.DB
}
