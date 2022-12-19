package Services

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"rental_easy.in/m/pkg/Mail"
	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
)

func (S *ServerSideImplementation) BookItem(ctx context.Context, in *rental.Booking) (*rental.BookingId, error) {
	startDate := Convert_to_String(strings.Split(in.GetStartDate(), "/"))
	endDate := Convert_to_String(strings.Split(in.GetEndDate(), "/"))
	Booking := models.Booking{
		Start_Date: time.Date(startDate[2], time.Month(startDate[1]), startDate[0], 0, 0, 0, 0, time.UTC),
		End_Date:   time.Date(endDate[2], time.Month(endDate[1]), endDate[0], 0, 0, 0, 0, time.UTC),
		UserID:     int(in.UserId),
		ItemID:     int(in.ItemId),
	}

	booking_id := S.Db.AddBooking(Booking)
	fmt.Println("Hello BookedItem")

	if booking_id != 0 {
		Mail.Mail(S.Db.GetUserEmail(int(in.UserId)), S.Db.Get_Item_Name(int(in.ItemId)), in.GetStartDate(), in.GetEndDate())
	}

	return &rental.BookingId{Id: int32(booking_id)}, nil

}

func (S *ServerSideImplementation) GetUserBookedItems(in *rental.UserId, stream rental.Rental_Easy_Functionalities_GetUserBookedItemsServer) error {
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
