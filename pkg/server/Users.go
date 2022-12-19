package Services

import (
	"context"
	"log"

	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
)

func (S *ServerSideImplementation) CreateUser(ctx context.Context, in *rental.User) (*rental.UserId, error) {
	log.Printf("Creating a New User")
	newUser := models.User{
		Name:         in.Name,
		Email:        in.Email,
		Phone_Number: &in.Phone_Number,
		Address:      in.Address,
		District:     in.District,
		Postal_Code:  in.Postal_Code,
		Country:      in.Country,
	}

	response := S.Db.CreateUser(newUser)

	return &rental.UserId{Id: int32(response)}, nil
}
