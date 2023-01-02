package client

import (
	"context"
	"io"
	"log"

	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func AddNewItem(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Creating a New Item")
	item, err := c.CreateItem(ctx, &rental.Item{})
	utils.CheckErr(err)

	if err == nil {
		log.Println("New Item has been Created with Id : ", item.Id)
	} else {
		log.Println(err)
	}

}

func GetAllItems(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Getting All the Items")

	stream, err := c.GetAllItems(ctx, &rental.ItemRequest{})
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

	item, err := c.GetItemById(ctx, &rental.Item{})
	utils.CheckErr(err)

	if err == nil {
		log.Println("Item with the Given Id is : ", item)
	}
}

func UpdateItem(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Updating the Item")

	item, err := c.UpdateItem(ctx, &rental.Item{Id: 2, Name: "IqooNeo6", Description: "This phone has snapdragon 870 processor with nice camera", Category: "MobilePhones", AvailableFrom: "16-01-2023", AvailableTo: "30-01-2023", AmountPerDay: 500, UserId: 5})
	utils.CheckErr(err)

	if err == nil {
		log.Println("Updated the Item with Id : ", item.Id)
	}
}

func DeleteItem(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Deleting an Item")

	item, err := c.DeleteItem(ctx, &rental.Item{Id: 100})
	utils.CheckErr(err)

	if err == nil {
		log.Println("Item has been deleted successfully with Id : ", item.Id)
	}
}

func GetItemsofOwner(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Getting the Items of Owner")

	stream, err := c.GetUserLeasedItems(ctx, &rental.User{})
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
