package server

import (
	"context"
	"log"

	"rental_easy.in/m/pkg/models"
	rental "rental_easy.in/m/pkg/rentalmgmt"
	"rental_easy.in/m/pkg/utils"
)

func (S *ServerSideImplementation) AddReview(ctx context.Context, in *rental.Review) (*rental.Review, error) {
	log.Println("Server : Creating a New Review")

	_, err := S.Db.GetItemById(int(in.GetItemId()))
	utils.CheckErr(err)
	if err != nil {
		return &rental.Review{}, err
	}

	User, err := S.Db.UserAlreadyBooked(int(in.UserId), int(in.ItemId))
	utils.CheckErr(err)
	if !User || err != nil {
		return &rental.Review{}, err
	}

	review, err := S.Db.UserAlreadyAddedReview(int(in.UserId), int(in.ItemId))
	utils.CheckErr(err)
	if review {
		return &rental.Review{}, err
	}

	newReview := models.Review{
		Comments: in.Comment,
		Rating:   int(in.Rating),
		UserId:   int(in.UserId),
		ItemId:   int(in.ItemId),
	}

	reviewId, err := S.Db.AddReview(newReview)
	utils.CheckErr(err)

	return &rental.Review{Id: int32(reviewId)}, err
}

func (S *ServerSideImplementation) UpdateReview(ctx context.Context, in *rental.Review) (*rental.Review, error) {
	log.Printf("Server : Updating the Review")

	review := models.Review{
		Comments: in.Comment,
		Rating:   int(in.Rating),
		UserId:   int(in.UserId),
		ItemId:   int(in.ItemId),
	}

	review.ID = uint(in.Id)

	id, err := S.Db.UpdateReview(review)
	utils.CheckErr(err)

	return &rental.Review{Id: int32(id)}, err
}

func (S *ServerSideImplementation) DeleteReview(ctx context.Context, in *rental.Review) (*rental.Review, error) {

	id, err := S.Db.DeleteReview(int(in.Id))
	utils.CheckErr(err)

	return &rental.Review{Id: int32(id)}, err
}

func (S *ServerSideImplementation) GetAllReviews(in *rental.Item, stream rental.Rental_Easy_Functionalities_GetAllReviewsServer) error {
	reviews, err := S.Db.GetReviews(int(in.Id))
	utils.CheckErr(err)
	if err != nil {
		return err
	}

	for _, review := range reviews {
		Review := rental.Review{
			Id:      int32(review.ID),
			Comment: review.Comments,
			Rating:  int32(review.Rating),
			UserId:  int32(review.UserId),
			ItemId:  int32(review.ItemId),
		}
		err := stream.Send(&Review)
		utils.CheckErr(err)
	}

	return nil
}
