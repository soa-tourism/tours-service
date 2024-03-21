package service

import (
	"fmt"
	"tours/model"
	"tours/repo"
)

type PublishedTourService struct {
	TourRepo *repo.TourRepository
}

func (service *PublishedTourService) FindAll() ([]model.Tour, error) {
	tours, err := service.TourRepo.FindAllPublished()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tours: %v", err)
	}
	return tours, nil
}

func (service *PublishedTourService) Find(id int64) (*model.Tour, error) {
	tour, err := service.TourRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tour: %v", err)
	}

	if tour.Status != model.Published {
		return nil, fmt.Errorf("tour with ID %d is not published", id)
	}

	return tour, nil
}
