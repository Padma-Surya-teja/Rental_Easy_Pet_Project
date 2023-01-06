package database

import (
	"rental_easy.in/m/pkg/models"
)

func (db DBClient) CreateUser(user models.User) (int, error) {

	result := db.Db.Create(&user)

	return int(user.ID), result.Error
}

func (db DBClient) GetUser(id int) (models.User, error) {
	var user models.User
	result := db.Db.First(&user, "id = ?", id)
	return user, result.Error
}

func (db DBClient) UpdateUser(user models.User) (int, error) {
	result := db.Db.Save(&user)

	return int(user.ID), result.Error
}
