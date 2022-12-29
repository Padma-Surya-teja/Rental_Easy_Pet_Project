package database

import (
	"github.com/jinzhu/gorm"
	"rental_easy.in/m/pkg/models"
)

type DataBase interface {
	CreateUser(models.User) (int, error)
	UpdateUser(models.User) (int, error)
	GetUser(int) (models.User, error)
	AddItem(models.Item) (int, error)
	GetItems() ([]models.Item, error)
	GetItemById(int) (models.Item, error)
	UpdateItem(item models.Item) (int, error)
	GetItemsofOwner(int) ([]models.Item, error)
	AddBooking(models.Booking) (models.Booking, error)
	GetBookings(int) ([]models.Booking, error)
	UserAlreadyBooked(int, int) (bool, error)
	AddReview(models.Review) (int, error)
	UpdateReview(models.Review) (int, error)
	UserAlreadyAddedReview(int, int) (bool, error)
	DeleteReview(int) (int, error)
	GetReviews(int) ([]models.Review, error)
	DeleteItem(int) (int, error)
	GetUserEmail(int) (string, error)
	GetItemName(int) (string, error)
	SearchItems(string, string) ([]models.Item, error)
}

type DBClient struct {
	Db *gorm.DB
}
