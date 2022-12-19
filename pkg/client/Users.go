package client

import (
	"context"
	"fmt"
	"io"
	"log"

	rental "rental_easy.in/m/pkg/rentalmgmt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Creating a New User
func Create_New_User(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	new_user_id, err := c.CreateUser(ctx, &rental.NewUser{Name: "Mani", Email: "Mani@gmail.com", Phone_Number: "8769806577", Address: "Secundrabad, Hyderabad", District: "Medchal", Postal_Code: "500106", Country: "India"})
	checkErr(err)
	if new_user_id.Id == 0 {
		panic("Phone Number Already Used")
	}

	fmt.Println("New User has been Created with Id : ", new_user_id.Id)
}

// Booking an Item
func Book_an_Item(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	booking_id, err := c.BookItem(ctx, &rental.Booking{StartDate: "20/12/2022", EndDate: "22/12/2022", UserId: 4, ItemId: 1})
	checkErr(err)
	if booking_id.Id == 0 {
		panic("Booking Failed")
	}

	fmt.Println("Booking Conformed with an Id : ", booking_id.Id)
}

// Get Booked Items Details
func Booked_Items(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	Booking_Details, err := c.GetBookedItems(ctx, &rental.UserId{Id: 5})
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
