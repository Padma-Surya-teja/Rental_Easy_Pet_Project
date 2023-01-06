package server

import (
	"context"
	"log"

	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

type User struct {
	Username       string
	HashedPassword string
}

func (S *ServerSideImplementation) CreateUser(ctx context.Context, in *rental.User) (*rental.User, error) {
	log.Printf("Server : Creating a New User")

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
	if err != nil {
		return &rental.User{}, err
	}

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

// func NewUser(username string, password string) (*User, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot hash password : %w", err)
// 	}

// 	user := &User{
// 		Username:       username,
// 		HashedPassword: string(hashedPassword),
// 	}

// 	return user, nil
// }

// func (user *User) IsCorrectPassword(password string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
// 	return err == nil
// }

// func (user *User) Clone() *User {
// 	return &User{
// 		Username:       user.Username,
// 		HashedPassword: user.HashedPassword,
// 	}
// }
