package server

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func (s *ServerSideImplementation) CreateItem(ctx context.Context, in *rental.Item) (*rental.Item, error) {
	log.Printf("Server : Creating a new Item")

	if in == nil || in.UserId == 0 {
		return &rental.Item{}, errors.New("please provide proper details")
	}
	Avail_From := utils.ConvertToInt(strings.Split(in.GetAvailableFrom(), "-"))
	Avail_To := utils.ConvertToInt(strings.Split(in.GetAvailableTo(), "-"))

	item := models.Item{
		Name:          in.GetName(),
		Description:   in.GetDescription(),
		Category:      in.GetCategory(),
		AvailableFrom: time.Date(Avail_From[2], time.Month(Avail_From[1]), Avail_From[0], 0, 0, 0, 0, time.UTC),
		AvailableTo:   time.Date(Avail_To[2], time.Month(Avail_To[1]), Avail_To[0], 0, 0, 0, 0, time.UTC),
		AmountPerDay:  int(in.AmountPerDay),
		UserId:        int(in.GetUserId()),
	}

	id, err := s.Db.AddItem(item)
	utils.CheckErr(err)

	return &rental.Item{
		Id: int32(id)}, err
}

func (s *ServerSideImplementation) GetAllItems(in *rental.ItemRequest, stream rental.Rental_Easy_Functionalities_GetAllItemsServer) error {
	log.Println("Server : Getting All the Items")

	items, err := s.Db.GetItems()
	if err != nil {
		return err
	}

	for _, item := range items {
		itm := rental.Item{
			Id:            int32(item.ID),
			Name:          item.Name,
			Description:   item.Description,
			Category:      item.Category,
			AvailableFrom: utils.DateFromTimeStamp(item.AvailableFrom.String()),
			AvailableTo:   utils.DateFromTimeStamp(item.AvailableTo.String()),
			AmountPerDay:  int32(item.AmountPerDay),
			UserId:        int32(item.UserId)}
		err := stream.Send(&itm)
		utils.CheckErr(err)
	}
	return nil
}

func (s *ServerSideImplementation) GetItemById(ctx context.Context, in *rental.Item) (*rental.Item, error) {
	log.Println("Server : Getting Item by Id")
	// use Item instead of Received_Item
	item, err := s.Db.GetItemById(int(in.GetId()))
	utils.CheckErr(err)
	if err != nil {
		return &rental.Item{}, err
	}

	return &rental.Item{
		Id:            int32(item.ID),
		Name:          item.Name,
		Description:   item.Description,
		Category:      item.Category,
		AvailableFrom: utils.DateFromTimeStamp(item.AvailableFrom.String()),
		AvailableTo:   utils.DateFromTimeStamp(item.AvailableTo.String()),
		AmountPerDay:  int32(item.AmountPerDay),
		UserId:        int32(item.UserId)}, nil
}

func (s *ServerSideImplementation) UpdateItem(ctx context.Context, in *rental.Item) (*rental.Item, error) {
	log.Println("Server : Updating the Item")

	item, err := s.Db.GetItemById(int(in.GetId()))
	utils.CheckErr(err)

	if err != nil {
		return &rental.Item{}, err
	}

	if item.UserId != int(in.GetUserId()) {
		return &rental.Item{}, errors.New("You Cannot Edit This Item")
	}

	Avail_From := utils.ConvertToInt(strings.Split(in.GetAvailableFrom(), "-"))
	Avail_To := utils.ConvertToInt(strings.Split(in.GetAvailableTo(), "-"))
	item = models.Item{
		Name:          in.GetName(),
		Description:   in.GetDescription(),
		Category:      in.GetCategory(),
		AvailableFrom: time.Date(Avail_From[2], time.Month(Avail_From[1]), Avail_From[0], 0, 0, 0, 0, time.UTC),
		AvailableTo:   time.Date(Avail_To[2], time.Month(Avail_To[1]), Avail_To[0], 0, 0, 0, 0, time.UTC),
		AmountPerDay:  int(in.GetAmountPerDay()),
		UserId:        int(in.GetUserId()),
	}
	item.ID = uint(in.GetId())

	id, err := s.Db.UpdateItem(item)
	utils.CheckErr(err)

	return &rental.Item{
		Id: int32(id)}, err

}

func (s *ServerSideImplementation) GetUserLeasedItems(in *rental.User, stream rental.Rental_Easy_Functionalities_GetUserLeasedItemsServer) error {
	log.Println("Getting the Owner Leased Items")

	items, err := s.Db.GetItemsofOwner(int(in.GetId()))
	utils.CheckErr(err)
	if err != nil {
		return err
	}

	for _, item := range items {
		itm := rental.Item{
			Id:            int32(item.ID),
			Name:          item.Name,
			Description:   item.Description,
			Category:      item.Category,
			AvailableFrom: utils.DateFromTimeStamp(item.AvailableFrom.String()),
			AvailableTo:   utils.DateFromTimeStamp(item.AvailableTo.String()),
			AmountPerDay:  int32(item.AmountPerDay),
			UserId:        int32(item.UserId)}
		err := stream.Send(&itm)
		utils.CheckErr(err)
	}

	return nil
}

func (S *ServerSideImplementation) DeleteItem(ctx context.Context, in *rental.Item) (*rental.Item, error) {
	log.Println("Server : Deleting an Item")

	id, err := S.Db.DeleteItem(int(in.GetId()))
	utils.CheckErr(err)

	return &rental.Item{Id: int32(id)}, err
}
