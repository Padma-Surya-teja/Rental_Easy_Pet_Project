package Services

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
)

func (S *ServerSideImplementation) AddReview(ctx context.Context, in *rental.Review) (*rental.ReviewId, error) {
	log.Printf("Creating a New Review")

	fmt.Println(int(in.UserId), int(in.ItemId))
	if !S.Db.UserAlreadyBooked(int(in.UserId), int(in.ItemId)) || S.Db.UserAlreadyAddedReview(int(in.UserId), int(in.ItemId)) {
		return &rental.ReviewId{Id: 0}, nil
	}

	newReview := models.Review{
		Comments: in.Comment,
		Rating:   int(in.Rating),
		UserID:   int(in.UserId),
		ItemID:   int(in.ItemId),
	}

	reviewId := S.Db.AddReview(newReview)

	return &rental.ReviewId{Id: int32(reviewId)}, nil
}

func (S *ServerSideImplementation) UpdateReview(ctx context.Context, in *rental.Review) (*rental.ReviewId, error) {
	log.Printf("Updating the Review")

	if !S.Db.UserAlreadyBooked(int(in.UserId), int(in.ItemId)) {
		return &rental.ReviewId{Id: 0}, nil
	}
	updatedReview := models.Review{
		Comments: in.Comment,
		Rating:   int(in.Rating),
		UserID:   int(in.UserId),
		ItemID:   int(in.ItemId),
	}

	updatedReview.ID = uint(in.Id)

	reviewId := S.Db.Update_Review(updatedReview)

	return &rental.ReviewId{Id: int32(reviewId)}, nil
}

func (S *ServerSideImplementation) DeleteReview(ctx context.Context, in *rental.ReviewId) (*rental.ReviewId, error) {

	deletedReviewId := S.Db.DeleteReview(int(in.Id))

	return &rental.ReviewId{Id: int32(deletedReviewId)}, nil
}

func (S *ServerSideImplementation) GetAllReviews(in *rental.ItemId, stream rental.Rental_Easy_Functionalities_GetAllReviewsServer) error {
	Reviews := S.Db.GetReviews(int(in.Id))

	var wg sync.WaitGroup
	for _, Review := range Reviews {
		wg.Add(1)
		go func(review models.Review) {
			defer wg.Done()
			time.Sleep(time.Duration(100) * time.Microsecond)
			Review := rental.Review{
				Id:      int32(review.ID),
				Comment: review.Comments,
				Rating:  int32(review.Rating),
				UserId:  int32(review.UserID),
				ItemId:  int32(review.ItemID),
			}
			err := stream.Send(&Review)
			checkErr(err)
		}(Review)
	}

	wg.Wait()
	return nil
}
