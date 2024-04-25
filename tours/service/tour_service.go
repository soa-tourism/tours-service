package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tours/model"
	"tours/repo"
)

type TourService struct {
	TourRepo      *repo.TourRepository
	EquipmentRepo *repo.EquipmentRepository
}

func (service *TourService) FindAllTours() ([]model.Tour, error) {
	tours, err := service.TourRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tours: %v", err)
	}
	return tours, nil
}

func (service *TourService) FindTour(id primitive.ObjectID) (*model.Tour, error) {
	tour, err := service.TourRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("tour with ID %s not found", id.Hex())
	}
	return tour, nil
}

func (service *TourService) CreateTour(tour *model.Tour) (*model.Tour, error) {
	createdTour, err := service.TourRepo.Create(tour)
	if err != nil {
		return nil, fmt.Errorf("failed to create tour: %v", err)
	}
	return createdTour, nil
}

func (service *TourService) UpdateTour(tour *model.Tour) (*model.Tour, error) {
	updatedTour, err := service.TourRepo.Update(tour)
	if err != nil {
		return nil, fmt.Errorf("failed to update tour: %v", err)
	}
	return updatedTour, nil
}

func (service *TourService) DeleteTour(id primitive.ObjectID) error {
	err := service.TourRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete tour with ID %s: %v", id.Hex(), err)
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

func (service *TourService) AddEquipment(tourId primitive.ObjectID, equipmentId primitive.ObjectID) error {
	equipment, err := service.EquipmentRepo.FindById(equipmentId)
	if err != nil {
		return err
	}

	tour, err := service.TourRepo.FindById(tourId)
	if err != nil {
		return err
	}

	tour.Equipment = append(tour.Equipment, *equipment)

	if _, err := service.TourRepo.Update(tour); err != nil {
		return err
	}

	return nil
}

func (service *TourService) RemoveEquipment(tourId primitive.ObjectID, equipmentId primitive.ObjectID) error {
	if err := service.TourRepo.RemoveEquipment(tourId, equipmentId); err != nil {
		return err
	}
	return nil
}
