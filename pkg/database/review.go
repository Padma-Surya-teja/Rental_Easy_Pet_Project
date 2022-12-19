package database

import "rental_easy.in/m/pkg/models"

func (db DBClient) AddReview(review models.Review) int {
	db.Db.Create(&review)

	return int(review.ID)
}

// Updating the Review
func (db DBClient) Update_Review(review models.Review) int {
	db.Db.Save(&review)

	return int(review.ID)
}

func (db DBClient) DeleteReview(reviewId int) int {
	db.Db.Delete(&models.Review{}, reviewId)

	return reviewId
}

func (db DBClient) GetReviews(item_id int) []models.Review {
	Reviews := []models.Review{}
	db.Db.Find(&Reviews, "item_id = ?", item_id)
	return Reviews
}
