package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

// Booking an Item
func BookAnItem(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Booking An Item")

	booking, err := c.BookItem(ctx, &rental.Booking{})
	utils.CheckErr(err)

	if err != nil {
		log.Println("Booking Failed", strings.Split(err.Error(), "desc = ")[1])
	} else if booking.Id == 0 {
		log.Println("Booking Failed!\n Please check with the Dates\n You have entered as the dates which have already booked or not in the available dates")
	} else {
		log.Println("Booking Conformed with an Id : ", booking.Id)
	}

}

// Get Booked Items Details
func GetUserBookedItems(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Getting User Booked Items")

	BookingDetails, err := c.GetUserBookedItems(ctx, &rental.User{Id: 24})
	utils.CheckErr(err)

	for {
		BookedItem, err := BookingDetails.Recv()

		if err == io.EOF {
			fmt.Println("finished")
			return
		}

		utils.CheckErr(err)

		log.Printf("Booking Details is: %s\n", BookedItem)
	}

}
