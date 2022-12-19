package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"rental_easy.in/m/pkg/database"
	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	Services "rental_easy.in/m/pkg/server"
)

// declaring the port number
const (
	port = "localhost:8080"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//Using command line flags to connect with the database
	db_user := flag.String("user", "postgres", "database user")
	db_password := flag.String("password", "root", "database password")
	db_name := flag.String("name", "Rental_Easy", "database name")
	db := models.SetupDB(db_user, db_password, db_name)

	listener, err := net.Listen("tcp", port)
	checkErr(err)

	defer db.Close()

	//Creating a New Server
	server := grpc.NewServer()
	log.Printf("server listening at %v", listener.Addr())

	rental.RegisterRental_Easy_FunctionalitiesServer(server, &Services.ServerSideImplementation{
		Db: database.DBClient{Db: db},
	})

	//if connection fails
	err = server.Serve(listener)
	checkErr(err)
}
