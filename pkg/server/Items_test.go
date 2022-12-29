package server

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"rental_easy.in/m/pkg/database/mocks"
	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func TestCreateItem(t *testing.T) {

	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := mocks.NewMockDataBase(controller)

	s := ServerSideImplementation{
		Db: mockDb,
	}

	ctx := context.Background()

	mockDb.EXPECT().AddItem(mockItem1).Return(1, nil)

	got, err := s.CreateItem(ctx, &inputItem)
	utils.CheckErr(err)

	expected := &rental.Item{Id: 1}

	if !reflect.DeepEqual(got, expected) && err == nil {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}

}

func TestGetItemById(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := mocks.NewMockDataBase(controller)

	s := ServerSideImplementation{
		Db: mockDb,
	}

	ctx := context.Background()

	mockDb.EXPECT().GetItemById(1).Return(mockItem1, nil)

	got, err := s.GetItemById(ctx, &rental.Item{Id: 1})
	utils.CheckErr(err)

	expected := &Item1

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestUpdateItem(t *testing.T) {
	cases := []struct {
		description string
		input       rental.Item
		updateditem models.Item
		item        models.Item
		expected    *rental.Item
	}{
		{
			description: "Success Case : When the User is the owner ",
			input: rental.Item{Id: 1, Name: "Iqoo Neo 6",
				Description:   "Phone is Superfast and Display is just amazing",
				Category:      "Mobile Phones",
				AvailableFrom: "01-01-2023",
				AvailableTo:   "03-01-2023",
				AmountPerDay:  500,
				UserId:        1},
			updateditem: models.Item{
				Name:          "Iqoo Neo 6",
				Description:   "Phone is Superfast and Display is just amazing",
				Category:      "Mobile Phones",
				AvailableFrom: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
				AvailableTo:   time.Date(2023, 01, 03, 0, 0, 0, 0, time.UTC),
				AmountPerDay:  500,
				UserId:        1,
			},
			item: models.Item{Name: "Iqoo Neo 6",
				Description:   "Phone is Superfast and Display is just amazing",
				Category:      "Mobile Phones",
				AvailableFrom: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
				AvailableTo:   time.Date(2023, 01, 03, 0, 0, 0, 0, time.UTC),
				AmountPerDay:  500,
				UserId:        1},
			expected: &rental.Item{Id: 1},
		},
		{
			description: "Fail Case : When the User is not the owner ",
			input: rental.Item{Id: 1, Name: "Iqoo Neo 6",
				Description:   "Phone is Superfast and Display is just amazing",
				Category:      "Mobile Phones",
				AvailableFrom: "01-01-2023",
				AvailableTo:   "03-01-2023",
				AmountPerDay:  500,
				UserId:        10},
			updateditem: models.Item{
				Name:          "Iqoo Neo 6",
				Description:   "Phone is Superfast and Display is just amazing",
				Category:      "Mobile Phones",
				AvailableFrom: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
				AvailableTo:   time.Date(2023, 01, 03, 0, 0, 0, 0, time.UTC),
				AmountPerDay:  500,
				UserId:        10,
			},
			item: models.Item{Name: "Iqoo Neo 6",
				Description:   "Phone is Superfast and Display is just amazing",
				Category:      "Mobile Phones",
				AvailableFrom: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
				AvailableTo:   time.Date(2023, 01, 03, 0, 0, 0, 0, time.UTC),
				AmountPerDay:  500,
				UserId:        1},
			expected: &rental.Item{},
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
		cases[0].updateditem.ID = 1
		mockDb.EXPECT().UpdateItem(cases[0].updateditem).Return(1, nil)
		cases[0].item.ID = 1
		mockDb.EXPECT().GetItemById(int(cases[0].input.Id)).Return(cases[0].item, nil)

		got, err := s.UpdateItem(ctx, &cases[0].input)
		utils.CheckErr(err)

		if !reflect.DeepEqual(got, cases[0].expected) {
			t.Errorf("The Function Retured is not expected one. got %v expected %v",
				got, cases[0].expected)
		}
	})
	t.Run(cases[1].description, func(t *testing.T) {
		// cases[1].updateditem.ID = 1
		// mockDb.EXPECT().UpdateItem(cases[1].updateditem).Return(1, nil)
		cases[1].item.ID = uint(cases[1].input.Id)
		mockDb.EXPECT().GetItemById(int(cases[1].input.Id)).Return(cases[1].item, nil)

		got, err := s.UpdateItem(ctx, &cases[1].input)
		utils.CheckErr(err)

		if !reflect.DeepEqual(got, cases[1].expected) {
			t.Errorf("The Function Retured is not expected one. got %v expected %v",
				got, cases[1].expected)
		}
	})
}

func TestDeleteItem(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := mocks.NewMockDataBase(controller)

	s := ServerSideImplementation{
		Db: mockDb,
	}

	ctx := context.Background()

	mockDb.EXPECT().DeleteItem(1).Return(1, nil)

	got, err := s.DeleteItem(ctx, &rental.Item{Id: 1})
	utils.CheckErr(err)

	expected := &rental.Item{Id: 1}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}
