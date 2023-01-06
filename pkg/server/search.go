package server

import (
	"log"

	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func (S *ServerSideImplementation) SearchItems(in *rental.ItemRequest, stream rental.Rental_Easy_Functionalities_SearchItemsServer) error {
	log.Println("Server : Search Items is Being Executed")

	var items []models.Item
	log.Println(in.Request)
	log.Println(in.Category)
	items, err := S.Db.SearchItems(in.Request, in.Category.String())
	utils.CheckErr(err)

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
