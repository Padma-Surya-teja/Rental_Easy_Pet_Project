package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"google.golang.org/grpc/status"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

// Function to Add Review
func AddReview(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Adding a Review")

	review, err := c.AddReview(ctx, &rental.Review{Comment: "The Phone is simply Superb and Processor is super fast", Rating: 3, UserId: 4, ItemId: 5})
	utils.CheckErr(err)

	if err != nil {
		log.Fatalf("Error : %v", strings.Split(err.Error(), "desc = ")[1])
	} else if review == nil {
		log.Fatalf("Sorry You are not allowed to add review to this Item as u have not used this item any time or You have already added the review")
	}

	log.Println("Your Review Is added Successfully with Id : ", review.Id)
}

func UpdateReview(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Updating Review")

	review, err := c.UpdateReview(ctx, &rental.Review{Id: 5, Comment: "The Phone is simply Superb Nice Camera and Processor is super fast", Rating: 5, UserId: 4, ItemId: 1})
	utils.CheckErr(err)

	if err != nil {
		log.Println("Error : ", err)
	} else if review.Id == 0 {
		log.Println("Sorry You are not allowed to add review to this Item as u have not used this item any time")
	} else {
		log.Println("Your Review Is updated Successfully with Id : ", review.Id)
	}
}

func DeleteReview(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Deleting the Review")
	review, err := c.DeleteReview(ctx, &rental.Review{Id: 5})
	utils.CheckErr(err)

	fmt.Println("Your Review Is deleted Successfully with Id : ", review.Id)
}

func GetAllReviews(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Getting Reviews")

	stream, err := c.GetAllReviews(ctx, &rental.Item{Id: 2})
	utils.CheckErr(err)
	st, ok := status.FromError(err)
	log.Println(st, ok)

	for {
		item, err := stream.Recv()

		if err == io.EOF {
			log.Printf("finished")
			return
		}

		utils.CheckErr(err)

		log.Printf("Item Received is: %s\n", item)
	}
}
