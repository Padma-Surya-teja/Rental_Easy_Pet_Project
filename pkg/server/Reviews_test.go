package server

import (
	"context"
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"rental_easy.in/m/pkg/database/mocks"
	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func TestAddReview(t *testing.T) {
	cases := []struct {
		description     string
		item            models.Item
		mockreview      models.Review
		review          rental.Review
		userbooked      bool
		alreadyreviewed bool
		bookingid       int
		expected        *rental.Review
	}{
		{
			description:     "Success Case : When the User Who booked adds a review",
			item:            mockItem1,
			mockreview:      mockReview1,
			review:          Review1,
			userbooked:      true,
			alreadyreviewed: false,
			bookingid:       1,
			expected:        &rental.Review{Id: 1},
		},
		{
			description:     "Failure Case : When the Item is not there",
			item:            models.Item{},
			mockreview:      mockReview2,
			review:          Review2,
			userbooked:      false,
			alreadyreviewed: false,
			bookingid:       0,
			expected:        &rental.Review{},
		},
		{
			description:     "Failure Case : When the Item is not there",
			item:            mockItem1,
			mockreview:      mockReview2,
			review:          Review2,
			userbooked:      false,
			alreadyreviewed: false,
			bookingid:       0,
			expected:        &rental.Review{},
		},
		{
			description:     "Failure Case : When the User Already Reviewed",
			item:            mockItem1,
			mockreview:      mockReview2,
			review:          Review2,
			userbooked:      true,
			alreadyreviewed: true,
			bookingid:       0,
			expected:        &rental.Review{},
		},
	}
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := mocks.NewMockDataBase(controller)

	s := ServerSideImplementation{
		Db: mockDb,
	}

	ctx := context.Background()

	t.Run(cases[0].description, func(t *testing.T) {
		mockDb.EXPECT().GetItemById(int(cases[0].review.ItemId)).Return(cases[0].item, nil)
		mockDb.EXPECT().AddReview(cases[0].mockreview).Return(cases[0].bookingid, nil)
		mockDb.EXPECT().UserAlreadyAddedReview(int(cases[0].review.UserId), int(cases[0].review.ItemId)).Return(cases[0].alreadyreviewed, nil)
		mockDb.EXPECT().UserAlreadyBooked(int(cases[0].review.UserId), int(cases[0].review.ItemId)).Return(cases[0].userbooked, nil)

		got, err := s.AddReview(ctx, &cases[0].review)
		utils.CheckErr(err)

		log.Println("got : ", got)
		log.Println("expected :", cases[0].expected)

		if !reflect.DeepEqual(got, cases[0].expected) {
			t.Errorf("The Function Retured is not expected one. got %v expected %v",
				got, cases[0].expected)
		}

	})

	t.Run(cases[1].description, func(t *testing.T) {
		mockDb.EXPECT().GetItemById(int(cases[1].review.ItemId)).Return(cases[1].item, errors.New("no record found"))
		// mockDb.EXPECT().AddReview(cases[1].mockreview).Return(cases[1].bookingid, nil)
		// mockDb.EXPECT().UserAlreadyAddedReview(int(cases[1].review.UserId), int(cases[1].review.ItemId)).Return(cases[0].alreadyreviewed, nil)
		// mockDb.EXPECT().UserAlreadyBooked(int(cases[1].review.UserId), int(cases[1].review.ItemId)).Return(cases[0].userbooked, nil)

		got, err := s.AddReview(ctx, &cases[1].review)
		utils.CheckErr(err)

		log.Println("got : ", got)
		log.Println("expected :", cases[1].expected)

		if !reflect.DeepEqual(got, cases[1].expected) {
			t.Errorf("The Function Retured is not expected one. got %v expected %v",
				got, cases[1].expected)
		}

	})

	t.Run(cases[2].description, func(t *testing.T) {
		mockDb.EXPECT().GetItemById(int(cases[2].review.ItemId)).Return(cases[2].item, nil)
		// mockDb.EXPECT().AddReview(cases[1].mockreview).Return(cases[1].bookingid, nil)
		// mockDb.EXPECT().UserAlreadyAddedReview(int(cases[1].review.UserId), int(cases[1].review.ItemId)).Return(cases[0].alreadyreviewed, nil)
		mockDb.EXPECT().UserAlreadyBooked(int(cases[2].review.UserId), int(cases[2].review.ItemId)).Return(cases[2].userbooked, nil)

		got, err := s.AddReview(ctx, &cases[2].review)
		utils.CheckErr(err)

		log.Println("got : ", got)
		log.Println("expected :", cases[2].expected)

		if !reflect.DeepEqual(got, cases[2].expected) {
			t.Errorf("The Function Retured is not expected one. got %v expected %v",
				got, cases[2].expected)
		}

	})

	t.Run(cases[3].description, func(t *testing.T) {
		mockDb.EXPECT().GetItemById(int(cases[3].review.ItemId)).Return(cases[3].item, nil)
		// mockDb.EXPECT().AddReview(cases[1].mockreview).Return(cases[1].bookingid, nil)
		mockDb.EXPECT().UserAlreadyAddedReview(int(cases[3].review.UserId), int(cases[3].review.ItemId)).Return(cases[3].alreadyreviewed, nil)
		mockDb.EXPECT().UserAlreadyBooked(int(cases[3].review.UserId), int(cases[3].review.ItemId)).Return(cases[3].userbooked, nil)

		got, err := s.AddReview(ctx, &cases[3].review)
		utils.CheckErr(err)

		log.Println("got : ", got)
		log.Println("expected :", cases[3].expected)

		if !reflect.DeepEqual(got, cases[3].expected) {
			t.Errorf("The Function Retured is not expected one. got %v expected %v",
				got, cases[3].expected)
		}

	})

}

func TestUpdateReview(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := mocks.NewMockDataBase(controller)

	s := ServerSideImplementation{
		Db: mockDb,
	}

	ctx := context.Background()

	mockReview1.ID = 1
	mockDb.EXPECT().UpdateReview(mockReview1).Return(1, nil)

	Review1.Id = 1
	got, err := s.UpdateReview(ctx, &Review1)
	utils.CheckErr(err)

	expected := &rental.Review{Id: 1}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestDeleteReview(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := mocks.NewMockDataBase(controller)

	s := ServerSideImplementation{
		Db: mockDb,
	}

	ctx := context.Background()

	mockDb.EXPECT().DeleteReview(1).Return(1, nil)

	got, err := s.DeleteReview(ctx, &rental.Review{Id: 1})
	utils.CheckErr(err)

	expected := &rental.Review{Id: 1}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}
