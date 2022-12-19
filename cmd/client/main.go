package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"rental_easy.in/m/pkg/client"
	rental "rental_easy.in/m/pkg/rentalmgmt"
)

// declaring port for the client to run on
const (
	address = "localhost:8080"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	connection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10*time.Second))
	log.Printf("Client Started")
	if err != nil {
		log.Fatal("Connection Failed", err.Error())
	}
	defer connection.Close()
	c := rental.NewRental_Easy_FunctionalitiesClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client.Booked_Items(c, ctx)
}
