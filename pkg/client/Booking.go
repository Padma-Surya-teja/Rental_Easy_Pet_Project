package client

import (
	"context"
	"fmt"
	"io"
	"log"

	rental "rental_easy.in/m/pkg/rentalmgmt"
)

// Booking an Item
func Book_an_Item(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	booking_id, err := c.BookItem(ctx, &rental.Booking{StartDate: "12/01/2023", EndDate: "12/01/2023", UserId: 8, ItemId: 4})
	checkErr(err)
	if booking_id.Id == 0 {
		panic("Booking Failed")
	}

	fmt.Println("Booking Conformed with an Id : ", booking_id.Id)
}

// Get Booked Items Details
func Get_User_Booked_Items(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	Booking_Details, err := c.GetUserBookedItems(ctx, &rental.UserId{Id: 4})
	checkErr(err)

	done := make(chan bool)

	go func() {
		for {
			Booked_Item, err := Booking_Details.Recv()

			if err == io.EOF {
				done <- true
				return
			}

			checkErr(err)

			log.Printf("Booking Details is: %s", Booked_Item)
		}
	}()

	<-done
	log.Printf("finished")
}
