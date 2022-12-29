package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"rental_easy.in/m/pkg/client"
	rental "rental_easy.in/m/pkg/rentalmgmt"
)

// declaring port for the client to run on
const (
	address = "localhost:8080"
)

func main() {

	connection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10*time.Second))
	log.Printf("Client Started")
	if err != nil {
		log.Fatal("Connection Failed", err.Error())
	}
	defer connection.Close()

	c := rental.NewRental_Easy_FunctionalitiesClient(connection)
	auth := rental.NewAuthServiceClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Do the Function Call Here
	token := client.CreateUser(auth, ctx)
	// client.GetAllItems(c, ctx)

	md := metadata.New(map[string]string{"authorization": token})
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	client.GetAllItems(c, ctx)

}
