package main

import (
	"flag"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"rental_easy.in/m/pkg/database"
	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/server"
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

// func createUser(username, password string) ( error) {
// 	_, err := server.NewUser(username, password)
// 	return err
// }

// func seedUsers() error {
// 	err := createUser("admin1", "secret")
// 	if err != nil {
// 		return err
// 	}
// 	return createUser("user1", "secret")
// }

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

func accessibleRoles() map[string][]string {
	const laptopServicePath = ""

	return map[string][]string{}
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
	jwtManager := server.NewJWTManager(secretKey, tokenDuration)
	authServer := server.NewAuthServer(jwtManager)
	if err != nil {
		log.Fatal("cannot seed users: ", err)
	}

	//Creating a New Server
	interceptor := server.NewAuthInterceptor(jwtManager, accessibleRoles())
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	log.Printf("server listening at %v", listener.Addr())

	rental.RegisterAuthServiceServer(grpcServer, authServer)
	rental.RegisterRental_Easy_FunctionalitiesServer(grpcServer, &Services.ServerSideImplementation{
		Db: database.DBClient{Db: db},
	})

	//if connection fails
	err = grpcServer.Serve(listener)
	checkErr(err)
}
