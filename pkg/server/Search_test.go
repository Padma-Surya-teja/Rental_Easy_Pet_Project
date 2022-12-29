package server

// import (
// 	"context"
// 	"io"
// 	"log"
// 	"testing"
// 	"time"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/golang/protobuf/proto"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/metadata"

// 	"rental_easy.in/m/pkg/client"
// 	"rental_easy.in/m/pkg/models"
// 	rental "rental_easy.in/m/pkg/rentalmgmt"
// 	"rental_easy.in/m/pkg/utils"
// )

// const (
// 	port = "localhost:8080"
// )

// const (
// 	address = "localhost:8080"
// )

// func TestSearchItems(t *testing.T) {
// 	log.Println("Testing Search Items")
// 	cases := []struct {
// 		description string
// 		request     string
// 		category    string
// 		input1      []models.Item
// 		input2      []models.Item
// 		req         *rental.ItemRequest
// 		want        []rental.Item
// 	}{
// 		{
// 			description: "Success Case : With Search Request",
// 			request:     "Iqoo neo 6",
// 			category:    "",
// 			input1:      search1,
// 			input2:      search2,
// 			req:         &req1,
// 			want:        result1,
// 		},
// 		{
// 			description: "Success Case : With Category Request",
// 			request:     "",
// 			category:    "Laptops",
// 			input1:      search1,
// 			input2:      search2,
// 			req:         &req2,
// 			want:        result2,
// 		},
// 	}

// 	//Client
// 	connection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10*time.Second))
// 	log.Printf("Client Started")
// 	if err != nil {
// 		log.Fatal("Connection Failed", err.Error())
// 	}
// 	defer connection.Close()
// 	c := rental.NewRental_Easy_FunctionalitiesClient(connection)
// 	auth := rental.NewAuthServiceClient(connection)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancel()
// 	token := client.CreateUser(auth, ctx)
// 	// client.GetAllItems(c, ctx)

// 	md := metadata.New(map[string]string{"authorization": token})
// 	ctx = metadata.NewOutgoingContext(context.Background(), md)

// 	i := 0
// 	for _, tt := range cases {
// 		t.Run(tt.description, func(t *testing.T) {

// 			tt.input1[0].ID = 5
// 			tt.input1[1].ID = 7
// 			tt.input2[0].ID = 6
// 			tt.input2[1].ID = 8

// 			myMockRows1 := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "description", "category", "available_from", "available_to", "amount_per_day", "user_id"}).
// 				AddRow(tt.input1[0].ID, tt.input1[0].CreatedAt, tt.input1[0].UpdatedAt, tt.input1[0].DeletedAt, tt.input1[0].Name, tt.input1[0].Description, tt.input1[0].Category, tt.input1[0].AvailableFrom, tt.input1[0].AvailableTo, tt.input1[0].AmountPerDay, tt.input1[0].UserId).
// 				AddRow(tt.input1[1].ID, tt.input1[1].CreatedAt, tt.input1[1].UpdatedAt, tt.input1[1].DeletedAt, tt.input1[1].Name, tt.input1[1].Description, tt.input1[1].Category, tt.input1[1].AvailableFrom, tt.input1[1].AvailableTo, tt.input1[1].AmountPerDay, tt.input1[1].UserId)

// 			myMockRows2 := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "description", "category", "available_from", "available_to", "amount_per_day", "user_id"}).
// 				AddRow(tt.input2[0].ID, tt.input2[0].CreatedAt, tt.input2[0].UpdatedAt, tt.input2[0].DeletedAt, tt.input2[0].Name, tt.input2[0].Description, tt.input2[0].Category, tt.input2[0].AvailableFrom, tt.input2[0].AvailableTo, tt.input2[0].AmountPerDay, tt.input2[0].UserId).
// 				AddRow(tt.input2[1].ID, tt.input2[1].CreatedAt, tt.input2[1].UpdatedAt, tt.input2[1].DeletedAt, tt.input2[1].Name, tt.input2[1].Description, tt.input2[1].Category, tt.input2[1].AvailableFrom, tt.input2[1].AvailableTo, tt.input2[1].AmountPerDay, tt.input2[1].UserId)

// 			mock.ExpectQuery(`SELECT * FROM "items"  WHERE "items"."deleted_at" IS NULL AND ((category LIKE $1))`).WithArgs(tt.category).WillReturnRows(myMockRows1)
// 			mock.ExpectQuery(`SELECT * FROM "items"  WHERE "items"."deleted_at" IS NULL AND ((name LIKE $1))`).WithArgs(tt.request).WillReturnRows(myMockRows2)

// 			stream, err := c.SearchItems(ctx, tt.req)
// 			utils.CheckErr(err)

// 			for _, expected := range tt.want {
// 				i++
// 				got, err := stream.Recv()

// 				if err == io.EOF {
// 					log.Printf("finished")
// 					return
// 				}
// 				utils.CheckErr(err)

// 				log.Println(i, " got : ", proto.Message(got))
// 				log.Println(i, " expected : ", proto.Message(&expected))
// 				if proto.Equal(proto.Message(got), proto.Message(&expected)) {
// 					t.Errorf("The Function Retured is not expected one. got %v expected %v",
// 						got, expected)
// 				}
// 			}
// 		})
// 	}

// }
