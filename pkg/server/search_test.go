package server

import (
	"context"
	"io"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"rental_easy.in/m/pkg/database/mocks"
	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func TestSearchItems(t *testing.T) {
	log.Println("Testing Search Items")
	search1[0].ID = 5
	search1[1].ID = 7
	search2[0].ID = 6
	search2[1].ID = 8

	cases := []struct {
		description string
		input       []models.Item
		req         rental.ItemRequest
		err         error
		expectedErr error
		want        []rental.Item
	}{
		{
			description: "Success Case : With Search Request",
			input:       search1,
			req:         rental.ItemRequest{Request: "Iqoo Neo 6"},
			err:         nil,
			expectedErr: nil,
			want:        result1,
		},
		{
			description: "Success Case : With Category Request",
			input:       search2,
			req:         rental.ItemRequest{Category: 1},
			err:         nil,
			expectedErr: nil,
			want:        result2,
		},
		{
			description: "Failure Case",
			input:       []models.Item{},
			req:         rental.ItemRequest{Category: 2},
			err:         nil,
			expectedErr: nil,
			want:        []rental.Item{},
		},
	}

	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := mocks.NewMockDataBase(controller)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	//Client
	c, closer := MockServer(mockDb, ctx)
	defer closer()

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			mockDb.EXPECT().SearchItems(tt.req.Request, tt.req.Category.String()).Return(tt.input, tt.err)
			stream, err := c.SearchItems(ctx, &tt.req)
			utils.CheckErr(err)

			for _, expected := range tt.want {
				got, err := stream.Recv()

				if err == io.EOF {
					log.Printf("finished")
					return
				}
				utils.CheckErr(err)

				if !reflect.DeepEqual(got, &expected) {
					t.Errorf("The Function Retured is not expected one. got %v expected %v",
						got, &expected)
				}
			}
		})
	}

}
