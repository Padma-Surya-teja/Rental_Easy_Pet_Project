package server

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"rental_easy.in/m/pkg/database/mocks"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func TestBookItem(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := mocks.NewMockDataBase(controller)

	s := ServerSideImplementation{
		Db: mockDb,
	}

	ctx := context.Background()

	mockDb.EXPECT().GetItemById(1).Return(mockItem1, nil)
	mockBooking2.ID = 1
	mockDb.EXPECT().AddBooking(mockBooking1).Return(mockBooking2, nil)
	mockuser1.ID = 1
	mockDb.EXPECT().GetUser(1).Return(mockuser1, nil)

	got, err := s.BookItem(ctx, &booking1)
	utils.CheckErr(err)

	log.Println(got)
	expected := &rental.Booking{Id: 1}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}
