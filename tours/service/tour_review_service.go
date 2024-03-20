package service

import (
	"fmt"
	"tours/model"
	"tours/repo"
)

type TourReviewService struct {
	TourReviewRepo *repo.TourReviewRepository
}

func (service *TourReviewService) FindAllTourReviewsByTourist(id int64) ([]model.TourReview, error) {
	tourReview, err := service.TourReviewRepo.FindAllByTourist(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tour reviews: %v", err)
	}
	return tourReview, nil
}

func (service *TourReviewService) FindAllTourReviewsByAuthor(id int64) ([]model.TourReview, error) {
	tourReview, err := service.TourReviewRepo.FindAllByAuthor(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tour reviews: %v", err)
	}
	return tourReview, nil
}

func (service *TourReviewService) FindTourReview(id int64) (*model.TourReview, error) {
	tourReview, err := service.TourReviewRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("tour review with ID %d not found", id)
	}
	return tourReview, nil
}

func (service *TourReviewService) CreateTourReview(tourReview *model.TourReview) (*model.TourReview, error) {
	createdTourReview, err := service.TourReviewRepo.Create(tourReview)
	if err != nil {
		return nil, fmt.Errorf("failed to create tour review: %v", err)
	}
	return &createdTourReview, nil
}
