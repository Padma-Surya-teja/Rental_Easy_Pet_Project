package database

import (
	"fmt"

	"rental_easy.in/m/pkg/models"
)

func (db DBClient) SearchItems(searchstring string) []models.Item {
	Items := []models.Item{}
	db.Db.Where("name LIKE ?", "%"+searchstring+"%").Find(&Items)

	fmt.Println(Items)
	return Items
}

func (db DBClient) SearchByCategory(category string) []models.Item {
	Items := []models.Item{}
	db.Db.Where("category LIKE ?", category).Find(&Items)

	fmt.Println(Items)
	return Items
}
