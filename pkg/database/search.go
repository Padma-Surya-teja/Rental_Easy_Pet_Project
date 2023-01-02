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
	if searchstring == "Get All Items" {
		result = db.Db.Find(&items)
	} else if searchstring != "" {
		result = db.Db.Where("name LIKE ?", "%"+searchstring+"%").Find(&items)
	} else {
		result = db.Db.Where("category LIKE ?", category).Find(&items)
	}

	return items, result.Error
}
