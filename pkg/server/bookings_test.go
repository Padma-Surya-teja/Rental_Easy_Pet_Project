package server

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"rental_easy.in/m/pkg/database/mocks"
	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func TestBookItem(t *testing.T) {

	cases := []struct {
		description       string
		input             rental.Booking
		mockitem          models.Item
		inputmockBooking  models.Booking
		outputmockBooking models.Booking
		mockUser          models.User
		err               error
		expectedError     error
		want              rental.Booking
	}{
		{
			description:       "Success Case : When all Items are retrived",
			input:             booking1,
			mockitem:          mockItem1,
			inputmockBooking:  mockBooking1,
			outputmockBooking: mockBooking1,
			mockUser:          mockuser1,
			err:               nil,
			expectedError:     nil,
			want:              rental.Booking{Id: 1},
		},
	}

	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := mocks.NewMockDataBase(controller)

	s := ServerSideImplementation{
		Db: mockDb,
	}

	ctx := context.Background()

	for _, tt := range cases {
		mockDb.EXPECT().GetItemById(int(tt.input.ItemId)).Return(tt.mockitem, nil)
		tt.outputmockBooking.ID = 1
		mockDb.EXPECT().AddBooking(tt.inputmockBooking, tt.mockitem).Return(tt.outputmockBooking, nil)

		mockDb.EXPECT().GetUser(int(tt.input.UserId)).Return(tt.mockUser, nil)

		got, err := s.BookItem(ctx, &tt.input)
		utils.CheckErr(err)

		log.Println(got)
		expected := &rental.Booking{Id: 1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("The Function Retured is not expected one. got %v expected %v",
				got, expected)
		}

	}
}
