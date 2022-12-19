package client

import (
	"context"
	"fmt"
	"io"
	"log"

	rental "rental_easy.in/m/pkg/rentalmgmt"
)

func Add_New_Item(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	new_item_id, err := c.CreateItem(ctx, &rental.NewItem{Name: "IqooNeo6", Description: "This phone has snapdragon 870 processor with nice camera", Category: "Mobile Phones", Available_From: "20/12/2022", Available_To: "30/12/2022", AmountPerDay: 500, UserID: 1})
	checkErr(err)
	if new_item_id.Id == 0 {
		panic("Error with Adding the Item")
	}

	fmt.Println("New Item has been Created with Id : ", new_item_id.Id)
}

func Get_All_Items(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	Items, err := c.GetAllItems(ctx, &rental.Request{Request: "Get All Items"})
	checkErr(err)

	fmt.Println("Items in the Database are : ")
	fmt.Println(Items)
}

func Get_Item_By_Id(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	Item, err := c.GetItemById(ctx, &rental.ItemId{Id: 2})
	checkErr(err)

	if Item.ID == 0 {
		fmt.Println("Item Not Found")
	} else {
		fmt.Println("Item with the Given Id is : ")
		fmt.Println(Item)
	}
}

func Update_Item(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	Updated_Item_id, err := c.UpdateItem(ctx, &rental.Item{ID: 1, Name: "IqooNeo6", Description: "This phone has snapdragon 870 processor with nice camera", Category: "Mobile Phones", Available_From: "16/12/2022", Available_To: "30/12/2022", AmountPerDay: 500, UserID: 1})
	checkErr(err)

	if Updated_Item_id.Id == 0 {
		fmt.Println("Error in Updating the Item")
	} else {
		fmt.Println("Item has been Updated Successfully with Id : ", Updated_Item_id.Id)
	}
}

func GetItemsofOwner(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	stream_of_items, err := c.GetUserLeasedItems(ctx, &rental.UserId{Id: 1})
	checkErr(err)

	done := make(chan bool)

	go func() {
		for {
			item, err := stream_of_items.Recv()

			if err == io.EOF {
				done <- true
				return
			}

			checkErr(err)

			log.Printf("Item Received is: %s", item)
		}
	}()

	<-done
	log.Printf("finished")
}
