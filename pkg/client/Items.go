package client

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc/status"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func AddNewItem(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Creating a New Item")
	item, err := c.CreateItem(ctx, &rental.Item{Name: "Asus ZenBook 14 OLED", Description: "Nice Display View Experience is amazing", Category: "Laptops", AvailableFrom: "01-01-2023", AvailableTo: "31-02-2023", AmountPerDay: 1000, UserId: 1})
	utils.CheckErr(err)

	if err != nil {
		log.Println("New Item has been Created with Id : ", item.Id)
	} else {
		log.Println("Error with Adding the Item")
	}

}

func GetAllItems(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Getting All the Items")

	stream, err := c.GetAllItems(ctx, &rental.ItemRequest{Request: "Get All Items"})
	utils.CheckErr(err)

	for {
		item, err := stream.Recv()

		if err == io.EOF {
			log.Fatalf("finished")
			return
		}

		utils.CheckErr(err)

		log.Println("\nItem Received is: ", item)
	}
}

func GetItemById(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Getting the Item")

	item, err := c.GetItemById(ctx, &rental.Item{Id: 2})
	utils.CheckErr(err)

	if err == nil {
		log.Println("Item with the Given Id is : ", item)
	}
}

func UpdateItem(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Updating the Item")

	item, err := c.UpdateItem(ctx, &rental.Item{Id: 2, Name: "IqooNeo6", Description: "This phone has snapdragon 870 processor with nice camera", Category: "MobilePhones", AvailableFrom: "16-01-2023", AvailableTo: "30-01-2023", AmountPerDay: 500, UserId: 3})
	utils.CheckErr(err)
	st, ok := status.FromError(err)
	log.Println(st, ok)

	if err != nil {
		log.Println("Updated the Item with Id : ", item.Id)
	}
}

func DeleteItem(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Deleting an Item")

	item, err := c.DeleteItem(ctx, &rental.Item{Id: 1})
	utils.CheckErr(err)

	if err == nil {
		log.Println("Item has been deleted successfully with Id : ", item.Id)
	}
}

func GetItemsofOwner(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Getting the Items of Owner")

	stream, err := c.GetUserLeasedItems(ctx, &rental.User{Id: 1})
	utils.CheckErr(err)

	for {
		item, err := stream.Recv()

		if err == io.EOF {
			log.Printf("finished")
			return
		}

		utils.CheckErr(err)

		log.Printf("\nItem Received is: %s", item)
	}

}
