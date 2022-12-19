package client

import (
	"context"
	"fmt"

	rental "rental_easy.in/m/pkg/rentalmgmt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Creating a New User
func Create_New_User(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	new_user_id, err := c.CreateUser(ctx, &rental.User{Name: "Surya", Email: "19bd1a057d@gmail.com", Phone_Number: "9398200123", Address: "Saidabad, Hyderabad", District: "Hyderabad", Postal_Code: "500059", Country: "India"})
	checkErr(err)
	if new_user_id.Id == 0 {
		panic("Phone Number Already Used")
	}

	fmt.Println("New User has been Created with Id : ", new_user_id.Id)
}
