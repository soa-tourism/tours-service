package repo

import (
	"database/sql"
	"tours/model"

	"gorm.io/gorm"
)

type TourReviewRepository struct {
	DB *gorm.DB
}

func (repo *TourReviewRepository) FindAllByTourist(id int64) ([]model.TourReview, error) {
	var tourReviews []model.TourReview
	dbResult := repo.DB.Where("tourist_id = ?", id).Find(&tourReviews)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tourReviews, nil
}

func (repo *TourReviewRepository) FindAllByAuthor(id int64) ([]model.TourReview, error) {
	var tourReviews []model.TourReview
	dbResult := repo.DB.
		Joins("JOIN tours ON tour_reviews.tour_id = tours.id").
		Where("tours.author_id = ?", id).
		Find(&tourReviews)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tourReviews, nil
}

func (repo *TourReviewRepository) FindAllByTour(id int64) ([]model.TourReview, error) {
	var tourReviews []model.TourReview
	dbResult := repo.DB.Where("tour_id = ?", id).Find(&tourReviews)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tourReviews, nil
}

func (repo *TourReviewRepository) FindById(id int64) (*model.TourReview, error) {
	tourReview := model.TourReview{}
	dbResult := repo.DB.First(&tourReview, id)
	if dbResult.Error != nil {
		return &tourReview, dbResult.Error
	}
	return &tourReview, nil
}

func (repo *TourReviewRepository) Create(tourReview *model.TourReview) (model.TourReview, error) {
	dbResult := repo.DB.Create(tourReview)
	if dbResult.Error != nil {
		return model.TourReview{}, dbResult.Error
	}
	return *tourReview, nil
}

func (repo *TourReviewRepository) GetAverageRating(id int64) (float64, error) {
	var averageRating sql.NullFloat64
	dbResult := repo.DB.Table("tour_reviews").
		Select("AVG(rating) AS average_rating").
		Where("tour_id = ?", id).
		Scan(&averageRating)
	if dbResult.Error != nil {
		return 0, dbResult.Error
	}

	if averageRating.Valid {
		return averageRating.Float64, nil
	}

	return 0, nil
}
