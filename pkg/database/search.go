package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"rental_easy.in/m/pkg/models"
)

func (db DBClient) SearchItems(searchstring string, category string) ([]models.Item, error) {
	log.Println("Search Items in db")
	items := []models.Item{}
	var result *gorm.DB
	if category == "" {
		result = db.Db.Where("category LIKE ?", category).Find(&items)
	} else {
		result = db.Db.Where("name LIKE ?", "%"+searchstring+"%").Find(&items)
	}

	return items, result.Error
}
