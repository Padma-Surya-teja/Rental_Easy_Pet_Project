package client

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc/status"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

// Function to Add Review
func AddReview(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Adding a Review")

	review, err := c.AddReview(ctx, &rental.Review{Comment: "The Phone is simply Superb and Processor is super fast", Rating: 5, UserId: 12, ItemId: 1})
	utils.CheckErr(err)

	if err != nil {
		log.Fatalf("record not found")
	} else if review == nil || review.Id == 0 {
		log.Fatalf("Sorry You are not allowed to add review to this Item as u have not used this item any time or You have already added the review")
	}

	log.Println("Your Review Is added Successfully with Id : ", review.Id)
}

func UpdateReview(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Updating Review")

	review, err := c.UpdateReview(ctx, &rental.Review{Id: 2, Comment: "The Phone is simply Superb Nice Camera and Processor is super fast", Rating: 5, UserId: 12, ItemId: 2})
	utils.CheckErr(err)
	st, ok := status.FromError(err)
	log.Println(st, ok)

	if review.Id == 0 {
		fmt.Println("Sorry You are not allowed to add review to this Item as u have not used this item any time")
		return
	}

	fmt.Println("Your Review Is added Successfully with Id : ", review.Id)
}

func DeleteReview(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Deleting the Review")
	review, err := c.DeleteReview(ctx, &rental.Review{Id: 2})
	utils.CheckErr(err)

	fmt.Println("Your Review Is deleted Successfully with Id : ", review.Id)
}

func GetAllReviews(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	log.Println("Client : Getting Reviews")

	stream, err := c.GetAllReviews(ctx, &rental.Item{Id: 3})
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
