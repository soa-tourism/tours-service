package service

import (
	"fmt"
	"tours/model"
	"tours/repo"
)

type TourService struct {
	TourRepo *repo.TourRepository
}

func (service *TourService) FindAllTours() ([]model.Tour, error) {
	tours, err := service.TourRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tours: %v", err)
	}
	return tours, nil
}

func (service *TourService) FindTour(id int64) (*model.Tour, error) {
	tour, err := service.TourRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("tour with ID %d not found", id)
	}
	return tour, nil
}

func (service *TourService) CreateTour(tour *model.Tour) (*model.Tour, error) {
	createdTour, err := service.TourRepo.Create(tour)
	if err != nil {
		return nil, fmt.Errorf("failed to create tour: %v", err)
	}
	return &createdTour, nil
}

func (service *TourService) UpdateTour(tour *model.Tour) (*model.Tour, error) {
	updatedTour, err := service.TourRepo.Update(tour)
	if err != nil {
		return nil, fmt.Errorf("failed to update tour: %v", err)
	}
	return &updatedTour, nil
}

func (service *TourService) DeleteTour(id int64) error {
	err := service.TourRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete tour with ID %d: %v", id, err)
	}
	return nil
}

func (service *TourService) FindToursByAuthor(id int64) ([]model.Tour, error) {
	tours, err := service.TourRepo.FindByAuthor(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tours: %v", err)
	}
	return tours, nil
}

func (service *TourService) AddEquipment(tourId int64, equipmentId int64) error {
	if err := service.TourRepo.AddEquipment(tourId, equipmentId); err != nil {
		return err
	}
	return nil
}

func (service *TourService) RemoveEquipment(tourId int64, equipmentId int64) error {
	if err := service.TourRepo.RemoveEquipment(tourId, equipmentId); err != nil {
		return err
	}
	return nil
}
