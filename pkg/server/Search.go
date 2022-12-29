package server

import (
	"fmt"
	"log"

	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func (S *ServerSideImplementation) SearchItems(in *rental.ItemRequest, stream rental.Rental_Easy_Functionalities_SearchItemsServer) error {
	fmt.Println("Server : Search Items is Being Executed")

	var items []models.Item
	items, err := S.Db.SearchItems(in.Request, in.Category.String())
	utils.CheckErr(err)
	if err != nil {
		return err
	}

	fmt.Println(len(items))
	for _, item := range items {
		log.Println("Items sent is", item.ID)
		itm := rental.Item{
			Id:            int32(item.ID),
			Name:          item.Name,
			Description:   item.Description,
			Category:      item.Category,
			AvailableFrom: utils.DateFromTimeStamp(item.AvailableFrom.String()),
			AvailableTo:   utils.DateFromTimeStamp(item.AvailableFrom.String()),
			AmountPerDay:  int32(item.AmountPerDay),
			UserId:        int32(item.UserId)}
		err := stream.Send(&itm)
		utils.CheckErr(err)
	}

	return nil
}
