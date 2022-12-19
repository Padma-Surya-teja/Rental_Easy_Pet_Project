package Services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
)

func (S *ServerSideImplementation) CreateUser(ctx context.Context, in *rental.NewUser) (*rental.UserId, error) {
	log.Printf("Creating a New User")
	newUser := models.User{
		Name:         &in.Name,
		Email:        in.Email,
		Phone_Number: in.Phone_Number,
		Address:      in.Address,
		District:     in.District,
		Postal_Code:  in.Postal_Code,
		Country:      in.Country,
	}

	fmt.Println("New User \n", newUser)
	response := S.Db.CreateUser(newUser)

	return &rental.UserId{Id: int32(response)}, nil
}

func (S *ServerSideImplementation) BookItem(ctx context.Context, in *rental.Booking) (*rental.BookingId, error) {
	startDate := Convert_to_String(strings.Split(in.GetStartDate(), "/"))
	endDate := Convert_to_String(strings.Split(in.GetEndDate(), "/"))
	Booking := models.Booking{
		Start_Date: time.Date(startDate[2], time.Month(startDate[1]), startDate[0], 0, 0, 0, 0, time.UTC),
		End_Date:   time.Date(endDate[2], time.Month(endDate[1]), endDate[0], 0, 0, 0, 0, time.UTC),
		UserID:     int(in.UserId),
		ItemID:     int(in.ItemId),
	}

	response := S.Db.AddBooking(Booking)

	return &rental.BookingId{Id: int32(response)}, nil

}

func (S *ServerSideImplementation) GetBookedItems(in *rental.UserId, stream rental.Rental_Easy_Functionalities_GetBookedItemsServer) error {
	Bookings := S.Db.GetBookings(int(in.Id))

	var wg sync.WaitGroup
	for _, Booking := range Bookings {
		wg.Add(1)
		go func(Booking models.Booking) {
			defer wg.Done()
			time.Sleep(time.Duration(100) * time.Microsecond)
			Item := S.Db.GetItemById(Booking.ItemID)
			Dates_Difference := Booking.End_Date.Sub(Booking.Start_Date)
			Booked := rental.BookedItems{
				Id:          int32(Booking.ID),
				ItemName:    Item.Name,
				StartDate:   Booking.Start_Date.String(),
				EndDate:     Booking.End_Date.String(),
				TotalAmount: int32(Dates_Difference.Hours()/24) * int32(Item.Amount_per_day),
			}
			err := stream.Send(&Booked)
			checkErr(err)
		}(Booking)
	}

	wg.Wait()
	return nil
}
