package client

import (
	"context"
	"fmt"
	"io"
	"log"

	rental "rental_easy.in/m/pkg/rentalmgmt"
)

// Function to Add Review
func Add_Review(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	new_review_id, err := c.AddReview(ctx, &rental.Review{Comment: "The Laptop is simply Superb and Processor is super fast", Rating: 5, UserId: 7, ItemId: 4})

	checkErr(err)

	if new_review_id.Id == 0 {
		fmt.Println("Sorry You are not allowed to add review to this Item as u have not used this item any time or You have already added the review")
		return
	}

	fmt.Println("Your Review Is added Successfully with Id : ", new_review_id.Id)
}

func Update_Review(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	updated_review_id, err := c.UpdateReview(ctx, &rental.Review{Id: 1, Comment: "The Phone is simply Superb and Processor is super fast", Rating: 5, UserId: 4, ItemId: 1})

	checkErr(err)
	if updated_review_id.Id == 0 {
		fmt.Println("Sorry You are not allowed to add review to this Item as u have not used this item any time")
		return
	}

	fmt.Println("Your Review Is added Successfully with Id : ", updated_review_id.Id)
}

func DeleteReview(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	deleted_review_id, err := c.DeleteReview(ctx, &rental.ReviewId{Id: 3})

	checkErr(err)

	fmt.Println("Your Review Is deleted Successfully with Id : ", deleted_review_id.Id)
}

func GetAllReviews(c rental.Rental_Easy_FunctionalitiesClient, ctx context.Context) {
	stream_of_reviews, err := c.GetAllReviews(ctx, &rental.ItemId{Id: 1})
	checkErr(err)

	done := make(chan bool)

	go func() {
		for {
			item, err := stream_of_reviews.Recv()

			if err == io.EOF {
				done <- true
				return
			}

			checkErr(err)

			log.Printf("Item Received is: %s", item)
		}
	}()

	<-done
	log.Printf("finished")
}
