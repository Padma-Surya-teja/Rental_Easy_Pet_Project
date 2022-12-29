package client

import (
	"context"
	"log"

	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func CreateUser(c rental.AuthServiceClient, ctx context.Context) string {
	log.Println("Client : Logging in the user")
	response, err := c.Login(ctx, &rental.LoginRequest{Username: "user1", Password: "secret"})
	utils.CheckErr(err)

	if response == nil {
		return ""
	}
	return response.AccessToken
}
