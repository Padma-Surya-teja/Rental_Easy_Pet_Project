package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"rental_easy.in/m/pkg/database/mocks"
	rental "rental_easy.in/m/pkg/rentalmgmt"
)

func MockServer(mockDb *mocks.MockDataBase, ctx context.Context) (rental.Rental_Easy_FunctionalitiesClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	s := ServerSideImplementation{
		Db: mockDb,
	}
	rental.RegisterRental_Easy_FunctionalitiesServer(baseServer, &s)
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := rental.NewRental_Easy_FunctionalitiesClient(conn)

	return client, closer
}
