package database

import "rental_easy.in/m/pkg/models"

func (db DBClient) AddReview(review models.Review) (int, error) {
	result := db.Db.Create(&review)

	return int(review.ID), result.Error
}

// Updating the Review
func (db DBClient) UpdateReview(review models.Review) (int, error) {
	result := db.Db.Save(&review)

	return int(review.ID), result.Error
}

func (db DBClient) DeleteReview(id int) (int, error) {
	result := db.Db.Delete(&models.Review{}, id)

	return id, result.Error
}

func (db DBClient) GetReviews(id int) ([]models.Review, error) {
	reviews := []models.Review{}
	result := db.Db.Find(&reviews, "item_id = ?", id)
	return reviews, result.Error
}
