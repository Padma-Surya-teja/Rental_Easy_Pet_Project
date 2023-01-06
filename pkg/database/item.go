package database

import (
	"rental_easy.in/m/pkg/models"
)

func (db DBClient) AddItem(item models.Item) (int, error) {
	result := db.Db.Create(&item)
	return int(item.ID), result.Error
}

func (db DBClient) GetItems() ([]models.Item, error) {
	items := []models.Item{}
	result := db.Db.Find(&items)
	return items, result.Error
}

func (db DBClient) GetItemById(id int) (models.Item, error) {
	var item models.Item
	result := db.Db.First(&item, "id = ?", id)
	return item, result.Error
}

func (db DBClient) GetItemName(id int) (string, error) {
	item := models.Item{}
	result := db.Db.Select("name").First(&item, id)

	return item.Name, result.Error
}

func (db DBClient) UpdateItem(item models.Item) (int, error) {
	result := db.Db.Save(&item)

	return int(item.ID), result.Error
}

func (db DBClient) DeleteItem(id int) (int, error) {
	result := db.Db.Delete(&models.Item{}, id)

	return id, result.Error
}

func (db DBClient) GetItemsofOwner(id int) ([]models.Item, error) {
	items := []models.Item{}
	result := db.Db.Find(&items, "user_id = ?", id)
	return items, result.Error
}
