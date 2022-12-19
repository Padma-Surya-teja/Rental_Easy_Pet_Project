package database

import (
	"rental_easy.in/m/pkg/models"
)

func (db DBClient) CreateUser(user models.User) (id int) {
	db.Db.Create(&user)

	return int(user.ID)
}

func (db DBClient) Get_User_Details(id int) models.User {
	var user models.User
	db.Db.First(&user, "id = ?", id)
	return user
}
