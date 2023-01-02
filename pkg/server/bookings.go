package server

import (
	"context"
	"log"
	"strings"
	"time"

	"rental_easy.in/m/pkg/Mail"
	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func (S *ServerSideImplementation) BookItem(ctx context.Context, in *rental.Booking) (*rental.Booking, error) {
	log.Println("Server : Booking an Item")

	item, err := S.Db.GetItemById(int(in.GetItemId()))
	utils.CheckErr(err)
	if err != nil {
		return &rental.Booking{}, err
	}

	startDate := utils.ConvertToInt(strings.Split(in.GetStartDate(), "-"))
	endDate := utils.ConvertToInt(strings.Split(in.GetEndDate(), "-"))

	start := time.Date(startDate[2], time.Month(startDate[1]), startDate[0], 0, 0, 0, 0, time.UTC)
	end := time.Date(endDate[2], time.Month(endDate[1]), endDate[0], 0, 0, 0, 0, time.UTC)
	NoofDays := (end.Sub(start)).Hours() / 24
	Booking := models.Booking{
		StartDate:   time.Date(startDate[2], time.Month(startDate[1]), startDate[0], 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(endDate[2], time.Month(endDate[1]), endDate[0], 0, 0, 0, 0, time.UTC),
		TotalAmount: (int(NoofDays) + 1) * item.AmountPerDay,
		UserId:      int(in.GetUserId()),
		ItemId:      int(in.GetItemId()),
	}

	booking, err := S.Db.AddBooking(Booking, item)
	utils.CheckErr(err)
	if err != nil {
		return &rental.Booking{}, err
	}

	user, err := S.Db.GetUser(int(in.GetUserId()))

	log.Println("Sending Booking Confirmation Mail to User")
	e := make(chan error)
	go Mail.BookingConfirmationMail(*user.Email, item.Name, in.GetStartDate(), in.GetEndDate(), e)

	if <-e == nil {
		log.Println("Mail has been send to user")
	}
	close(e)

	return &rental.Booking{Id: int32(booking.ID)}, err

}

func (S *ServerSideImplementation) GetUserBookedItems(in *rental.User, stream rental.Rental_Easy_Functionalities_GetUserBookedItemsServer) error {
	log.Println("Getting User Booked Items")

	bookings, err := S.Db.GetBookings(int(in.Id))
	utils.CheckErr(err)
	if err != nil {
		return err
	}

	for _, booking := range bookings {
		Item, _ := S.Db.GetItemById(booking.ItemId)
		if Item.Name == "" {
			Item.Name = "Item has been removed by the owner"
		}
		bookingDetails := rental.Booking{
			Id:          int32(booking.ID),
			ItemName:    Item.Name,
			StartDate:   booking.StartDate.String(),
			EndDate:     booking.EndDate.String(),
			TotalAmount: int32(booking.TotalAmount),
		}
		err = stream.Send(&bookingDetails)
		utils.CheckErr(err)
	}
	return nil
}
