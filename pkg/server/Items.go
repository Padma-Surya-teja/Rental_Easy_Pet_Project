package Services

import (
	"context"
	"log"
	"strconv"
	"strings"
	"sync"
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

func (s *ServerSideImplementation) CreateItem(ctx context.Context, in *rental.Item) (*rental.ItemId, error) {
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

func (s *ServerSideImplementation) GetAllItems(in *rental.Request, stream rental.Rental_Easy_Functionalities_GetAllItemsServer) error {
	Received_Items := s.Db.GetItems()

	var wg sync.WaitGroup
	for _, Item := range Received_Items {
		wg.Add(1)
		go func(item models.Item) {
			defer wg.Done()
			time.Sleep(time.Duration(100) * time.Microsecond)
			itm := rental.Item{
				ID:             int32(item.ID),
				Name:           item.Name,
				Description:    item.Description,
				Category:       item.Category,
				Available_From: item.Available_From.String(),
				Available_To:   item.Available_From.String(),
				AmountPerDay:   int32(item.Amount_per_day),
				UserID:         int32(item.UserID)}
			err := stream.Send(&itm)
			checkErr(err)
		}(Item)
	}

	wg.Wait()
	return nil
}

func (s *ServerSideImplementation) GetItemById(ctx context.Context, in *rental.ItemId) (*rental.Item, error) {
	Received_Item := s.Db.GetItemById(int(in.Id))

	return &rental.Item{
		ID:             int32(Received_Item.ID),
		Name:           Received_Item.Name,
		Description:    Received_Item.Description,
		Category:       Received_Item.Category,
		Available_From: Received_Item.Available_From.String(),
		Available_To:   Received_Item.Available_From.String(),
		AmountPerDay:   int32(Received_Item.Amount_per_day),
		UserID:         int32(Received_Item.UserID)}, nil
}

func (s *ServerSideImplementation) UpdateItem(ctx context.Context, in *rental.Item) (*rental.ItemId, error) {
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

func (s *ServerSideImplementation) GetUserLeasedItems(in *rental.UserId, stream rental.Rental_Easy_Functionalities_GetUserLeasedItemsServer) error {
	Received_Items := s.Db.GetItemsofOwner(int(in.Id))

	var wg sync.WaitGroup
	for _, Item := range Received_Items {
		wg.Add(1)
		go func(item models.Item) {
			defer wg.Done()
			time.Sleep(time.Duration(100) * time.Microsecond)
			itm := rental.Item{
				ID:             int32(item.ID),
				Name:           item.Name,
				Description:    item.Description,
				Category:       item.Category,
				Available_From: item.Available_From.String(),
				Available_To:   item.Available_From.String(),
				AmountPerDay:   int32(item.Amount_per_day),
				UserID:         int32(item.UserID)}
			err := stream.Send(&itm)
			checkErr(err)
		}(Item)
	}

	wg.Wait()
	return nil
}

func (S *ServerSideImplementation) DeleteItem(ctx context.Context, in *rental.ItemId) (*rental.ItemId, error) {

	deletedItemId := S.Db.DeleteItem(int(in.Id))

	return &rental.ItemId{Id: int32(deletedItemId)}, nil
}
