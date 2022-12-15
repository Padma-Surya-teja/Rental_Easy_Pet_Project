package Services

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
)

func Convert_to_String(array []string) []int {
	var result = []int{}
	for _, str := range array {
		num, err := strconv.Atoi(str)
		checkErr(err)

		result = append(result, num)
	}

	return result
}

func (s *ServerSideImplementation) CreateItem(ctx context.Context, in *rental.NewItem) (*rental.ItemId, error) {
	log.Printf("Creating a new Item")
	Avail_From := Convert_to_String(strings.Split(in.GetAvailable_From(), "/"))
	Avail_To := Convert_to_String(strings.Split(in.GetAvailable_To(), "/"))
	item := models.Item{
		Name:           in.GetName(),
		Description:    in.GetDescription(),
		Category:       in.GetCategory(),
		Available_From: time.Date(Avail_From[2], time.Month(Avail_From[1]), Avail_From[0], 0, 0, 0, 0, time.UTC),
		Available_To:   time.Date(Avail_To[2], time.Month(Avail_To[1]), Avail_To[0], 0, 0, 0, 0, time.UTC),
		Amount_per_day: int(in.AmountPerDay),
		UserID:         int(in.GetUserID()),
	}

	response := s.Db.AddItem(item)

	return &rental.ItemId{
		Id: int32(response)}, nil
}

func (s *ServerSideImplementation) GetAllItems(ctx context.Context, in *rental.Request) (*rental.Items, error) {
	Items := []*rental.Item{}
	All_Items := s.Db.GetItems()
	for _, Item := range All_Items {
		Items = append(Items, &rental.Item{Id: int32(Item.ID), Name: Item.Name, Description: Item.Description, AmountPerDay: int32(Item.Amount_per_day)})
	}
	return &rental.Items{
		Items: Items}, nil
}

func (s *ServerSideImplementation) GetItem(ctx context.Context, in *rental.ItemId) (*rental.DetailedItem, error) {
	Received_Item := s.Db.GetItem(int(in.Id))

	return &rental.DetailedItem{
		ID:             int32(Received_Item.ID),
		Name:           Received_Item.Name,
		Description:    Received_Item.Description,
		Category:       Received_Item.Category,
		Available_From: Received_Item.Available_From.String(),
		Available_To:   Received_Item.Available_From.String(),
		AmountPerDay:   int32(Received_Item.Amount_per_day),
		UserID:         int32(Received_Item.UserID)}, nil
}

func (s *ServerSideImplementation) UpdateItem(ctx context.Context, in *rental.DetailedItem) (*rental.ItemId, error) {
	Avail_From := Convert_to_String(strings.Split(in.GetAvailable_From(), "/"))
	Avail_To := Convert_to_String(strings.Split(in.GetAvailable_To(), "/"))
	Updated_Item := models.Item{
		Name:           in.GetName(),
		Description:    in.GetDescription(),
		Category:       in.GetCategory(),
		Available_From: time.Date(Avail_From[2], time.Month(Avail_From[1]), Avail_From[0], 0, 0, 0, 0, time.UTC),
		Available_To:   time.Date(Avail_To[2], time.Month(Avail_To[1]), Avail_To[0], 0, 0, 0, 0, time.UTC),
		Amount_per_day: int(in.AmountPerDay),
		UserID:         int(in.UserID),
	}
	Updated_Item.ID = uint(in.ID)

	Updated_Item_id := s.Db.Update_Item(Updated_Item)

	return &rental.ItemId{
		Id: int32(Updated_Item_id)}, nil

}
