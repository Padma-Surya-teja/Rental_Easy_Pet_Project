package client

import (
	"context"
	"io"
	"log"

	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func SearchItems(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Searching The Items")

	req := rental.ItemRequest{
		Request: "IqooNeo6",
	}
	stream, err := c.SearchItems(ctx, &req)
	utils.CheckErr(err)

	for {
		item, err := stream.Recv()

		if err == io.EOF {
			log.Printf("finished")
			return
		}
		utils.CheckErr(err)

		log.Printf("Item Received is: %s", item)
	}
}
