package server

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"rental_easy.in/m/pkg/database/mocks"
	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func TestCreateUser(t *testing.T) {

	// mockuser1.ID = 1
	User1.Id = 1
	cases := []struct {
		description string
		input1      models.User
		input2      rental.User
		err         error
		expectedErr error
		expected    *rental.User
	}{
		{
			description: "Success Case : When the User with all values sent ",
			input1:      mockuser1,
			input2:      User1,
			err:         nil,
			expectedErr: nil,
			expected:    &rental.User{Id: int32(1)},
		},
		{
			description: "Fail Case : When the User Gets Repeated",
			input1:      mockuser2,
			input2:      User2,
			err:         errors.New(`pq: duplicate key value violates unique constraint "users_email_key"`),
			expectedErr: errors.New(`pq: duplicate key value violates unique constraint "users_email_key"`),
			expected:    &rental.User{},
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
		mockDb.EXPECT().CreateUser(cases[0].input1).Return(1, nil)

		got, err := s.CreateUser(ctx, &cases[0].input2)
		utils.CheckErr(err)
		if err != nil {
			if !errors.Is(err, cases[0].expectedErr) {
				t.Errorf("The Function Retured is not expected one. got %v expected %v",
					got, cases[0].expected)
			}
		}
		if !reflect.DeepEqual(got, cases[0].expected) {
			t.Errorf("The Function Retured is not expected one. got %v expected %v",
				got, cases[0].expected)
		}
	})
	t.Run(cases[1].description, func(t *testing.T) {
		mockDb.EXPECT().CreateUser(cases[1].input1).Return(0, cases[1].err)

		got, err := s.CreateUser(ctx, &cases[1].input2)
		utils.CheckErr(err)
		if err != nil {
			if !reflect.DeepEqual(got, cases[1].expected) {
				t.Errorf("The Function Retured is not expected one. got %v expected %v",
					got, cases[0].expected)
			}
		}
		if !reflect.DeepEqual(got, cases[1].expected) {
			t.Errorf("The Function Retured is not expected one. got %v expected %v",
				got, cases[1].expected)
		}
	})

}

func TestUpdateUser(t *testing.T) {
	cases := []struct {
		description string
		input1      models.User
		input2      rental.User
		expected    *rental.User
	}{
		{
			description: "Success Case : When the User with all values sent ",
			input1:      mockuser1,
			input2:      User1,
			expected:    &rental.User{Id: 1},
		},
		{
			description: "Fail Case : When the Empty User is sent ",
			input1:      models.User{},
			input2:      rental.User{},
			expected:    &rental.User{Id: 0},
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
		cases[0].input1.ID = 1
		mockDb.EXPECT().UpdateUser(cases[0].input1).Return(1, nil)

		cases[0].input2.Id = 1
		got, err := s.UpdateUser(ctx, &cases[0].input2)
		utils.CheckErr(err)
		if !reflect.DeepEqual(got, cases[0].expected) {
			t.Errorf("The Function Retured is not expected one. got %v expected %v",
				got, cases[0].expected)
		}
	})

	t.Run(cases[1].description, func(t *testing.T) {

		got, err := s.UpdateUser(ctx, &cases[1].input2)
		utils.CheckErr(err)

		if !reflect.DeepEqual(got, cases[1].expected) {
			t.Errorf("The Function Retured is not expected one. got %v expected %v",
				got, cases[0].expected)
		}
	})
}

func TestGetUser(t *testing.T) {
	cases := []struct {
		description string
		input1      int
		input2      rental.User
		expected    *rental.User
	}{
		{
			description: "Success Case : When the User with all values sent ",
			input1:      1,
			input2:      rental.User{Id: 1},
			expected:    &User1,
		},
		{
			description: "Fail Case : When the Empty User is sent ",
			input1:      0,
			input2:      rental.User{},
			expected:    &rental.User{},
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

		mockuser1.ID = 1
		mockDb.EXPECT().GetUser(cases[0].input1).Return(mockuser1, nil)

		got, err := s.GetUser(ctx, &cases[0].input2)
		utils.CheckErr(err)
		cases[0].expected.Id = 1
		if !reflect.DeepEqual(got, cases[0].expected) {
			t.Errorf("The Function Retured is not expected one. got %v expected %v",
				got, cases[0].expected)
		}
	})
}
