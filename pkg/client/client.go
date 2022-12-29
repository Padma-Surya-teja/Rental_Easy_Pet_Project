package client

import (
	"google.golang.org/grpc"
	"rental_easy.in/m/pkg/rentalmgmt"
)

type RentalClient struct {
	service rentalmgmt.Rental_Easy_FunctionalitiesClient
}

func NewClient(cc *grpc.ClientConn) *RentalClient {
	service := rentalmgmt.NewRental_Easy_FunctionalitiesClient(cc)
	return &RentalClient{service}
}
