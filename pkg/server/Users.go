package server

import (
	"context"
	"log"

	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func (S *ServerSideImplementation) CreateUser(ctx context.Context, in *rental.User) (*rental.User, error) {
	log.Printf("Server : Creating a New User")
	if in.GetId() == 0 {
		return &rental.User{}, nil
	}
	newUser := models.User{
		Name:        in.Name,
		Email:       &in.Email,
		PhoneNumber: in.PhoneNumber,
		Address:     in.Address,
		District:    in.District,
		PostalCode:  in.PostalCode,
		Country:     in.Country,
	}

	id, err := S.Db.CreateUser(newUser)
	utils.CheckErr(err)

	return &rental.User{Id: int32(id)}, err
}

func (S *ServerSideImplementation) UpdateUser(ctx context.Context, in *rental.User) (*rental.User, error) {
	log.Println("Server : Updating the User")
	if in.GetId() == 0 {
		return &rental.User{}, nil
	}
	user := models.User{
		Name:        in.Name,
		Email:       &in.Email,
		PhoneNumber: in.PhoneNumber,
		Address:     in.Address,
		District:    in.District,
		PostalCode:  in.PostalCode,
		Country:     in.Country,
	}
	user.ID = uint(in.GetId())

	id, err := S.Db.UpdateUser(user)
	utils.CheckErr(err)

	return &rental.User{Id: int32(id)}, err
}

func (S *ServerSideImplementation) GetUser(ctx context.Context, in *rental.User) (*rental.User, error) {
	log.Println("Server : Getting the User")

	user, err := S.Db.GetUser(int(in.GetId()))
	utils.CheckErr(err)

	return &rental.User{
		Id:          int32(user.ID),
		Name:        user.Name,
		Email:       *user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		District:    user.District,
		PostalCode:  user.PostalCode,
		Country:     user.Country,
	}, err
}
