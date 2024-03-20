package repo

import (
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

func (repo *TourReviewRepository) Update(tourReview *model.TourReview) (model.TourReview, error) {
	dbResult := repo.DB.Save(tourReview)
	if dbResult.Error != nil {
		return model.TourReview{}, dbResult.Error
	}
	return *tourReview, nil
}
