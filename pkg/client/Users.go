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
	new_user_id, err := c.CreateUser(ctx, &rental.NewUser{Name: "Sai", Email: "Sai@gmail.com", Phone_Number: "8769800987", Address: "Secundrabad, Hyderabad", District: "Medchal", Postal_Code: "500106", Country: "India"})
	checkErr(err)
	if new_user_id.Id == 0 {
		panic("Please Change the Username")
	}

	fmt.Println("New User has been Created with Id : ", new_user_id.Id)
}
