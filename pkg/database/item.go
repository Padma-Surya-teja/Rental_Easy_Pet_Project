package database

import "rental_easy.in/m/pkg/models"

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

func (db DBClient) Get_Item_Name(item_id int) string {
	Item_name := models.Item{}
	db.Db.Select("name").First(&Item_name, item_id)

	return Item_name.Name
}

func (db DBClient) Update_Item(item models.Item) (id int) {
	db.Db.Save(&item)

	return int(item.ID)
}

func (db DBClient) DeleteItem(itemId int) int {
	db.Db.Delete(&models.Item{}, itemId)

	return itemId
}

func (db DBClient) GetItemsofOwner(id int) []models.Item {
	Items := []models.Item{}
	db.Db.Find(&Items, "user_id = ?", id)
	return Items
}
