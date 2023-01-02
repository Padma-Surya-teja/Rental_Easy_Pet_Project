package client

import (
	"context"
	"log"

	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

// Creating a New User
func CreateNewUser(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Creating a User From Client Side")
	user, err := c.CreateUser(ctx, &rental.User{})
	utils.CheckErr(err)

	if user == nil {
		log.Fatalf("Please Enter all values in the database")
	}
	if err != nil {
		log.Fatalf("Email Id Already Used")
	}

	log.Println("New User has been Created with Id : ", user.Id)
}

func GetUser(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Getting a User From Client Side")
	user, err := c.GetUser(ctx, &rental.User{Id: 23})
	utils.CheckErr(err)

	if err == nil {
		log.Println("User Details : \n ", user)
	}
}

func UpdateUser(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Updating a User From Client Side")
	User2.Id = 2
	user, err := c.UpdateUser(ctx, &User2)
	utils.CheckErr(err)

	if err == nil {
		log.Println("User has been Updated with Id : ", user.Id)
	}
}
